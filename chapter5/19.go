package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"hello"
	"strconv"
)

type Mydata struct {
	ID   int
	Name string
	Mail string
	Age  int
}

func (m Mydata) Str() string {
	return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\" " + m.Mail + "," + strconv.Itoa(m.Age) + ">"
}

var qry string = "select * from mydata where id = $1"

func main() {
	con, er := sql.Open("postgres", "dbname=postgres sslmode=disable")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	for true {
		s := hello.Input("id")
		if s == "" {
			break
		}
		n, er := strconv.Atoi(s)
		if er != nil {
			panic(er)
		}
		rs := con.QueryRow(qry, n)
		var md Mydata
		er2 := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er2 != nil {
			panic(er2)
		}
		fmt.Println(md.Str())
	}

	fmt.Println("***end***")
}
