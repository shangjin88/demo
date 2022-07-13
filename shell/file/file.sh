#!/bin/bash

# 按行读取
cat /etc/passwd | while read line
do
    echo "$line"
done