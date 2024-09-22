#!/bin/bash
# $ sh queuedocker.sh
docker stop simple-rabbit
docker rm simple-rabbit
docker run -d --hostname my-rabbit --name simple-rabbit -p 15672:15672 -p 5672:5672 -e RABBITMQ_DEFAULT_USER=user -e RABBITMQ_DEFAULT_PASS=password rabbitmq:3-management
