package trustedtext

const default_config_path = "config.json"

type Config_struct struct {
	Peerlist_path string
	Chain_path    string
	Port_used     int
}
