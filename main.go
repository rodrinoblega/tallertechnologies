package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrinoblega/tallertechnologies/adapters/handlers"
	"github.com/rodrinoblega/tallertechnologies/adapters/repo"
	"github.com/rodrinoblega/tallertechnologies/frameworks/db"
	usecases "github.com/rodrinoblega/tallertechnologies/use_cases"
)

func main() {
	//This password must go to Secrets
	dsn := "postgresql://neondb_owner:npg_ahKC9ALugP4n@ep-nameless-water-adxtcdet-pooler.c-2.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
	pg, err := db.NewPostgresDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	repo := repo.NewPostgresEventRepository(pg)
	uc := usecases.NewEventUseCase(repo)
	handler := handlers.NewEventHandler(uc)

	r := mux.NewRouter()
	handler.RegisterRoutes(r)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
