package middlewares

import (
	"github.com/joho/godotenv"
)

func init() {
    Setup()
}
func Setup() {
    err := godotenv.Load("../.env.testing")
    if err != nil {
        panic(err)
    }
}

