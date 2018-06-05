# sqldata：Database operation  #

[![Build Status](https://travis-ci.org/dofound/sqldata.svg?branch=master)](https://travis-ci.org/dofound/sqldata) [![Coverage Status](https://coveralls.io/repos/github/dofound/sqldata/badge.svg?branch=master)](https://coveralls.io/github/dofound/sqlx?branch=master) [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/dofound/sqldata) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/dofound/sqldata/master/LICENSE)

This library mainly solves the problem that go reads MySQL data and unify the operation of `SQL & MySQL` library, which supports `addr/driver_name` client configuration.

The lazy external library needs to be installed：
```go
go get github.com/go-sql-driver/mysql
go get github.com/BurntSushi/toml

```

Query database operation, data operation process is very simple, first look at the table structure.：

```go
CREATE TABLE `infos` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `age` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8

```

Write GO code, only one line of code：

```go
datas,err := sqlHand.FetchMapFromSql("SELECT * FROM infos where id=?",2)
```

The result of datas is a two-dimensional map. Try to make him look like the PHP calls Mysql to return. The result is very convenient, as shown below.:

```go
//The value can be taken out of this：
datas[0]["id"], datas[0]["age"], datas[0]["name"] 
datas[1]["id"], datas[1]["age"], datas[2]["name"] 
```

Package document：[sqldata]()
  

# Read the data #

 First, the client set up the configuration file, the new `my.conf` file, and add the code as follows：
 
```go 
[database]
addr="xiaojianhe:123456@tcp(127.0.0.1:3306)/my?charset=utf8"
driver_name="mysql"
    
```


 Second: the new file `op.go`, get a table of data, the code is as follows：

```go

var sysconfig Config
var configPath string
flag.StringVar(&configPath, "config", "my.conf", "server config.")
flag.Parse()

configPath = "my.conf"

if _, err := toml.DecodeFile(configPath, &sysconfig); err != nil {
    t.Fatalf("decode err:%v", err)
}
newSql := NewFactory(&sysconfig)
ctx:=context.Background()
sqlHand := newSql.New(ctx)
```

Read the data

```go
datas,err := sqlHand.FetchMapFromSql("SELECT * FROM infos where age=? limit 3",30)
if err!=nil {
    fmt.Printf("get data. [err:%v]", err)
}
fmt.Printf("gat data : %v",datas)


```
 
 # Insert the data  #

Batch proposal：

```go
lastId, err := sqlHand.PrepareInsert("INSERT INTO `infos` (`name`, `age`) VALUES (?,?),(?,?)",
		"xiaojianhe2", 28,"xiaojianhe3", 30)
if err != nil {
    t.Fatalf("get data. [err:%v]", err)
}
```

Single record insertion：

```go
lastId, err := sqlHand.Insert("INSERT INTO `infos` (`name`, `age`) VALUES (?,?)",
		"xiaojh4", 30)
if err != nil {
    t.Fatalf("get data. [err:%v]", err)
}
```

 # Update the data  #

```go
affect, err := sqlHand.PrepareOpAffected("UPDATE `my`.`infos` SET `name`=? WHERE `id`=?",
		"xiaojh12", 26)
if err != nil {
    t.Fatalf("get error. [err:%v]", err)
}
```


```go
affect, err := sqlHand.OpAffected("UPDATE `my`.`infos` SET `name`=? WHERE `id`=?",
		"xiaojh22", 22)
if err != nil {
    t.Fatalf("get error. [err:%v]", err)
}
```

# Transaction processing data  #

```go
sqlHand.Begin()
stmt,err := sqlHand.TxPrepare("UPDATE `my`.`infos` SET `name`=? WHERE `id`=?")
result,err:=sqlHand.TxExec(stmt,"hahha",2)
//reid,err := result.RowsAffected()
err = sqlHand.Commit()
if err != nil {
    sqlHand.Rollback()
    t.Fatalf("InterfaceCommit error. [err:%v]", err)
}
```

If it is a supporting transaction engine, such as InnoDB, there is a system parameter setting automatically commit.
```go
mysql> show variables like '%autocommit%';
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| autocommit    | ON    |
+---------------+-------+
1 row in set (0.04 sec)
```


# Finally  #

It is suggested that `sqlHand` be made into `a single example factory`.


Comments are welcome..

If it's really good, give a star. 

Thanks.
