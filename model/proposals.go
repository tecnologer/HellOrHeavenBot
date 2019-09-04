package model

//Proposals struct for proposals to be used as responses
type Proposals struct {
	Downvote uint     `json:"downvote"`
	Upvote   uint     `json:"upvote"`
	Proposal Response `json:"proposal"`
	Voters   []int    `json:"voters"`
}
