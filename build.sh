#!/bin/bash
VERSION=latest
docker build -t download_service .
docker tag download_service timoreymann/download_service:$VERSION
docker push timoreymann/download_service:$VERSION
