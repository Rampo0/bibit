#!/bin/bash

setup_postgres()
{
    echo "SETUP POSTGRES"
    docker stop $(docker ps -aq)
    docker rm $(docker ps -aq)
    docker run --name postgresdb --mount type=bind,source="$(pwd)"/init-scripts,target=/docker-entrypoint-initdb.d -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
}

main()
{
    setup_postgres
}

main