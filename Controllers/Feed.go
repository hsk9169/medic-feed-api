package Controllers

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/medic-basic/s3-test/Services"
	"github.com/medic-basic/s3-test/Utils"
	"github.com/medic-basic/s3-test/Validations"
)

func PostFeed(c *gin.Context) {
	var err error
	req := Validations.PostFeedRequest{}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		Services.ErrorResponse(c, 400, err)
		return
	}

	curTime, err := Utils.GetDatetimeMillisecond()
	if err != nil {
		fmt.Println("Failed to get current time")
		Services.ErrorResponse(c, 500, "Internal server error")
		return
	}

	err = Services.UploadOriginImages(req.PatientId, req.Files, curTime)
	if err != nil {
		Services.ErrorResponse(c, 500, "Failed to post feeds")
		return
	}

	var imgUrls []string
	imgUrls, err = Services.UploadFeedImages(req.PatientId, req.Files, curTime)
	if err != nil {
		Services.ErrorResponse(c, 500, "Failed to post feeds")
		return
	}
	fmt.Println(imgUrls)

	postFeedToBasic := Validations.PostFeedToBasicApi{
		PatientId:   req.PatientId,
		AuthorId:    req.FeedMeta.AuthorId,
		PostContent: req.FeedMeta.PostContent,
		ImageUrls:   imgUrls,
		TagIDs:      req.FeedMeta.TagIDs,
	}
	if err = Services.CreateBasicFeed(postFeedToBasic); err != nil {
		fmt.Println(err)
		Services.ErrorResponse(c, 500, "Failed to create feed data")
		// Add file removal job to Kafka server
		return
	}

	Services.Created(c, "Create Feed Success", nil)
}

func GetFeedList(c *gin.Context) {
	var err error
	patientId := c.Param("patientId")
	fmt.Printf("PatientId from Uri: %s", patientId)
	req := Validations.GetFeedListRequest{}
	if err = c.ShouldBindQuery(&req); err != nil {
		fmt.Println(err)
		Services.ErrorResponse(c, 400, err)
		return
	}

	body, err := Services.GetFeedList(patientId, req)
	if err != nil {
		fmt.Println(err)
		Services.ErrorResponse(c, 400, err)
		return
	}
	var feedResponseFromBasicApi Validations.GetFeedListFromBasicApi
	if err = json.Unmarshal(body, &feedResponseFromBasicApi); err != nil {
		fmt.Println("Failed to unmarshal response json body")
		Services.ErrorResponse(c, 500, "internal server error")
		return
	}
	feedListFromBasicApi := feedResponseFromBasicApi.FeedList

	var feedList []Validations.FeedListResponseData
	var feedListResp Validations.GetFeedListResponse

	var fileData []byte
	for _, feed := range feedListFromBasicApi {
		if (len(feed.ImageUrls) > 0) {
			fileData, err = Services.DownloadFeedImage(patientId, feed.ImageUrls[0])
		} else {
			fileData = nil
			err = nil
		}
		if err != nil {
			fmt.Println("Failed to download file")
			Services.ErrorResponse(c, 500, err)
			return
		}
		feedList = append(feedList, Validations.FeedListResponseData{
			CreateDate:      feed.CreateDate,
			AuthorName:      feed.AuthorName,
			PostContent:     feed.PostContent,
			ImageUrls:		feed.ImageUrls,
			ImageBase64: 	fileData})
	}
	feedListResp = Validations.GetFeedListResponse{
		NextCursor:   feedResponseFromBasicApi.NextCursor,
		FeedDataList: feedList,
	}

	Services.Success(c, "Success", feedListResp)
}

func GetFeedImageList(c *gin.Context) {
	patientId := c.Param("patientId")
	imageUrls, flag := c.GetQueryArray("imageUrls")
	if !flag {
		Services.ErrorResponse(c, 400, "invalid request")
		return
	}

	var imgList [][]byte
	for _, url := range imageUrls {
		fileData, err := Services.DownloadFeedImage(patientId, url)
		if err != nil {
			fmt.Println("Failed to download file")
			Services.ErrorResponse(c, 500, err)
			return
		}
		imgList = append(imgList, fileData)
	}
	
	Services.Success(c, "Success", imgList)
}