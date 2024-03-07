#!/bin/bash

cn=$1
dnsName=$2
certsDir="certs"

if ! [ -e $certsDir ]; then
  mkdir $certsDir
fi

openssl req -new -nodes -newkey rsa:2048 \
  -keyout "$certsDir/$cn.key" -out "$certsDir/$cn.csr" \
  -subj "/C=XW/ST=Default State/L=Default/O=Development/OU=Local Development/CN=$cn" || exit 1

echo "authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names

[alt_names]
DNS.1 = $dnsName" > "$certsDir/$cn.ext"

openssl x509 -req -in "$certsDir/$cn.csr" -CA root.pem -CAkey root.key \
  -CAcreateserial -out "$certsDir/$cn.crt" -days 30 -sha256 -extfile "$certsDir/$cn.ext"

rm "$certsDir/$cn.csr" "$certsDir/$cn.ext"
