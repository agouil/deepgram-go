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

type CheckBalanceRequest struct {
	Action string `json:"action"`
	UserId string `json:"userID"`
}

type CheckBalanceResponse struct {
	Balance float32 `json:"balance"`
	UserId  string  `json:"userID"`
}
