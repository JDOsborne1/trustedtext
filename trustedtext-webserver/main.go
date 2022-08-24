package main

import (
	"fmt"
	"log"
	"net/http"
)


const default_config_path = "config.json"

// const test_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
// const test_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"

func webservice_main(_port_to_use int) {

	test_handler := new(generic_handler)

	log.Fatal(http.ListenAndServe(":" + fmt.Sprint(_port_to_use), test_handler))
}

func main() {

	webservice_main(8080)
}
