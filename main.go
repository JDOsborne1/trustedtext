package main

import (
	"trustedtext/trustedtextlib"
	"log"
	"net/http"
)

// const test_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
// const test_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"


func webservice_main() {

	test_handler := new(trustedtext.Generic_handler)

	log.Fatal(http.ListenAndServe(":8080", test_handler))
}

func main() {
	webservice_main()
}
