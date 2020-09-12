# log
sh ~/server-script/refresh_log.sh

# app restart
cd /home/isucon/isuumo/webapp/go && make && cd /home/isucon
sudo systemctl disable isuumo.go.service

# nginx restart 
sudo systemctl restart nginx

# mysql restart 
