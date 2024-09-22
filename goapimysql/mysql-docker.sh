#!/bin/bash
# $ sh mysql-docker.sh
docker stop test-mysql
docker rm test-mysql
docker run --name test-mysql -p 33061:3306 -e MYSQL_ROOT_PASSWORD=strong_password -e MYSQL_DATABASE=yourdb -d mysql