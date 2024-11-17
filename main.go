package main

import (
	"flag"
)

// Command line arguments passed upon execution
var (
	host          = flag.String("host", "localhost", "database host")
	port          = flag.Int("port", 5432, "port where host is running")
	dbname        = flag.String("dbname", "", "database name")
	user          = flag.String("user", "", "database user")
	password      = flag.String("password", "", "database user password (if necessary)")
	dataDirectory = flag.String("data-dir", "", "Directory where the three corresponding data files are stored, named 'estado.csv' 'embalses.csv' 'embalses_extra_info.tsv'")
	loadData      = flag.Bool("load-data", false, "WARNING: this will delete any data in tables 'embalse' or 'estado_agua' and generate new tables with data in corresponding dataDirectory provided")
)

func main() {
	flag.Parse()

	server := NewServer()
	server.Run()
}
