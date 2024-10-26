#!/bin/bash

# Set up variables
CA_KEY="ca.key"
CA_CERT="ca.crt"
CLIENT_KEY="client.key"
CLIENT_CSR="client.csr"
CLIENT_CERT="client.crt"
DAYS_VALID=365

# Generate client private key
openssl genpkey -algorithm RSA -out $CLIENT_KEY -pkeyopt rsa_keygen_bits:2048

# Generate client certificate signing request (CSR)
openssl req -new -key $CLIENT_KEY -out $CLIENT_CSR -subj "/C=US/ST=State/L=City/O=Organization/OU=OrgUnit/CN=client.example.com"

# Generate client certificate signed by CA
openssl x509 -req -in $CLIENT_CSR -CA $CA_CERT -CAkey $CA_KEY -CAcreateserial -out $CLIENT_CERT -days $DAYS_VALID -sha256

echo "Certificates generated successfully."
