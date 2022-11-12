package main

import (
	"file"
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
	persistence file.Storage
}

func (h generic_handler) Return_storage() file.Storage {
	return h.persistence

}

// ServeHTTP is a custom replacement for the default handler from the http package.
// It makes use of the shift path strategy to walk through the route and then delegate
// the processing to the appropriate sub handler or sub strategy. 
func (h generic_handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shift_path(r.URL.Path)

	used_storage := h.Return_storage()
	
	if head == "block" {
		block_handler(w, r, used_storage)
	} else if head == "all_blocks" {
		give_known_blocks(w, r, used_storage)
	} else if head == "head_block" {
		head_block_handler(w, r, used_storage)
	} else if head == "peer" {
		peer_handler(w, r, used_storage)
	} else if head == "all_peers" {
		share_peerlist(w, r, used_storage)
	} else if head == "check" {
		peer_check(w, r, used_storage)
	} else {
		http.Error(w, "service not implemented", http.StatusServiceUnavailable)
	}
}

// block_handler is a strategy for dispatching block handlers, this allows the block route to 
// handle retrievals, submissions, and special categories of blocks.
func block_handler(w http.ResponseWriter, r *http.Request, _store file.Storage) {
	var head string
	head, r.URL.Path = shift_path(r.URL.Path)
	if r.Method == "GET" {
		give_block(w, r, _store, head)
	}
	if r.Method == "POST" {
		submit_block(w, r, _store)
	}
}


// peer_handler is a dedicated strategy for peers, which currently just does method restricton
func peer_handler(w http.ResponseWriter, r *http.Request, _store file.Storage) {
	if r.Method != "POST" {
		http.Error(w, "Only post handling for peers", http.StatusMethodNotAllowed)
	return
	}
	
	add_peer(w, r, _store)
}

func head_block_handler(w http.ResponseWriter, r *http.Request, _store file.Storage) {
	var head string
	head, r.URL.Path = shift_path(r.URL.Path)

	if head == "hash" {
		give_head_block_hash(w, r, _store)
		return
	}
	if head == "raw" {
		give_head_block_unprocessed(w, r, _store)
		return
	}

	give_head_block_md_processed(w, r, _store)
}