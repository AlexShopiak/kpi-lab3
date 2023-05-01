#!/bin/bash

curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "white"
curl -X POST http://localhost:17000 -d "figure 100 100"


for i in {1..40}
do
    curl -X POST http://localhost:17000 -d "move 15 15"
    curl -X POST http://localhost:17000 -d "update"
done

for i in {1..40}
do
    curl -X POST http://localhost:17000 -d "move -15 -15"
    curl -X POST http://localhost:17000 -d "update"
done


