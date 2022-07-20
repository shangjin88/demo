#!/bin/bash

# 死循环 加1
count=1
while true
do
    touch "file-$count"
#    echo "shell-$count"
    let count++
#    break
done

