#!/bin/bash

openssl req -x509 -new -nodes -newkey rsa:4096 -days 365 \
  -keyout root.key -out root.pem \
  -subj "/C=XW/ST=Default State/L=Default/O=Development/OU=Local Development/CN=Local CA" || exit 1
