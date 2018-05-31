# sqldata：Database operation  #

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
datas,err := sqlHand.MysqlFetchMap("SELECT * FROM infos where id=?",2)
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
datas,err := sqlHand.MysqlFetchMap("SELECT * FROM infos where age=? limit 3",30)
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

waiting..*[]: 



# Finally  #

It is suggested that `sqlHand` be made into `a single example factory`.


Comments are welcome..

If it's really good, give a star. 

Thanks.
