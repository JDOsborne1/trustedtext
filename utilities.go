package trustedtext

// Util_slice_to_bool_map is a function which takes a slice of strings, and turns it into
// a convenient shape for easy checking
func Util_slice_to_bool_map(input_slice_of_keys []string) map[string]bool {
	mapped_keys_to_keep := make(map[string]bool)

	for _, key := range input_slice_of_keys {
		mapped_keys_to_keep[key] = true
	}

	return mapped_keys_to_keep
}

// Util_subset_map is a function to return a subset of the input map, based on a slice of keys to keep.
func Util_subset_map(_original_map map[string]Trustedtext_s, _keys_to_keep []string) map[string]Trustedtext_s {
	mapped_keys_to_keep := Util_slice_to_bool_map(_keys_to_keep)

	new_map := make(map[string]Trustedtext_s)
	for key, value := range _original_map {
		if mapped_keys_to_keep[key] {
			new_map[key] = value
		}
	}

	return new_map
}

func Util_anti_set_map[Res Trustedtext_s | bool](_original_map map[string]Res, _keys_to_compare []string) map[string]Res {
	mapped_keys_to_keep := Util_slice_to_bool_map(_keys_to_compare)

	new_map := make(map[string]Res)
	for key, value := range _original_map {
		if !mapped_keys_to_keep[key] {
			new_map[key] = value
		}
	}

	return new_map
}
