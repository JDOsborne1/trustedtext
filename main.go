package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func read_peerlist(_config config_struct) ([]peer_detail, error) {
	bytefile, err := ioutil.ReadFile(_config.Peerlist_path)
	if err != nil {
		return []peer_detail{}, err
	}
	peerlist := &[]peer_detail{}
	err = json.Unmarshal(bytefile, peerlist)
	if err != nil {
		return []peer_detail{}, err
	}
	return *peerlist, nil
}

func write_peerlist(_peerlist []peer_detail, _config config_struct) error {
	marshalled_peerlist, err := json.MarshalIndent(_peerlist, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(_config.Peerlist_path, marshalled_peerlist, 0644) 
	
	if err != nil {
		return err
	}

	return nil 
}

func write_config(_config config_struct) error {
	marshalled_config, err := json.MarshalIndent(
		_config,
		"",
		"  ",
	)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(
		"config.json",
		marshalled_config,
		0644,
	)
	if err != nil {
		return err
	}
	return nil
} 

func read_config(_config_path string) (config_struct, error) {
	bytefile, err := ioutil.ReadFile(_config_path)
	if err != nil {
		return config_struct{}, err
	}
	config := &config_struct{}
	err = json.Unmarshal(bytefile, config)
	if err != nil {
		return config_struct{}, err
	}

	return *config, nil
}

func write_chain(_chain trustedtext_chain_s, _config config_struct) error {
	marshalled_chain, err := json.MarshalIndent(_chain, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_config.Chain_path, marshalled_chain, 0644)
	if err != nil {
		return err
	}
	return err
}

func read_chain(_config config_struct) (trustedtext_chain_s, error) {
	bytefile, err := ioutil.ReadFile(_config.Chain_path)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	chain := &trustedtext_chain_s{}
	err = json.Unmarshal(bytefile, chain)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	return *chain, nil
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