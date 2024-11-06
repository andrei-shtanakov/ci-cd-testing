package service

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestUserService_HandleUsers(t *testing.T) {
    service := NewUserService()
    
    // Создаём тестовый HTTP запрос
    req := httptest.NewRequest(http.MethodGet, "/users", nil)
    w := httptest.NewRecorder()
    
    // Вызываем тестируемый обработчик
    service.HandleUsers(w, req)
    
    // Проверяем результат
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }
}