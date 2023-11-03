package Services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/medic-basic/s3-test/Validations"
)

var BasicImpl Basic

type Basic struct {
	url            string
	scheme         string
	createFeedPath string
	getFeedPath    string
}

func InitBasic() {
	BasicImpl = Basic{
		url:            "{{basic_api_dns}}",
		scheme:         "http",
		createFeedPath: "/feed/create",
		getFeedPath:    "/feed",
	}
}

func CreateBasicFeed(reqData Validations.PostFeedToBasicApi) error {
	var err error
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	resp, err := http.Post(BasicImpl.url+BasicImpl.createFeedPath, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	str := string(respBody)
	fmt.Printf("Basic Api response status code: %d", resp.StatusCode)
	fmt.Printf("Body: %s", str)
	if resp.StatusCode >= 400 {
		fmt.Println(str)
		return errors.New("internal server error")
	}
	return nil
}

func GetFeedList(patientId string, reqData Validations.GetFeedListRequest) ([]byte, error) {
	var err error
	params := url.Values{}
	params.Add("limit", strconv.Itoa(reqData.Limit))
	params.Add("cursor", reqData.Cursor)
	u, _ := url.ParseRequestURI(BasicImpl.url)
	u.Path = BasicImpl.getFeedPath + "/" + patientId
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	resp, err := http.Get(urlStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	str := string(respBody)
	fmt.Printf("Basic Api response status code: %d", resp.StatusCode)
	fmt.Printf("Body: %s", str)
	if resp.StatusCode >= 400 {
		fmt.Println(str)
		return nil, errors.New("internal server error")
	}
	return respBody, nil
}
