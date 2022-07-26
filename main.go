package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
)


var test_chain trustedtext_chain_s
// const test_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
// const test_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const default_config_path = "config.json"


type config_struct struct {
	Peerlist_path string 
	Chain_path string
}

// shift_path splits the given path into the first segment (head) and
// the rest (tail). For example, "/foo/bar/baz" gives "foo", "/bar/baz".
func shift_path(p string) (head, tail string) {
	p = path.Clean("/" + p)
    i := strings.Index(p[1:], "/") + 1
    if i <= 0 {
		return p[1:], "/"
    }
    return p[1:i], p[i:]
}

type generic_handler struct {
	
}

func (generic_handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shift_path(r.URL.Path)
	if head == "block" {
		block_handler(w, r)
	}
}

func block_handler(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shift_path(r.URL.Path)
	if r.Method == "GET" {
		give_block(w, r, head)
	}
	if r.Method == "POST" {
		submit_block(w, r)
	}
}

func give_block(w http.ResponseWriter, r *http.Request, _block_hash string) {
	requested_block, _ := Return_specified_hash(test_chain, _block_hash)
	text_block, _ := json.Marshal(requested_block)
	fmt.Fprint(w, string(text_block))
}

func submit_block(w http.ResponseWriter, r *http.Request) {
	var post_deposit []byte
	var err error
	post_deposit, err = ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	resultant_block := &trustedtext_s{}
	err = json.Unmarshal(post_deposit, resultant_block)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	new_chain, err := Process_incoming_block(test_chain, *resultant_block)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	if err != nil {
		test_chain = new_chain
	}

	w.WriteHeader(http.StatusCreated)
}


func main() {

	used_config, _ := read_config(default_config_path)
	test_chain, _ = read_chain(used_config)

	http.HandleFunc("/known_blocks", give_known_blocks)
	http.HandleFunc("/share_peerlist", share_peerlist)
	http.HandleFunc("/add_peer", add_peer)
	http.HandleFunc("/check_with_peers", peer_check_handler)
	
	test_handler := new(generic_handler)

	log.Fatal(http.ListenAndServe(":8080", test_handler))
}