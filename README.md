# sqldata：mysql数据库操作  #

这个库主要是解决获取mysql数据，统一封装了 `sql & mysql` 库的操作，这个库支持addr/driver_name客户端配置

实现查询数据库操作,数据操作过程非常简间，就一句语代码：

```go
condId:=2
datas,err := sqlHand.MysqlFetchMap("SELECT * FROM infos where id=?",condId)
```


接口文档：[sqldata]()
  

# 客户端使用 conf  #

 首先：客户端设置配置文件， 在`my.conf`里添加 如下信息：
 
```go 
[database]
addr="xiaojianhe:123456@tcp(127.0.0.1:3306)/my?charset=utf8"
driver_name="mysql"
    
```


 其次：新建文件 `client/email.go` 编写代码如下：

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
		t.Fatalf("get data. [err:%v]", err)
	}
	t.Logf("gat data : %v",datas)
	t.Run("get connect", func(t *testing.T) {
		//fmt.Println("ok")
	})

```

 最后：建议 把db 做成一个单例的factory 来操作。
 
 