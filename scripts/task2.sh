#!/bin/bash

totalx=600
totaly=150
steps=20
interval=0.01

dx=$((totalx/steps))
dy=$((totaly/steps))

sleep $interval
curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "white"
curl -X POST http://localhost:17000 -d "figure 100 100"

for i in {1..20}
do
    curl -X POST http://localhost:17000 -d "move $dx $dy"
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
done

for i in {1..20}
do
    curl -X POST http://localhost:17000 -d "move -$dx $dy"
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
done

for i in {1..20}
do
    curl -X POST http://localhost:17000 -d "move $dx $dy"
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
done

for i in {1..20}
do
    curl -X POST http://localhost:17000 -d "move -$dx $dy"
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
done

for i in {1..30}
do
    curl -X POST http://localhost:17000 -d "move $dx $dy"
    curl -X POST http://localhost:17000 -d "update"
    sleep $interval
done