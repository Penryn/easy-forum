package model

type Report struct{
	ID int `json:"id"`
	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
	Reason string `json:"reason"`
	Status int `json:"status"`
}