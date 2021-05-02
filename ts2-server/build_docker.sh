#!/usr/bin/env bash
cp ts2-server docker/
cp simulation/testdata/demo.json docker/
docker build -t ts2simulator/server --no-cache docker
