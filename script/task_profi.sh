#!/bin/bash

#CONFIG-----------------------------------------------------
min=100
avg=400
max=700

interval=0.01
posX=100
posY=100
isBorder=0

stepX=25
stepY=25
steps=(15 20 25 30)


#FUNCTIONS-------------------------------------------------------
changeSteps () {
    index1=$(( $RANDOM % 3 + 1 )) #?
    index2=$(( $RANDOM % 3 + 1 )) #?
    stepX=${steps[$index1]}
    stepY=${steps[$index2]}
}

checkBorder () {
    if [ $posX -ge $max ]; then
    isBorder=1
    elif [ $posX -le $min ]; then
    isBorder=1
    elif [ $posY -ge $max ]; then
    isBorder=1
    elif [ $posY -le $min ]; then
    isBorder=1
    else
    isBorder=0
    fi
}

startAnim++ () {
    while true;
    do

    curl -X POST http://localhost:17000 -d "move $stepX $stepY"
    curl -X POST http://localhost:17000 -d "update"
    posX=$((posX+stepX))
    posY=$((posY+stepY))

    checkBorder
    if [ $isBorder -eq 1 ];then
    break
    fi

    done
}

startAnim+- () {
    while true;
    do

    curl -X POST http://localhost:17000 -d "move $stepX -$stepY"
    curl -X POST http://localhost:17000 -d "update"
    posX=$((posX+stepX))
    posY=$((posY-stepY))

    checkBorder
    if [ $isBorder -eq 1 ];then
    break
    fi

    done
}

startAnim-+ () {
    while true;
    do

    curl -X POST http://localhost:17000 -d "move -$stepX $stepY"
    curl -X POST http://localhost:17000 -d "update"
    posX=$((posX-stepX))
    posY=$((posY+stepY))

    checkBorder
    if [ $isBorder -eq 1 ];then
    break
    fi

    done
}

startAnim-- () {
    while true;
    do

    curl -X POST http://localhost:17000 -d "move -$stepX -$stepY"
    curl -X POST http://localhost:17000 -d "update"
    posX=$((posX-stepX))
    posY=$((posY-stepY))

    checkBorder
    if [ $isBorder -eq 1 ];then
    break
    fi

    done
}

#START_ANIMATION----------------------------------------------------
curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "white"
curl -X POST http://localhost:17000 -d "figure $posX $posY"
curl -X POST http://localhost:17000 -d "update"
sleep $interval

startAnim++

startAnim () {
    if [ $posX -ge $max ]; then
        if [ $posY -ge $avg ]; then
        changeSteps
        startAnim--
        isBorder=0
        startAnim
        else
        changeSteps
        startAnim-+
        isBorder=0
        startAnim
        fi
    fi

    if [ $posX -le $min ]; then
        if [ $posY -ge $avg ]; then
        changeSteps
        startAnim+-
        isBorder=0
        startAnim
        else
        changeSteps
        startAnim+-
        isBorder=0
        startAnim
        fi
    fi

    if [ $posY -ge $max ]; then
        if [ $posX -ge $avg ]; then
        changeSteps
        startAnim+-
        isBorder=0
        startAnim
        else
        changeSteps
        startAnim-- 
        isBorder=0
        startAnim
        fi
    fi

    if [ $posY -le $min ]; then
        if [ $posX -ge $avg ]; then
        changeSteps
        startAnim-+ 
        isBorder=0
        startAnim
        else
        changeSteps
        startAnim++
        isBorder=0
        startAnim
        fi
    fi
}
startAnim

