package resource

import "golang/app/modules/auth/model"

type UserResource struct {
    ID    uint   `json:"id"`
    Name  *string `json:"name"`
    Email string `json:"email"`
    CreatedAt string `json:"created_at"` 
    UpdatedAt string `json:"updated_at"` 
}

func NewUserResource(user *model.User) UserResource {
    return UserResource{
        ID:    user.ID,
        Name:  user.Name,
        Email: user.Email,
        CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
        UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
    }
}
