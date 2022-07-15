package types

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type LoginDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
type TaskDto struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
type LoginResponse struct {
	Response
	Result LoginDto
}
type ListResponse struct {
	Response
	Result []TaskDto
}
type CreateResponse struct {
	Response
	Result struct {
		Id string `json:"id"`
	}
}
type UpdateResponse struct {
	Response
	Result struct {
	}
}
