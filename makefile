mysql:
	docker run --name mysqlbank -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql
migrateup:
	migrate  -path db/migration -database "mysql://root:123456@tcp(127.0.0.1:3307)/bank" -verbose up
migratedown:
	migrate  -path db/migration -database "mysql://root:123456@tcp(127.0.0.1:3307)/bank" -verbose down
test:
	go test -v -cover ./...
server:
	go run main.go
.PHONY: mysql migrateup migratedown test server