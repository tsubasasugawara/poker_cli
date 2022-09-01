FROM mysql

ADD ./my.cnf /etc/mysql/conf.d/my.conf

CMD ["mysqld"]
