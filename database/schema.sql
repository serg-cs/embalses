CREATE EXTENSION IF NOT EXISTS postgis; -- Requires postgis installed on the datbase's server

CREATE TABLE IF NOT EXISTS embalse (
	id INTEGER PRIMARY KEY,
	nombre VARCHAR(40),
	ubicacion POINT,
	capacidad INTEGER,
	embalse_hidroelectrico BOOLEAN,
	demarcacion VARCHAR(40),
	cauce VARCHAR(60),
	provincia VARCHAR(20),
	ccaa VARCHAR(40),
	tipo VARCHAR(60),
	cota_coronacion REAL,
	altitud_cimientos REAL,
	google VARCHAR(100),
	openstreetmap VARCHAR(100),
	wikidata VARCHAR(100),
	informe VARCHAR(140)
);

CREATE TABLE IF NOT EXISTS estado_agua (
	id INTEGER,
	fecha DATE,
	cantidad INTEGER
);
