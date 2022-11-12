package main

import (
	"file"
	"fmt"
	"log"
	"net/http"
)

const default_config_path = "config.json"

// const test_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
// const test_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"

func webservice_main(_store file.Storage) {

	config, err := _store.Config.Read_config()

	if err != nil {
		fmt.Println("Failed to load config")
		return
	}
	test_handler := new(generic_handler)
	test_handler.persistence = _store

	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(config.Port_used), test_handler))
}

func main() {
	store, err := file.Storage_from_file(default_config_path)
	if err != nil {
		fmt.Println(err)
		return
	}

	webservice_main(store)
}
