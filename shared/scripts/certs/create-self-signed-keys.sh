#!/usr/bin/env bash

openssl req \
  -newkey rsa:2048 -nodes -keyout ./shared/certs/crane.key \
  -x509 -sha256 -days 365 -out ./shared/certs/crane.crt