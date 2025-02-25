#!/bin/bash

# 使用方法：
# ./genModel.sh [数据库名] [表名]
# 示例 ./genModel.sh upms base_option
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改【 package 】

#生成的表名
tablename=$2
#表生成的genmodel目录
modeldir=./

# 数据库配置
host=106.15.74.79
port=3310
dbname=$1
username=dev
passwd=sq654321!
#ignoreColumns=deleted,created_date,created_user_name,created_user_id,modified_date,modified_user_name,modified_user_id
#if [ -z"$dbname" -o -z"$tablename" ];then
#    read -p "请输入数据库名和表明:" var1 var2
#    dbname=$var1
#    tablename=$var2
#fi

echo "开始创建库：$dbname 的表：$tablename"
# 可以使用-home指定模版路径，默认是~/.goctl
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tablename}"  -dir="${modeldir}" --home=./1.5.6 --style=goZero #-i="${ignoreColumns}"