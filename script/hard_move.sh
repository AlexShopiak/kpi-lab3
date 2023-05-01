#!/bin/bash

curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "white"
curl -X POST http://localhost:17000 -d "figure 100 100"


for i in {1..34}
do
    curl -X POST http://localhost:17000 -d "move $i 5"
    curl -X POST http://localhost:17000 -d "update"
done

for i in {1..34}
do
    curl -X POST http://localhost:17000 -d "move -$i 5"
    curl -X POST http://localhost:17000 -d "update"
done

for i in {1..34}
do
    curl -X POST http://localhost:17000 -d "move $i -5"
    curl -X POST http://localhost:17000 -d "update"
done

for i in {1..34}
do
    curl -X POST http://localhost:17000 -d "move -$i -gi5"
    curl -X POST http://localhost:17000 -d "update"
done