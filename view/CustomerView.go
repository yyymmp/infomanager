package main

import (
	"fmt"
	"gocode/infomanager/model"
	"gocode/infomanager/service"
)

type customerView struct {
	key  string                   //接受用户输入
	loop bool                     //是否循环
	cs   *service.CustomerService //操作
}

//index 列表方法  调用CustomerService的list方法
func (this *customerView) index() {
	list := this.cs.Index()
	//打印表头
	fmt.Println("----------客户列表-------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱\t")
	for _, val := range list {
		fmt.Println(val.Getinfo())
	}
	fmt.Println("---------客户列表end------")

}

//添加客户  输出页面并且调用service层方法
func (this *customerView) insert() {
	var c model.Customer
	c = model.Customer{
		Name:   "",
		Gender: "",
		Age:    0,
		Phone:  "",
		Email:  "",
	}
	fmt.Println("请输入客户姓名")
	fmt.Scanln(&c.Name)
	fmt.Println("请输入客户性别")
	fmt.Scanln(&c.Gender)
	fmt.Println("请输入客户年龄")
	fmt.Scanln(&c.Age)
	fmt.Println("请输入客户电话")
	fmt.Scanln(&c.Phone)
	fmt.Println("请输入客户邮箱")
	fmt.Scanln(&c.Email)
	this.cs.Insert(c)
	fmt.Println("添加成功")
}

//删除功能
func (this *customerView) delete() {
	var num int
	var key string
	fmt.Println("请输入要删除的客户编号")
	fmt.Scanln(&num)
	fmt.Println(num)
	fmt.Printf("确定要删除客户%d吗?y or no", num)
	fmt.Scanln(&key)
	if key != "y" && key != "n" {
		fmt.Println("非法输入")
		return
	}
	//查询id存在
	if this.cs.Delete(this.cs.FindById(num)) {
		fmt.Println("删除客户")
	} else {
		fmt.Println("不存在客户")
	}
}

//退出
func (this *customerView) exit() {
	for {
		var key string
		fmt.Println("确定要退出吗 y or n")
		fmt.Scanln(&key)
		if key != "y" && key != "n" {
			fmt.Println("非法输入")
		} else {
			if key == "y" {
				this.loop = false
			}
			break
		}
	}
}

//更新
func (this *customerView) update() {
	var id int
	fmt.Println("请选择客户id")
	fmt.Scanln(&id)
	if this.cs.Info(id) {
		fmt.Println("修改成功")
	}
}

//view 显示主菜单
func (this *customerView) menu() {
	for {
		fmt.Println("----------客户信息管理系统---------")
		fmt.Println("            1 添加客户")
		fmt.Println("            2 修改客户")
		fmt.Println("            3 删除客户")
		fmt.Println("            4 客户列表")
		fmt.Println("            5 退出")
		fmt.Println("请选择1-5进行对应操作\n")
		var key int
		fmt.Scanln(&key)
		switch key {
		case 1:
			this.insert()
		case 2:
			this.update()
		case 3:
			this.delete()
		case 4:
			this.index()
		case 5:
			this.exit()
		default:
			fmt.Println("非法输入")
		}
		if !this.loop {
			fmt.Println("已退出")
			break
		}

	}
}
func main() {
	fmt.Println("ok")
	var c customerView
	c = customerView{loop: true}
	c.cs = service.NewCustomerService()
	c.menu()
}
