package model

type Comment struct {
	MID     int64  `json:"Mid"`
	RecMID  string `json:"Rec_Mid"`
	Sender  string `json:"Sender"`
	Details string `json:"Details"`
}
