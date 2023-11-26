#!/bin/bash

nombre_servicio=$1
cantidad=$2

if ! docker-compose -f ../docker-compose-dev.yml config --services | grep -qw "$nombre_servicio"; then
    echo "El servicio $nombre_servicio no existe en el archivo de configuración."
    exit 1
fi

docker-compose -f ../docker-compose-dev.yml up --scale $nombre_servicio=$cantidad -d
docker-compose -f ../docker-compose-dev.yml restart nginx_$nombre_servicio