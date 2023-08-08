package model

type CreateStudentResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
}
