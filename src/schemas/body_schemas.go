package schemas

type CreateUserSchema struct {
	Name     string `json:"name" validate:"required,min=3,max=10,alphanum"`
	Password string `json:"password" validate:"required,min=3,max=10,alphanum"`
}

type SignInSchema struct {
	Name           string
	Password       string
	EnterpriseName string
	Email          string `json:"email" validate:"required,email"`
}
