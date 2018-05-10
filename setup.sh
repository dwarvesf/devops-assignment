#!/bin/bash

docker exec -t accident_db sh -c "mysql -utest -ptest accident < /sql/seed.sql"