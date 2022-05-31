package model

type Crypto struct {
	Name     int  `json:"name,omitempty"`
	Upvote   bool `json:"upvote,omitempty"`
	Downvote bool `json:"downvote,omitempty"`
}

type Votes struct {
	ID      string    `json:"id,omitempty"`
	Author  string    `json:"author,omitempty"`
	MyVotes [5]Crypto `json:"cryptos,omitempty"`
}
