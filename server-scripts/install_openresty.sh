wget -qO - https://openresty.org/package/pubkey.gpg | apt-key add -
apt-get -y install software-properties-common
add-apt-repository -y "deb http://openresty.org/package/ubuntu $(lsb_release -sc) main"
apt-get update
apt-get install -y libpcre3-dev \
    libssl-dev perl make build-essential curl
apt-get install -y openresty

# sudo service openresty start
