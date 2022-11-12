package main

import (
	"net/http"
	"testing"
	"trustedtext"

	"golang.org/x/exp/maps"
)

const first_test_env = "http://localhost:8081"
const second_test_env = "http://localhost:8082"

const junk_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const junk_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"


func Test_integration_environment_setup(t *testing.T) {
	resp, err := http.Get(first_test_env +  "/all_blocks")
	if err != nil {
		t.Log("Fails to get local test env")
		t.Fail()
	}

	if resp.StatusCode != 200 {
		t.Log("Testing env response with failing status")
		t.Fail()
	}

}

func Test_block_list_helper(t *testing.T) {
	_, err := helper_format_external_block_list(first_test_env)

	if err != nil {
		t.Log("Fails to collect and format external block list")
		t.Fail()
	}
}


func Test_publish_submission_works(t *testing.T) {

	
	all_blocks_check_map, _ := helper_format_external_block_list(first_test_env)

	test_block, _  := trustedtext.Test_helper_generate_standard_test_block()

	if all_blocks_check_map[test_block.Hash] {
		t.Log("attempting to use test block which is already in test chain")
		t.Fail()
	}

	response, err := test_helper_post_block_to_path(test_block, first_test_env + "/block")

	if err != nil {
		t.Log("Submission request fails, with error", err)
		t.Fail()
	}

	if response.StatusCode != 201 {
		t.Log("Did not successfully create resource, instead received code: ", response.StatusCode)
		t.Fail()
	}

	new_all_blocks_check_map, _ := helper_format_external_block_list(first_test_env)

	if !new_all_blocks_check_map[test_block.Hash] {
		t.Log("Resource not available to be retrieved, block", test_block.Hash, "not in list")
		t.Fail()
	}

	restrieved_block, err := helper_retrieve_and_format_external_block(first_test_env + "/block/" + test_block.Hash)

	if err != nil {
		t.Log("Fails on retrieving newly created block")
		t.Fail()
	}

	if restrieved_block.Hash_signature != test_block.Hash_signature {
		t.Log("submitted block, and resultant block have mismatched signatures")
		t.Fail()
	}


}


func Test_head_move_submission_works(t *testing.T) {
	head_move_block, _ := trustedtext.Generate_head_move_block(
		junk_pub_key,
		"de2ad46ae2a00b8bf758e29285037a715f0fe033",
		junk_pri_key,		
	)

	all_blocks_check_map, _ := helper_format_external_block_list(first_test_env)

	if all_blocks_check_map[head_move_block.Hash] {
		t.Log("attempting to use test block which is already in test chain")
		t.Fail()
	}

	response, err := test_helper_post_block_to_path(head_move_block, first_test_env + "/block")

	if err != nil {
		t.Log("Submission request fails, with error", err)
		t.Fail()
	}

	if response.StatusCode != 201 {
		t.Log("Did not successfully create resource, instead received code: ", response.StatusCode)
		t.Fail()
	}

	new_head_block, err := helper_retrieve_and_format_external_block(first_test_env + "/head_block/raw" )

	if err != nil {
		t.Log("Failed to retrieve new head block")
		t.Fail()
	}
	
	if new_head_block.Hash != "de2ad46ae2a00b8bf758e29285037a715f0fe033" {
		t.Log("Head change submission failed to change head hash")
		t.Fail()
	}


}

var second_test_env_details = Peer_detail{
	Claimed_name: "second_test_env",
	Path: "trustedtext-test_beta-1:8080",
}

func Test_peer_setup(t *testing.T) {
		// submit second peer to first env

		second_to_first_response, err := test_helper_post_peer_to_path(second_test_env_details, first_test_env + "/peer")

		if err != nil {
			t.Log("Error in attempting to post second env to the first as a peer", err)
			t.Fail()
		}
	
		if second_to_first_response.StatusCode != http.StatusCreated {
			t.Log("Fails to create peer resource in 1st env, status:", second_to_first_response.StatusCode)
			t.Fail()
		}

		peer_names, err := helper_format_external_peer_names(first_test_env)

		if err != nil {
			t.Log("Error in retrieviving the peerlist", err)
			t.Fail()
		}

		new_peer_in_list := peer_names[second_test_env_details.Claimed_name]

		if !new_peer_in_list {
			t.Log("Failed find newly created peer in list, only found", maps.Keys(peer_names))
			t.Fail()
		}

		
}


func Test_peer_alignment(t *testing.T) {
	// submit new block to second env
	new_block, _ := trustedtext.Test_helper_generate_standard_test_block() 

	new_block_to_second_response, err := test_helper_post_block_to_path(new_block, second_test_env + "/block")

	if err != nil {
		t.Log("Error in attempting to post new block to second env", err)
		t.Fail()
	}

	if new_block_to_second_response.StatusCode != http.StatusCreated {
		t.Log("Fails to create new block in 2nd env, status: ", new_block_to_second_response.StatusCode)
		t.Fail()
	}

	// Trigger alignment

	check_response, err := http.Get(first_test_env + "/check")

	if err != nil {
		t.Log("Error in calling check endpoint", err)
		t.Fail()
	}

	if check_response.StatusCode != http.StatusAccepted {
		t.Log("Unsuccessful call to check endpoint, status:", check_response.StatusCode)
		t.Fail()
	}

	// Check alignment successful
	composed_new_block_string := first_test_env + "/block/" + new_block.Hash
	retrieved_block, err :=  helper_retrieve_and_format_external_block(composed_new_block_string)


	if err != nil {
		t.Log("Error in retrieving block from first env", err)
		t.Fail()
	}

	if retrieved_block.Hash != new_block.Hash {
		t.Log("Shared block doesn't match sent block")
		t.Fail()
	}


}

