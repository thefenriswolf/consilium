package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func LaunchServer(address string) {
	// router setup
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// routes
	router.Get("/dragon", handleImage)
	router.Get("/", handleMainPage)
	// router.Get("/test", handleTestPage)
	// router.Get("/itest", handleInteractiveTestPage)
	// router.Get("/temp", templateHandler)
	router.Get("/dragon", handleImage)

	// annouce server start
	go func() {
		log.Printf("\n\nListening on localhost%s\n\n", address)
	}()
	// actually start server
	err := http.ListenAndServe(address, router)
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
func handleMainPage(w http.ResponseWriter, r *http.Request) {
	var template *template.Template
	welcome, _ := template.ParseFS(resources, "resources/templates/welcome.html")
	welcome.Execute(w, "to insomniplan")
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	buf, err := fs.ReadFile(resources, "resources/assets/dragon.png")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "image/png")
	log.Println("loading dragon page ...")
	w.Write(buf)
}

// func tmp() {
// 	// shutdown logic
// 	connClosed := make(chan struct{})
// 	go func() {
// 		sigint := make(chan os.Signal, 1)
// 		signal.Notify(sigint, os.Interrupt)
// 		signal.Notify(sigint, syscall.SIGTERM)
// 		<-sigint
// 		log.Println("service interrupt recieved")
// 		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
// 		defer cancel()
// 		err := server.Shutdown(ctx)
// 		if err != nil {
// 			log.Printf("server shutdown error: %v", err)
// 		}
// 		log.Println("shutdown completed successfully")
// 		close(connClosed)

// 	}()

// 	log.Println("loading resources ...")
// 	tpl, _ = tpl.ParseFS(resources, "resources/templates/welcome.html")

// 	mux.HandleFunc("/", handleMainPage)
// 	mux.HandleFunc("/test", handleTestPage)
// 	mux.HandleFunc("/itest", handleInteractiveTestPage)
// 	mux.HandleFunc("/temp", templateHandler)
// 	mux.HandleFunc("/dragon", handleImage)

// 	log.Println("starting up server ...")
// 	log.Printf("\nServer is ready!\nListening on: http://localhost%s\nn", PORT)

// 	err := server.ListenAndServe()
// 	if err != http.ErrServerClosed {
// 		log.Fatalf("Failed to start server: %v", err)
// 	}

// 	<-connClosed
// 	log.Println("server stopped successfully")
// }
