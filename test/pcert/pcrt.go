package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

func ParseCert(cert []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(cert)
	if block == nil {
		return nil, nil
	}
	return x509.ParseCertificate(block.Bytes)
}

func main(){
	cert , err := ioutil.ReadFile("cacert/DYNRO.crt")
	if err != nil {
		panic(err)
	}
	crt , err := ParseCert(cert)
	println(crt)
}