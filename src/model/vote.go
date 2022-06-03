package model

type Crypto struct {
	Name     string `json:"name"`
	Upvote   bool   `json:"upvote"`
	Downvote bool   `json:"downvote"`
}

type Data struct {
	Name     string `json:"name"`
	Upvote   int    `json:"upvote"`
	Downvote int    `json:"downvote"`
}

type Record struct {
	Crypto   string `json:"crypto"`
	Upvote   int    `json:"upvote"`
	Downvote int    `json:"downvote"`
}

func UnmarshalVote(name string, vote string) *Crypto {
	unmarshal := &Crypto{
		Name: name,
	}
	if vote == "upvote" {
		unmarshal.Upvote = true
		unmarshal.Downvote = false
		return unmarshal
	} else {
		unmarshal.Upvote = false
		unmarshal.Downvote = true
		return unmarshal
	}
}

func UnmarshalCrypto(name string) *Data {
	unmarshal := &Data{
		Name:     name,
		Upvote:   0,
		Downvote: 0,
	}
	return unmarshal
}
