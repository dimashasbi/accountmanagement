package engine

import (
	"AccountManagement/model"
	"encoding/base64"
	"fmt"
	"time"
)

type (
	// AddUserReq for Request Create users
	AddUserReq struct {
		ID           string
		UserName     string
		Password     string
		UserFullname string
		RoleID       int
		Email        string
	}
)

func (u *users) AddUsers(h *AddUserReq) *UsersDefaultResp {
	sha.Write([]byte(h.UserName + "|" + h.Password))
	encrPass := base64.StdEncoding.EncodeToString(sha.Sum([]byte(key)))

	usermod := model.NewUsers(
		h.UserName,
		encrPass,
		h.UserFullname,
		h.Email,
		h.RoleID,
		0, // login fail count
		time.Now(),
		true, // active
	)
	err := u.repository.Insert(usermod)
	if err != nil {
		fmt.Printf("%+v", err)

		return &UsersDefaultResp{
			ID:    string(h.ID),
			Error: "Error input to Users Table",
		}
	}

	return &UsersDefaultResp{
		ID:    h.ID,
		Error: "",
	}
}
