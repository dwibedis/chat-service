package controllers

import (
	"context"
	"encoding/json"
	"github.com/dwibedis/chat-service/app/entities"
	"github.com/dwibedis/chat-service/app/service"
	"github.com/dwibedis/chat-service/app/util"
	"net/http"
)

type UserRegister struct {
	userService *service.User
}

func NewUserRegister(userService *service.User) *UserRegister {
	return &UserRegister{userService:userService}
}

func (u * UserRegister) RegisterUser(w http.ResponseWriter, r *http.Request)  {
	var user *entities.User
	var resp interface{}
	reqBody, err := util.ParseRequest(r)
	if err != nil {
		resp = util.GetDefaultFailureResponse(500, "failed", "Something wrong with request body")
	}

	if resp != nil {
		util.WriteResponseIntoOutputStream(w, resp)
		return
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		resp = util.GetDefaultFailureResponse(500, "failed", "Something wrong with request body")
	}

	if resp != nil {
		util.WriteResponseIntoOutputStream(w, resp)
		return
	}
	resp = u.userService.RegisterUser(context.Background(), user)
	util.WriteResponseIntoOutputStream(w, resp)
	return
}
