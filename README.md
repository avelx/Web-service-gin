== Start mySQL via docker

docker run -p 127.0.0.1:3306:3306 --name store-mysql -e MYSQL_ROOT_PASSWORD=password123 -d mysql

How to get a new external module ⇒
go get github.com/go-sql-driver/mysql