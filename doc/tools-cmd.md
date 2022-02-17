### MySql
```
docker run -d --name mysql-master -p 13306:3306 -v ~/data/mysql/conf/master.conf:/etc/mysql/mysql.conf.d/mysqld.cnf -v ~/data/mysql/datam:/var/lib/mysql  -e MYSQL_ROOT_PASSWORD=password mysql:latest

docker run -d --name mysql-master -p 13306:3306 -v /data/mysql/conf/master.conf:/etc/mysql/mysql.conf.d/mysqld.cnf -v /data/mysql/datam:/var/lib/mysql  -e MYSQL_ROOT_PASSWORD=password mysql:latest

docker exec -it mysql-master mysql -u root -p
```

### protobuf
```
docker run -v $(pwd):/go/src/app -w /go/src/app rain/proto-builder:g1.14-p3.12.0 sh -c 'protoc *.proto --go_out=. --micro_out=.;'
```