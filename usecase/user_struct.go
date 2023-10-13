package usecase

type GetOneUserByIDReq struct {
	ID string `json:"id"`
}

type CreateUserReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
