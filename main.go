package main

import (
	"awesomeProject/controllers"
	"awesomeProject/dao"
	"awesomeProject/services"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func main() {
	//services.GenerateKeys()
	//certFile := flag.String("certfile", "cert.pem", "certificate PEM file")
	//keyFile := flag.String("keyfile", "key.pem", "key PEM file")
	//flag.Parse()

	db, err := dao.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect db")
	}
	err = db.Migrate()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to run migrations")
	}

	roleRoutes := make(map[int16][]string)
	defaultRoutes := []string{"/", "/login", "/signup", "/logout", "/static", "/error"}

	roleRoutes[services.Student] = append(defaultRoutes, []string{"/test/get", "/result/get", "/teacher/add",
		"/teacher/list", "/teacher/get", "/teacher/delete", "/teacher/test"}...)

	roleRoutes[services.Teacher] = append(defaultRoutes, []string{"/test/create", "/test/update", "/test/delete",
		"/test/get", "/result/get", "/teacher/students"}...)

	roleRoutes[services.Admin] = append(defaultRoutes, []string{"/set", "/list", "/record"}...)

	handler := services.NewHandler(db, roleRoutes)
	authHandler := controllers.AuthHandler{Handler: handler}
	adminHandler := controllers.AdminHandler{Handler: handler}
	indexHandler := controllers.IndexHandler{Handler: handler}
	testHandler := controllers.TestHandler{Handler: handler}
	resultHandler := controllers.ResultHandler{Handler: handler}

	fileServer := http.FileServer(http.Dir("./view/static"))
	handler.PathPrefix("/static").Handler(http.StripPrefix("/static", fileServer))

	handler.Use(authHandler.AuthMiddleware)
	handler.Use(authHandler.RoleMiddleware)

	handler.Handle("/", &indexHandler)

	handler.Handle("/set", &adminHandler)
	handler.Handle("/list", &adminHandler)
	handler.Handle("/record", &adminHandler)

	handler.Handle("/login", &authHandler)
	handler.Handle("/signup", &authHandler)
	handler.Handle("/logout", &authHandler)

	teacherRouter := handler.PathPrefix("/teacher").Subrouter()
	teacherRouter.Handle("/add", &testHandler)
	teacherRouter.Handle("/list", &testHandler)
	teacherRouter.Handle("/get", &testHandler)
	teacherRouter.Handle("/delete", &testHandler)
	teacherRouter.Handle("/test", &testHandler)
	teacherRouter.Handle("/students", &testHandler)

	testRouter := handler.PathPrefix("/test").Subrouter()
	testRouter.Handle("/get", &testHandler)
	testRouter.Handle("/create", &testHandler)
	testRouter.Handle("/update", &testHandler)
	testRouter.Handle("/delete", &testHandler)

	resultRouter := handler.PathPrefix("/result").Subrouter()
	resultRouter.Handle("/get", &resultHandler)

	handler.HandleFunc("/error", controllers.ErrorHandler)

	srv := &http.Server{
		Handler:      handler,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		//TLSConfig: &tls.Config{
		//	MinVersion: tls.VersionTLS13,
		//},
	}

	log.Info().Msg("running server")
	//err = srv.ListenAndServeTLS(*certFile, *keyFile)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to run server")
	}
}
