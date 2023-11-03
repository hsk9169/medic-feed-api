package Validations

import (
	"mime/multipart"
)

type PostAuthImageRequest struct {
	File    *multipart.FileHeader `form:"file" binding:"required"`
	MedicId string                `form:"medicId" binding:"required"`
}

type GetAuthIamgeRequest struct {
	ImgUrl string `json:"imgUrl"`
}

type GetAuthImageResponse struct {
	ImageBase64 []byte
}
