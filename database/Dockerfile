FROM mysql
COPY ./init.sql /tmp
CMD [ "mysqld", "--init-file=/tmp/init.sql" ]