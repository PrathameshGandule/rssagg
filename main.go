package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/prathameshgandule/rssagg/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	// load .env file
	godotenv.Load(".env")

	// port error handling
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found")
	}

	// database url error handling
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	// db connection
	conn, connerr := sql.Open("postgres", dbURL)
	if connerr != nil {
		log.Fatal("Can't connect to database", connerr)
	}
		 
	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	// creating chi router
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// new router
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	// mounting v1 router on main router
	router.Mount("/v1", v1Router)

	// creating server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("server started on port %v", portString)

	// server error handling
	srverr := srv.ListenAndServe()
	if srverr != nil {
		log.Fatal(srverr)
	}

}
