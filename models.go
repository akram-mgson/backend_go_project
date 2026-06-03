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

type Document struct {
	ID    int 		`json:"id"`
	Title string	`json:"title"`
	Body  string	`json:"body"`
}



type Comment struct{
	Text	string `json:"text"`

	}
type Document struct {
	ID    int
	Title string
	Body  string
}

type Auth struct{
	Name	string	`json:"name"`
	Email	string	`json:"i_ivanov@test.com"`	
}
