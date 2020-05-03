package controllers

import (
	"context"
	"encoding/json"
	"github.com/dwibedis/chat-service/app/entities"
	"github.com/dwibedis/chat-service/app/service"
	"github.com/dwibedis/chat-service/app/util"
	"net/http"
)

type UserLogin struct {
	userService *service.User
}

func (u *UserLogin) Login(w http.ResponseWriter, r *http.Request)  {
	var userOtp *entities.UserOtp
	var resp interface{}
	reqBody, err := util.ParseRequest(r)
	if err != nil {
		resp = util.GetDefaultFailureResponse(500, "failed", "Something wrong with request body")
	}

	if resp != nil {
		util.WriteResponseIntoOutputStream(w, resp)
		return
	}

	err = json.Unmarshal(reqBody, &userOtp)
	if err != nil {
		resp = util.GetDefaultFailureResponse(500, "failed", "Something wrong with request body")
	}

	if resp != nil {
		util.WriteResponseIntoOutputStream(w, resp)
		return
	}
	resp = u.userService.VerifyOtp(context.Background(), userOtp)
	util.WriteResponseIntoOutputStream(w, resp)
	return
}

func (u *UserLogin) SendOtp(w http.ResponseWriter, r *http.Request) {
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

	err = u.userService.GenerateAndSendUserOtp(context.Background(), user)
	if err != nil {
		resp = util.GetDefaultFailureResponse(500, "failed", err.Error())
	}
	resp = util.GetDefaultFailureResponse(200, "success", "SMS sent successfully")
	util.WriteResponseIntoOutputStream(w, resp)
	return
}


