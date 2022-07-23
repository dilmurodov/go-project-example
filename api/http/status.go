package http

type Status struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

var (
	OK = Status{
		Code: 200,
		Status: "OK",
		Description: "The request has successed",
	}
	Created = Status{
		Code: 201,
		Status: "CREATED",
		Description: "The request fullfilled and has resulted in one or more new resources being created",
	}
	BadRequest = Status{
		Code:        400,
		Status:      "BAD_REQUEST",
		Description: "The server could not understand the request due to invalid syntax",
	}
	InvalidArgument = Status{
		Code:        400,
		Status:      "INVALID_ARGUMENT",
		Description: "Invalid argument value passed",
	}
)