package repository

import (
	"testing"
	"fmt"
	"time"
	"strconv"
)

func TestGetMySQLEngine(t *testing.T){
	GetMySQLEngine(true)
}


func TestSaveCustomer(t *testing.T) {
	customer := new(Customer)
	customer.Phone = "15913313959"
	customer.Sex = "女"
	customer.Username = "李嘉洁"
	customer.Birthday =time.Date(1990, time.February,
		19, 0, 0, 0, 0, time.UTC)
	result, err := SaveCustomer(customer)
	if err != nil || customer.Id == 0{
		t.Error("error:", err.Error())
	} else {
		fmt.Println(fmt.Sprintf("affected: %d, Id: %d", result, customer.Id))
	}

}

func TestSaveCustomers(t *testing.T) {
	cs := []Customer{
		{Username:"李强", Sex:"男", Phone: "15918877539", Address:"广州天河区", Birthday:time.Date(1990, time.February, 19, 0, 0, 0, 0, time.UTC)},
		{Username:"张学森", Sex:"男", Phone: "15918899539", Address:"广州天河区", Birthday:time.Date(1987, time.February, 19, 0, 0, 0, 0, time.UTC)},
		{Username:"吴琪", Sex:"男", Phone: "15918899599", Address:"广州天河区", Birthday:time.Date(1988, time.May, 19, 0, 0, 0, 0, time.UTC)},
		{Username:"吴梦如", Sex:"女", Phone: "15918899599", Address:"广州天河区", Birthday:time.Date(1989, time.May, 19, 0, 0, 0, 0, time.UTC)},
		{Username:"王梦茹", Sex:"女", Phone: "15918899599", Address:"广州天河区", Birthday:time.Date(1992, time.October, 29, 0, 0, 0, 0, time.UTC)},
	}
	affected, err := SaveCustomers(cs)
	if err != nil {
		t.Error(err.Error())
	}else {
		fmt.Printf("成功添加记录：%d\n", affected)
	}

}

func TestXORMFind( t *testing.T){
	var cs []Customer
	err := GetMySQLEngine(false).Find(&cs)
	if err != nil {
		t.Error(err.Error())
	}else {
		for i, c := range cs {
			fmt.Println(fmt.Sprintf("i: %d, c:%v", i, c))
		}
	}
	fmt.Println("--------------WHERE, LIMIT-----------------")
	//Where, Limit
	var cs1 []Customer
	err = GetMySQLEngine(false).Where(" id > ? AND sex = ?", 5, "男").Limit(3, 2).Find(&cs1)
	if err != nil {
		t.Error(err.Error())
	}else {
		for i, c := range cs1 {
			fmt.Println(fmt.Sprintf("i: %d, c:%v", i, c))
		}
	}
	fmt.Println("--------------In-----------------")
	//in opts
	var cs2 []Customer
	err = GetMySQLEngine(false).In("username", "王梦茹", "吴梦如").Find(&cs2)
	if err != nil {
		t.Error(err.Error())
	}else {
		for i, c := range cs2 {
			fmt.Println(fmt.Sprintf("i: %d, c:%v", i, c))
		}
	}
}

func TestUpdateCustomer(t *testing.T) {
	c := Customer{Id:43}
	c.Phone = "15918899591"
	affected, err := GetMySQLEngine(false).Id(c.Id).Update(c)
	if err != nil {
		t.Error(err.Error())
	}else {
		fmt.Printf("成功修改的记录：%d\n", affected)
		c1 := Customer{Id: 43}
		isExist, err :=GetMySQLEngine(false).Get(&c1)
		if err != nil {
			t.Error(err.Error())
		}else if isExist {
			fmt.Println(fmt.Sprintf("customer:%v", c1))
		}
	}
}

func TestUpdateCustomerByMap(t *testing.T) {
	props := make(map[string]interface{}, 2)
	props["phone"] = "15919918138"
	props["address"] = "广州越秀区"
	affected, err := UpdateCustomerByMap(43, props)
	if err != nil {
		t.Error(err.Error())
	}else {
		fmt.Printf("成功修改的记录：%d\n", affected)
		c1 := Customer{Id: 43}
		isExist, err :=GetMySQLEngine(false).Get(&c1)
		if err != nil {
			t.Error(err.Error())
		}else if isExist {
			fmt.Println(fmt.Sprintf("customer:%v", c1))
		}
	}

}

func TestDeleteRecord(t *testing.T){
	affected, err := GetMySQLEngine(false).Delete(&Customer{Username:"张学森"})
	if err != nil {
		t.Error(err.Error())
	}else {
		fmt.Printf("成功删除的记录：%d\n", affected)
	}
}

func TestGetCurrentFileName(t *testing.T){
	fmt.Println(GetCurrentFileName() )
}


func TestExecuteSQL(t *testing.T) {
	//creation operation
	sql := "INSERT INTO customer(username, phone, address, sex) VALUES(?, ?, ?, ?)"
	affected, err := GetMySQLEngine(false).Exec(sql, "刘继刚", "13913413545", "广州天河区", "男")
	if err != nil {
		t.Error(fmt.Sprintf("%v", err.Error()))
	}else {
		fmt.Printf("成功创建的记录：%d\n", affected)
		c1 := Customer{Username: "刘继刚"}
		isExist, err :=GetMySQLEngine(false).Get(&c1)
		if err != nil {
			t.Error(err.Error())
		}else if isExist {
			fmt.Println(fmt.Sprintf("customer:%v", c1))
		}
	}

	//update operation
	sql = "UPDATE customer SET phone=? WHERE username=?"
	affected, err = GetMySQLEngine(false).Exec(sql, "13914313949", "刘继刚")
	if err != nil {
		t.Error(fmt.Sprintf("%v", err.Error()))
	}else {
		fmt.Printf("成功修改的记录：%d\n", affected)
		c1 := Customer{Username: "刘继刚"}
		isExist, err :=GetMySQLEngine(false).Get(&c1)
		if err != nil {
			t.Error(err.Error())
		}else if isExist {
			fmt.Println(fmt.Sprintf("customer:%v", c1))
		}
	}

	//deletion operation
	sql = "DELETE FROM customer WHERE username=?"
	affected, err = GetMySQLEngine(false).Exec(sql, "刘继刚")
	if err != nil {
		t.Error(fmt.Sprintf("%v", err.Error()))
	}else {
		fmt.Printf("成功删除的记录：%d\n", affected)
	}
}

func TestQuerySQL(t *testing.T){
	sql := "SELECT id, username, phone, address, sex  FROM customer ORDER BY id desc"
	result, err := GetMySQLEngine(false).Query(sql)
	customers := make([]Customer, 0)
	if err != nil {
		t.Error(err.Error())
	}else {
		for _, row := range result {
			if len(row) > 0 {
				c := Customer{}
				id , _ := strconv.ParseInt(string(row["id"]), 10, 64)
				c.Id = id
				c.Username = string(row["username"])
				c.Phone = string(row["phone"])
				c.Address = string(row["address"])
				c.Sex = string(row["sex"])
				customers = append(customers, c)
			}
		}
	}
	if len(customers) > 0 {
		for i, v := range customers {
			fmt.Println(fmt.Sprintf("i: %d, v: %v", i, v))
		}
	}
}