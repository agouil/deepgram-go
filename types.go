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
