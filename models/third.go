package models


type UserInfo struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone []string `json:"phone"`
}

type GetUserResp struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Id    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	} `json:"data"`
}

type CreateUserResp struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	} `json:"data"`
}

type T struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message struct {
		Email []string `json:"email"`
		Phone []string `json:"phone"`
	} `json:"message"`
	Data []interface{} `json:"data"`
}

type UpdateUserResp struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type DeleteUserResp struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ListUserResp struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Meta    struct {
		HasPages    bool `json:"has_pages"`
		Total       int  `json:"total"`
		LastPage    int  `json:"last_page"`
		CurrentPage int  `json:"current_page"`
		PerPage     int  `json:"per_page"`
	} `json:"meta"`
	Data []struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Phone     string `json:"phone"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	} `json:"data"`
}

//type CreateUserResp struct {
//	Code    int      `json:"code"`
//	Status  string   `json:"status"`
//	Message UserInfo `json:"message"`
//	Data    UserInfo `json:"data"`
//}
