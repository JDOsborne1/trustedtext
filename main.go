package main

import (
	"fmt"
	"log"
	"net/http"
)


var test_chain trustedtext_chain_s
// const test_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
// const test_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const default_config_path = "config.json"

func test_handler(w http.ResponseWriter, r *http.Request) {

	used_config, _ := read_config(default_config_path)
	peerlist, _ := read_peerlist(used_config)
	err := check_with_peers(peerlist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
}

type config_struct struct {
	Peerlist_path string 
	Chain_path string
}



func main() {

	used_config, _ := read_config(default_config_path)
	test_chain, _ = read_chain(used_config)

	http.HandleFunc("/block", give_block)
	http.HandleFunc("/known_blocks", give_known_blocks)
	http.HandleFunc("/submit_block", submit_block)
	http.HandleFunc("/share_peerlist", share_peerlist)
	http.HandleFunc("/add_peer", add_peer)
	
	http.HandleFunc("/test", test_handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}