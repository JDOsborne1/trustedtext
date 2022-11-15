package main

import (
	"file"
	"io"
	"net/http"
	"trustedtext"
)


// Borrowing the junk keys from the test scripts to hardcode for now, need to define internal keys later / make keygen public
const temp_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const temp_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"


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

func peer_consensus(_store file.Storage) (string, error) {
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
