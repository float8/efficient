# efficient for web framework

_efficient是基于Golang gin框架编写的web、command富框架。灵活的应用目录结构，可根据需要自定义目录结构。_

_efficient 主要基于下面的包进行了封装, 基本保留了原有包的用法_

[gin github 地址点击这里](https://github.com/gin-gonic/gin)

# 目录

- [安装](#安装)
- [快速开始](#快速开始)
  - [WEB站点](#WEB站点)
  - [代码生成器](#代码生成器)
- [系统配置](#系统配置)
- [路由规则](#路由规则)
- [控制器](#控制器)
- [数据层](#数据层)
  - [数据库连接](#数据库连接)
  - [数据层设计](#数据层设计)
  - [Model](#Model)
  - [Dao](#Dao)
  - [应用](#应用)
- [数据验证](#数据验证)
- [日志](#日志)
- [环境变量](#环境变量)
- [生成器](#生成器)
- [插件注册](#插件注册)
  - [注册数据类型](#注册数据类型)
  - [注册验证器](#注册验证器)
  - [注册代码生成器](#注册代码生成器)
  - [注册SQL生成器](#注册SQL生成器)
- [工具包](#工具包)

# 安装

#### 1. 首先需要安装 [Go](https://golang.org/) (**version 1.10+**), 可以使用下面的命令进行安装 efficient.

```sh
$ go get github.com/whf-sky/efficient
```

#### 2. 导入你的代码

```go
import "github.com/whf-sky/efficient"
```

如使用go mod包依赖管理工具,请参考下面命令

#### Windows 下开启 GO111MODULE 的命令为：
```sh
$ set GO111MODULE=on
```

#### MacOS 或者 Linux 下开启 GO111MODULE 的命令为：
```sh
$ export GO111MODULE=on
```

#### Windows 下设置 GOPROXY 的命令为：
```sh
$ go env -w GOPROXY=https://goproxy.cn,direct
```

#### MacOS 或 Linux 下设置 GOPROXY 的命令为：
```sh
$ export GOPROXY=https://goproxy.cn
```



# 快速开始

## WEB站点

```sh
$ cat main.go
```

```go
package main

import (
  "github.com/whf-sky/efficient"
  "net/http"
)

type TestController struct {
  efficient.Controller
}

func (this *TestController) Get(cxt efficient.Context) {
  id := cxt.Query("id")
  cxt.String(200, "get:"+id)
}

func (this *TestController) Post(cxt efficient.Context) {
  id := cxt.PostForm("id")
  cxt.String(200, "post:"+id)
}

func main(){
  efficient.Routers.Add("/test", &TestController{}, http.MethodGet, http.MethodPost)
  efficient.Run()
}

```

## 代码生成器

```sh
$ cat generate.go
```

```go
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/whf-sky/efficient/database"
	"github.com/whf-sky/efficient/generate"
)

func main() {
	//实例化代码生成器
	generate := generate.NewGenerate()
	//应用生成
	generate.Application()
	//设置数据库并生成dao和model文件
	generate.SetDb("mysql", database.NewDb().
		MysqlDsn("127.0.0.1", "3306", "root", "123456", "test", "utf8mb4").
		Open()).Database("test")
}
```

# 系统配置

#### 设置web服务的端口地址

```go
efficient.Config.Addr = ":80"
```

#### 设置Debug模式

```go
efficient.Config.Debug = true
```

#### 设置APP名称

```go
efficient.Config.AppName = "efficient"
```

#### 设置中间件

```go
efficient.Config.Middleware = []Middleware{
	func(ctx *Context){},
}
```

#### 设置语言包

```go
//参考框架 widget/lang
efficient.Config.Lang = Lang{
    "validation_no_lang":"不存在此语言KEY'%s'，请检查是否拼写错误！",
}
```

# 路由规则

`参数1 地址目录的相对目录`

`参数2 执行的实例化的控制器，控制器必须嵌入efficient.Controller`

`参数3... 请求方式，请参考http包，可以添加多个，Restful风格`

```go
efficient.Routers.
	Add("/test", &TestController{}, http.MethodGet, http.MethodPost).
	Add("/demo", &DemoController{}, http.MethodGet, http.MethodPut)
```

# 控制器

控制器必须内嵌`efficient.Controller`, 控制器的方法大小写请与http包中的方法名称保持一致，首字母大写即可。

`efficient.Context` 用法与gin的`*Context`用法一致。

[gin github 地址点击这里](https://github.com/gin-gonic/gin)

#### 示例代码

```go
package main

import (
  "github.com/whf-sky/efficient"
)

type TestController struct {
  efficient.Controller
}

func (this *TestController) Get(cxt efficient.Context) {
  id := cxt.Query("id")
  cxt.String(200, "get:"+id)
}

func (this *TestController) Post(cxt efficient.Context) {
  id := cxt.PostForm("id")
  cxt.String(200, "post:"+id)
}
```

# 数据层

## 数据库连接

#### 设置数据库连接池

```go
database.NewDb().MysqlDsn("地址", "端口", "账户", "密码", "数据库", "数据库编码").Open(func(db *sql.DB) {
    db.SetConnMaxIdleTime(time.Minute * 4) //设置最大连接生存时间
    db.SetMaxOpenConns(5)                  //设置最大打开链接数
    db.SetMaxIdleConns(5)                  //设置最大空闲链接
    db.SetConnMaxLifetime(time.Minute * 2) //设置连接最大空闲时间
})
```

## 数据层设计

```
数据层采用了model和dao的模式，model和dao可以通过代码生成器进行生成，建议直接采用生成器生成，在此基础上进行增加内容。

model对应的数据表结构。

dao层与model相对，对数据表进行操作。
```

## Model

### 1.定义Model

#### 示例代码

```go
package model

import (
	"encoding/json"
	"github.com/whf-sky/efficient/database"
)

//实例化model
func NewUsers() *Users {
	m := &Users{}
	//mysql 数据库驱动
	//test 数据库
	//设置model的唯一标识，即当前表名即可
    m.Model.Init("mysql","test", "users")
	return m
}
//必须内嵌database.Model
//结构体内对应的是表字段和字段信息，最终会把mysql字段类型转化为golang数据类型和自定义的数据类型
type Users struct {
	database.Model
	Id         uint32          `column:"id" type:"int" auto_increment:"true" primary_key:"true" unsigned:"true" precision:"10" scale:"0" comment:"用户编号"`
	Username   string          `column:"username" type:"varchar" unique:"true" length:"100" octet_length:"400" comment:"账户"`
	Passwd     string          `column:"passwd" type:"varchar" length:"100" octet_length:"400" comment:"密码"`
}

//TableName 数据库表明方法
func (u *Users) TableName() string {
	return "users"
}

//Ptrs 所有数据表字段对应映射的接口体字段的指针
func (u *Users) Ptrs() map[string]interface{} {
	return map[string]interface{}{
		"id":          &u.Id,
		"username":    &u.Username,
		"passwd":      &u.Passwd,
	}
}

//Get 根据数据表字段进行获取model字段值
func (u *Users) Get(key string) interface{} {
	switch key {
	case "id":
		return u.Id
	case "username":
		return u.Username
	case "passwd":
		return u.Passwd
	}
	return nil
}

//Set 根据数据表字段进行设置model字段值
func (u *Users) Set(key string, val interface{}) {
	switch key {
	case "id":
		u.Id = val.(uint32)
	case "username":
		u.Username = val.(string)
	case "passwd":
		u.Passwd = val.(string)
	}
}

//以下是model所有字段的get方法

func (u *Users) GetId() uint32 {
	return u.Id
}

func (u *Users) GetUsername() string {
	return u.Username
}

func (u *Users) GetPasswd() string {
	return u.Passwd
}

//以下是model所有字段的set方法

func (u *Users) SetId(value uint32) *Users {
	//添加赋值的字段，因为golang 结构体不能区分是否已经赋值所以使用u.AddAssColumns进行判断是否赋值
	u.AddAssColumns("id")
	u.Id = value
	return u
}

func (u *Users) SetUsername(value string) *Users {
	u.AddAssColumns("username")
	u.Username = value
	return u
}

func (u *Users) SetPasswd(value string) *Users {
	u.AddAssColumns("passwd")
	u.Passwd = value
	return u
}

//InsertEvent Insert时自动触发此事件，在Insert执行之前触发
func (u *Users) InsertEvent() {}


//UpdateEvent Update时自动触发此事件，在Update执行之前触发
func (u *Users) UpdateEvent() {}

//把model格式化成json字符串
func (u *Users) ToString() string {
	bytes, _ := json.Marshal(u)
	return string(bytes)
}
```

#### Model struct 

`必须内嵌 database.Model`

```go
type Users struct {
	database.Model
	Id         uint32          `column:"id" type:"int" auto_increment:"true" primary_key:"true" unsigned:"true" precision:"10" scale:"0" comment:"用户编号"`
	Username   string          `column:"username" type:"varchar" unique:"true" length:"100" octet_length:"400" comment:"账户"`
	Passwd     string          `column:"passwd" type:"varchar" length:"100" octet_length:"400" comment:"密码"`
}
```

#### New 

`实例化Model`

```go
func NewUsers() *Users {
  m := &Users{}
  m.init("users")  //设置model的唯一标识，即当前表名即可
  return m
}
```

#### TableName  

`表名的方法；使用方法可以根据字段信息执行自定义分表`

```go
func (u *Users) TableName() string {
	return "表名"
}
```

#### Ptrs 

`所有数据表字段对应Model字段的指针；主要用于dao的封装中`

```go
func (u *Users) Ptrs() map[string]interface{} {
	return map[string]interface{}{
		"id":          &u.Id,
		"username":    &u.Username,
		"passwd":      &u.Passwd,
	}
}
```

#### Get 

`根据数据表字段获取Model字段的值`

```go
func (u *Users) Get(key string) interface{} {
  switch key {
  case "id":
    return u.Id
  case "username":
    return u.Username
  case "passwd":
    return u.Passwd
  }
  return nil
}
```

#### Set 

`根据数据表字段设置Model字段的值`

```go
func (u *Users) Set(key string, val interface{}) {
  switch key {
  case "id":
    u.Id = val.(uint32)
  case "username":
    u.Username = val.(string)
  case "passwd":
  	u.Passwd = val.(string)
  }
}
```

#### InsertEvent 

`Insert时自动触发此事件，在Insert执行之前触发`

```go
func (u *Users) InsertEvent() {}
```

#### UpdateEvent 

`Update时自动触发此事件，在Update执行之前触发`

```go
func (u *Users) UpdateEvent() {}
```

#### ToString 

`把model格式化成json字符串`

```go
func (u *Users) ToString() string {
  bytes, _ := json.Marshal(u)
  return string(bytes)
}
```

#### ToString 

`把model格式化成json字符串`

```go
func (u *Users) ToString() string {
  bytes, _ := json.Marshal(u)
  return string(bytes)
}
```

### 2.标签注解

#### 标签

```
column 表字段，字符串

type 表类型，字符串

enum 枚举类型的选项值，值的格式：{'北京','上海','重庆'}"

auto_increment 是否自增，字符串的布尔型，true：代表自增

primary_key 是否是主键，字符串的布尔型，true：代表主键

unique 是否是唯一索引，字符串的布尔型，true：代表是唯一索引

unsigned 是否有符号，字符串的布尔型，true：没有符号

null 字段是否为空，字符串的布尔型，true：空

precision 数字精度

scale 小数位精度

length 字符串长度

octet_length 字符串字节长度

default 默认值

comment 字段注释，可以进行修改作为验证器的提示文字

on_insert_time insert时自动添加时间

on_update_time update时自动添加时间

validators 验证器，验证器请参考 widget/validation
```

#### validators 示例 

`数据类型请保持与字段的原生即golang原始类型一致`

```
validators="gt,lt" v-gt="int:5" v-lt="int:20"
```

### 3.Model字段的数据类型

#### 字段非NULL类型 

`字段默认不为空的情况下mysql类型对应的golang类型，包含有符号和无符号类型`

```go
var types = map[string]goType{
	"time": {
		name: "string",
	},
	"date": {
		name: "database.Date",
		pkg:  "github.com/whf-sky/efficient/database",
	},
	"datetime": {
		name: "database.Time",
		pkg:  "github.com/whf-sky/efficient/database",
	},
	"timestamp": {
		name: "database.Time",
		pkg:  "github.com/whf-sky/efficient/database",
	},
	"year": {
		name: "string",
	},
	"decimal": {
		name: "decimal.Decimal",
		pkg:  "github.com/shopspring/decimal",
	},
	"numeric": {
		name: "decimal.Decimal",
		pkg:  "github.com/shopspring/decimal",
	},
	"bit": {
		name: "database.Bit",
		pkg:  "github.com/whf-sky/efficient/database",
	},
	"bits": {
		name: "database.Bits",
		pkg:  "github.com/whf-sky/efficient/database",
	},
	"char": {
		name: "string",
	},
	"varchar": {
		name: "string",
	},
	"tinytext": {
		name: "string",
	},
	"text": {
		name: "string",
	},
	"mediumtext": {
		name: "string",
	},
	"longtext": {
		name: "string",
	},
	"enum": {
		name: "string",
	},
	"set": {
		name: "database.Set",
		pkg:  "github.com/whf-sky/efficient/database",
	},
	"tinyint": {
		name:     "int8",
		unsigned: goUtype{name: "uint8"},
	},
	"smallint": {
		name:     "int16",
		unsigned: goUtype{name: "uint16"},
	},
	"mediumint": {
		name:     "int32",
		unsigned: goUtype{name: "uint32"},
	},
	"int": {
		name:     "int32",
		unsigned: goUtype{name: "uint32"},
	},
	"integer": {
		name:     "int32",
		unsigned: goUtype{name: "uint32"},
	},
	"bigint": {
		name:     "int64",
		unsigned: goUtype{name: "uint64"},
	},
	"double": {
		name: "float64",
	},
	"float": {
		name: "float64",
	},
	"real": {
		name: "float64",
	},
}
```

#### 字段NULL类型 

`字段默认为空的情况下mysql类型对应的golang类型，包含有符号和无符号类型`

```go
var nullTypes = map[string]goType{
  "date": {
    name: "database.NullDate",
    pkg:  "github.com/whf-sky/efficient/database",
  },
  "datetime": {
    name: "database.NullTime",
    pkg:  "github.com/whf-sky/efficient/database",
  },
  "timestamp": {
    name: "database.Time",
    pkg:  "github.com/whf-sky/efficient/database",
  },
  "time": {
    name: "sql.NullString",
    pkg:  "database/sql",
  },
  "year": {
    name: "sql.NullString",
    pkg:  "database/sql",
  },
  "bit": {
    name: "database.NullBit",
    pkg:  "github.com/whf-sky/efficient/database",
  },
  "bits": {
    name: "database.NullBits",
    pkg:  "github.com/whf-sky/efficient/database",
  },
  "char": {
    name: "sql.NullString",
    pkg:  "database/sql",
  },
  "varchar": {
    name: "sql.NullString",
    pkg:  "database/sql",
  },
  "tinytext": {
    name: "sql.NullString",
    pkg:  "database/sql",
  },
  "text": {
    name: "sql.NullString",
    pkg:  "database/sql",
  },
  "mediumtext": {
    name: "sql.NullString",
    pkg:  "database/sql",
  },
  "longtext": {
    name: "sql.NullString",
    pkg:  "database/sql",
  },
  "enum": {
    name: "sql.NullString",
    pkg:  "database/sql",
  },
  "set": {
    name: "database.NullSet",
    pkg:  "github.com/whf-sky/efficient/database",
  },
  "double": {
    name:     "sql.NullFloat64",
    pkg:      "database/sql",
  },
  "float": {
    name: "sql.NullFloat64",
    pkg:  "database/sql",
  },
  "real": {
    name: "sql.NullFloat64",
    pkg:  "database/sql",
  },
  "decimal": {
    name: "decimal.Decimal",
    pkg:  "github.com/shopspring/decimal",
  },
  "numeric": {
    name: "decimal.Decimal",
    pkg:  "github.com/shopspring/decimal",
  },
  "tinyint": {
    name:     "database.NullInt8",
    pkg:      "github.com/whf-sky/efficient/database",
    unsigned: goUtype{
      name:     "database.NullUint8",
      pkg:      "github.com/whf-sky/efficient/database",
    },
  },
  "smallint": {
    name:     "sql.NullInt16",
    pkg:      "database/sql",
    unsigned: goUtype{
      name:     "database.NullUint16",
      pkg:      "github.com/whf-sky/efficient/database",
    },
  },
  "mediumint": {
    name:     "sql.NullInt32",
    pkg:      "database/sql",
    unsigned: goUtype{
      name:     "database.NullUint32",
      pkg:      "github.com/whf-sky/efficient/database",
    },
  },
  "int": {
    name: "sql.NullInt32",
    pkg:  "database/sql",
    unsigned: goUtype{
      name:     "database.NullUint32",
      pkg:      "github.com/whf-sky/efficient/database",
    },
  },
  "integer": {
    name:     "sql.NullInt32",
    pkg:      "database/sql",
    unsigned: goUtype{
      name:     "database.NullUint32",
      pkg:      "github.com/whf-sky/efficient/database",
    },
  },
  "bigint": {
    name: "sql.NullInt64",
    pkg:  "database/sql",
    unsigned: goUtype{
      name:     "database.NullUint64",
      pkg:      "github.com/whf-sky/efficient/database",
    },
  },
  }
```

### 4.注册Model

`因为框架初始化时会解析结构体标签所以必须对Model进行注册`

#### Register

```go
database.RegisterModel(NewUsers())
```

#### 父Model 

`在应用时可设置父Model，表Model可以直接嵌入父Model`

```go
package model

import "github.com/whf-sky/efficient/database"

type Model struct {
  database.Model
}

func (m *Model) init(key string) *Model {
  m.Model.Init("mysql","test", key)
  return m
}

func init() {
  database.RegisterModel(NewUsers())
}


func NewUsers() *Users {
  m := &Users{}
  m.init("users")
  return m
}

type Users struct {
  Model
  Id         uint32          `column:"id" type:"int" auto_increment:"true" primary_key:"true" unsigned:"true" precision:"10" scale:"0" comment:"用户编号"`
  Username   string          `column:"username" type:"varchar" unique:"true" length:"100" octet_length:"400" comment:"账户"`
  Passwd     string          `column:"passwd" type:"varchar" length:"100" octet_length:"400" comment:"密码"`
}

```

## Dao

### 1.定义Dao 

#### 示例代码

```go
package dao

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/whf-sky/efficient.demo/application/model"
  "github.com/whf-sky/efficient/database"
  "time"
)

//请重新配置数据连接信息
var db = func() *sql.DB {
  return database.NewDb().MysqlDsn("127.0.0.1", "3306", "root", "123456", "test", "utf8mb4").Open(func(db *sql.DB) {
    db.SetConnMaxIdleTime(time.Minute * 4) //设置最大连接生存时间
    db.SetMaxOpenConns(5)                  //设置最大打开链接数
    db.SetMaxIdleConns(5)                  //设置最大空闲链接
    db.SetConnMaxLifetime(time.Minute * 2) //设置连接最大空闲时间
  })
}()

func NewUsersDao(db *sql.DB) *UsersDao {
  d := &UsersDao{}
  d.SetDb("mysql", db)
  d.SetModel(func() database.ModelInterface {
    return model.NewUsers()
  })
  return d
}

type UsersDao struct {
  database.Dao
}

```

#### Dao struct 

`必须内嵌 database.Dao `

```go
type UsersDao struct {
  database.Dao
}

```

#### New

` d.SetDb(db) 设置数据库连接，数据类型：*sql.DB`

` d.DriverName("mysql") 设置数据库驱动`

` d.SetModel(func() database.ModelInterface { return model.NewUsers()}) 设置Dao使用的模型`

```go
func NewUsersDao(db *sql.DB) *UsersDao {
  d := &UsersDao{}
  d.SetDb("mysql", db)
  d.SetModel(func() database.ModelInterface {
    return model.NewUsers()
  })
  return d
}
```

### 2.Dao应用

#### 父Dao 

`应用时可以设置父Dao`

```go
package dao

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/whf-sky/efficient.demo/config"
  "github.com/whf-sky/efficient/database"
)

var db = func() *sql.DB {
  return database.NewDb().MysqlDsn(
    config.DbConfig.Addr,
    config.DbConfig.Port,
    config.DbConfig.Account,
    config.DbConfig.Passwd,
    config.DbConfig.Dbname,
    config.DbConfig.Charset,
  ).Open(func(db *sql.DB) {
    db.SetConnMaxLifetime(config.DbConfig.SetConnMaxLifetime) //设置最大连接生存时间
    db.SetConnMaxIdleTime(config.DbConfig.ConnMaxIdleTime) //设置连接最大空闲时间
    db.SetMaxOpenConns(config.DbConfig.MaxOpenConns) //设置最大打开链接数
    db.SetMaxIdleConns(config.DbConfig.MaxIdleConns) //设置最大空闲链接
  })
}()

type Dao struct {
  database.Dao
}

func (d *Dao) Init(fun func() database.ModelInterface) *Dao {
  d.SetDb("mysql", db)
  d.SetModel(fun)
  return d
}
```

#### 父Dao的嵌入使用

```go
package dao

import (
	"efficient.demo/application/model"
	"github.com/whf-sky/efficient/database"
)

func NewUsersDao() *UsersDao {
	d := &UsersDao{}
	d.Init(func() database.ModelInterface {return model.NewUsers()})
	return d
}

type UsersDao struct {
	Dao
}

```

## 应用

### Insert

```go
func (d *Dao) Insert() (result sql.Result, err error)
```

#### 示例代码

```go
package main

import (
  "fmt"
  "github.com/whf-sky/efficient.demo/application/dao"
  "github.com/whf-sky/efficient.demo/application/model"
  "math/rand"
  "strconv"
  "time"
)

func main()  {
  rand.Seed(time.Now().UnixNano())
  users := model.NewUsers()
  users.SetUsername("zhangsan-" + strconv.FormatUint(rand.Uint64(), 10)).
    SetPasswd("123456").
    SetAge(10)

  userDao := dao.NewUsersDao()
  ret,_ := userDao.SetData(users).Insert()

  fmt.Println(ret.LastInsertId())
}
```

#### SetData 

`设置数据`

```go
func (d *Dao) SetData(data ...ModelInterface) *Dao
```

#### Insert 

`执行插入操作`

```go
func (d *Dao) Insert() (result sql.Result, err error)
```

### update

#### 示例代码

```go
package main

import (
	"fmt"
	"github.com/whf-sky/efficient.demo/application/dao"
	"github.com/whf-sky/efficient.demo/application/model"
)

func main()  {
	users := model.NewUsers()

	users.SetUsername("lisi").
		SetPasswd("456789").
		SetAge(5)

  usersDao := dao.NewUsersDao()
  rowsAffected, err := usersDao.SetData(users).Where("id=?", 1342).Update()
  fmt.Println("Update:", rowsAffected, err)
}
```

#### SetData 

`设置数据`

```go
func (d *Dao) SetData(data ...ModelInterface) *Dao
```

#### Where 

`设置条件`

```go
func (d *Dao) Where(where string, args ...interface{}) *Dao 
```

#### Update 

`执行修改操作`

```go
func (d *Dao) Update() (rowsAffected int64, err error) 
```

### Delete

#### 示例代码

```go
package main

import (
	"fmt"
	"github.com/whf-sky/efficient.demo/application/dao"
	"github.com/whf-sky/efficient.demo/application/model"
)

func main() {
  usersDao := dao.NewUsersDao()

  rowsAffected, err := usersDao.Where("id=?", 1342).Delete()

  fmt.Println("Delete:", rowsAffected, err)
}
```

#### Where 

`设置条件`

```go
func (d *Dao) Update() (rowsAffected int64, err error) 
```

#### Delete 

`执行删除操作`

```go
func (d *Dao) Delete() (rowsAffected int64, err error) 
```

### QueryRow

#### 示例代码

```go
package main

import (
  "fmt"
  "github.com/whf-sky/efficient.demo/application/dao"
)

func main()  {

  userDao := dao.NewUsersDao()
  row := userDao.QueryRow("select id,username,passwd from users where id=?", 100)
  var id int64
  var username string
  var passwd string
  row.Scan(&id, &username, &passwd)

  fmt.Println(id, username, passwd)
}

```

#### QueryRow 

`单行查询`

```go
func (d *Dao) QueryRow(query string, args ...interface{}) *sql.Row 
```

#### Scan 

`获取行结果`

```go
var id int64
var username string
var passwd string
row.Scan(&id, &username, &passwd)
```

### Query

#### 示例代码

```go
package main

import (
  "fmt"
  "github.com/whf-sky/efficient.demo/application/dao"
)

func main()  {
  userDao := dao.NewUsersDao()
  result := userDao.Query("select id,username,passwd from users where passwd=?", "123")
  rows, err := result.ToModels()
  fmt.Println(rows, err)
}
```

#### Query 

`执行多行查询`

```go
func (d *Dao) Query(query string, args ...interface{}) *query 
```

#### Rows 

`原生获取多行数据`

```go
func (q *query) Rows() (*sql.Rows, error) 
```

`Rows 示例代码`

```go
rows, err := result.Rows()
defer func() {
  err := rows.Close()
  if err != nil {
    fmt.Println(err)
  }
}()
if err != nil {
  fmt.Println(err)
}
if err != nil {
  fmt.Println(err)
}
ret := []map[string]interface{}{}
for rows.Next() {
  var id int64
  var username string
  var passwd string
  rows.Scan(&id, &username, &passwd)
  ret = append(ret, map[string]interface{}{
    "id":id,
    "username":username,
    "passwd":passwd,
  })
}
err = rows.Close()
if err != nil {
  fmt.Println(err)
}
fmt.Println("rows=============", ret)
```

#### ToMaps 

`把结果集放入 []map[string]interface{}`

```go
func (q *query) ToMaps() (result []map[string]interface{}, err error)
```

#### ToMap 

`获取单行数据,把结果集放入map[string]interface{}`

`使用Query()获取单行数据一定要加limit 1`

```go
func (q *query) ToMap() (result map[string]interface{}, err error)
```

#### ToModels 

`把结果集放入 []ModelInterface`

```go
func (q *query) ToModels() (models []ModelInterface, err error)
```

#### ToModel 

`获取单行数据,把结果集放入 ModelInterface`

`使用Query()获取单行数据一定要加limit 1`

```go
func (q *query) ToModel() (model ModelInterface, err error) 
```

### Exec

#### 示例代码

```go
result, err := userDao.Exec("delete from users where id=?", 10)
if err != nil {
    return 0, err
}
result.RowsAffected()
```

#### Exec 

`执行SQL`

```go
func (d *Dao) Exec(query string, args ...interface{}) (result sql.Result, err error)
```

### transaction

#### 示例代码

```go
package main

import (
	"fmt"
	"github.com/whf-sky/efficient.demo/application/dao"
)

func main()  {
	var err error
	var rowsAffected int64
	usersDao := dao.NewUsersDao()

	usersDao.Begin()

	defer func() {
		if err != nil {
			_ = usersDao.Rollback()
		}
	}()

	rowsAffected, err = usersDao.Where("id=?", 1342).Delete()

	fmt.Println("rowsAffected:", rowsAffected)

	if err != nil {
		err = usersDao.Rollback()
	} else {
		err = usersDao.Commit()
	}

	fmt.Println("Delete:", err)
}

```


#### Begin 

`创建事务`

```go
func (d *Dao) Begin() *Dao
```

#### Commit 

`提交事务`

```go
func (d *Dao) Commit() error
```


#### Rollback 

`回滚事务`

```go
func (d *Dao) Rollback() error
```


# 数据验证

#### 示例代码

```go
package main

import (
	"fmt"
	"github.com/whf-sky/efficient/validation"
)

func main() {
	errs := validation.NewValidation().
		Validator("num", 101.1, validation.V{"gte": 5.1, "lte": 100.1}, "数量").
		Validator("num1", 101, validation.V{"gte": 5, "lte": 100}, "数量").
        Errors()

	fmt.Println(errs)
}
```

## 使用

#### import

```go
import "github.com/whf-sky/efficient/validation"
```

#### New

`做数据验证前，先对验证器进行实例化`

```go
validation.NewValidation()
```

#### Validator

`添加验证规则`

`key 需要验证的数据标识，如：验证数据为用户名，key即可为username`

`data 需要验证的数据`

`validators 验证规则`

`comments 要验证的数据的注释，如：key为id，comment为编号`

```go
func (v *Validation) Validator(key string, data interface{}, validators V, comments ...string) *Validation 
```

#### V

`验证规则，底层数据结构为 map[string]interface{}`

`key 为验证规则的英文名称，即验证器`

`value 为规则的数据，规则数据要与验证的数据的类型保持一致，否则验证失败报类型不匹配错误`

```go
validation.V{"gte": 5.1, "lte": 100.1}
```

## 验证规则

### eq `等于`

`支持的数据类型`

```
int8 uint8 int16 uint16 int uint int32 uint32 int64 uint64 float32 float64 string bool

```

### neq `不等于`

`支持的数据类型`

```
int8 uint8 int16 uint16 int uint int32 uint32 int64 uint64 float32 float64 string bool

```

### gt `大于`

`支持的数据类型`

```
int8 uint8 int16 uint16 int uint int32 uint32 int64 uint64 float32 float64 string

```

### gte `大于等于`

`支持的数据类型`

```
int8 uint8 int16 uint16 int uint int32 uint32 int64 uint64 float32 float64 string

```

### lt `小于`

`支持的数据类型`

```
int8 uint8 int16 uint16 int uint int32 uint32 int64 uint64 float32 float64 string

```

### lte `小于等于`

`支持的数据类型`

```
int8 uint8 int16 uint16 int uint int32 uint32 int64 uint64 float32 float64 string

```

### in `数据在验证规则数据中`

`验证数据支持的数据类型`

```
int8 uint8 int16 uint16 int uint int32 uint32 int64 uint64 float32 float64 string

```

`验证规则数据为验证数据的数组形式，如：[]int8`

### in-multi `多个数据在验证规则数据中`

`验证数据和验证规则数据支持的数据类型为下面类型的数组形式，如：[]int8`

```
int8 uint8 int16 uint16 int uint int32 uint32 int64 uint64 float32 float64 string

```

### empty `空`

支持的数据类型 `string`

### email `邮箱`

支持的数据类型 `string`

### regexp `正则`

支持的数据类型 `string`

### required `必须的`

支持的数据类型 `nil`


# 日志

使用的是日志插件logrus,对外暴露的`efficient.Log`,错误级别调用的方法都与logrus一致

https://github.com/sirupsen/logrus

gin的 `gin.DefaultWriter` 和 `gin.DefaultErrorWriter` 可以分别设置

示例代码

```go
package config

import (
  "github.com/gin-gonic/gin"
  "github.com/sirupsen/logrus"
  "github.com/whf-sky/efficient"
  "os"
)

func init()  {
  env := efficient.GetEnv()
  if env == "production" {
    gin.DefaultWriter = os.Stdout
    gin.DefaultErrorWriter = os.Stderr
    efficient.SetLogger(func(logger *logrus.Logger, log *logrus.Entry) {
      logger.SetReportCaller(false)
      log = logrus.NewEntry(logger)
      logger.Out = os.Stdout
    })
  }
}
```

## import

```go
import "github.com/whf-sky/efficient"
```

## Logger的设置

```go
package config

import (
	"github.com/sirupsen/logrus"
	"github.com/whf-sky/efficient"
)

func init()  {
	env := efficient.GetEnv()
	if env == "production" {
		efficient.SetLogger(func(logger *logrus.Logger, log *logrus.Entry) {
			logger.SetReportCaller(false)
			log = logrus.NewEntry(logger)
		})
	}
}
```

#### SetLogger

可对`*logrus.Logger`进行设置，对`*logrus.Entry`进行初始化

```go
func SetLogger(fun func(logger *logrus.Logger, log *logrus.Entry))
```



# 环境变量

## import

```go
import "github.com/whf-sky/efficient"
```

## 获取环境变量 

`默认production`

```go
efficient.GetEnv()
```

## 设置系统环境变量 

#### Windows 下设置 EFFICIENT_ENV 的命令为：
```sh
$ set EFFICIENT_ENV=test
```

#### MacOS 或者 Linux 下设置 EFFICIENT_ENV 的命令为：
```sh
$ export EFFICIENT_ENV=test
```

# 生成器

#### 示例代码 

`可以把配置信息提取到配置文件中`

```go
//实例化代码生成器
generate := generate.NewGenerate()
//应用生成
generate.Application()
//设置数据库并生成dao和model文件
db : = database.NewDb().MysqlDsn("127.0.0.1", "3306", "root", "123456", "test", "utf8mb4").Open()
generate.SetDb("mysql", db).Database("test")
```

#### import

`暂只支持Mysql模型生成，可根据情况导入下面的包`

`github.com/go-sql-driver/mysql mysql连接驱动包`

`github.com/whf-sky/efficient/database 数据层包`

`github.com/whf-sky/efficient/generate 生成器的包`

```go
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/whf-sky/efficient/database"
	"github.com/whf-sky/efficient/generate"
)
```

#### 实例化生成器

```go
generate.NewGenerate()
```

#### 应用目录设置

```go
//此目录是默认目录，无特殊需要可以不设置，需要在调用Application()前调用此方法
appdirs := map[string]string{
  "dao":     "application/dao",
  "model":   "application/model",
  "service": "application/service",
  "config":  "config",
  "cmd":     "cmd",
}
generate.SetAppDir(appdirs)
```

#### 生成应用文件 

`应用目录参考'应用目录设置'`

```go
NewGenerate().Application()
```

#### 生成数据层文件

```go
generate.
	//设置数据库连接
	SetDb("驱动", database.NewDb().MysqlDsn("127.0.0.1", "3306", "root", "123456", "test", "utf8mb4").Open()).
	//生成数据层文件
	Database("数据库")
```

# 插件注册

## 注册数据类型

`注册的数据类型主要用于代码生成的类型注册，一般情况用不到`

#### import

```go
import "github.com/whf-sky/efficient/generate/database/mysql/generate"
```

#### RegisterType 

`注册mysql对应的golang数据类型，非NULL值的`

`unsigned 是否有符号`

`mType mysql数据类型`

`gType golang数据类型`

`pkgs golang的类型包，只能写一个`

```go
func RegisterType(unsigned bool, mType, gType string, pkgs ...string)
```

#### RegisterNullType 

`注册mysql对应的golang数据类型，可NULL值的`

`unsigned 是否有符号`

`mType mysql数据类型`

`gType golang数据类型`

`pkgs golang的类型包，只能写一个`

```go
func RegisterNullType(unsigned bool, mType, gType string, pkgs ...string) 
```

## 注册验证器

#### import

```go
import "github.com/whf-sky/efficient/validation"
```

#### RegisterValidation 

`注册验证器`

`name 验证器名称`

`validation 验证器`

```go
func RegisterValidation(name string, validation ValidationHandle)
```

## 注册代码生成器

#### import

```go
import "github.com/whf-sky/efficient/generate"
```

#### RegisterGenerate 

`注册生成器`

`driver 数据库驱动`

`generate 生成器`

```go
func RegisterGenerate(driver string, generate GenerateInterface) 
```

## 注册SQL生成器

#### import

```go
import "github.com/whf-sky/efficient/database"
```

#### RegisterSQLS 

`注册SQL生成器`

`driver 数据库驱动`

`sql sql生成器`

```go
func RegisterSQLS(driver string, sql SQLInterface) 
```

# 工具包

参见 `widget/tools` 包
