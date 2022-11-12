package main

import (
	"file"
	"fmt"
)

const default_config_path = "config.json"

func main() {
	store, err := file.Storage_from_file(default_config_path)
	if err != nil {
		fmt.Println(err)
		return
	}

	localapp(store)
}
