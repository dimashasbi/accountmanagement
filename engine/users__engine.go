package engine

import (
	"crypto/sha1"
)

type (
	// Users is the interface for interactor
	Users interface {
		AddUsers(h *AddUserReq) *UsersDefaultResp
		// UpdateUsers()
		// DeleteUsers()
		SelectUsers(h *SelectUserReq) *SelectUserResp
	}

	users struct {
		repository UsersRepository
	}

	// UsersDefaultResp for default Respon after success Respon
	UsersDefaultResp struct {
		ID    string
		Error string
	}
)

var (
	sha = sha1.New()
	key = "DOTA"
)

func (f *engineFactory) NewUsersEngine() Users {
	repostruc := &users{
		repository: f.NewUsersRespository(),
	}

	repostruc.InitializationUsers()
	return repostruc
}

func (s users) InitializationUsers() {
	// create new default users administrator if not exist
	firstUser := SelectUserReq{
		ID:       "0",
		UserName: "adminSkywalker",
	}
	val := s.SelectUsers(&firstUser)
	if val.Error != "" || !val.Active && val.UserName == "adminSkywalker" {
		// no users = no data record = error
		// no rows = val.active is false and role ID is 0
		// lets Add Users
		inp := &AddUserReq{
			ID:           "0",
			UserName:     firstUser.UserName,
			Password:     "jaquest",
			UserFullname: "Dota MMR nya cuma 4k, Ancient",
			RoleID:       1,
			Email:        "dimashasbi.habibi@gmail.com",
		}
		sil := s.AddUsers(inp)
		if sil.Error != "" {
			panic("Error input First users")
		}
	}
}
