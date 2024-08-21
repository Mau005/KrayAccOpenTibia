package models

type AnswerExpected struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Token    string `json:"token"`
}
