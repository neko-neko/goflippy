package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neko-neko/goflippy/cmd/goflippy-api/ctx"
	"github.com/neko-neko/goflippy/cmd/goflippy-api/service"
	"github.com/neko-neko/goflippy/collection"
	"github.com/neko-neko/goflippy/util"
)

// UserHandler is users resource handler
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler returns new instance
func NewUserHandler(u *service.UserService) *UserHandler {
	return &UserHandler{
		service: u,
	}
}

// postUsersRequest is PostUsers request parameter
// swagger:parameters
type postUsersRequest struct {
	FirstName string `json:"first_name" validate:"-"`
	LastName  string `json:"last_name" validate:"-"`
	UUID      string `json:"uuid" validate:"required"`
	Email     string `json:"email" validate:"omitempty,email"`
	Groups    []struct {
		Name string `json:"name"`
	} `json:"groups" validate:"omitempty,unique"`
	Attributes []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"attributes" validate:"omitempty,unique"`
}

// postUsersResponse is PostUsers response
// swagger:parameters postUsersResponse
type postUsersResponse struct {
	User collection.User `json:"user"`
}

// PostUsers register a user
//
// swagger:route POST /users users registerUser
//
// Register a user
//
// Responses:
//  200: postUsersResponse
//  400: errorResponse
func (u *UserHandler) PostUsers(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	param := postUsersRequest{}
	if err := util.BindJSONParam(r, &param); err != nil {
		return http.StatusBadRequest, nil, err
	}
	if err := util.ValidateStruct(&param); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user := collection.NewUser()
	user.SetProjectID(ctx.GetProjectID(r.Context()))
	user.FirstName = param.FirstName
	user.LastName = param.LastName
	user.UUID = param.UUID
	user.Email = param.Email
	for _, v := range param.Groups {
		user.AppendGroup(v.Name)
	}
	for _, v := range param.Attributes {
		user.AppendAttribute(v.Key, v.Value)
	}

	err := u.service.RegisterUser(user)
	if aerr, ok := err.(*service.ResourceAlreadyExistsError); ok {
		return http.StatusConflict, nil, aerr
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, postUsersResponse{
		User: *user,
	}, nil
}

// patchUsersRequest is patchUsers request parameter
// swagger:parameters
type patchUsersRequest struct {
	UUID      string `validate:"required"`
	FirstName string `json:"first_name" validate:"-"`
	LastName  string `json:"last_name" validate:"-"`
	Email     string `json:"email" validate:"omitempty,email"`
	Groups    []struct {
		Name string `json:"name"`
	} `json:"groups" validate:"omitempty,unique"`
	Attributes []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"attributes" validate:"omitempty,unique"`
}

// patchUsersResponse is PatchUsers response
// swagger:parameters patchUsersResponse
type patchUsersResponse struct {
	User collection.User `json:"user"`
}

// PatchUsers update a user
//
// swagger:route PATCH /users users update
//
// Update a user
//
// Responses:
//  200: patchUserResponse
//  400: errorResponse
func (u *UserHandler) PatchUsers(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	param := patchUsersRequest{
		UUID: vars["uuid"],
	}
	if err := util.BindJSONParam(r, &param); err != nil {
		return http.StatusBadRequest, nil, err
	}
	if err := util.ValidateStruct(&param); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user := collection.NewUser()
	user.SetProjectID(ctx.GetProjectID(r.Context()))
	user.FirstName = param.FirstName
	user.LastName = param.LastName
	user.UUID = param.UUID
	user.Email = param.Email
	user.ResetGroups()
	for _, v := range param.Groups {
		user.AppendGroup(v.Name)
	}
	user.ResetAttributes()
	for _, v := range param.Attributes {
		user.AppendAttribute(v.Key, v.Value)
	}

	err := u.service.UpdateUserInfo(user)
	if rerr, ok := err.(*service.ResourceNotFoundError); ok {
		return http.StatusNotFound, nil, rerr
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, patchUsersResponse{
		User: *user,
	}, nil
}

// deleteUsersRequestParameter is deleteUsers request parameter
// swagger:parameters
type deleteUsersRequest struct {
	UUID string `validate:"required"`
}

// DeleteUsers delete a user
//
// swagger:route DELETE /users users delete
//
// Delete a user
//
// Responses:
//  204: noContent
//  400: errorResponse
func (u *UserHandler) DeleteUsers(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	param := deleteUsersRequest{
		UUID: vars["uuid"],
	}
	if err := util.ValidateStruct(&param); err != nil {
		return http.StatusBadRequest, nil, err
	}

	err := u.service.DeleteUser(param.UUID, ctx.GetProjectID(r.Context()))
	if rerr, ok := err.(*service.ResourceNotFoundError); ok {
		return http.StatusNotFound, nil, rerr
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}
