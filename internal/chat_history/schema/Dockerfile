FROM postgres:latest
ADD 01_init_up.sql /docker-entrypoint-initdb.d/
RUN chmod a+r /docker-entrypoint-initdb.d/*