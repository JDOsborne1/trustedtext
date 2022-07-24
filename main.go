package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/exp/maps"
)

var peerlist []peer_detail
var test_chain trustedtext_chain_s
const test_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const test_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"

func check_with_peers() error{
	if len(peerlist) == 0 {
		return errors.New("cant validate against empty peerlist")
	}
	resp, err := http.Get("http://" + peerlist[0].Path + "/known_blocks")
	if err != nil {
		return err
	}
	response_decoder := json.NewDecoder(resp.Body)
	known_blocks_of_peer := &[]string{}
	response_decoder.Decode(known_blocks_of_peer)
	peer_blocks_map := util_make_boolean_map_from_slice(*known_blocks_of_peer)
	my_blocks := maps.Keys(test_chain.tt_chain)

	new_keys_of_peer := util_anti_set_map(peer_blocks_map, my_blocks)

	if len(new_keys_of_peer) != 0 {
		fmt.Println("Missing items on the web")
	}
	return nil
}

func test_handler(w http.ResponseWriter, r *http.Request) {
	err := check_with_peers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
}


func main() {

	new_peer := peer_detail{
		Claimed_name: "self",
		Path: "localhost:8080",
	}
	peerlist = append(peerlist, new_peer)

	test_chain, _ = Genesis(
		test_pub_key,
		[]string{"test", "not-a-blog"},
		test_pri_key,
	)
	test_block_1, _ := Instantiate(
		test_pub_key,
		tt_body{
			Instruction_type: "publish",
			Instruction: "My First ever message",
		},
		test_pri_key,
	)
	test_block_2, _ := Instantiate(
		test_pub_key,
		tt_body{
			Instruction_type: "publish",
			Instruction: "My Second ever message",
		},
		test_pri_key,
	)
	test_chain, _ = Process_incoming_block(test_chain, test_block_1)
	test_chain, _ = Process_incoming_block(test_chain, test_block_2)

	http.HandleFunc("/block", give_block)
	http.HandleFunc("/known_blocks", give_known_blocks)
	http.HandleFunc("/submit_block", submit_block)
	http.HandleFunc("/share_peerlist", share_peerlist)
	http.HandleFunc("/add_peer", add_peer)
	
	http.HandleFunc("/test", test_handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}