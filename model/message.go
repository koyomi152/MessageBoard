package model

type Message struct {
	MID      int64  `json:"MID"`
	SendName string `json:"SendName"`
	Receiver string `json:"Receiver"`
	Details  string `json:"Details"`
}
