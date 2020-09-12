# log
#sudo sh ~/server-scripts/refresh_log.sh

# app restart
cd /home/isucon/isuumo/webapp/go && make && cd /home/isucon
sudo systemctl restart isuumo.go.service

# nginx restart 
sudo systemctl restart nginx

# mysql restart 
