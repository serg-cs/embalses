package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Server struct {
	httpServer *http.Server
	templates  *template.Template
	db         *sql.DB
}

//go:embed static/*
var staticFiles embed.FS

func NewServer() *Server {
	log.Println("Creating server...")

	s := &Server{
		templates: template.Must(template.ParseFS(htmlFiles,
			"html/index.html",
			"html/navbar.html",
			"html/list.html")),
	}

	// Setup http server multiplexer
	multiplexer := http.NewServeMux()
	multiplexer.HandleFunc("/", s.index)
	multiplexer.HandleFunc("/list/", s.list)
	multiplexer.Handle("/static/", http.FileServerFS(staticFiles))

	// Setup http server
	s.httpServer = &http.Server{
		Addr:    ":8000",
		Handler: multiplexer,
	}
	log.Println("Server created")

	s.db = connectToDB()
	s.pingDB()

	s.setupDB()

	if *loadData {
		s.insertData()
	}

	return s
}

// Returns sql.DB connection to database from flag arguments
func connectToDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", *host, *port, *user, *dbname)
	if *password != "" {
		connStr += "password=" + *password
	}

	log.Println("Opening DB connection...")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}
	return db
}

// Ping DB (checks connection)
func (s *Server) pingDB() {
	if err := s.db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("DB pinged successfully")
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

//go:embed database/schema.sql
var schema embed.FS

// Setup DB using corresponding schema
func (s *Server) setupDB() {
	log.Println("Setting up DB...")

	// Create tables
	schema, err := schema.ReadFile("database/schema.sql")
	if err != nil {
		log.Println(err)
	}
	if _, err := s.db.Exec(string(schema)); err != nil {
		log.Println(err)
	}
	log.Println("DB schema set")
}

// Inserts data from files in flag dataDirectory to tables in DB
func (s *Server) insertData() {
	// WARNING: this will erase all rows in tables 'embalse' and 'estado_agua'
	if _, err := s.db.Exec("TRUNCATE TABLE embalse;"); err != nil {
		log.Println(err)
	}
	if _, err := s.db.Exec("TRUNCATE TABLE estado_agua;"); err != nil {
		log.Println(err)
	}
	log.Println("Truncated tables in DB")

	embalses := generateEmbalseData()
	for _, e := range embalses {
		if e.Informe != "" {
			if _, err := s.db.Exec(`
			INSERT INTO embalse (id, nombre, ubicacion, capacidad, embalse_hidroelectrico, demarcacion, cauce, provincia, ccaa, tipo, cota_coronacion, altitud_cimientos, google, openstreetmap, wikidata, informe)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);`, e.ID, e.Nombre, fmt.Sprint("(", e.Ubicacion.x, ", ", e.Ubicacion.y, ")"), e.Capacidad, e.Embalse_hidroelectrico, e.Demarcacion, e.Cauce, e.Provincia, e.CCAA, e.Tipo, e.Cota_coronacion, e.Altitud_cimientos, e.Google, e.Openstreetmap, e.Wikidata, e.Informe); err != nil {
				log.Println(err)
			}
		}
	}

	// TO DO
	// prepare []estado_agua
	// insert []estado_agua

	log.Println("Data loaded")
}
