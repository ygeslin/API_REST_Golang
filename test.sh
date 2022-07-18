#!/bin/zsh

#Inital post that fill the database with data set
echo 'Inital post that fill the database with data set'
curl -X POST --header "Content-Type: application/json" -d @data/DataSet.json  http://localhost:8080/add/users

echo ' curl -X POST http://localhost:8080/add/users' "\n"
curl -X POST http://localhost:8080/add/users

echo ' curl -X POST http://localhost:8080/login'
curl -X POST http://localhost:8080/login

echo ' curl -X DELETE http://localhost:8080/delete/user/:id'
curl -X DELETE http://localhost:8080/delete/user/1qS9OI4YX8daKvHpwvhrUt6PVnG6MLQMemeFirBdqzEjwibcE1y1EZJELvXWi6w7hU9GwHMQ0RgVc3uWEOEJBbwolVD7rqIUgcwN

curl -X DELETE http://localhost:8080/delete/user/1t5VsIBXpGl4s8C4CAXTsAlIZISYEOlicj14obz3CwFXCCvaRyuhDI10fah5IfdMS3VblW51my8xt6aQvJI3qNg5as0yqoTCvdZd

echo ' curl -X GET http://localhost:8080/users/list'
curl -X GET http://localhost:8080/users/list

echo ' curl -X GET http://localhost:8080/user/:id'
curl -X GET http://localhost:8080/user/1t5VsIBXpGl4s8C4CAXTsAlIZISYEOlicj14obz3CwFXCCvaRyuhDI10fah5IfdMS3VblW51my8xt6aQvJI3qNg5as0yqoTCvdZd

curl -X GET http://localhost:8080/user/mAdo98L1nvCzdw4CanEmah8PrPqzFNQmO509HrdKCxcLEcluy1zwm9PLvSIOWhZxDFtnM3rLvjK4cVKVpC0BjYbeU6KpDe14Eh0T

echo ' curl -X PATCH http://localhost:8080/user/:id'
curl -X PATCH --header "Content-Type: application/json" \
-d @data/DataSetUpdate1.json \
http://localhost:8080/user/1qS9OI4YX8daKvHpwvhrUt6PVnG6MLQMemeFirBdqzEjwibcE1y1EZJELvXWi6w7hU9GwHMQ0RgVc3uWEOEJBbwolVD7rqIUgcwN

curl -X PATCH --header "Content-Type: application/json" \
-d @data/DataSetUpdate2.json http://localhost:8080/user/1t5VsIBXpGl4s8C4CAXTsAlIZISYEOlicj14obz3CwFXCCvaRyuhDI10fah5IfdMS3VblW51my8xt6aQvJI3qNg5as0yqoTCvdZd