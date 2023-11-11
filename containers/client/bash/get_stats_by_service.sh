#!/bin/bash

nombre_servicio=$1

if ! docker-compose -f ../docker-compose-dev.yml config --services | grep -qw "$nombre_servicio"; then
    echo "El servicio $nombre_servicio no existe en el archivo de configuración."
    exit 1
fi

# Obtener IDs de contenedores del servicio
container_ids=$(docker-compose -f ../docker-compose-dev.yml ps -q $nombre_servicio)

# Verificar si hay contenedores para el servicio
if [ -z "$container_ids" ]; then
    echo "{ \"containers_stats\": [] }"
    exit 0
fi

# Obtener estadísticas de los contenedores del servicio
stats_output=$(docker stats --no-stream $container_ids)

# Construir un JSON a partir de la salida de docker stats
json_output="{ \"containers_stats\": ["
json_output+=$(echo "$stats_output" | awk 'NR>1 {printf "{\"container_id\":\"%s\",\"name\":\"%s\",\"cpu\":\"%s\",\"memory_usage\":\"%s\",\"memory_limit\":\"%s\",\"memory\":\"%s\",\"net_i\":\"%s\",\"net_o\":\"%s\",\"block_i\":\"%s\",\"block_o\":\"%s\"},",$1,$2,$3,$4,$6,$7,$8,$10,$11,$13}')
json_output=${json_output%,} # Eliminar la coma final
json_output+="]}"

echo $json_output
