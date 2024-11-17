package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed html/*.html
var htmlFiles embed.FS

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	if err := s.templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Println(err)
	}
}

func (s *Server) list(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(5 * time.Second) // Force slow down for testing
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	params := r.Form
	lat := params.Get("lat")
	lon := params.Get("lon")

	rows, err := s.db.Query(`
		SELECT
			id,
			nombre,
			capacidad,
			embalse_hidroelectrico,
			demarcacion,
			cauce,
			provincia,
			ccaa,
			tipo,
			cota_coronacion,
			altitud_cimientos,
			google,
			openstreetmap,
			wikidata,
			informe,
			(ST_Distance(
				('SRID=4326;POINT(' || $1 || ' ' || $2 || ')')::geography,
				ST_SetSRID(ST_Point(ubicacion[0], ubicacion[1]), 4326)::geography
			) / 1000) AS distance
		FROM embalse
		ORDER BY distance;`, lat, lon)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var embalses []Embalse
	var e Embalse
	for rows.Next() {
		err := rows.Scan(
			&e.ID,
			&e.Nombre,
			&e.Capacidad,
			&e.Embalse_hidroelectrico,
			&e.Demarcacion,
			&e.Cauce,
			&e.Provincia,
			&e.CCAA,
			&e.Tipo,
			&e.Cota_coronacion,
			&e.Altitud_cimientos,
			&e.Google,
			&e.Openstreetmap,
			&e.Wikidata,
			&e.Informe,
			&e.DistanceFromLocation)
		if err != nil {
			log.Println(err)
		}

		embalses = append(embalses, e)
	}

	if err := s.templates.ExecuteTemplate(w, "list.html", embalses); err != nil {
		log.Println(err)
	}
}
