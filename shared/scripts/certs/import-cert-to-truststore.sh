#!/usr/bin/env bash

keytool -import -alias selfsigned -file ./shared/certs/crane.crt -keystore truststore.jks