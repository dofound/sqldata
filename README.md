# sqldata：数据库操作  #

这个库主要是解决go读取mysql数据，统一封装了 `sql & mysql` 库的操作，这个库支持addr/driver_name客户端配置

实现查询数据库操作,数据操作过程非常简单，就一句语代码：

```go
condId:=2
datas,err := sqlHand.MysqlFetchMap("SELECT * FROM infos where id=?",condId)
```

datas的结果是一个二维的map,尽量让他长得像php返回的样子， 他的结果是这样的：:

```go
CREATE TABLE `infos` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `age` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8

//上面的datas返回值 可以这样取出来：

datas[0]["id"], datas[0]["age"], datas[0]["name"] 
datas[1]["id"], datas[1]["age"], datas[2]["name"] 
```

接口文档：[sqldata]()
  

# 客户端使用 conf  #

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
 
 