package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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

func main() {
	con, er := sql.Open("postgres", "dbname=postgres sslmode=disable")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	q := "select * from mydata"
	rs, er := con.Query(q)
	if er != nil {
		panic(er)
	}
	for rs.Next() {
		var md Mydata
		er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
		if er != nil {
			panic(er)
		}

		fmt.Println(md.Str())
	}

}
