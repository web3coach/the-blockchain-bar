package wallet

import "path/filepath"

const keystoreDirName = "keystore"
const AndrejAccount = "0x22ba1F80452E6220c7cc6ea2D1e3EEDDaC5F694A"
const BabaYagaAccount = "0x21973d33e048f5ce006fd7b41f51725c30e4b76b"
const CaesarAccount = "0x84470a31D271ea400f34e7A697F36bE0e866a716"

func GetKeystoreDirPath(dataDir string) string {
	return filepath.Join(dataDir, keystoreDirName)
}
