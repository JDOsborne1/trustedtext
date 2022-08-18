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

func (generic_handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shift_path(r.URL.Path)
	if head == "block" {
		block_handler(w, r)
	} else if head == "all_blocks" {
		give_known_blocks(w, r)
	} else if head == "head_block" {
		give_head_block(w, r)
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

func peer_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		add_peer(w, r)
	}
	http.Error(w, "Only post handling for peers", http.StatusMethodNotAllowed)
}
