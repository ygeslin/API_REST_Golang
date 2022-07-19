#!/bin/bash

docker pull mongo
docker run --name mongo-db -p 27017:27017 -d mongo:latest