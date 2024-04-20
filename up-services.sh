#!/bin/bash

docker compose up --build | tee logs/services.log
