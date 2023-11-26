#!/bin/bash

container_ids=$(docker-compose -f ../docker-compose-dev.yml ps -q)
stats_output=$(docker stats --no-stream $container_ids)

# Construir un JSON a partir de la salida de docker stats
json_output="{ \"containers_stats\": ["
json_output+=$(echo "$stats_output" | awk 'NR>1 {printf "{\"container_id\":\"%s\",\"name\":\"%s\",\"cpu\":\"%s\",\"memory_usage\":\"%s\",\"memory_limit\":\"%s\",\"memory\":\"%s\",\"net_i\":\"%s\",\"net_o\":\"%s\",\"block_i\":\"%s\",\"block_o\":\"%s\"},",$1,$2,$3,$4,$6,$7,$8,$10,$11,$13}')
json_output=${json_output%,} # Eliminar la coma final
json_output+="]}"

echo $json_output
