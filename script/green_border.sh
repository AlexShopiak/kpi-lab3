#!/bin/bash

curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "white"
curl -X POST http://localhost:17000 -d "bgrect 50 50 750 750"
curl -X POST http://localhost:17000 -d "figure 300 300"
curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "figure 500 500"
curl -X POST http://localhost:17000 -d "update"