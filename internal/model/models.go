package model

import "time"

type QuestionAnswer struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Signature struct {
	User      string           `json:"user"`
	Signature string           `json:"signature"`
	Answers   []QuestionAnswer `json:"answers"`
	Timestamp time.Time        `json:"timestamp"`
}