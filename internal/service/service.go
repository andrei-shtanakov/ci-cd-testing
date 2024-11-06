package service

import "net/http"

type UserService struct{}

func NewUserService() *UserService {
    return &UserService{}
}

func (s *UserService) HandleUsers(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello from UserService!"))
}