package main

import (
	"file"
	"io"
	"net/http"
	"trustedtext"
)

func retrieve_head_hash_from_peer(peer trustedtext.Peer_detail) (string, error) {
	composed_path := "http://" + peer.Path + "head_block/hash"

	resp, err := http.Get(composed_path)
	if err != nil {
		return "", err
	}

	read_resp, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(read_resp), nil
}

func first_come_first_believed(input_hashes []string) string {
	return input_hashes[0]
}

func Peer_consensus(_store file.Storage) (string, error) {
	peerlist, err := _store.Peerlist.Read_peerlist()

	if err != nil {
		return "", err
	}

	head_hashes := []string{}

	for _, peer := range peerlist {
		peer_head_hash, err := retrieve_head_hash_from_peer(peer)
		if err != nil {
			return "", err
		}
		head_hashes = append(head_hashes, peer_head_hash)
	}

	return first_come_first_believed(head_hashes), nil
}
