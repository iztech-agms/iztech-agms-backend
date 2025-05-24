package server

import (
	"graduation-system/endpoints/handlers/auth"
	gradStatus "graduation-system/endpoints/handlers/graduationStatus"
	handlers "graduation-system/endpoints/handlers/notification"
	"graduation-system/endpoints/handlers/student"
	"graduation-system/endpoints/handlers/test"
	"graduation-system/endpoints/middleware"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func RunDBHttpServer(port string) {
	router := fasthttprouter.New()

	// Endpoints handlers
	router.POST("/test", test.ExecuteTestHandler)
	router.POST("/test_getusers", student.GetAllUsersTest) // Delete after use
	router.POST("/auth/login", auth.AuthLoginHandler)

	router.POST("/notifications/get/user-id/:user-id", wrapWithJWTMiddleware(handlers.GetNotificationsByUserIDHandler))
	router.POST("/notifications/update", wrapWithJWTMiddleware(handlers.UpdateNotificationHandler))
	router.POST("/notifications/delete/id/:id", wrapWithJWTMiddleware(handlers.DeleteNotificationHandler))

	router.POST("/graduation-status/std/:id", wrapWithJWTMiddleware(gradStatus.GetGradStatusByUserIDHandler))
	router.POST("/graduation-status/list/:year", wrapWithJWTMiddleware(gradStatus.GetGradStatusByGradYear))
	router.POST("/graduation-status/update", wrapWithJWTMiddleware(gradStatus.UpdateGradStatus))

	router.POST("/student/get/list/userid/:userid", wrapWithJWTMiddleware(student.GetStudentListDetailedByUserIDHandler))

	srv := &fasthttp.Server{
		Handler: router.Handler,
	}

	log.Printf("DB HTTP Server is started listening on port : %s", port)
	if err := srv.ListenAndServe("127.0.0.1:" + port); err != nil {
		log.Fatalf("(Error): error running the server : %v", err)
	}
	log.Println("DB HTTP Server is closed gracefully.")
}

func wrapWithJWTMiddleware(requestHandler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return middleware.JWTMiddleware(requestHandler)
}
