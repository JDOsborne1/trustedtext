package trustedtext

type Config_struct struct {
	Peerlist_path string
	Chain_path    string
	Port_used     int
	Authoritative_mode bool
}

type Config interface {
	Write_config(Config_struct) error
	Read_peerlist() (Config_struct, error)
}
