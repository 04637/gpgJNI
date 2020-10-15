package main

import (
	"io"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func main() {

}

// EncryptByPubKey 使用公钥加密文件
// pubKey 公钥文件路径
// toEnc 待加密文件路径
// toSave 加密文件保存路径
func EncryptByPubKey(pubKey, toEnc, toSave string) Resp {
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

// EncryptByPsw 使用密码加密文件
// psw 密码
// toEnc 待加密文件路径
// toSave 加密文件保存路径
func EncryptByPsw(psw, toEnc, toSave string) Resp {
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
	code int
	msg  string
	data interface{}
}

func ok(data interface{}) Resp {
	return Resp{
		code: 200,
		msg:  "ok",
		data: data,
	}
}

func fail(msg string) Resp {
	return Resp{
		code: 500,
		msg:  msg,
		data: nil,
	}
}
