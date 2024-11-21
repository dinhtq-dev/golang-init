package request

import "golang/app/base"

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (r *UpdateUserRequest) Validate() map[string]string {
	return base.ValidateStruct(r)
}