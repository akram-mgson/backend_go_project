package main


type Order struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	StageID string `json:"stageId"`
}


type Info struct {
	PhoneNumber string `json:"phone_number"`
	Email        string `json:"email"`
	Manager      string `json:"manager"`
}

type ErrorResp struct{
	Code	string `json:"code"`
	Message string `json:"message"`
}

type SuccessResp struct{
	Code	string `json:"code"`
	Message string `json:"message"`
}


type Comment struct{
	Text	string `json:"text"`

	}
