package main

import (
	"log"
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)


var test_chain trustedtext_chain_s
// const test_pub_key = "faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
// const test_pri_key = "366c15a87d86f7a6fe6f7509ecaab3d453f0488b414aef12175a870cc5d1b124faa372113c86e434298d3c2c76c230c41f8ec890d165ef0d124c62758d89a66a"
const default_config_path = "config.json"


type config_struct struct {
	Peerlist_path string 
	Chain_path string
}

func webservice_main() {
	
		used_config, _ := read_config(default_config_path)
		test_chain, _ = read_chain(used_config)
		
		test_handler := new(generic_handler)
	
		log.Fatal(http.ListenAndServe(":8080", test_handler))
	}

func generate_block(_instruction_type string, _instruction_body string, _public_key string, _private_key string) (trustedtext_s, error) {
	new_instruction := tt_body{
		Instruction_type: _instruction_type,
		Instruction: _instruction_body,
	}
	new_block, err := instantiate(_public_key, new_instruction, _private_key)
	if err != nil {
		return trustedtext_s{}, err
	}

	return new_block, nil

}

func announce_block_generation(_instruction_type string, _instruction_body string, _public_key string, _private_key string)  {
	block, err := generate_block(_instruction_type, _instruction_body, _public_key, _private_key)
	if err != nil {
		log.Println("Failed to initiate block, with error:", err)
		return
	}
	log.Println("Successfully created block, with hash:", block.Hash)
}
	
	
	
func main() {
	// webservice_main()

	tt_app := app.New()
	main_window := tt_app.NewWindow("Hello World Window!")
	main_window.SetFullScreen(true)

	body_input := widget.NewMultiLineEntry()
	body_input.SetMinRowsVisible(10)

	private_key_input := widget.NewPasswordEntry()
	private_key_input.SetPlaceHolder("private key")
	public_key_input := widget.NewEntry()
	public_key_input.SetPlaceHolder("public key")

	save_button := widget.NewButton("Save", func() {announce_block_generation("publish", body_input.Text, public_key_input.Text, private_key_input.Text)})

	

	content := container.NewVBox(body_input, private_key_input, public_key_input, save_button)


	main_window.SetContent(content)
	main_window.ShowAndRun()
}