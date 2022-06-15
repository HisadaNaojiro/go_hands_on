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

func main() {
	con, er := sql.Open("postgres", "dbname=postgres sslmode=disable")
	if er != nil {
		panic(er)
	}
	defer con.Close()

	nm := hello.Input("name")
	ml := hello.Input("mail")
	age := hello.Input("age")
	ag, _ := strconv.Atoi(age)

	qry := "insert into mydata (name, mail, age) values ($1, $2, $3)"

	con.Exec(qry, nm, ml, ag)
	showRecord(con)
}

func showRecord(con *sql.DB) {
	qry := "select * from mydata"
	rs, _ := con.Query(qry)
	for rs.Next() {
		fmt.Println(mydatafmRws(rs).Str())
	}
}

func mydatafmRws(rs *sql.Rows) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}

func mydatafmRw(rs *sql.Row) *Mydata {
	var md Mydata
	er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
	if er != nil {
		panic(er)
	}
	return &md
}
