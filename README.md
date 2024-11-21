# embalses

This project originates from the **2024 Malackathon at UMA Higher Technical School of Computer Engineering**. 
The project consists of designing a web page that displays information about dams in Spain. 
The data is provided by a series of .csv and .tsv files, and their data is stored, accessed, and manipulated from a database. 

My objective with this project is to see all aspects of a fully functional web page. 
Once finished, I hope to have a global image, although still humble and incomplete, of everything that takes place to have a working project running.
Along the way, I plan to learn new technologies and explore areas outside of what I have yet seen.

### Technologies and languages used:
* Go
* HTML
* HTMX
* JavaScript
* CSS
* SQL
* PostgreSQL database

## Use

- A server with a running PostgreSQL database server is required (I used a local Postgres running on my machine).
- The original data files are stored in the `originalDataFiles` directory of the project. NOTE: the file names are expected as originally provided.
- When running, you must pass some parameters in the command line as explained below.
- All necessary files are prepared with `go:embed` to form part of the Golang executable if compiled, except original data files which may not be used as given or be stored elsewhere.

## Command line

* **[host]** by default is "localhost"
* **[port]** where the database host is running (5432 is the default for PostgreSQL)
* **dbname** the name of a running database on the host server is required
* **user** the username to access the database
* **[password]** password for the given user if required
* **[load-data]** by default false, but the ***first time it is required***, this prepares the database from the data files in the provided data-dir path
* **[data-dir]** directory where data files are stored. Only required when load-data is active. NOTE: can be absolute or relative path

## Installation
You can install the program using [Go](https://go.dev) with the following command:

```go install github.com/serg-cs/embalses@latest```

And then using the command embalses with the corresponding flags. An example could be:

```embalses -dbname embalses -user serg-cs -data-dir pathToDataFilesDirectory -load-data```

NOTE: If you wish to use the original data files you must download them from the repository separately and change `pathToDataFilesDirectory` to the path where your data files are stored. 
