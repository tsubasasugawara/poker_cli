FROM mysql

ADD ./docker/mysql/my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]
