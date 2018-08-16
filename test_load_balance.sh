#!/bin/bash

for n in $(seq 1 1 10)
do
    nohup curl -XGET -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzExMjUxMjUsImlkIjo5LCJuYmYiOjE1MzExMjUxMjUsInVzZXJuYW1lIjoiYWRtaW4ifQ.PZpiRQ_HAXfFtyKqOfnvrHTNvYgQ2tKJKlX19WA7KIY" http://goapi.com/v1/user &>/dev/null
done
