FROM postgres

ENV POSTGRES_PASSWORD=password

COPY ./db.sql /docker-entrypoint-initdb.d/db.sql
