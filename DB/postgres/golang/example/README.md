[comment]: <> (This is a generated file please edit source in ./templates)
[comment]: <> (All modification will be lost, you have been warned)
[comment]: <> ()
### Sample CRUD API for the postgres database postgresql://nomad_logistics_dev:mIPqf6kC0eqinBOFevwvgArLc0I487@pgm-6nn7xj3p540344s2167570.pg.rds.aliyuncs.com/nomad_logistics_dev?sslmode=disable

## Example
The project is a RESTful api for accessing the postgres database postgresql://nomad_logistics_dev:mIPqf6kC0eqinBOFevwvgArLc0I487@pgm-6nn7xj3p540344s2167570.pg.rds.aliyuncs.com/nomad_logistics_dev?sslmode=disable.

## Project Files
The generated project will contain the following code under the `./example` directory.
* Makefile
  * useful Makefile for installing tools building project etc. Issue `make` to display help
* .gitignore
  * git ignore for go project
* go.mod
  * go module setup, pass `--module` flag for setting the project module default `example.com/example`
* README.md
  * Project readme
* app/server/main.go
  * Sample Gin Server, with swagger init and comments
* api/*.go
  * REST crud controllers
* dao/*.go
  * DAO functions providing CRUD access to database
* model/*.go
  * Structs representing a row for each database table

The REST api server utilizes the Gin framework, GORM db api and Swag for providing swagger documentation
* [Gin](https://github.com/gin-gonic/gin)
* [Swaggo](https://github.com/swaggo/swag)
* [Gorm](https://github.com/jinzhu/gorm)

## Building
```.bash
make example
```
Will create a binary `./bin/example`

## Running
```.bash
./bin/example
```
This will launch the web server on localhost:8080

## Swagger
The swagger web ui contains the documentation for the http server, it also provides an interactive interface to exercise the api and view results.
http://localhost:8080/swagger/index.html

## REST urls for fetching data


* http://localhost:8080/fcconfigurationtemplate

## Project Generated Details
```.bash
/Users/song/go/bin/gen \
    --sqltype=postgres \
    --connstr \
    postgresql://nomad_logistics_dev:mIPqf6kC0eqinBOFevwvgArLc0I487@pgm-6nn7xj3p540344s2167570.pg.rds.aliyuncs.com/nomad_logistics_dev?sslmode=disable \
    --database \
    main \
    --table \
    fc_configuration_template \
    --json \
    --gorm \
    --guregu \
    --rest \
    --out \
    ./example \
    --module \
    example.com/rest/example \
    --mod \
    --server \
    --makefile \
    --json-fmt=snake \
    --generate-dao \
    --generate-proj \
    --overwrite
```











