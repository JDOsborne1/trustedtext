package trustedtext

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

const first_test_env = "http://localhost:8081"


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

func Test_block_submission_works(t *testing.T) {

	
	all_blocks_check_map, _ := helper_format_external_block_list(first_test_env)

	test_block, _  := generate_standard_test_block()

	if all_blocks_check_map[test_block.Hash] {
		t.Log("attempting to use test block which is already in test chain")
		t.Fail()
	}

	marshalled_test_block, err := json.MarshalIndent(test_block, "", " ")
	if err != nil {
		t.Log("Fails to marshal block")
		t.Fail()
	}

	test_block_reader := bytes.NewReader(marshalled_test_block)

	submission_request, err := http.NewRequest("POST", first_test_env + "/block", test_block_reader)

	if err != nil {
		t.Log("Fails to compose http request")
		t.Fail()
	}

	submission_request.Header.Set("Content-Type", "application/json")

	sending_client := &http.Client{}

	response, err := sending_client.Do(submission_request)

	if err != nil {
		t.Log("Submission request fails")
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