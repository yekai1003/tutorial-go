GoLang，习惯被人成为go语言，是目前后端程序员比较喜欢的一款语言，在区块链行业内，go语言更被称为是第一编程语言。go语言的优势有很多，比如内存回收、高并发、语法简洁、开发效率和运行效率都高等。本文通过一个例子，给大家介绍如何用go语言来实现一个访问数据库（mysql）的客户端，也就是我们经常用到的那种命令行窗口！
在go语言中为我们提供了sql操作的api，需要引用包：database/sql，不过如果要使用mysql，还需要引用包（驱动包）：github.com/go-sql-driver/mysql ，而且这个包在引用的时候要求是匿名引用，也就是说database/sql在实现时需要借助驱动包的内容。在实现我们的小目标之前，我们先要分析一下都需要做哪些事情。

首先我们要做数据库的开发，肯定要知道如何调用curd（增删改查）接口，当然在调用这些接口前，你需要先知道如何连接或者说登陆到数据库，毕竟针对数据库的操作，都是在登陆之后才可以操作的！

连接数据库，我们使用sql包内的Open函数，顾名思义是打开一个数据库，它的函数原型如下：


```
func Open(driverName, dataSourceName string) (*DB, error) 
```

driverName显然就是驱动的名字，在这里我们填写“mysql”就可以，dataSourceName是数据源，需要填写mysql的连接串。


```
username:password@protocol(address)/dbname?param=value
```

protocol是代表连接mysql的方式，可以用tcp，也可以用本地unix，dbname就是要连接数据库的名字，param=value则是在登陆时的设置，比如登陆时限定字符集为utf8。参考例子如下：


```
root:abc123@tcp(10.211.55.3:3306)/yekai?charset=utf8
```

在搞清楚数据源如何填写后，我们就可以使用Open函数了，它的返回值是一个数据库连接句柄，此外还有一个错误信息提示，当没有错误时，此值为nil。


```
    db, err := sql.Open("mysql", "root:abc123@tcp(10.211.55.3:3306)/yekai?charset=utf8")
	if err != nil {
		log.Panic("failed to open mysql ", err)
	}
```

这样我们就获得了一个数据库的连接句柄，但是小编惊奇的发现，当yekai这个数据库不存在的时候，该函数并不会报错，但是如果ip地址填错了，则访问超时报错。因此为了确保连接确实没问题，确实可用，我们可以在用Sql.DB结构体内部的Ping函数来测试一下，如果Ping没有报错，则代表真的连接成功，可以将此部分代码放在init函数中，这样代码在初始化时会自动运行。


```
var dbconn *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:abc123@tcp(10.211.55.3:3306)/yekai?charset=utf8")
	if err != nil {
		log.Panic("failed to open mysql ", err)
	}
	if err = db.Ping(); err != nil {
		log.Panic("failed to ping mysql ", err)
	}
	dbconn = db
}
```

连接到mysql数据库后，接下来就可以做增删改查操作了。其实说是增删改查操作，那是指sql层面的，对我们开发客户端来说，可以认为主要是两类数据库操作，一类是有结果集返回的，一类是没有结果集返回的。比如create、drop、insert、update等这些都是无结果集返回的语句，而show、desc、select则是有结果集返回的sql，其中当然用的最多的也就是select这样的sql。

在go语言开发中，我们实际用的也可以认为多是两类接口，一个是Exec，一个是Query，这两个接口都是Sql.DB结构内的函数，我们在Open之后得到的结果刚好就是调用的入口。


```
func exec_sql(xsql string) {
	result, err := db.Exec(xsql)
	if err != nil {
		fmt.Println("failed to exec sql:", xsql, err)
		return
	}
	rowsaff, _ := result.RowsAffected()
	fmt.Println("RowsAffected:", rowsaff)
}
```

从返回的result中，可以查询到影响的记录数，以上就是没有结果集的函数操作。对于有结果集的函数，操作起来就要麻烦一些，主要就是针对结果集的处理。


```
    rows, err := db.Query(xsql)
	if err != nil {
		fmt.Println("failed for query sql:", err)
		return
	}
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("failed for get columns:", err)
		return
	}
	colCount := len(cols)

	//fmt.Println(colCount, "\n", cols)
	for _, v := range cols {
		fmt.Printf("%s\t", v)
	}
```

使用Query函数，可以进行查询操作，返回一个结果集的rows，我们可以把他理解为结果集的多行记录。首先在rows这个结构集结构体中，我们可以用Columns()获得全部的列名，如果要打印结果集，可以先把列名打印出来，当然我们也可以根据Columns()获得对应的字段个数。

如果要获得每一行结果集以及具体字段的value值，那么就需要遍历结果集以及扫描结果记录了。主要使用rows.Next()以及rows.Scan()相配合，当Next返回为真时，可以调用Scan去获得该条记录的结果集，对于明确结果集的查询，比较好办，我们可以直接定义好要接收的变量，将它传入到Scan中去获得相应的值，因为Scan接口是这样的：


```
Scan(dest ...interface{}) 
```
我们可以把要接收的多个目标传进去，这样就可以获得对应的该条记录的不同字段的值，比如代码可以写成这样：


```
rows, err := db.Query("SELECT name FROM users WHERE age = ?", age)
if err != nil {
    log.Fatal(err)
}
for rows.Next() {
    var name string
    if err := rows.Scan(&name); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s is %d\n", name, age)
}
if err := rows.Err(); err != nil {
    log.Fatal(err)
}
```

对于我们要写一个sql是用人随便输入的客户端来说，显然不能用这样的方式，在这里我们可以利用绑定接口值的方式，提前定义好两个map进行接口变量的绑定。


