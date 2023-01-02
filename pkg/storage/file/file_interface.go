package file

import (
	"encoding/json"
	"os"
	"trustedtext"

	"github.com/pkg/errors"
)

type Storage struct {
	Peerlist File_peerlist
	Chain    File_chain
	Config   File_config
}

func Storage_from_file(_config_path string) (Storage, error) {
	config_entity := File_config{
		path: _config_path,
	}

	tt_config, err := config_entity.Read_config()

	if err != nil {
		return Storage{}, err
	}

	chain_entity := File_chain{
		path: tt_config.Chain_path,
	}

	peerlist_entity := File_peerlist{
		path: tt_config.Peerlist_path,
	}

	storage_entity := Storage{
		Peerlist: peerlist_entity,
		Chain:    chain_entity,
		Config:   config_entity,
	}

	return storage_entity, nil
}

type File_peerlist struct {
	path string
}

func Generate_peerlist_from_config(_config trustedtext.Config_struct) File_peerlist {
	needed_path := _config.Peerlist_path
	return File_peerlist{path: needed_path}
}

func (_peerlist_store File_peerlist) Read_peerlist() ([]trustedtext.Peer_detail, error) {
	bytefile, err := os.ReadFile(_peerlist_store.path)

	if err != nil {
		return []trustedtext.Peer_detail{}, errors.Wrap(err, "Fails to generate bytefile")
	}
	peerlist := &[]trustedtext.Peer_detail{}
	err = json.Unmarshal(bytefile, peerlist)
	if err != nil {
		return []trustedtext.Peer_detail{}, errors.Wrap(err, "Fails to unmarshal bytefile")
	}
	return *peerlist, nil
}

func (_peerlist_store File_peerlist) Write_peerlist(_peerlist []trustedtext.Peer_detail) error {
	marshalled_peerlist, err := json.MarshalIndent(_peerlist, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(_peerlist_store.path, marshalled_peerlist, 0644)

	if err != nil {
		return err
	}

	return nil
}

type File_config struct {
	path string
}

//TODO - Fix the config writer to actually satisfy the interface

func (_config File_config) Write_config() error {
	marshalled_config, err := json.MarshalIndent(
		_config,
		"",
		"  ",
	)
	if err != nil {
		return err
	}
	err = os.WriteFile(
		"config.json",
		marshalled_config,
		0644,
	)
	if err != nil {
		return err
	}
	return nil
}

func (_config File_config) Read_config() (trustedtext.Config_struct, error) {
	bytefile, err := os.ReadFile(_config.path)
	if err != nil {
		return trustedtext.Config_struct{}, err
	}
	config := &trustedtext.Config_struct{}
	err = json.Unmarshal(bytefile, config)
	if err != nil {
		return trustedtext.Config_struct{}, err
	}

	return *config, nil
}

type File_chain struct {
	path string
}

func (_file_chain File_chain) Write_chain(_chain trustedtext.Chain) error {
	marshalled_chain, err := json.MarshalIndent(_chain, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(_file_chain.path, marshalled_chain, 0644)
	if err != nil {
		return err
	}
	return err
}

func (_file_chain File_chain) Read_chain() (trustedtext.Chain, error) {
	bytefile, err := os.ReadFile(_file_chain.path)
	if err != nil {
		return trustedtext.Chain{}, err
	}
	chain := &trustedtext.Chain{}
	err = json.Unmarshal(bytefile, chain)
	if err != nil {
		return trustedtext.Chain{}, err
	}
	return *chain, nil
}
