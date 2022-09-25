package trustedtext

type Peer_detail struct {
	Claimed_name string
	Path         string
}

type Peerlist interface {
	Write_peerlist([]Peer_detail) error
	Read_peerlist() ([]Peer_detail, error)
}
