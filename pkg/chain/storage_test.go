package trustedtext_test

import (
	"file"
	"testing"
)

const default_config_path = "config.json"
func Test_storage_establishment(t *testing.T) {
	test_storage, err := file.Storage_from_file(default_config_path)

	if err != nil {
		t.Log("Cannot generate storage fron file", err)
		t.Fail()
	}

	_, err = test_storage.Chain.Read_chain()

	if err != nil {
		t.Log("cannot read from chain store", err)
		t.Fail()
	}

	_, err = test_storage.Config.Read_config()

	if err != nil {
		t.Log("cannot read config from store", err)
		t.Fail()
	}

	_, err = test_storage.Peerlist.Read_peerlist()

	if err != nil {
		t.Log("cannot read peerlist from store", err)
		t.Fail()
	}

}


func Test_store_retrieve_peers(t *testing.T) {
	t.Fail()
}

func Test_store_retrieve_chain(t *testing.T) {
	t.Fail()
}
