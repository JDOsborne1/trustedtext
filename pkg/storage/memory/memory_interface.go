package memory

import "trustedtext"

type Storage struct {
	Peerlist Memory_peerlist
	Chain Memory_chain
	Config Memory_config
}

type Memory_peerlist struct {
	peers []trustedtext.Peer_detail
}

func (m Memory_peerlist) Write_peerlist(_input []trustedtext.Peer_detail) error {
	m.peers = _input
	
	return nil
}

func (m Memory_peerlist) Read_peerlist() ([]trustedtext.Peer_detail, error) {
	return m.peers, nil
}

type Memory_chain struct {
	chain trustedtext.Trustedtext_chain_s
}

func (m Memory_chain) Write_chain(_input trustedtext.Trustedtext_chain_s) error {
	m.chain = _input
	return nil
}

func (m Memory_chain) Read_chain() (trustedtext.Trustedtext_chain_s, error) {
	return m.chain, nil
}


type Memory_config struct {}
