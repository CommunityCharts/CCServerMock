package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/CommunityCharts/CCServerMock/pkg/db"
	"github.com/CommunityCharts/CCServerMock/pkg/router"
	muxHandlers "github.com/gorilla/handlers"
)

const PORT = 4000

func main() {
	r := router.CreateMuxRouter()

	headersOk := muxHandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := muxHandlers.AllowedOrigins([]string{"http://localhost:3000"})
	methodsOk := muxHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	handler := muxHandlers.CORS(headersOk, originsOk, methodsOk)(r)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + strconv.Itoa(PORT),
	}

	db.InitDB()

	fmt.Printf("Binding to :%v\n", PORT)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
