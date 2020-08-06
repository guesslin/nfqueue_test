#!/usr/bin/env bash

TIME=$(date +%Y%m%d-%H%M%S)

docker build . -t guesslin/nfqueue:latest
docker tag guesslin/nfqueue:latest guesslin/nfqueue:$TIME
docker push guesslin/nfqueue:$TIME

echo
echo "Successfully pushed docker image 'guesslin/nfqueue:$TIME'"
