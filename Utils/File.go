package Utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"mime/multipart"
	"strconv"
	"time"
)

func GetDatetimeMillisecond() (string, error) {
	loc, err := time.LoadLocation("Asia/Seoul")
    if err != nil {
        panic(err)
    }
	now := time.Now()
	t := now.In(loc)
	t = t.Add(time.Millisecond)
	y, m, d := t.Date()
	td := (int64(y)*100+int64(m))*100 + int64(d)
	hh, mm, ss := t.Clock()
	tc := (int64(hh)*100+int64(mm))*100 + int64(ss)
	ms := t.Nanosecond() / int(time.Millisecond)
	tm := (td*1000000+tc)*1000 + int64(ms)
	return strconv.FormatInt(tm, 10), nil
}

func GetFileReader(file *multipart.FileHeader) (*bytes.Reader, error) {
	openedFile, _ := file.Open()
	binaryFile, err := ioutil.ReadAll(openedFile)
	if err != nil {
		return nil, err
	}

	defer func(openedFile multipart.File) {
		err := openedFile.Close()
		if err != nil {
			log.Fatalf("Failed closing file %v", file.Filename)
		}
	}(openedFile)

	return bytes.NewReader(binaryFile), nil
}

func GetHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}
