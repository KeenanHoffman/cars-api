FROM golang:1.14.2-stretch

RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get install -y wget
RUN apt-get install -y jq
RUN apt-get install -y curl
RUN apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8
# Add PostgreSQL's repository. It contains the most recent stable release
#     of PostgreSQL, ``9.3``.
RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ precise-pgdg main" > /etc/apt/sources.list.d/pgdg.list

# Install ``python-software-properties``, ``software-properties-common`` and PostgreSQL 9.3
#  There are some warnings (in red) that show up during the build. You can hide
#  them by prefixing each apt-get statement with DEBIAN_FRONTEND=noninteractive
#RUN apt-get update && apt-get install -y postgresql-9.3 postgresql-client-9.3 postgresql-contrib-9.3
RUN apt-get update && apt-get install -y postgresql


# Note: The official Debian and Ubuntu images automatically ``apt-get clean``
# after each ``apt-get``

# Run the rest of the commands as the ``postgres`` user created by the ``postgres-9.3`` package when it was ``apt-get installed``
USER postgres

## Create a PostgreSQL role named ``docker`` with ``docker`` as the password and
## then create a database `docker` owned by the ``docker`` role.
## Note: here we use ``&&\`` to run commands one after the other - the ``\``
##       allows the RUN command to span multiple lines.
RUN    /etc/init.d/postgresql start &&\
    psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';"

# Adjust PostgreSQL configuration so that remote connections to the
# database are possible.
#RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/9.3/main/pg_hba.conf

# And add ``listen_addresses`` to ``/etc/postgresql/9.3/main/postgresql.conf``
#RUN echo "listen_addresses='*'" >> /etc/postgresql/9.3/main/postgresql.conf

# Expose the PostgreSQL port
#EXPOSE 5432

# Add VOLUMEs to allow backup of config, logs and databases
#VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

# Set the default command to run when starting the container
#CMD ["/usr/lib/postgresql/9.3/bin/postgres", "-D", "/var/lib/postgresql/9.3/main", "-c", "config_file=/etc/postgresql/9.3/main/postgresql.conf"]

RUN service postgresql start
COPY . /go/src/github.com/keenanhoffman/cars-api
WORKDIR /go/src/github.com/keenanhoffman/cars-api

RUN cd /go/src/github.com/keenanhoffman/cars-api/
CMD ./test.sh