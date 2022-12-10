#!/usr/bin/env bash

docker image rm truongnh28/spotify
docker build --tag truongnh28/spotify .
docker push truongnh28/spotify