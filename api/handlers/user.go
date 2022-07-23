package handlers

import (
	"project/api/http"
	"project/storage"

	"github.com/gin-gonic/gin"
)


func (h *Handler) CreateUser (c *gin.Context){

	user := &storage.User{}
	
	err := c.ShouldBindJSON(user)
	if err != nil{
		h.handleResponse(c, http.InvalidArgument, err)
		return
	}

	resp, err := h.repo.User().Create(c.Request.Context(), &storage.User{
		Name: user.Name,
		Age: user.Age,
	})
	if err != nil{
		h.handleResponse(c, http.InvalidArgument, err)
		return
	}
	
	h.handleResponse(c, http.Created, resp)
}

func (h *Handler) GetUser(c *gin.Context){

	id := c.Param("id")

	resp, err := h.repo.User().Get(c.Request.Context(), id)
	if err != nil{
		h.handleResponse(c, http.BadRequest, err)
		return
	}

	h.handleResponse(c, http.OK, resp)
}