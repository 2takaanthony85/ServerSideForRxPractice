FROM mysql

## EXPOSE 3306

COPY ./my.cnf /etc/mysql/conf.d/
COPY ./mysql/init.sql /docker-entrypoint-initdb.d

CMD ["mysqld"]