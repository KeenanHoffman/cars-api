#!/usr/bin/env bash
set -eu

export SERVER_DB_NAME=cars_test
export SERVER_DB_URL=localhost
export SERVER_DB_PORT=5432
export SERVER_URL=localhost
export SERVER_PORT=8080
export CLIENT_PORT=4040
tmp_dir=$(mktemp -d)
cleanup() {
  rm -rf $tmp_dir
  kill -9 $server_pid
  kill -9 $client_pid
  psql $SERVER_DB_NAME <<EOF
DROP TABLE cars;
EOF
  psql postgres <<EOF
DROP DATABASE $SERVER_DB_NAME;
EOF
}
trap "cleanup" QUIT TERM EXIT INT

psql postgres <<EOF
CREATE DATABASE $SERVER_DB_NAME;
EOF
psql $SERVER_DB_NAME -f cars-test.sql

pushd server > /dev/null
go build -o $tmp_dir > /dev/null
$tmp_dir/server &
server_pid=$!
popd > /dev/null

pushd client > /dev/null
go build -o $tmp_dir > /dev/null
$tmp_dir/client &
client_pid=$!
popd > /dev/null

sleep 3

echo "======== Create/Get By ID Test ========"
curl -X POST localhost:4040/cars -H "Content-type:application/json" -d '{"make":"Dodge","model":"charger","vin":"007"}'
echo
if [[ $(curl localhost:4040/cars/1 -H "Accept:application/json" | jq -r '.make') == "Dodge" ]]; then
  echo "Create/Get By ID Test Passed"
else
  echo "Create/Get By ID Test Failed"
  exit 0
fi

echo "======== Get All Test  ========"
if [[ $(curl localhost:4040/cars -H "Accept:application/json" | jq -r '.[0].make') == "Dodge" ]]; then
  echo "Get All Test Passed"
else
  echo "Get All Test Failed"
  exit 0
fi

echo "======== Update Test ========"
curl -X PATCH localhost:4040/cars/1 -H "Content-type:application/json" -d '{"vin":"64"}'
echo
if [[ $(curl localhost:4040/cars/1 -H "Accept:application/json" | jq -r '.vin') == "64" ]]; then
  echo "Update Test Passed"
else
  echo "Update Test Failed"
  exit 0
fi

echo "======== Replace Test ========"
curl -X PUT localhost:4040/cars/1 -H "Content-type:application/json" -d '{"make":"Chevy","model":"terrain","vin":"123"}'
echo
if [[ $(curl localhost:4040/cars/1 -H "Accept:application/json") == '{"id":1,"make":"Chevy","model":"terrain","vin":"123"}' ]]; then
  echo "Replace Test Passed"
else
  echo "Replace Test Failed"
  exit 0
fi

echo "======== Delete Test ========"
curl -X DELETE localhost:4040/cars/1
echo
if [[ $(curl localhost:4040/cars -H "Accept:application/json") == "null" ]]; then
  echo "Delete Test Passed"
else
  echo "Delete Test Failed"
  exit 0
fi

echo "======== Tests Passed ========"
