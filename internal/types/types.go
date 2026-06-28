package types

type Student struct {
	Id     int64  `json:"id"`
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required"`
	Enroll string `json:"enroll" validate:"required"`
	Age    int    `json:"age" validate:"required"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
