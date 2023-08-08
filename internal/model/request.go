package model

type CreateStudentRequest struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
}
