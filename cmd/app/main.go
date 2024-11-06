package main

import (
    "log"
    "net/http"
    
    "github.com/andrei-shtanakov/ci-cd-testing/internal/service"
    "github.com/andrei-shtanakov/ci-cd-testing/pkg/logger"
)

func main() {
    log := logger.New()
    userService := service.NewUserService()
    
    http.HandleFunc("/users", userService.HandleUsers)
    log.Info("Starting server on :8080")
    http.ListenAndServe(":8080", nil)
}