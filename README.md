# sqldata：数据库操作  #

这个库主要是解决go读取mysql数据，统一封装了 `sql & mysql` 库的操作，这个库支持addr/driver_name客户端配置

实现查询数据库操作,数据操作过程非常简单，先看表结构：

```go
CREATE TABLE `infos` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `age` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8

```

编写go代码：

```go
condId:=2
datas,err := sqlHand.MysqlFetchMap("SELECT * FROM infos where id=?",condId)
```

datas的结果是一个二维的map,尽量让他长得像php调用mysql返回的样子， 他的结果取值很方便 如下所示:

```go
//上面的datas返回值 可以这样取出来：
datas[0]["id"], datas[0]["age"], datas[0]["name"] 
datas[1]["id"], datas[1]["age"], datas[2]["name"] 
```

接口文档：[sqldata]()
  

# 读取数据.客户端  #

 首先：客户端设置配置文件， 新建`my.conf`文件，添加代码 如下：
 
```go 
[database]
addr="xiaojianhe:123456@tcp(127.0.0.1:3306)/my?charset=utf8"
driver_name="mysql"
    
```


 其次：新建文件 `op.go`，获取某个表里的数据，编写代码如下：

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

datas,err := sqlHand.MysqlFetchMap("SELECT * FROM infos limit 3")
if err!=nil {
    fmt.Printf("get data. [err:%v]", err)
}
fmt.Printf("gat data : %v",datas)


```

 最后：建议 把 sqlHand 做成一个单例的factory 来操作。
 
 # 插入数据  #

批量建议：

```go
lastId, err := sqlHand.PrepareInsert("INSERT INTO `infos` (`name`, `age`) VALUES (?,?),(?,?)",
		"肖2", 30,"肖2", 30)
if err != nil {
    t.Fatalf("get data. [err:%v]", err)
}
```

单记录插入：

```go
lastId, err := sqlHand.Insert("INSERT INTO `infos` (`name`, `age`) VALUES (?,?),(?,?)",
		"肖2", 30,"肖2", 30)
if err != nil {
    t.Fatalf("get data. [err:%v]", err)
}
```

 # 修改数据  #

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