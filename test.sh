#!/bin/bash


# Check OS to print with colors
if [[ "$OSTYPE" == "darwin"* ]]; then
red="\033[31m"
blue="\033[34m"
end="\033[0m"
bold="\033[1m"
else
red="\e[31m"
blue="\e[34m"
end="\e[0m"
bold="\e[1m"
fi

# Variables
DataSetPath="./data/DataSet.json"

Id1="1qS9OI4YX8daKvHpwvhrUt6PVnG6MLQMemeFirBdqzEjwibcE1y1EZJELvXWi6w7hU9GwHMQ0RgVc3uWEOEJBbwolVD7rqIUgcwN"
Password1="xZinMPI3moup6ymz3gVn"
Id2="1t5VsIBXpGl4s8C4CAXTsAlIZISYEOlicj14obz3CwFXCCvaRyuhDI10fah5IfdMS3VblW51my8xt6aQvJI3qNg5as0yqoTCvdZd"
Password2="CGUsfQkS06lo7LM2guSV"
Id3="buhnVfyzGfe4qfuRsqwjGhzoF8ubXcM4JXqmLRmvPmJdLXtITBY99SI0rER0bX0INj2XkAyMkLUEkQrxNG46PDXHt00jGURsnjHy"
Password3="o6sRZkUSroEP4o51npXb"

# check if Docker is installed
if [ -x "$(command -v docker)" ]; then
    echo "Docker is installed"
else
    echo "Install docker"
fi


# ! Inital post that fill the database with data set
echo -e "$red $bold TEST 1$end"
echo '--------------------------------------------------------------------'
echo -e "$blue $bold Inital post that fill the database with data set $end"
echo '--------------------------------------------------------------------'
curl -X POST --header "Content-Type: application/json" -d @$DataSetPath  http://localhost:8080/add/users
echo 'Refer to API log to see the results'
echo ''
echo ''

sleep 5

# ! Login
echo -e "$red $bold TEST 2$end"
echo '--------------------------------------------------------------------'
echo -e "$blue $bold Login with Id1 : $Id1 $end"
echo '--------------------------------------------------------------------'
token1=$( curl -s -X POST -H "Content-Type: application/json" -d "{\"id\":\"$Id1\",\"password\": \"$Password1\"}" http://localhost:8080/login | sed -e "s/{\"access_token\":\"//" | sed -e "s/\",\"status\":\"success\"}//")
token1=$(echo -n "$token1")
echo "JwT for Id1 is $token1\tr"
echo ''
sleep 1
token2=$( curl -s -X POST -H "Content-Type: application/json" -d "{\"id\":\"$Id2\",\"password\": \"$Password2\"}" http://localhost:8080/login | sed -e "s/{\"access_token\":\"//" | sed -e "s/\",\"status\":\"success\"}//")
token2=$(echo -n "$token2")
echo "JwT token for Id2 is $token2\tr"
echo ''
sleep 1
token3=$( curl -s -X POST -H "Content-Type: application/json" -d "{\"id\":\"$Id3\",\"password\": \"$Password3\"}" http://localhost:8080/login | sed -e "s/{\"access_token\":\"//" | sed -e "s/\",\"status\":\"success\"}//")
token3=$(echo -n "$token3")
echo "JwT token for Id2 is $token3\tr"
echo ''
sleep 1

# ! Get userList and get single User
echo -e "$red $bold TEST 3$end"
echo '--------------------------------------------------------------------'
echo -e "$blue $bold Get userList and single user $end"
echo '--------------------------------------------------------------------'
echo 'result for userList:'
curl -H "Authorization: Bearer $token1\tr" http://localhost:8080/users/list
echo ''
sleep 1
echo 'result for Id1:'
curl -H "Authorization: Bearer $token1\tr" http://localhost:8080/user/$Id1
echo ''
sleep 1
echo 'result for Id2:'
curl -H "Authorization: Bearer $token2\tr" http://localhost:8080/user/$Id2
echo ''
sleep 1
echo 'result for Id3:'
curl -H "Authorization: Bearer $token3\tr" http://localhost:8080/user/$Id3
echo ''

# # ! Update in dataBase
echo -e "$red $bold TEST 4$end"
echo '--------------------------------------------------------------------'
echo -e "$blue $bold Update user in Database and modify file if data has changed $end"
echo '--------------------------------------------------------------------'
echo 'result for update1:'
curl -X PATCH --header "Content-Type: application/json" \
-H "Authorization: Bearer $token1\tr" \
-d @data/DataSetUpdate1.json \
http://localhost:8080/user/$Id1
echo ''
sleep 1
echo 'result for update2:'
# echo ' curl -X PATCH http://localhost:8080/user/:id'
curl -X PATCH --header "Content-Type: application/json" \
-H "Authorization: Bearer $token2\tr" \
-d @data/DataSetUpdate2.json \
http://localhost:8080/user/$Id2
sleep 1
echo ''

# # ! Delete in dataBase
echo -e "$red $bold TEST 5$end"
echo '--------------------------------------------------------------------'
echo -e "$blue $bold Update user in Database and modify file if data has changed $end"
echo '--------------------------------------------------------------------'
echo 'result for delete Id1'
curl -X DELETE \
-H "Authorization: Bearer $token1\tr" \
http://localhost:8080/delete/user/$Id1
echo ''
sleep 1
echo 'result for delete Id2'
# echo ' curl -X PATCH http://localhost:8080/user/:id'
curl -X DELETE \
-H "Authorization: Bearer $token2\tr" \
http://localhost:8080/delete/user/$Id2
sleep 1
echo '--------------------------------------------------------------------'
echo -e "$red $bold END OF THE TESTS $end"
echo '--------------------------------------------------------------------'