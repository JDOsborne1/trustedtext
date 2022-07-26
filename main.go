package main

import (
	"log"
	"net/http"
)


var test_chain trustedtext_chain_s
// const test_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
// const test_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const default_config_path = "config.json"


type config_struct struct {
	Peerlist_path string 
	Chain_path string
}




func main() {

	used_config, _ := read_config(default_config_path)
	test_chain, _ = read_chain(used_config)

	http.HandleFunc("/share_peerlist", share_peerlist)
	http.HandleFunc("/add_peer", add_peer)
	http.HandleFunc("/check_with_peers", peer_check_handler)
	
	test_handler := new(generic_handler)

	log.Fatal(http.ListenAndServe(":8080", test_handler))
}