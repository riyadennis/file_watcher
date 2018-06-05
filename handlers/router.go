package handlers

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/sirupsen/logrus"
)


type ApiResponse struct {
	Status int
	Detail string
	Title  string
}
func Run() {
	route := httprouter.New()
	port := ":8080"
	route.GET("/watch", Server)

	fmt.Printf("Listenning to port %s \n", port)
	logrus.Fatal(http.ListenAndServe(port, nil))
}

func jsonResponseDecorator(response *ApiResponse, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), response.Status)
		return
	}
}
func createResponse(detail, title string, status int) *ApiResponse {
	return &ApiResponse{
		Status: status,
		Detail: detail,
		Title:  title,
	}
}