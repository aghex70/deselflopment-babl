package ports

type CreateCalendarRequest struct {
	Name       		string 		`json:"name"`
}

type UpdateCalendarRequest struct {
	Id					string  	`json:"id"`
	Name       			string 		`json:"name"`
}

type CreateEntryRequest struct {
	Name       		string 		`json:"name"`
}

type UpdateEntryRequest struct {
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
