package server

import (
	"graduation-system/endpoints/handlers/auth"
	handlers "graduation-system/endpoints/handlers/notification"
	"graduation-system/endpoints/handlers/test"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func RunDBHttpServer(port string) {
	router := fasthttprouter.New()

	// Endpoints handlers
	router.POST("/test", test.ExecuteTestHandler)
	router.POST("/auth/login", auth.AuthLoginHandler)

	router.POST("/notifications/get/user-id/:user-id", handlers.GetNotificationsByUserIDHandler)
	router.POST("/notifications/update", handlers.UpdateNotificationHandler)
	router.POST("/notifications/delete/id/:id", handlers.DeleteNotificationHandler)

	srv := &fasthttp.Server{
		Handler: router.Handler,
	}

	log.Printf("DB HTTP Server is started listening on port : %s", port)
	if err := srv.ListenAndServe("127.0.0.1:" + port); err != nil {
		log.Fatalf("(Error): error running the server : %v", err)
	}
	log.Println("DB HTTP Server is closed gracefully.")
}
