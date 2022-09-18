package main

import (
	"net/http"
	"path"
	"strings"
)

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

// ServeHTTP is a custom replacement for the default handler from the http package.
// It makes use of the shift path strategy to walk through the route and then delegate
// the processing to the appropriate sub handler or sub strategy. 
func (generic_handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shift_path(r.URL.Path)
	if head == "block" {
		block_handler(w, r)
	} else if head == "all_blocks" {
		give_known_blocks(w, r)
	} else if head == "head_block" {
		give_head_block_md_processed(w, r)
	} else if head == "peer" {
		peer_handler(w, r)
	} else if head == "all_peers" {
		share_peerlist(w, r)
	} else if head == "check" {
		peer_check(w, r)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
}

// block_handler is a strategy for dispatching block handlers, this allows the block route to 
// handle retrievals, submissions, and special categories of blocks.
func block_handler(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shift_path(r.URL.Path)
	if r.Method == "GET" && head == "head" {
		give_head_block_unprocessed(w, r)
	}
	if r.Method == "GET" && head != "head" {
		give_block(w, r, head)
	}
	if r.Method == "POST" {
		submit_block(w, r)
	}
}


// peer_handler is a dedicated strategy for peers, which currently just does method restricton
func peer_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only post handling for peers", http.StatusMethodNotAllowed)
	return
	}
	
	add_peer(w, r)
}
