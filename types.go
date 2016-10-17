package deepgram

// Deepgram represents a Deepgram API client which holds information
// about the user's API key and the API endpoints
type Deepgram struct {
	ApiKey string
}

// Host returns the main Deepgram API endpoint
func (dg *Deepgram) Host() string {
	return "http://api.deepgram.com"
}

// GroupSearchHost returns the Deepgram API endpoint for group
// and parallel search requests
func (dg *Deepgram) GroupSearchHost() string {
	return "http://groupsearch.api.deepgram.com"
}

// ResponseError represents an error returned by the API.
// This error will return a HTTP 200 code and the error will
// be encoded in the response JSON.
type ResponseError struct {
	Error string `json:"error"`
}

// getObjectInfoRequest represents a request payload to get
// the object's information from the API
type getObjectInfoRequest struct {
	Action    string `json:"action"`
	UserId    string `json:"userID"`
	ContentId string `json:"contentID"`
}

// FilterParameters are the parameters used to filter a query request
type FilterParameters struct {
	Nmax int32   `json:"Nmax"`
	Pmin float32 `json:"Pmin"`
}

// GroupFilterParameters are the parameters used to filter a
// group search request
type GroupFilterParameters struct {
	Nmax int32 `json:"Nmax"`
}

// checkBalanceRequest represents a request payload to get the
// user's balance from the API
type checkBalanceRequest struct {
	Action string `json:"action"`
	UserId string `json:"userID"`
}

// CheckBalanceResponse represents the response for the
// "get_balance" API action
type CheckBalanceResponse struct {
	Balance float32 `json:"balance"`
	UserId  string  `json:"userID"`
}

// CheckStatusResponse represents the response for the
// "get_object_status" API action
type CheckStatusResponse struct {
	Status string `json:"status"`
}

// uploadRequest represents the request payload to upload
// a new object to the API
type uploadRequest struct {
	Action  string   `json:"action"`
	UserId  string   `json:"userID"`
	DataUrl string   `json:"data_url"`
	Tags    []string `json:"tags"`
}

// UploadResponse represents the response for the "index_content" API action
type UploadResponse struct {
	ContentId string `json:"contentID"`
}

// uploadListRequest represents the request payload to upload a list
// of objectst to the API
type uploadListRequest struct {
	Action  string   `json:"action"`
	UserId  string   `json:"userID"`
	DataUrl []string `json:"data_url"`
}

// UploadListResponse represents the response for the "index_content_list"
// API action
type UploadListResponse struct {
	ContentId []string `json:"contentID"`
}

// tagRequest represents the request payload to add a tag to an object
type tagRequest struct {
	Action    string `json:"action"`
	UserId    string `json:"userID"`
	ContentId string `json:"contentID"`
	Tag       string `json:"tag"`
}

// TagResponse represents the response for the "tag_object" API action
type TagResponse struct {
	Result string `json:"result"`
}

// GetTagsResponse represents the response for the
// "get_object_tags" API action
type GetTagsResponse struct {
	ContentId string   `json:"contentID"`
	Tags      []string `json:"tags"`
}

// TranscriptResponse represents the response for the
// "get_object_transcript" API action
type TranscriptResponse struct {
	ContentId           string    `json:"contentID"`
	Paragraphs          []string  `json:"paragraphs"`
	ParagraphStartTimes []float32 `json:"paragraphStartTimes"`
}

// QueryRequestParameters represents the user defined
// request parameters passed to the query function
type QueryRequestParameters struct {
	Snippet *bool
	Nmax    *int32
	Pmin    *float32
	Sort    *string
}

// querySearchRequest represents the request payload for a query request
type querySearchRequest struct {
	Action    string           `json:"action"`
	UserId    string           `json:"userID"`
	ContentId string           `json:"contentID"`
	Query     string           `json:"query"`
	Sort      string           `json:"sort"`
	Snippet   bool             `json:"snippet"`
	Filter    FilterParameters `json:"filter"`
}

// QueryResponse represents the response for the
// "object_search" API action
type QueryResponse struct {
	Snippet   []string  `json:"snippet"`
	StartTime []float32 `json:"startTime"`
	EndTime   []float32 `json:"endTime"`
	P         []float32 `json:"P"`
	N         []int32   `json:"N"`
}

// groupSearchRequest represents the request payload for a group search request
type groupSearchRequest struct {
	Action string `json:"action"`
	UserId string `json:"userID"`
	Tag    string `json:"tag"`
	Query  string `json:"query"`
}

// GroupSearchResponse represents the response for the "group_search" API action
type GroupSearchResponse struct {
	ContentId []string  `json:"contentID"`
	P         []float32 `json:"P"`
	N         []int32   `json:"N"`
}

// ParallelSearchParameters represents the user defined
// request parameters passed to the parallel search function
type ParallelSearchParameters struct {
	Snippet    *bool
	Tag        *string
	GroupNmax  *int32
	ObjectNmax *int32
	ObjectPmin *float32
	Sort       *string
}

// parallelSearchRequest represents the request payload for the
// parallel search request
type parallelSearchRequest struct {
	Action       string                `json:"action"`
	UserId       string                `json:"userID"`
	Query        string                `json:"query"`
	Tag          string                `json:"tag"`
	GroupFilter  GroupFilterParameters `json:"group_filter"`
	ObjectFilter FilterParameters      `json:"object_filter"`
	Snippet      bool                  `json:"snippet"`
	Sort         string                `json:"sort"`
}

// ObjectResult represents an object that matched the parallel search criteria
type ObjectResult struct {
	ContentId string    `json:"contentID"`
	Snippet   []string  `json:"snippet"`
	StartTime []float32 `json:"startTime"`
	EndTime   []float32 `json:"endTime"`
	N         []int32   `json:"N"`
	P         []float32 `json:"P"`
}

// ParallelSearchResponse represents the response for the
// "parallel_search" API action
type ParallelSearchResponse struct {
	ObjectResult []ObjectResult `json:"object_result"`
}
