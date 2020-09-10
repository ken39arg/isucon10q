#!/usr/bin/env sh
apt-get update \
	&& apt-get install -y --no-install-recommends \
		bash \
		curl \
		gzip \
		unzip \
		wget \
		make \
		git \
		jq \
		dstat \
		iotop \
		iftop \
		socat \
		vim \
		sed \
		rsync \
		ngrep \
		tcpdump \
		percona-toolkit \
		default-mysql-client \
	&& rm -rf /var/lib/apt/lists/*

# 違ったらしょうが無い
wget --no-check-certificate https://github.com/tkuchiki/alp/releases/download/v1.0.3/alp_linux_amd64.zip
unzip alp_linux_amd64.zip
install ./alp ./alp /usr/local/bin
