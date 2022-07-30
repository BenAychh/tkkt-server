FROM postgres:14.4 as base
ADD migrations/*.up.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_DB tkkt
ENV POSTGRES_USER tkkt
ENV POSTGRES_PASSWORD password