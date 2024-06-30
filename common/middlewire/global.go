package middlewire

import (
	"fmt"
	"net/http"
)

type GlobelMiddleware struct {
}

func NewGlobelMiddleware() *GlobelMiddleware {
	return &GlobelMiddleware{}
}

func (m *GlobelMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handle handel")
		next(w, r)
		fmt.Println("after handle handel")
	}
}