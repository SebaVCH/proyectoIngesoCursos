package main

import (
	"log"
	"net/http"
	"proyectoIngesoCursos/graph"
	"proyectoIngesoCursos/models"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors" // Importar el middleware CORS
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var bd *gorm.DB

// Inicialización de la base de datos dentro de la función init
func init() {
	var err error
	bd, err = gorm.Open(sqlite.Open("cursos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos", err)
	}

	// Migrar el esquema de Curso y Usuario
	err = bd.AutoMigrate(&models.Curso{})
	if err != nil {
		log.Fatal("Error al migrar la base de datos", err)
	}
}

func main() {
	// Resolver
	resolver := graph.Resolver{DB: bd}

	// Servidor GraphQL
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver}))

	// Middleware CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Cambia esto si tu frontend está en otro dominio o puerto
		AllowCredentials: true,
	}).Handler(srv)

	http.Handle("/graphql", corsHandler)
	http.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

	log.Println("Iniciando servidor en :8081...")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %s\n", err)
	}
}
