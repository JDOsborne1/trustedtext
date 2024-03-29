package trustedtext

import "testing"

func Test_boolean_mapping_as_string_membership(t *testing.T) {
	test_slice := []string{"testkey1", "testkey2"}

	test_map := Util_slice_to_bool_map(test_slice)

	for _, val := range test_slice {
		if !test_map[val] {
			t.Log("Not all elements of slice found in map")
			t.Fail()
		}
	}
}

func Test_subsetting_maps(t *testing.T) {
	test_map := make(map[string]Block)

	test_map["testkey1"] = Block{}
	test_map["testkey2"] = Block{}

	test_map2 := Util_subset_map(test_map, []string{"testkey2"})

	if len(test_map2) != 1 {
		t.Log("resultant map incorrect size")
		t.Fail()
	}
}
