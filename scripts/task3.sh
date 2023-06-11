#!/bin/bash

#=======================CONFIG===============================
#============================================================

#Borders for figure center
min=100
avg=400
max=700

#Current status
interval=0.01
posX=100
posY=100
elser=0 #just for else

#Length of step
stepX=25
stepY=25

#Backgroung color
R=51
G=200 #const
B=51
A=0.7

maxClr=255
minClr=50


#Color changing
toIncR=1
toDecR=0
toIncB=0
toDecB=0

clrStep=10

#=======================FUNCTIONS============================
#============================================================

changeRGBA () {
    #change colors
    if [ $toIncR -eq 1 ]; then
        R=$((R+clrStep))
    elif [ $toDecR -eq 1 ]; then
        R=$((R-clrStep))
    elif [ $toIncB -eq 1 ]; then
        B=$((B+clrStep))
    elif [ $toDecB -eq 1 ]; then
        B=$((B-clrStep))
    else
        elser=0
    fi
    #change flags
    if [ $R -ge $maxClr ]; then
        R=$((maxClr-1))
        toIncR=0
        toDecR=0
        toIncB=1
        toDecB=0
    elif [ $R -le $minClr ]; then
        R=$((minClr+1))
        toIncR=0
        toDecR=0
        toIncB=0
        toDecB=1
    elif [ $B -ge $maxClr ]; then
        B=$((maxClr-1))
        toIncR=0
        toDecR=1
        toIncB=0
        toDecB=0
    elif [ $B -le $minClr ]; then
        B=$((minClr+1))
        toIncR=1
        toDecR=0
        toIncB=0
        toDecB=0
    else
        elser=0
    fi
}

changeSteps () {
    stepX=$(( $RANDOM % 30 + 20 ))
    stepY=$(( $RANDOM % 20 + 10 ))
}

movepp () {
    while true;
    do
        changeRGBA
        curl -X POST http://localhost:17000 -d "fill $R $G $B $A"
        curl -X POST http://localhost:17000 -d "move $stepX $stepY"
        curl -X POST http://localhost:17000 -d "update"
        posX=$((posX+stepX))
        posY=$((posY+stepY))

        if [ $posX -ge $max ]; then
        break
        elif [ $posX -le $min ]; then
        break
        elif [ $posY -ge $max ]; then
        break
        elif [ $posY -le $min ]; then
        break
        else
            elser=0
        fi

        sleep $interval
    done
}

movepm () {
    while true;
    do
        changeRGBA
        curl -X POST http://localhost:17000 -d "fill $R $G $B $A"
        curl -X POST http://localhost:17000 -d "move $stepX -$stepY"
        curl -X POST http://localhost:17000 -d "update"
        posX=$((posX+stepX))
        posY=$((posY-stepY))

        if [ $posX -ge $max ]; then
        break
        elif [ $posX -le $min ]; then
        break
        elif [ $posY -ge $max ]; then
        break
        elif [ $posY -le $min ]; then
        break
        else
            elser=0
        fi

        sleep $interval
    done
}

movemp () {
    while true;
    do
        changeRGBA
        curl -X POST http://localhost:17000 -d "fill $R $G $B $A"
        curl -X POST http://localhost:17000 -d "move -$stepX $stepY"
        curl -X POST http://localhost:17000 -d "update"
        posX=$((posX-stepX))
        posY=$((posY+stepY))

        if [ $posX -ge $max ]; then
        break
        elif [ $posX -le $min ]; then
        break
        elif [ $posY -ge $max ]; then
        break
        elif [ $posY -le $min ]; then
        break
        else
            elser=0
        fi

        sleep $interval
    done
}

movemm () {
    while true;
    do
        changeRGBA
        curl -X POST http://localhost:17000 -d "fill $R $G $B $A"
        curl -X POST http://localhost:17000 -d "move -$stepX -$stepY"
        curl -X POST http://localhost:17000 -d "update"
        posX=$((posX-stepX))
        posY=$((posY-stepY))

        if [ $posX -ge $max ]; then
        break
        elif [ $posX -le $min ]; then
        break
        elif [ $posY -ge $max ]; then
        break
        elif [ $posY -le $min ]; then
        break
        else
            elser=0
        fi

        sleep $interval
    done
}

startAnim () {
    if [ $posX -ge $max ]; then
        if [ $posY -ge $avg ]; then
            changeSteps
            movemm
            startAnim
        else
            changeSteps
            movemp
            startAnim
        fi
    elif [ $posX -le $min ]; then
        if [ $posY -ge $avg ]; then
            changeSteps
            movepm
            startAnim
        else
            changeSteps
            movepm
            startAnim
        fi
    elif [ $posY -ge $max ]; then
        if [ $posX -ge $avg ]; then
            changeSteps
            movepm
            startAnim
        else
            changeSteps
            movemm
            startAnim
        fi
    elif [ $posY -le $min ]; then
        if [ $posX -ge $avg ]; then
            changeSteps
            movemp
            startAnim
        else
            changeSteps
            movepp
            startAnim
        fi
    else 
        elser=0
    fi
}



#=======================START_ANIMATION======================
#============================================================
curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "fill $R $G $B $A"
curl -X POST http://localhost:17000 -d "figure $posX $posY"
curl -X POST http://localhost:17000 -d "update"
sleep $interval

movepp
startAnim