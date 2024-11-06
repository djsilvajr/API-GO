# Usa a imagem base do MySQL 5.7
FROM mysql:5.7

# Configurações do MySQL
ENV MYSQL_ROOT_PASSWORD=123456
ENV MYSQL_DATABASE=db
ENV MYSQL_USER=djsilvajr
ENV MYSQL_PASSWORD=123456


# Copia um arquivo SQL para configurar permissões
COPY ./init-db.sql /docker-entrypoint-initdb.d/

# Desabilita uso de AIO nativo, caso seja necessário
CMD ["mysqld", "--innodb-use-native-aio=0"]
