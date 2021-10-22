FROM mysql:latest

ADD ./create_db.sql /docker-entrypoint-initdb.d