package controllers

import (
	"0922duoproject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//1.获取name`age`sex
	//2.做数据对比
	//3.根据对比结果进行判断处理
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "h14770983071@163.com"
	c.TplName = "index.tpl"
}
//编写一个post方法,
/*func (c *MainController) Post(){
	name := c.Ctx.Request.FormValue("name")
	age := c.Ctx.Request.FormValue("age")
	sex := c.Ctx.Request.FormValue("sex")
	fmt.Println(name,age,sex)
	c.Ctx.WriteString("返回数据")
}*/
func (c *MainController) Post()  {
	var person models.Personer
	dataBytes,err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err !=nil{
		c.Ctx.WriteString("数据解析失败，请重试")
	}
	err = json.Unmarshal(dataBytes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败")
		return
	}
	fmt.Println("姓名",person.Name)
	fmt.Println("生日",person.Birthday)
	fmt.Println("地址",person.Address)
	fmt.Println("昵称",person.Nick)
	c.Ctx.WriteString("数据解析成功")
}