#!/bin/bash

red="\e[31m"
blue="\e[34m"
end="\e[0m"
bold="\e[1mBold"
# Variables
DataSetPath="./data/DataSet.json"
Id1="1t5VsIBXpGl4s8C4CAXTsAlIZISYEOlicj14obz3CwFXCCvaRyuhDI10fah5IfdMS3VblW51my8xt6aQvJI3qNg5as0yqoTCvdZd"
Password1="CGUsfQkS06lo7LM2guSV"
Id2="mAdo98L1nvCzdw4CanEmah8PrPqzFNQmO509HrdKCxcLEcluy1zwm9PLvSIOWhZxDFtnM3rLvjK4cVKVpC0BjYbeU6KpDe14Eh0T"
Password2="VMRmmHi0NvcGHQHI5wrA"
Id3="buhnVfyzGfe4qfuRsqwjGhzoF8ubXcM4JXqmLRmvPmJdLXtITBY99SI0rER0bX0INj2XkAyMkLUEkQrxNG46PDXHt00jGURsnjHy"
Password3="o6sRZkUSroEP4o51npXb"

function jsonValue() {
KEY=$1
num=$2
awk -F"[,:}]" '{for(i=1;i<=NF;i++){if($i~/'$KEY'\042/){print $(i+1)}}}' | tr -d '"' | sed -n ${num}p

}


#Inital post that fill the database with data set
echo '--------------------------------------------------------------------'
echo -e "$blue $bold Inital post that fill the database with data set $end"
echo '--------------------------------------------------------------------'
# curl -X POST --header "Content-Type: application/json" -d @$DataSetPath  http://localhost:8080/add/users
echo 'Refer to API log to see the results'
echo ''

echo '--------------------------------------------------------------------'
echo -e "$blue $bold Login with Id1 : $Id1 $end"
echo '--------------------------------------------------------------------'
# echo ' curl -X POST http://localhost:8080/login'
# curl -X POST --header "Content-Type: application/json" -d "{\"id\":\"$Id1\",\"password\": \"$Password1\"}" http://localhost:8080/login
token1=$( curl -X POST --header "Content-Type: application/json" -d "{\"id\":\"$Id1\",\"password\": \"$Password1\"}" http://localhost:8080/login | sed -e "s/{\"access_token\":\"//" | sed -e "s/\",\"status\":\"success\"+$}//")
# echo '--------------------------------------------------------------------'
# echo $token1


# echo ' curl -X DELETE http://localhost:8080/delete/user/:id'
# curl -X DELETE http://localhost:8080/delete/user/1qS9OI4YX8daKvHpwvhrUt6PVnG6MLQMemeFirBdqzEjwibcE1y1EZJELvXWi6w7hU9GwHMQ0RgVc3uWEOEJBbwolVD7rqIUgcwN

# curl -X DELETE http://localhost:8080/delete/user/1t5VsIBXpGl4s8C4CAXTsAlIZISYEOlicj14obz3CwFXCCvaRyuhDI10fah5IfdMS3VblW51my8xt6aQvJI3qNg5as0yqoTCvdZd

# echo ' curl -X GET http://localhost:8080/users/list'
 curl -H "Authorization: Bearer $token1" http://localhost:8080/users/list
 curl -H "Authorization: Bearer $token1" http://localhost:8080/user/$Id1

# echo ' curl -X GET http://localhost:8080/user/:id'
# curl -X GET http://localhost:8080/user/1t5VsIBXpGl4s8C4CAXTsAlIZISYEOlicj14obz3CwFXCCvaRyuhDI10fah5IfdMS3VblW51my8xt6aQvJI3qNg5as0yqoTCvdZd

# curl -X GET http://localhost:8080/user/mAdo98L1nvCzdw4CanEmah8PrPqzFNQmO509HrdKCxcLEcluy1zwm9PLvSIOWhZxDFtnM3rLvjK4cVKVpC0BjYbeU6KpDe14Eh0T

# echo ' curl -X PATCH http://localhost:8080/user/:id'
# curl -X PATCH --header "Content-Type: application/json" \
# -d @data/DataSetUpdate1.json \
# http://localhost:8080/user/1qS9OI4YX8daKvHpwvhrUt6PVnG6MLQMemeFirBdqzEjwibcE1y1EZJELvXWi6w7hU9GwHMQ0RgVc3uWEOEJBbwolVD7rqIUgcwN

# curl -X PATCH --header "Content-Type: application/json" \
# -d @data/DataSetUpdate2.json http://localhost:8080/user/1t5VsIBXpGl4s8C4CAXTsAlIZISYEOlicj14obz3CwFXCCvaRyuhDI10fah5IfdMS3VblW51my8xt6aQvJI3qNg5as0yqoTCvdZd