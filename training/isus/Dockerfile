FROM ubuntu:latest

ENV NOTVISIBLE "in users profile"
RUN apt-get update && apt-get install -y --no-install-recommends \
		bash \
		curl \
		openssh-server \
		sudo \
	&& rm -rf /var/lib/apt/lists/* \
	&& echo 'root:screencast' | chpasswd \
	&& mkdir -p /var/run/sshd

RUN useradd -m isucon \
	&& echo 'isucon:isucon' | chpasswd \
	&& usermod -aG sudo isucon \
	&& mkdir -p /home/isucon/.ssh \
    && echo "newuser    ALL=(ALL)       ALL" >> /etc/sudoers \
	&& touch /home/isucon/.ssh/authorized_keys \
	&& curl -k https://github.com/ken39arg.keys >> /home/isucon/.ssh/authorized_keys \
	&& chown -R isucon:isucon /home/isucon/.ssh \
	&& chmod 700 /home/isucon/.ssh \
	&& chmod 600 /home/isucon/.ssh/authorized_keys

EXPOSE 22
CMD ["/usr/sbin/sshd", "-D"]
