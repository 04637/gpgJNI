package main

import "testing"

func TestEncryptByPubKey(t *testing.T) {
	pubKey := "E:/qq2/1.pub"
	fileToEnc := "E:/qq2/1.rar"
	resp := EncryptByPubKey(pubKey, fileToEnc, fileToEnc+".pgp")
	if resp.code != 200 {
		t.Error(resp)
	} else {
		t.Log(resp)
	}
}

func TestEncryptByPsw(t *testing.T) {
	psw := ""
	fileToEnc := "E:/qq2/1.rar"
	fileToSave := "E:/qq2/psw.rar2.pgp"
	resp := EncryptByPsw(psw, fileToEnc, fileToSave)
	if resp.code != 200 {
		t.Error(resp)
	} else {
		t.Log(resp)
	}
}
