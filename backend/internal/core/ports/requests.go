package ports

type CreateCalendarRequest struct {
	Name       		string 		`json:"name"`
}

type UpdateCalendarRequest struct {
	Id					string  	`json:"id"`
	Name       			string 		`json:"name"`
}

type CreateEventRequest struct {
	Name       		string 		`json:"name"`
}

type UpdateEventRequest struct {
	Id					string  	`json:"id"`
	Name       			string 		`json:"name"`
}

type CreateUserRequest struct {
	Name       		string 		`json:"name"`
}

type UpdateUserRequest struct {
	Id					string  	`json:"id"`
	Name       			string 		`json:"name"`
}
