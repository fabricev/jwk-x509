package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"

	jose "gopkg.in/square/go-jose.v2"
)

func main() {
	bodyBytes, err := ioutil.ReadAll(os.Stdin)

	// Unmarshall input into JWK object
	var key jose.JSONWebKey
	if err := json.Unmarshal(bodyBytes, &key); err != nil {
		log.Fatal("Can't read stdin")
		panic(err)
	}

	// Unmarshall public key data
	data, err := x509.MarshalPKIXPublicKey(key.Key)
	if err != nil {
		log.Fatal("Wrong JWK format")
		panic(err)
	}

	// X509 public key conversion
	block := pem.Block{Type: "PUBLIC KEY", Bytes: data}
	if err := pem.Encode(os.Stdout, &block); err != nil {
		log.Fatal("Could not encode public key in X509")
		panic(err)
	}
}
