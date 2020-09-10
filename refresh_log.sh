LOGDIR=${HOME}/log
ID=`date '+%Y%m%d%H%M%S'`

mkdir -p ${LOGDIR}/${ID}

ACCESSLOG=/var/log/nginx/access.log
cp ${ACCESSLOG} ${LOGDIR}/${ID}/access.log
cp /dev/null ${ACCESSLOG}

SLOWLOG=/var/log/slow.log
cp ${SLOWLOG} ${LOGDIR}/${ID}/slow.log
cp /dev/null ${SLOWLOG}

# ほしければ
# APPLOG=/var/log/app.log
# cp ${APPLOG} ${LOGDIR}/${ID}/app.log
# cp /dev/null ${APPLOG}

