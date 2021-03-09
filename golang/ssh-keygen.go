package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
	"syscall/js"

	"golang.org/x/crypto/ssh"
)

type SSHKeyGen struct {
}

func (g *SSHKeyGen) GenerateRSAKeyPair(this js.Value, args []js.Value) interface{} {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil
	}

	// generate and write private key as PEM
	var privKeyBuf strings.Builder

	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err := pem.Encode(&privKeyBuf, privateKeyPEM); err != nil {
		return nil
	}

	// generate and write public key
	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil
	}

	var pubKeyBuf strings.Builder
	pubKeyBuf.Write(ssh.MarshalAuthorizedKey(pub))

	ret := map[string]interface{}{
		"privateKey": privKeyBuf.String(),
		"publicKey":  pubKeyBuf.String(),
	}
	return ret
}

func main() {
	fmt.Println("Hello, wasm!")
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	var g = &SSHKeyGen{}
	js.Global().Set("sshKeyGen", js.ValueOf(
		map[string]interface{}{
			"GenerateRSAKeyPair": js.FuncOf(g.GenerateRSAKeyPair),
		},
	))
}
