== Start mySQL via docker

docker run -p 127.0.0.1:3306:3306 --name store-mysql -e MYSQL_ROOT_PASSWORD=password123 -d mysql


====== Database access samples ====
==> https://go.dev/doc/tutorial/database-access
 How to get a new external module ⇒
 go get github.com/go-sql-driver/mysql

Endpoint to return data extracted from mySQL:
http://localhost:8080/albumsByName/name=John


How to use GoFmt:
==> https://go.dev/blog/gofmt