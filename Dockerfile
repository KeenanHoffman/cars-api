FROM golang:1.14.2-stretch

RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get install -y wget
RUN apt-get install -y jq
RUN apt-get install -y curl
RUN apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8
RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ precise-pgdg main" > /etc/apt/sources.list.d/pgdg.list
RUN apt-get update && apt-get install -y postgresql
USER postgres

RUN    /etc/init.d/postgresql start &&\
    psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';"

COPY . /go/src/github.com/keenanhoffman/cars-api
WORKDIR /go/src/github.com/keenanhoffman/cars-api

RUN cd /go/src/github.com/keenanhoffman/cars-api/
