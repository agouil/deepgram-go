package deepgram

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//
// Helper functions
//

func makeRequest(url string, payload string) (string, error) {
	request, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	//TODO: check the response code first and if OK then continue

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func parseResponse(response string, t interface{}) error {
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(response))
	err := jsonDecoder.Decode(t)
	if err != nil {
		return err
	}
	return nil
}

//
// API Methods
//

func (dg *Deepgram) CheckBalance() (*CheckBalanceResponse, error) {
	return nil, nil
}

func (dg *Deepgram) CheckStatus(obj string) (*CheckStatusResponse, error) {
	return nil, nil
}

func (dg *Deepgram) Upload(mediaUrl string, tags []string) (*UploadResponse, error) {
	return nil, nil
}

func (dg *Deepgram) UploadList(mediaUrls []string) (*UploadListResponse, error) {
	return nil, nil
}

func (dg *Deepgram) Query(obj string, query string, options *QueryRequestParameters) (*QueryResponse, error) {
	return nil, nil
}

func (dg *Deepgram) GroupSearch(query string, tag string) (*GroupSearchResponse, error) {
	return nil, nil
}

func (dg *Deepgram) ParallelSearch(query string, options *ParallelSearchParameters) (*ParallelSearchResponse, error) {
	return nil, nil
}

func (dg *Deepgram) Tag(obj string, tag string) (*TagResponse, error) {
	return nil, nil
}

func (dg *Deepgram) GetTags(obj string) (*GetTagsResponse, error) {
	return nil, nil
}

func (dg *Deepgram) Transcript(obj string) (*TranscriptResponse, error) {
	return nil, nil
}
