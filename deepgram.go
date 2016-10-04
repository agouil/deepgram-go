package deepgram

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

//
// Helper functions
//

func makeRequest(url string, payload interface{}) ([]byte, error) {
	reqJson, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(reqJson))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	//TODO: check the response code first and if OK then continue

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func parseResponse(response []byte, t interface{}) error {
	respErr := new(ResponseError)
	err := json.Unmarshal(response, respErr)
	if err != nil {
		return err
	}
	if respErr.Error != "" {
		return errors.New(respErr.Error)
	}
	err = json.Unmarshal(response, t)
	if err != nil {
		return err
	}
	return nil
}

//
// API Methods
//

func (dg *Deepgram) CheckBalance() (*CheckBalanceResponse, error) {
	req := CheckBalanceRequest{
		Action: "get_balance",
		UserId: dg.ApiKey,
	}
	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := makeRequest(dg.Host(), string(reqJson))
	if err != nil {
		return nil, err
	}
	result := new(CheckBalanceResponse)
	err = parseResponse(resp, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// func (dg *Deepgram) CheckStatus(obj string) (*CheckStatusResponse, error) {
// 	return nil, nil
// }

// func (dg *Deepgram) Upload(mediaUrl string, tags []string) (*UploadResponse, error) {
// 	return nil, nil
// }

// func (dg *Deepgram) UploadList(mediaUrls []string) (*UploadListResponse, error) {
// 	return nil, nil
// }

// func (dg *Deepgram) Query(obj, query string, options *QueryRequestParameters) (*QueryResponse, error) {
// 	return nil, nil
// }

// func (dg *Deepgram) GroupSearch(query, tag string) (*GroupSearchResponse, error) {
// 	return nil, nil
// }

// func (dg *Deepgram) ParallelSearch(query string, options *ParallelSearchParameters) (*ParallelSearchResponse, error) {
// 	return nil, nil
// }

// func (dg *Deepgram) Tag(obj, tag string) (*TagResponse, error) {
// 	return nil, nil
// }

// func (dg *Deepgram) GetTags(obj string) (*GetTagsResponse, error) {
// 	return nil, nil
// }

// func (dg *Deepgram) Transcript(obj string) (*TranscriptResponse, error) {
// 	return nil, nil
// }
