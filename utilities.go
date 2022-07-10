package main 

// util_make_boolean_map_from_slice is a function which takes a slice of strings, and turns it into 
// a convenient shape for easy checking
func util_make_boolean_map_from_slice(input_slice_of_keys []string) map[string]bool {
	mapped_keys_to_keep := make(map[string]bool)

	for _, key := range input_slice_of_keys {
		mapped_keys_to_keep[key] = true
	}	

	return mapped_keys_to_keep
}

// util_subset_map is a function to return a subset of the input map, based on a slice of keys to keep. 
func util_subset_map(_original_map map[string]trustedtext_s, _keys_to_keep []string) map[string]trustedtext_s {
	mapped_keys_to_keep := util_make_boolean_map_from_slice(_keys_to_keep)

	new_map := make(map[string]trustedtext_s)
	for key , value := range _original_map {
		if mapped_keys_to_keep[key] {
			new_map[key] = value
		}
	} 

	return new_map
}