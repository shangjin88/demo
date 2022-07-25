#!/bin/bash

# 字符串转数组
ips="10.16.1.1,10.16.1.2"
ip_array=(${ips//,/ })

# 遍历数组
for ip in ${ip_array[@]}; do
    echo "$ip"
done

# 遍历数组
for ip in $ips; do
    echo "$ip"
done



