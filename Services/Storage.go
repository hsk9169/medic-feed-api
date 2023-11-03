package Services

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/medic-basic/s3-test/Utils"
)

func ListBuckets() error {
	buckets, err := StorageImpl.s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return err
	}
	fmt.Println(buckets)
	return nil
}

func UploadFeedImages(patientId string, files []*multipart.FileHeader, curTime string) ([]string, error) {
	basePath := StorageImpl.feedPrefix + patientId
	var url string
	var imgUrls []string
	for idx, file := range files {
		fileBuffer, err := Utils.GetFileReader(file)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Failed to get file reader from file: %d", idx)
			return nil, errors.New("file read error")
		}
		folderPath := Utils.GetHash(patientId + curTime)
		filePath := Utils.GetHash(curTime + strconv.Itoa(idx))
		url = "/" + folderPath + "/" + filePath
		imgUrls = append(imgUrls, url)
		bucketName := StorageImpl.feedBucketName
		_, err = StorageImpl.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(basePath + url),
			Body:   fileBuffer,
		})
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Failed to upload file to S3: %d", idx)
			return nil, errors.New("s3 upload error")
		}
	}

	return imgUrls, nil
}

func UploadOriginImages(patientId string, originFiles []*multipart.FileHeader, curTime string) error {
	basePath := StorageImpl.feedPrefix + patientId
	var url string
	for idx, file := range originFiles {
		fileBuffer, err := Utils.GetFileReader(file)
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Failed to get file reader from file: %d", idx)
			return errors.New("file read error")
		}
		folderPath := Utils.GetHash(patientId + curTime)
		filePath := Utils.GetHash(curTime + strconv.Itoa(idx))
		url = "/origin/" + folderPath + "/" + filePath
		bucketName := StorageImpl.feedBucketName
		_, err = StorageImpl.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(basePath + url),
			Body:   fileBuffer,
		})
		if err != nil {
			fmt.Println(err)
			fmt.Printf("Failed to upload file to S3: %d", idx)
			return errors.New("s3 upload error")
		}
	}

	return nil
}

func DownloadFeedImage(patientId string, imgUrl string) ([]byte, error) {
	bucketName := StorageImpl.feedBucketName
	basePath := StorageImpl.feedPrefix + patientId
	var fileData []byte

	result, err := StorageImpl.s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(basePath + imgUrl),
	})
	if err != nil {
		fmt.Printf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, imgUrl, err)
		fileData = nil
	} else {
		defer result.Body.Close()
		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			fmt.Printf("Couldn't read object body from %v. Here's why: %v\n", imgUrl, err)
			fileData = nil
		}
		fileData = body
	}

	return fileData, nil
}

func UploadAuthImage(medicId string, file *multipart.FileHeader, curTime string) (string, error) {
	basePath := StorageImpl.authPrefix + medicId
	var url string
	var imgUrl string
	fileBuffer, err := Utils.GetFileReader(file)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Failed to get file reader from file")
		return "", errors.New("file read error")
	}
	filePath := Utils.GetHash(medicId + curTime)
	url = "/" + filePath
	imgUrl = url
	bucketName := StorageImpl.authBucketName
	_, err = StorageImpl.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(basePath + url),
		Body:   fileBuffer,
	})
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Failed to upload file to S3")
		return "", errors.New("s3 upload error")
	}

	return imgUrl, nil
}

func DownloadAuthImage(medicId string, imgUrl string) ([]byte, error) {
	bucketName := StorageImpl.authBucketName
	basePath := StorageImpl.authPrefix + medicId
	var fileData []byte

	result, err := StorageImpl.s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(basePath + imgUrl),
	})
	if err != nil {
		fmt.Printf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, imgUrl, err)
		return nil, errors.New("failed to get s3 file")
	} else {
		defer result.Body.Close()
		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			fmt.Printf("Couldn't read object body from %v. Here's why: %v\n", imgUrl, err)
			return nil, errors.New("failed to read data from body")
		}
		fileData = body
	}

	return fileData, nil
}
