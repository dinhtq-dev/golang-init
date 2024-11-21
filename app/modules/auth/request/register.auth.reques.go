package request

import "golang/app/base"

// CreateUserRequest struct với các tag validation
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"required,email,email_custom"`
}

func (r *CreateUserRequest) Validate() map[string]string {
	return base.ValidateStruct(r)
}