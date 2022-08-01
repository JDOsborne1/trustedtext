package main

import (
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2/app"
)

// const test_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
// const test_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const default_config_path = "config.json"


type config_struct struct {
	Peerlist_path string 
	Chain_path string
}

func webservice_main() {
		
		test_handler := new(generic_handler)
	
		log.Fatal(http.ListenAndServe(":8080", test_handler))
	}



	

	
func localapp_main() {
	tt_app := app.New()
	
	main_window := block_generator_window(tt_app)
	main_window.ShowAndRun()

}	
func main() {
	if os.Getenv("TT_INTERACTIVE") == "TRUE" {
		go webservice_main()
		localapp_main()
	}

	webservice_main()
}