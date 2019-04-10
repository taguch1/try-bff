package model

// TodoSaveRequest save request
type TodoSaveRequest struct {
	Titile string `json:"title"`
}

// TodoGetRequest get request
type TodoGetRequest struct {
	ID string `json:"id"`
}

// TodoListRequest list request
type TodoListRequest struct {
	Offset int `json:"offset"`
	Length int `json:"langth"`
}

// TodoUpdateRequest update request
type TodoUpdateRequest struct {
	ID    string `json:"id"`
	Title string `json:"titile"`
}

// TodoDeleteRequest delete request
type TodoDeleteRequest struct {
	ID string `json:"id"`
}

// TodoResponse response
type TodoResponse struct {
	ID    string `json:"id"`
	Title string `json:"titile"`
}

// TodoListResponse list response
type TodoListResponse struct {
	Todos []*TodoResponse `json:"todos"`
}
