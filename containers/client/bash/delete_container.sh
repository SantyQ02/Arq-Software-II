#!/bin/bash

container_id=$1

service_name=$(docker inspect --format '{{ index .Config.Labels "com.docker.compose.service" }}' $container_id)

docker rm -f $container_id
docker-compose -f ../docker-compose-dev.yml restart nginx_$service_name