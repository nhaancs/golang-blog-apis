#!/usr/bin/env bash

if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

echo "Setting up netdata..."

ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} \
  "docker rm -f netdata && docker run -d --name=netdata \
      -p 19999:19999 \
      -v netdataconfig:/etc/netdata \
      -v netdatalib:/var/lib/netdata \
      -v netdatacache:/var/cache/netdata \
      -v /etc/passwd:/host/etc/passwd:ro \
      -v /etc/group:/host/etc/group:ro \
      -v /proc:/host/proc:ro \
      -v /sys:/host/sys:ro \
      -v /etc/os-release:/host/etc/os-release:ro \
      --restart unless-stopped \
      --cap-add SYS_PTRACE \
      --security-opt apparmor=unconfined \
      netdata/netdata"

# -e VIRTUAL_HOST=${MONITORING_VIRTUAL_HOST} \
# -e VIRTUAL_PORT=${MONITORING_VIRTUAL_PORT} \
# -e LETSENCRYPT_HOST=${MONITORING_LETSENCRYPT_HOST} \
# -e LETSENCRYPT_EMAIL=${MONITORING_LETSENCRYPT_EMAIL} \

echo "Done"
