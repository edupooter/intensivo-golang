# intensivo-golang

Implementation for this 2 hours long tutorial, using Go, Kafka, MySQL and Docker:

[Intensivão Go Lang: A linguagem queridinha das grandes empresas](https://www.youtube.com/watch?v=Bq2XTGzV8yc)

## Requirements
Docker Desktop / Docker Engine

## How to use it

Build Docker containers on host:
`$ docker-compose up -d`

Setup the table on MySQL:
```
$ docker-compose exec mysql bash
root@20d30b06338a:/# mysql -u root -p
Enter password: 
[...]
mysql> use golang;
mysql> create table courses (id varchar (255), name varchar (255), description varchar (255), state varchar (255));
```

Create Kafka topics and producer inside container:
```
$ docker-compose exec kafka bash
[appuser@aaaaa ~]$ kafka-topics --bootstrap-server=localhost:9092 --create --topic courses --partitions=3 --replication-factor=1
[appuser@aaaaa ~]$ kafka-console-producer --bootstrap-server=localhost:9092 --topic=courses
```

The Kafka input CLI will be available like this: `>`

Run the application. It will be listening to new messages at kafka instance:
```
$ docker-compose exec app bash
root@424242:/go/src# go run cmd/main.go
```

*Note: MySQL dependency will be downloaded at the first time, but the console output will be ready after that.*

Input a valid JSON as expected on Kafka bash:
`{"name":"Golang","description":"Go, Kafka and MySQL","status":"Finished"}`

An output will be shown in the app console to confirm it worked:
```
root@4588e8452047:/go/src# go run cmd/main.go
{63aec930-a862-4775-bc08-1d477fe6fee2 Golang Go, Kafka and MySQL Finished}
```

Finally you can confirm that the message was persisted on the database:
```
mysql> select * from courses;
+--------------------------------------+------------------+-------------------------+----------+
| id                                   | name             | description             | status   |
+--------------------------------------+------------------+-------------------------+----------+
| 63aec930-a862-4775-bc08-1d477fe6fee2 | Intensivo Golang | Go, Kafka and MySQL     | Finished |
+--------------------------------------+------------------+-------------------------+----------+
1 row in set (0.00 sec)
```

### Optional
Building the binary file for Linux:
`root@424242:/go/src# go build cmd/main.go`
