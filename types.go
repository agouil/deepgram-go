package deepgram

// package to define all custom structs

type Deepgram struct {
	ApiKey string
}

func (dg *Deepgram) Host() string {
	return "http://api.deepgram.com"
}

func (dg *Deepgram) GroupSearchHost() string {
	return "http://groupsearch.api.deepgram.com"
}

type ResponseError struct {
	Error string `json:"error"`
}

type GetObjectInfoRequest struct {
	Action    string `json:"action"`
	UserId    string `json:"userID"`
	ContentId string `json:"contentID"`
}

type CheckBalanceRequest struct {
	Action string `json:"action"`
	UserId string `json:"userID"`
}

type CheckBalanceResponse struct {
	Balance float32 `json:"balance"`
	UserId  string  `json:"userID"`
}

type CheckStatusResponse struct {
	Status string `json:"status"`
}

type TagRequest struct {
	Action    string `json:"action"`
	UserId    string `json:"userID"`
	ContentId string `json:"contentID"`
	Tag       string `json:"tag"`
}

type TagResponse struct {
	Result string `json:"result"`
}

type GetTagsResponse struct {
	ContentId string   `json:"contentID"`
	Tags      []string `json:"tags"`
}
