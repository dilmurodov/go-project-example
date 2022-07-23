package handlers

import (
	"project/api/http"
	"project/storage"

	"github.com/gin-gonic/gin"
)

// Create user
// @ID           create user
// @Router       /user [POST]
// @Summary      Add user
// @Description  Add user data
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  http.Response{data=storage.User}  "user"
// @Response     400  {object}  http.Response{data=string}        "Invalid Argument"
// @Failure      500  {object}  http.Response{data=string}        "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {

	user := &storage.User{}

	err := c.ShouldBindJSON(user)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err)
		return
	}

	resp, err := h.repo.User().Create(c.Request.Context(), &storage.User{
		Name: user.Name,
		Age:  user.Age,
	})
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err)
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// Get cutomers list
// @ID           Get user
// @Param        id  path  string  true  "user ID"
// @Router       /user/{id} [GET]
// @Summary      get user
// @Description  get user data
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  http.Response{data=storage.User}  "user"
// @Response     400  {object}  http.Response{data=string}       "Invalid Argument"
// @Failure      500  {object}  http.Response{data=string}       "Server Error"
func (h *Handler) GetUser(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.repo.User().Get(c.Request.Context(), id)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err)
		return
	}

	h.handleResponse(c, http.OK, resp)
}
