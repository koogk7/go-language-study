#!bin/bash
echo "docker container running...."

go run ./webserver/main.go > log.txt

echo "docker conainner completes!"

vim logginout > lt.txt
