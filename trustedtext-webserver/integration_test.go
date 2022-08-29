package main

import (
	"net/http"
	"testing"
	"trustedtext"
)

const first_test_env = "http://localhost:8081"

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

	response, err := trustedtext.Test_helper_post_block_to_path(test_block, first_test_env + "/block")

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

	response, err := trustedtext.Test_helper_post_block_to_path(head_move_block, first_test_env + "/block")

	if err != nil {
		t.Log("Submission request fails, with error", err)
		t.Fail()
	}

	if response.StatusCode != 201 {
		t.Log("Did not successfully create resource, instead received code: ", response.StatusCode)
		t.Fail()
	}

	new_head_block, err := helper_retrieve_and_format_external_block(first_test_env + "/block/head" )

	if err != nil {
		t.Log("Failed to retrieve new head block")
		t.Fail()
	}
	
	if new_head_block.Hash != "de2ad46ae2a00b8bf758e29285037a715f0fe033" {
		t.Log("Head change submission failed to change head hash")
		t.Fail()
	}


}