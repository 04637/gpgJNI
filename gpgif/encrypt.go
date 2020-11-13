package main

import (
	"C"
	"encoding/json"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
	"io"
	"os"
)

func main() {

}

//export EncryptByPubKey
func EncryptByPubKey(_pubKey, _toEnc, _toSave *C.char) *C.char {
	pubKey := C.GoString(_pubKey)
	toEnc := C.GoString(_toEnc)
	toSave := C.GoString(_toSave)
	recipient, err := readEntity(pubKey)
	if err != nil {
		return fail(err.Error())
	}

	f, err := os.Open(toEnc)
	if err != nil {
		return fail(err.Error())
	}
	defer f.Close()

	dst, err := os.Create(toSave)
	if err != nil {
		return fail(err.Error())
	}
	defer dst.Close()
	err = encrypt([]*openpgp.Entity{recipient}, nil, f, dst)
	if err != nil {
		return fail(err.Error())
	}
	return ok(toSave)
}

//export EncryptByPsw
func EncryptByPsw(_psw, _toEnc, _toSave *C.char) *C.char {
	psw := C.GoString(_psw)
	toEnc := C.GoString(_toEnc)
	toSave := C.GoString(_toSave)
	f, err := os.Open(toEnc)
	if err != nil {
		return fail(err.Error())
	}
	defer f.Close()

	dst, err := os.Create(toSave)
	if err != nil {
		return fail(err.Error())
	}
	err = symEncrypt([]byte(psw), f, dst)
	if err != nil {
		return fail(err.Error())
	}
	return ok(toSave)
}

func encrypt(recipient []*openpgp.Entity, signer *openpgp.Entity, r io.Reader, w io.Writer) error {
	wc, err := openpgp.Encrypt(w, recipient, signer, &openpgp.FileHints{IsBinary: true}, nil)
	if err != nil {
		return err
	}
	if _, err := io.Copy(wc, r); err != nil {
		return err
	}
	return wc.Close()
}

func symEncrypt(psw []byte, r io.Reader, w io.Writer) error {
	wc, err := openpgp.SymmetricallyEncrypt(w, psw, &openpgp.FileHints{IsBinary: true}, nil)
	if err != nil {
		return err
	}
	if _, err := io.Copy(wc, r); err != nil {
		return err
	}
	return wc.Close()
}

func readEntity(name string) (*openpgp.Entity, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	block, err := armor.Decode(f)
	if err != nil {
		return nil, err
	}
	return openpgp.ReadEntity(packet.NewReader(block.Body))
}

type Resp struct {
	Code int         `json:"Code"`
	Msg  string      `json:"Msg"`
	Data interface{} `json:"Data"`
}

func ok(data interface{}) *C.char {
	resp, _ := json.Marshal(Resp{
		Code: 200,
		Msg: "ok",
		Data: data,
	})
	return C.CString(string(resp))
}

func fail(msg string) *C.char {
	resp, _ := json.Marshal(Resp{
		Code: 500,
		Msg: msg,
		Data: nil,
	})
	return C.CString(string(resp))
}
