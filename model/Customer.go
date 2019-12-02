package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//工厂模式 返回实例  接受传入的信息
func NewCustomer(id int, name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//显示一个客户列表
func (customer Customer) Getinfo() string {
	line := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", customer.Id, customer.Name, customer.Gender, customer.Age, customer.Phone, customer.Email)
	return line
}
