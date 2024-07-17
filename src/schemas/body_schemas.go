package schemas

type CreateUserSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=3,max=10,alphanum"`
	Password string `json:"password" validate:"required,min=3,max=10,alphanum"`
}

type GetTokenSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=10,alphanum"`
}
