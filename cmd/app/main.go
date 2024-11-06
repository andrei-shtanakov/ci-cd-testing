package main

import (
    "net/http"
    
    "github.com/andrei-shtanakov/ci-cd-testing/internal/service"
    "github.com/andrei-shtanakov/ci-cd-testing/pkg/logger"
)

func main() {
    log := logger.New()
    userService := service.NewUserService()
    
    http.HandleFunc("/users", userService.HandleUsers)
    log.Println("Starting server on :8080")  // используем Println вместо Info
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Printf("Server error: %v", err)
    }
}