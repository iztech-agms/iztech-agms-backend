package server

import (
	"graduation-system/endpoints/handlers/test"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func RunDBHttpServer(port string) {
	router := fasthttprouter.New()

	// Endpoints handlers
	router.POST("/test", test.ExecuteTestHandler)

	srv := &fasthttp.Server{
		Handler: router.Handler,
	}

	log.Printf("DB HTTP Server is started listening on port : %s", port)
	if err := srv.ListenAndServe("127.0.0.1:" + port); err != nil {
		log.Fatalf("(Error): error running the server : %v", err)
	}
	log.Println("DB HTTP Server is closed gracefully.")
}
