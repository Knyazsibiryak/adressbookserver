package main

import (
	"adreesbook\controller\stdhttp"
	"adreesbook\psg"
	"net\http"
)

func main() {
	db := psg.NewPsg("psql://user:password@localhost/dbname", "login", "password")
	controller := stdhttp.NewController(":8080", db)

	http.HandleFunc("/record/add", controller.RecordAdd)
	http.HandleFunc("/records", controller.RecordsGet)
	http.HandleFunc("/record/update", controller.RecordUpdate)
	http.HandleFunc("/record/delete", controller.RecordDeleteByPhone)

	http.ListenAndServe(":8080", nil)
}
