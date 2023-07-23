package main

import (
	"errors"
	"fmt"
	"strconv"
)

type employee struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Salary   int    `json:"salary"`
	Exp      int    `json:"exp"`
	About    string `json:"about"`
}

var EmpList = map[int]employee{}

func InfById(InId string) (error, *employee) { //finds employee with input id
	id, err := strconv.Atoi(InId)
	if err != nil {
		panic(err)
	}
	for _, val := range EmpList {
		if val.Id == id {
			return nil, &val
		}
	}
	return errors.New("No such employee"), nil
}

////////////DB logic

func DeleteById(InId string) error { //deletes employee with input id
	id, err := strconv.Atoi(InId)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec("DELETE FROM employee where id = $1", id)
	if err == nil {
		for key, _ := range EmpList {
			if key == id {
				delete(EmpList, key)
				return nil
			}
		}
	}
	return err
}

func DataFDb() { // loads data from database to map
	rows, err := DB.Query("select * from employee")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		exmp := employee{}
		err := rows.Scan(&exmp.Id, &exmp.Name, &exmp.Position, &exmp.Salary, &exmp.Exp, &exmp.About)
		if err != nil {
			fmt.Println(err)
			continue
		}
		EmpList[exmp.Id] = exmp
	}
}
func DataTDb(name, pos, about string, salary, exp int) { // loads new row to database
	_, err := DB.Exec("insert into employee(name,position,salary,experience,about) values ($1,$2,$3,$4,$5)", name,
		pos, salary, exp, about)
	if err != nil {
		fmt.Println(err)
	}
}
