FROM postgres:alpine
ENV POSTGRES_USER=root
ENV POSTGRES_PASSWORD=example
ENV POSTGRES_DB=psqldb
COPY ./psql.sql /docker-entrypoint-initdb.d/