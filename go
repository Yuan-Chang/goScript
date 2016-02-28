#! /bin/bash
source global

filePath="${SCRIPT_PATH}/goScript/path/$1"
path=$(cat $filePath)

cd $path

