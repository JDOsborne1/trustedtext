package trustedtext

import (
	"encoding/json"
	"os"
)

type Peer_detail struct {
	Claimed_name string
	Path         string
}

func Read_peerlist(_config Config_struct) ([]Peer_detail, error) {
	bytefile, err := os.ReadFile(_config.Peerlist_path)
	if err != nil {
		return []Peer_detail{}, err
	}
	peerlist := &[]Peer_detail{}
	err = json.Unmarshal(bytefile, peerlist)
	if err != nil {
		return []Peer_detail{}, err
	}
	return *peerlist, nil
}

func Write_peerlist(_peerlist []Peer_detail, _config Config_struct) error {
	marshalled_peerlist, err := json.MarshalIndent(_peerlist, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(_config.Peerlist_path, marshalled_peerlist, 0644)

	if err != nil {
		return err
	}

	return nil
}

func write_config(_config Config_struct) error {
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

func Read_config(_config_path string) (Config_struct, error) {
	bytefile, err := os.ReadFile(_config_path)
	if err != nil {
		return Config_struct{}, err
	}
	config := &Config_struct{}
	err = json.Unmarshal(bytefile, config)
	if err != nil {
		return Config_struct{}, err
	}

	return *config, nil
}

func Write_chain(_chain Trustedtext_chain_s, _config Config_struct) error {
	marshalled_chain, err := json.MarshalIndent(_chain, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(_config.Chain_path, marshalled_chain, 0644)
	if err != nil {
		return err
	}
	return err
}

func Read_chain(_config Config_struct) (Trustedtext_chain_s, error) {
	bytefile, err := os.ReadFile(_config.Chain_path)
	if err != nil {
		return Trustedtext_chain_s{}, err
	}
	chain := &Trustedtext_chain_s{}
	err = json.Unmarshal(bytefile, chain)
	if err != nil {
		return Trustedtext_chain_s{}, err
	}
	return *chain, nil
}
