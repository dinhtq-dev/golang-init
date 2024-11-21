// UserController handles user-related HTTP requests.
package controller

import (
	"golang/app/base"
	"golang/app/modules/auth/request"
	"golang/app/modules/auth/resource"
	"golang/app/modules/auth/service"
	"golang/langs"
	"golang/pkg/utils/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.BaseApiController
	userService service.UserService
}

// NewUserController creates a new user controller.
func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

// CreateUser creates a new user. Handles request binding, validation, and user creation.
func (uc *UserController) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		uc.SendError(c, langs.Messages["AppName"], http.StatusBadRequest)
		return
	}
	
	if err := req.Validate(); err != nil {
		uc.SendValidator(c, err)
		return
	}

	// Create user
	user, err := uc.userService.Create(req)
	if err != nil {
		errors, _ := validator.HandleMySQLError(err)
		uc.SendValidator(c, errors)

		return
	}

	userResource := resource.NewUserResource(user)
	uc.SendResponse(c, userResource, "User created successfully")
}

// GetUserDetail retrieves user details by ID and sends the response.
func (uc *UserController) GetUserDetail(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		uc.SendError(c, "User ID is required", http.StatusBadRequest)
		return
	}
	
	user, err := uc.userService.GetByID(userID)
	if err != nil {
		uc.SendError(c, "User not found", http.StatusNotFound)
		return
	}

	userResource := resource.NewUserResource(user)
	uc.SendResponse(c, userResource)
}

// GetUserList retrieves a list of users and sends the response.
func (uc *UserController) GetUserList(c *gin.Context) {
    pageNum, perPageNum := uc.ParsePaginationParams(c)

	result, err := uc.userService.GetAll(pageNum, perPageNum)
    if err != nil {
        uc.SendError(c, "Failed to retrieve users", http.StatusInternalServerError)
        return
    }

    var userResources []interface{}
    for _, user := range result.Users {
        userResources = append(userResources, resource.NewUserResource(&user))
    }

    uc.SendPaginationArrayResponse(c, userResources, result.Total, pageNum, perPageNum, "Users fetched successfully")
}

// UpdateUser updates a user by ID.
func (uc *UserController) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		uc.SendError(c, "User ID is required", http.StatusBadRequest)
		return
	}

	var req request.UpdateUserRequest

	// Bind and validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		uc.SendError(c, langs.Messages["AppName"], http.StatusBadRequest)
		return
	}
	
	if err := req.Validate(); err != nil {
		uc.SendError(c, "Validation failed", http.StatusBadRequest)
		return
	}

	// Update user
	updatedUser, err := uc.userService.Update(userID, req)
	if err != nil {
		uc.SendError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	userResource := resource.NewUserResource(updatedUser)
	uc.SendResponse(c, userResource, "User updated successfully")
}
