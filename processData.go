package main

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Point struct {
	x float64
	y float64
}

type Embalse struct {
	ID                     int
	Nombre                 string
	Ubicacion              Point
	Capacidad              int
	Embalse_hidroelectrico bool
	Demarcacion            string
	Cauce                  string
	Provincia              string
	CCAA                   string
	Tipo                   string
	Cota_coronacion        float64
	Altitud_cimientos      float64
	Google                 string
	Openstreetmap          string
	Wikidata               string
	Informe                string
	DistanceFromLocation   float64
}

// Reads corresponding data files to insert data into tables
func generateEmbalseData() []Embalse {
	path, err := filepath.Abs(*dataDirectory)
	if err != nil {
		log.Println(err)
	}

	var embalses []Embalse
	addEmbalsesCSV(&embalses, path)
	addEmbalsesExtraInfoTSV(&embalses, path)

	return embalses
}

// Formats embalse name so that they have the same format accross tables
func formatStr(str string) string {
	var formatedStr string
	formatedStr = strings.ToUpper(str)

	// Equivalent to converting "LAST_NAME, NAME" to "NAME LAST_NAME"
	if strings.Contains(formatedStr, ",") {
		subStrings := strings.Split(formatedStr, ",")
		for i, s := range subStrings {
			subStrings[i] = strings.TrimSpace(s)
		}

		// Remove articles "LA", "EL", "LOS", "LAS", "DE"
		for i, s := range subStrings {
			if s == "LA" || s == "EL" || s == "LOS" || s == "LAS" || s == "DE" || s == "A" {
				subStrings = append(subStrings[:i], subStrings[i+1:]...)
			}
		}

		formatedStr = ""
		for i := 0; i < len(subStrings); i++ {
			formatedStr += subStrings[i] + " "
		}

		// Remove accents, parentheses and hyphens
		accentsReplacer := strings.NewReplacer("Á", "A", "É", "E", "Í", "I", "Ó", "O", "Ú", "U", "(", "", ")", "", "-", " ")
		formatedStr = accentsReplacer.Replace(formatedStr)

		formatedStr = strings.TrimSpace(formatedStr)
	}

	return formatedStr
}

// Reads embalses in {dataDirectory}/embalses.csv
func addEmbalsesCSV(embalses *[]Embalse, path string) {
	file, err := os.Open(path + "/embalses.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println(err)
	}
	rows = rows[1:] // Remove CSV header

	var e Embalse
	for _, row := range rows {
		if e.ID, err = strconv.Atoi(row[0]); err != nil {
			continue
		}
		e.Demarcacion = row[1]
		e.Nombre = formatStr(row[2])
		if e.Capacidad, err = strconv.Atoi(row[3]); err != nil {
			continue
		}
		if e.Embalse_hidroelectrico, err = strconv.ParseBool(row[4]); err != nil {
			continue
		}

		*embalses = append(*embalses, e)
	}
}

// Reads embalses extra info in {dataDirectory}/embalses_extra_info.tsv
func addEmbalsesExtraInfoTSV(embalses *[]Embalse, path string) {
	file, err := os.Open(path + "/embalses_extra_info.tsv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	csvReader.Comma = '\t'

	rows, err := csvReader.ReadAll()
	if err != nil {
		log.Println(err)
	}
	rows = rows[1:] // Remove CSV header

	for _, row := range rows {
		for i, embalse := range *embalses {
			if embalse.Nombre == formatStr(row[1]) { // TODO: check data inserted is not being removed by second registry
				if (*embalses)[i].Ubicacion.x, err = strconv.ParseFloat(strings.Replace(row[3], ",", ".", 1), 64); err != nil {
					log.Println(err)
				}
				if (*embalses)[i].Ubicacion.y, err = strconv.ParseFloat(strings.Replace(row[4], ",", ".", 1), 64); err != nil {
					log.Println(err)
				}
				(*embalses)[i].Cauce = row[6]
				(*embalses)[i].Google = row[7]
				(*embalses)[i].Openstreetmap = row[8]
				(*embalses)[i].Wikidata = row[9]
				(*embalses)[i].Provincia = row[10]
				(*embalses)[i].CCAA = row[11]
				(*embalses)[i].Tipo = row[12]
				// ERROR: reads row[13] and row[14] as ""
				// if (*embalses)[i].cota_coronacion, err = strconv.ParseFloat(strings.Replace(row[13], ",", ".", 1), 64); err != nil {
				// 	fmt.Println(row[13], ": ")
				// 	log.Println(err)
				// }
				// if (*embalses)[i].altitud_cimientos, err = strconv.ParseFloat(strings.Replace(row[14], ",", ".", 1), 64); err != nil {
				// 	fmt.Println(row[14], ": ")
				// 	log.Println(err)
				// }
				(*embalses)[i].Informe = row[15]
			}
		}
	}
}