```
    values := make([]String, colCount)
	oneRows := make([]interface{}, colCount)
	for k, _ := range values {
		oneRows[k] = &values[k] //将查询结果的返回地址绑定，这样才能变参获取数据
	}
```

然后扫描的代码就可以直接使用oneRows了，当oneRows被扫描后，values的结果也填充完成了。


```
for rows.Next() {

		//扫描结果集，一定要在Next调用后，方可使用
		err = rows.Scan(oneRows...)
		if err != nil {
			fmt.Println("failed to Scan result set", err)
			break
		}
		//fmt.Println(values)
		for _, v := range values {
			if v.Valid {
				fmt.Printf("%s\t", v.String)
			} else {
				fmt.Printf("%s\t", "NULL")
			}
		}
		fmt.Println()
	}
```

分析到此，我们可以具备编写客户端的能力了，只需要编写一个循环接收输入的命令终端，将命令分类，如果是查询类，也就是有结果集这一类的操作时我们调用Query相关的处理，当调用无结果集的操作时，我们直接调用Exec即可。

整理下来的步骤应该是这样：
- 1.连接到数据库
- 2.循环等待命令输入
- 3.判断是查询还是非查询
- - 3.1 如果是查询，调用Query，打印结果集
- - 3.2 如果非查询，调用Exec，打印影响记录数


参考代码如下：

```
/*
   file    : client.go
   author  : yekai
   company : pdj(pdjedu.com)
*/

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	connstr string
	driver  string
	dbconn  *sql.DB
}

func NewClient(user, pass, dbname, protocol string) *Client {
	connstr := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", user, pass, protocol, dbname)
	return &Client{connstr, "mysql", nil}
}

func (cli *Client) Conn() error {
	db, err := sql.Open(cli.driver, cli.connstr)
	if err != nil {
		fmt.Println("failed to open database ", err)
		return err
	}
	cli.dbconn = db
	return cli.dbconn.Ping()
}

func (cli *Client) Run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to mysql client ")
	for {
		fmt.Printf("yekai-mysql>")
		sqlstr, err := reader.ReadString('\n')
		if err != nil {
			log.Panic("failed to ReadString ", err)
		}
		sqlstr = strings.Trim(sqlstr, "\r\n")
		sqls := []byte(sqlstr)
		if len(sqls) > 6 {
			if string(sqls[:6]) == "select" || string(sqls[:4]) == "show" || string(sqls[:4]) == "desc" {
				//result set sql
				cli.query_sql(sqlstr)
			} else {
				//no result set sql
				cli.exec_sql(sqlstr)
			}
		}
		if sqlstr == "quit" {
			fmt.Println("bye bye ")
			break
		}
	}
}

func (cli *Client) exec_sql(xsql string) {
	result, err := cli.dbconn.Exec(xsql)
	if err != nil {
		fmt.Println("failed to exec sql:", xsql, err)
		return
	}
	rowsaff, _ := result.RowsAffected()
	fmt.Println("RowsAffected:", rowsaff)
}

func (cli *Client) query_sql(xsql string) {
	rows, err := cli.dbconn.Query(xsql)
	if err != nil {
		fmt.Println("failed for query sql:", err)
		return
	}
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("failed for get columns:", err)
		return
	}
	colCount := len(cols)

	//fmt.Println(colCount, "\n", cols)
	for _, v := range cols {
		fmt.Printf("%s\t", v)
	}
	fmt.Println("\n----------------------------------------")

	values := make([]String, colCount)
	oneRows := make([]interface{}, colCount)
	for k, _ := range values {
		oneRows[k] = &values[k] //将查询结果的返回地址绑定，这样才能变参获取数据
	}

	for rows.Next() {

		//扫描结果集，一定要在Next调用后，方可使用
		err = rows.Scan(oneRows...)
		if err != nil {
			fmt.Println("failed to Scan result set", err)
			break
		}
		//fmt.Println(values)
		for _, v := range values {
			if v.Valid {
				fmt.Printf("%s\t", v.String)
			} else {
				fmt.Printf("%s\t", "NULL")
			}
		}
		fmt.Println()
	}
}

```



```
/*
   file    : main.go
   author  : yekai
   company : pdj(pdjedu.com)
*/
package main

import (
	"log"
)

func main() {
	cli := NewClient("root", "abc123", "yekai", "tcp(10.211.55.3:3306)")
	if cli.Conn() != nil {
		log.Panic("failed to conn to mysql ")
	}
	cli.Run()
}
```

上述代码还有点小问题，因为数据库里还有一个非常坑人的小玩意儿--NULL，在查询到空值的时候，我们的普通String类型处理不了，这时需要携带指示器类型的结构体，在sql包内给我们提供了sql.NullString，我们直接使用即可，也就是将上述代码稍稍替换一下。


```
    //使用NullString可以很好的支持空值
	values := make([]sql.NullString, colCount)
	oneRows := make([]interface{}, colCount)
```

好，终于写完了，我们不妨来测试一下效果！


```
localhost:sql yekai$ go run main.go client.go 
Welcome to mysql client 
yekai-mysql>show databses
failed for query sql: Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'databses' at line 1
yekai-mysql>show databases
Database	
----------------------------------------
information_schema	
mysql	
performance_schema	
sys	
yekai	
yekai-mysql>show tables
Tables_in_yekai	
----------------------------------------
person	
person2	
yekai-mysql>desc person
Field	Type	Null	Key	Default	Extra	
----------------------------------------
name	varchar(30)	YES		NULL		
age	int(11)	YES		NULL		
yekai-mysql>select * from person
name	age	
----------------------------------------
luffy	18	
zero	30	
nami	20	
sanji	25	
robin	35	
usopp	20	
yekai-mysql>quit
bye bye 
localhost:sql yekai$ 
```

