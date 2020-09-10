#!/usr/bin/env sh

apt-get update
apt-get -yV upgrade
apt -yV install redis-doc
apt -yV install redis-server
