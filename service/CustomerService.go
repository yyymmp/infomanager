package service

import (
	"fmt"
	"gocode/infomanager/model"
)

//业务逻辑层，增删改查
type CustomerService struct {
	customer    []model.Customer
	customerNum int //长度
}

//客户
func NewCustomerService() *CustomerService {
	c := model.NewCustomer(1, "jiaa", "nan", 90, "5545", "dad@qq.com") //初始化一个用户
	cs := CustomerService{}
	cs.customerNum = 1
	cs.customer = append(cs.customer, *c)
	return &cs

}

//添加用户
func (cs *CustomerService) Insert(c model.Customer) {
	cs.customerNum++
	c.Id = cs.customerNum //id系统分配
	cs.customer = append(cs.customer, c)
}

//用户列表
func (cs *CustomerService) Index() []model.Customer {
	return cs.customer
}

//根据索引删除用户
func (cs *CustomerService) Delete(index int) bool {
	if index == -1 {
		return false
	}
	cs.customer = append(cs.customer[:index], cs.customer[index+1:]...)
	return true
}

//查询是否有函数存在
func (cs *CustomerService) FindById(index int) int {
	flag := -1
	for inde, in := range cs.customer {
		if in.Id == index {
			flag = inde
			break
		}
	}
	return flag
}

//输出信息
func (cs *CustomerService) Info(id int) bool {
	index := cs.FindById(id) //切片索引
	if index == -1 {
		fmt.Println("客户不存在")
		return false
	}
	newData := model.Customer{
		Id:     id,
		Name:   "",
		Gender: "",
		Age:    0,
		Phone:  "",
		Email:  "",
	}
	fmt.Printf("客户姓名(%v):\n", cs.customer[index].Name)
	fmt.Scanln(&newData.Name)
	fmt.Printf("客户性别(%v):\n", cs.customer[index].Gender)
	fmt.Scanln(&newData.Gender)
	fmt.Printf("客户年龄(%v):\n", cs.customer[index].Age)
	fmt.Scanln(&newData.Age)
	fmt.Printf("客户电话(%v):\n", cs.customer[index].Age)
	fmt.Scanln(&newData.Phone)
	fmt.Printf("客户邮箱(%v):\n", cs.customer[index].Age)
	fmt.Scanln(&newData.Email)
	return cs.Update(index, newData)
}

//更新
func (cs *CustomerService) Update(index int, c model.Customer) bool {
	cs.customer[index].Id = c.Id
	cs.customer[index].Email = c.Email
	cs.customer[index].Phone = c.Phone
	cs.customer[index].Gender = c.Gender
	cs.customer[index].Name = c.Name
	return true
}
