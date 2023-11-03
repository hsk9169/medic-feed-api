package Validations

import (
	"mime/multipart"
	"time"
)

type FeedMeta struct {
	AuthorId    string   `json:"authorId" binding:"required"`
	PostContent string   `json:"postContent" binding:"required"`
	TagIDs      []string `json:"tagIDs" binding:"required"`
}

type PostFeedRequest struct {
	Files       []*multipart.FileHeader `form:"files" binding:"required"`
	OriginFiles []*multipart.FileHeader `form:"originFiles" binding:"required"`
	PatientId   string                  `form:"patientId" binding:"required"`
	FeedMeta    FeedMeta                `form:"feedMeta" binding:"required"`
}

type PostFeedToBasicApi struct {
	PatientId   string   `json:"patientId"`
	AuthorId    string   `json:"authorId"`
	PostContent string   `json:"postContent"`
	ImageUrls   []string `json:"imageUrls"`
	TagIDs      []string `json:"tagIDs"`
}

type FeedFromBasicApi struct {
	CreateDate  time.Time
	AuthorName  string
	PostContent string
	ImageUrls   []string
}

type FeedListResponseData struct {
	CreateDate      time.Time
	AuthorName      string
	PostContent     string
	ImageUrls		[]string
	ImageBase64 	[]byte
}

type GetFeedListRequest struct {
	Limit  int    `form:"limit,default=100"`
	Cursor string `form:"cursor"`
}

type GetFeedListToBasicApi struct {
	Limit  int
	Cursor string
}

type GetFeedListFromBasicApi struct {
	NextCursor string             `json:"nextCursor"`
	FeedList   []FeedFromBasicApi `json:"feedList"`
}

type GetFeedListResponse struct {
	NextCursor   string
	FeedDataList []FeedListResponseData
}