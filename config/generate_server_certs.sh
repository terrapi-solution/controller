#!/bin/bash

# Set up variables
CA_KEY="ca.key"
CA_CERT="ca.crt"
SERVER_KEY="server.key"
SERVER_CSR="server.csr"
SERVER_CERT="server.crt"
DAYS_VALID=365

# Generate CA private key
openssl genpkey -algorithm RSA -out $CA_KEY -pkeyopt rsa_keygen_bits:2048

# Generate CA certificate
openssl req -x509 -new -nodes -key $CA_KEY -sha256 -days $DAYS_VALID -out $CA_CERT -subj "/C=US/ST=State/L=City/O=Organization/OU=OrgUnit/CN=example.com"

# Generate server private key
openssl genpkey -algorithm RSA -out $SERVER_KEY -pkeyopt rsa_keygen_bits:2048

# Generate server certificate signing request (CSR)
openssl req -new -key $SERVER_KEY -out $SERVER_CSR -subj "/C=US/ST=State/L=City/O=Organization/OU=OrgUnit/CN=server.example.com"

# Generate server certificate signed by CA
openssl x509 -req -in $SERVER_CSR -CA $CA_CERT -CAkey $CA_KEY -CAcreateserial -out $SERVER_CERT -days $DAYS_VALID -sha256

echo "Certificates generated successfully."
