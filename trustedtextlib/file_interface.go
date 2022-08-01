package trustedtext

import (
	"encoding/json"
	"io/ioutil"
)

func read_peerlist(_config config_struct) ([]peer_detail, error) {
	bytefile, err := ioutil.ReadFile(_config.Peerlist_path)
	if err != nil {
		return []peer_detail{}, err
	}
	peerlist := &[]peer_detail{}
	err = json.Unmarshal(bytefile, peerlist)
	if err != nil {
		return []peer_detail{}, err
	}
	return *peerlist, nil
}

func write_peerlist(_peerlist []peer_detail, _config config_struct) error {
	marshalled_peerlist, err := json.MarshalIndent(_peerlist, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(_config.Peerlist_path, marshalled_peerlist, 0644)

	if err != nil {
		return err
	}

	return nil
}

func write_config(_config config_struct) error {
	marshalled_config, err := json.MarshalIndent(
		_config,
		"",
		"  ",
	)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(
		"config.json",
		marshalled_config,
		0644,
	)
	if err != nil {
		return err
	}
	return nil
}

func Read_config(_config_path string) (config_struct, error) {
	bytefile, err := ioutil.ReadFile(_config_path)
	if err != nil {
		return config_struct{}, err
	}
	config := &config_struct{}
	err = json.Unmarshal(bytefile, config)
	if err != nil {
		return config_struct{}, err
	}

	return *config, nil
}

func Write_chain(_chain trustedtext_chain_s, _config config_struct) error {
	marshalled_chain, err := json.MarshalIndent(_chain, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_config.Chain_path, marshalled_chain, 0644)
	if err != nil {
		return err
	}
	return err
}

func Read_chain(_config config_struct) (trustedtext_chain_s, error) {
	bytefile, err := ioutil.ReadFile(_config.Chain_path)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	chain := &trustedtext_chain_s{}
	err = json.Unmarshal(bytefile, chain)
	if err != nil {
		return trustedtext_chain_s{}, err
	}
	return *chain, nil
}
