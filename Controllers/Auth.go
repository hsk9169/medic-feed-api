package Controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/medic-basic/s3-test/Services"
	"github.com/medic-basic/s3-test/Utils"
	"github.com/medic-basic/s3-test/Validations"
)

func PostAuthImage(c *gin.Context) {
	var err error
	req := Validations.PostAuthImageRequest{}
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

	var imgUrl string
	imgUrl, err = Services.UploadAuthImage(req.MedicId, req.File, curTime)
	if err != nil {
		Services.ErrorResponse(c, 500, "Failed to post auth image")
		return
	}
	fmt.Println(imgUrl)

	Services.Created(c, "Create Feed Success", imgUrl)
}

func GetAuthImage(c *gin.Context) {
	var err error
	medicId := c.Param("medicId")
	fmt.Printf("MedicId from Uri: %s", medicId)

	req := Validations.GetAuthIamgeRequest{}
	if err = c.ShouldBindQuery(&req); err != nil {
		fmt.Println(err)
		Services.ErrorResponse(c, 400, err)
		return
	}

	var fileData []byte
	fileData, err = Services.DownloadAuthImage(medicId, req.ImgUrl)
	if err != nil {
		fmt.Println("Failed to download file")
		Services.ErrorResponse(c, 404, err)
		return
	}

	authImageResp := Validations.GetAuthImageResponse{ImageBase64: fileData}

	Services.Success(c, "Success", authImageResp)
}
