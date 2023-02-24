#!/bin/bash

# $1 - domain name, i.g. "tesseract.club"

domain=$1
pem_loc="/etc/letsencrypt/live"

cat $pem_loc/$domain/privkey.pem $pem_loc/$domain/fullchain.pem > /etc/monit/pemfile-$domain.pem
chmod 600 /etc/monit/pemfile-$domain.pem
