Start docker mysql container: `docker-compose up`

First time setup test DB: `go run setup/main.go`

Run database script: `go run main.go`

To allow phpmyadmin to connect to the latest MySQL version, run:
```shell
docker exec -it database_db_1 bash
bash> mysql -u root -pasdf1234
mysql> ALTER USER root IDENTIFIED WITH mysql_native_password BY 'asdf1234';
```