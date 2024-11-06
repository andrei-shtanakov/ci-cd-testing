package logger

import "log"

func New() *log.Logger {
    return log.New(log.Writer(), "app: ", log.LstdFlags)
}