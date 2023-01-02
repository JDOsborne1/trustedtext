package main

import (
	"file"
	"log"
	"trustedtext"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/pkg/errors"
)

func announce_block_generation(_instruction_type string, _instruction_body string, _public_key string, _private_key string, _store file.Storage) {
	block, err := trustedtext.Generate_block(_instruction_type, _instruction_body, _public_key, _private_key)
	if err != nil {
		log.Println("Failed to initiate block, with error:", err)
		return
	}

	log.Println("Successfully created block, with hash:", block.Hash)

	err = amend_chain_with(_store, block)
	if err != nil {
		log.Println(err)
		return
	}

}

func amend_chain_with(_store file.Storage, _block trustedtext.Block) error {
	existing_chain, err := _store.Chain.Read_chain()
	if err != nil {
		errors.Wrap(err, "Failed to load chain")
		return err
	}

	new_chain, err := trustedtext.Process_incoming_block(existing_chain, _block)
	if err != nil {
		errors.Wrap(err, "Failed to process new block")
		return err
	}

	err = _store.Chain.Write_chain(new_chain)
	if err != nil {
		errors.Wrap(err, "Failed to write chain")
		return err
	}

	return nil
}

func announce_head_change_block(_new_head_hash string, _public_key string, _private_key string, _store file.Storage) {
	move_block, err := trustedtext.Generate_head_move_block(_public_key, _new_head_hash, _private_key)
	if err != nil {
		log.Println("Failed to generate new head block")
		return
	}

	err = amend_chain_with(_store, move_block)
	if err != nil {
		log.Println(err)
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

	save_button := widget.NewButton("Save", func() {
		announce_block_generation("publish", body_input.Text, public_key_input.Text, private_key_input.Text, _store)
	})
	content := container.NewVBox(body_input, private_key_input, public_key_input, save_button)

	main_window.SetContent(content)

	return main_window
}

func head_hash_change_window(_app_to_launch_in fyne.App, _store file.Storage) fyne.Window {
	change_window := _app_to_launch_in.NewWindow("Change head block")
	change_window.SetFullScreen(false)

	new_head_hash_input := widget.NewEntry()
	new_head_hash_input.SetPlaceHolder("Hash of new head block")

	private_key_input := widget.NewPasswordEntry()
	private_key_input.SetPlaceHolder("private key")
	public_key_input := widget.NewEntry()
	public_key_input.SetPlaceHolder("public key")

	save_button := widget.NewButton("Save", func() {
		announce_head_change_block(new_head_hash_input.Text, public_key_input.Text, private_key_input.Text, _store)
	})
	content := container.NewVBox(new_head_hash_input, private_key_input, public_key_input, save_button)

	change_window.SetContent(content)

	return change_window
}

func localapp(_given_store file.Storage) {
	tt_app := app.New()

	main_window := head_hash_change_window(tt_app, _given_store)
	main_window.ShowAndRun()

}
