## Motivation

While playinf with Dex and using jwt.io to inspect the obtained tokens, I struggled to get my tokens validated easily. Looking around, found https://play.golang.org/p/wVusucNGDI and modified it to get a simple command line tool that can pe piped easily.

```
go build
```

Usage with dex:

```shell
$ curl -s https://dex.tools.local/keys | jq '.keys[] | select(.kid == "0ac8777d24a570e8fb8cbdb5c98e030aebe7b3be")' | ./jwk-x509
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxH7M8K7eNpnS26VSgJml
YS9XE9N34Cr91LnIDysaUE+UlQUOO4YG/SALTbZ6vEcWVcGb46qBvawU+uFT/jqd
uD/VGcltIvgQK+N9eqJtZjb8nuTYGtmB8bavWNO48LRtuCylAjA+6IEBH8RF4hSN
dbzq5kbEixkWyTt7ZgQWks3Rng92XepaAHPUYwpQrFVKPgxVEbfzSZkr1BCAPb7V
GdPTCRiCTxAQplqu5KzqSZkbI119cYuQY4vDwbyXZy+N5f9ANYBD6SOzg//nzcbu
BoKeVRMVPKRV111bvsYcT09n+oWRt3Tqo44enZFLcx8E/cxICQmXM7N8FblBlmG/
0QIDAQAB
-----END PUBLIC KEY-----
```
