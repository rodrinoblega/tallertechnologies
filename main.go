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

	dsn := "postgresql://postgres:P4zcYFKexte6hkVn@db.awevsesovideiduznuzi.supabase.co:5432/postgres"
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
