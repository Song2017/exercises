package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://nomad_logistics_dev?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// age := 21
	rows, err := db.Query("SELECT gateway_app_id, name, description, connector_ids FROM fc_gateway_app where gateway_app_id = '110704655'")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(rows)

	for rows.Next() {
		var gateway_app_id string
		var name string
		var description string
		var connector_ids string

		rows.Scan(&gateway_app_id, &name, &description, &connector_ids)
		fmt.Println(name, gateway_app_id, description, connector_ids)

		var listoflists []string
		dec := json.NewDecoder(strings.NewReader(connector_ids))
		err := dec.Decode(&listoflists)
		fmt.Println(err, listoflists[0])
	}
}
