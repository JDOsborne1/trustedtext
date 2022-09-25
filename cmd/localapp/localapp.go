package main

import (
	"file"
	"log"
	"trustedtext"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)


func announce_block_generation(_instruction_type string, _instruction_body string, _public_key string, _private_key string, _store file.Storage)  {
	block, err := trustedtext.Generate_block(_instruction_type, _instruction_body, _public_key, _private_key)
	if err != nil {
		log.Println("Failed to initiate block, with error:", err)
		return
	}
	log.Println("Successfully created block, with hash:", block.Hash)
	
	existing_chain, err := _store.Chain.Read_chain()
	if err != nil {
		log.Println("Failed to load chain, with error:", err)
		return
	}
	
	new_chain, err := trustedtext.Process_incoming_block(existing_chain, block)
	if err != nil {
		log.Println("Failed to process new block, with error:", err)
		return
	}
	
	err = _store.Chain.Write_chain(new_chain)
	if err != nil {
		log.Println("Failed to write chain, with error:", err)
		return
	}
}



func block_generator_window(_app_to_launch_in fyne.App, _store file.Storage) fyne.Window {
	main_window := _app_to_launch_in.NewWindow("Block Generator window")
	main_window.SetFullScreen(false)
	
	body_input := widget.NewMultiLineEntry()
	body_input.SetMinRowsVisible(10)
	
	private_key_input := widget.NewPasswordEntry()
	private_key_input.SetPlaceHolder("private key")
	public_key_input := widget.NewEntry()
	public_key_input.SetPlaceHolder("public key")
	
	save_button := widget.NewButton("Save", func() {announce_block_generation("publish", body_input.Text, public_key_input.Text, private_key_input.Text, _store)})
	content := container.NewVBox(body_input, private_key_input, public_key_input, save_button)
	
	main_window.SetContent(content)

	return main_window
}


func localapp() {
	tt_app := app.New()

	store := file.Storage{}
	
	main_window := block_generator_window(tt_app, store)
	main_window.ShowAndRun()

}
