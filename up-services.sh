#!/bin/bash

name=hackaton-2024

docker-compose up | tee logs/services.log
container_name=docker ps -a | grep $name | awk '{print$1}'

docker wait $container_name
docker stop $container_name && docker rm $contsiner_name

