#!/bin/bash
# @File:         shell.sh
# @Aim:          常用shell 命令 
# @Author:       foyon
# @Created Time: 2020-11-12

ps auxwww | grep goland | awk -F " " '{a[$1]++}END{for(i in a){print a[i],i}}' | sort | uniq | sort -nr

## nginx acces_log 求每个请求的qps
cat acc.log| awk -F " "  '{a[$2 $3]++}END{for(i in a){ print a[i], i}}' | sort | uniq | sort -nr

