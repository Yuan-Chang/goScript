#!/bin/zsh

source global

filePath="${SCRIPT_PATH}/goScript/path/$1"
myPath=$(cat $filePath)

cd $myPath

