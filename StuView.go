package view
//该包的作用是将显示的界面显示出来，其中的方法主要是来处理界面，而不是真正功能的实际代码
import (
    "fmt"
    "github.com/XIE_3/Service"
    "github.com/XIE_3/Student"
    _"net"
    _"io"
)
type studentview struct {
     flag int
     studentservice *service.StudentService
}
func Newstudentview() *studentview{
     return &studentview{
        flag : 0,
     }
}
func (this *studentview) show() {
     student := this.studentservice.List()
     fmt.Println("--------------学生信息列表---------------")
     fmt.Println("学号\t姓名\t专业名\t性别\t出生日期\t总学分\t备注")
     for i:=0;i<len(student);i++ {
         fmt.Println(student[i].Show())
     }
     fmt.Println("-------------学生信息列表完成--------------")
     fmt.Println()
}
func (this *studentview) add() {
     fmt.Println("-------------添加学生信息------------")
     fmt.Printf("ID:")
     id := 0
     fmt.Scanf("%d",&id)
     fmt.Printf("Name:")
     name := ""
     fmt.Scanf("%s",&name)
     fmt.Printf("Major:")
     major := ""
     fmt.Scanf("%s",&major)
     fmt.Printf("Sex:")
     sex := 0
     fmt.Scanf("%d",&sex)
     fmt.Printf("Birthday:")
     birthday := ""
     fmt.Scanf("%s",&birthday)
     fmt.Printf("Socre:")
     socre := 0
     fmt.Scanf("%d",&socre)
     fmt.Printf("Note:")
     note := ""
     fmt.Scanf("%s",&note)
     student := model.Newstudent(id,name,major,sex,birthday,socre,note)
     if this.studentservice.Add(student) {
        fmt.Println("-------------------添加成功---------------")
     } else {
        fmt.Println("-------------------添加失败---------------")
     }
}
func (this *studentview) deletes() {
     fmt.Println("------------------删除学生信息-----------------")
     fmt.Printf("请选择删除学生的编号(-1退出):")
     id := -1
     fmt.Scanf("%d",&id)
     if id == -1 {
        return
     }
     fmt.Println("确认是否删除(y/n):")
     flag := ""
     fmt.Scanf("%s",&flag)
     if flag == "y" {
        if this.studentservice.Deletes(id) {
           fmt.Println("------------------删除完成-----------------")
        } else {
           fmt.Println("----------------- 删除失败-----------------")
        }
     }
}
func (this *studentview) update() {
     fmt.Println("----------------修改学生信息----------------")
     fmt.Printf("请选择待修改客户编号(-1退出):")
     id := -1
     fmt.Scanf("%d",&id)
     if id == -1 {
        return
     }
     fmt.Println("请输入以下修改信息:")
     socre := 0
     major := ""
     note := ""
     fmt.Printf("请输入分数:")
     fmt.Scanf("%d",&socre)
     fmt.Printf("请输入专业名:")
     fmt.Scanf("%s",&major)
     fmt.Printf("请输入如备注:")
     fmt.Scanf("%s",&note)
     if this.studentservice.Modify(id,socre,major,note) {
        fmt.Println("-----------------修改信息成功---------------")
     } else {
        fmt.Println("-----------------修改信息失败---------------")
     }
}
func (this *studentview) sort() {
     var flag int
     fmt.Printf("请输入排序的规则(顺序-1/倒序-0):")
     fmt.Scanf("%d",&flag)
     students := this.studentservice.Sorts(flag)
     fmt.Println("-------------学生信息列表-------------")
     fmt.Println("学号\t姓名\t专业名\t性别\t出生日期\t总学分\t备注")
     for _, val := range students {
          fmt.Println(val.Show())
     }
     fmt.Println("-------------学生信息列表完成----------")
     fmt.Println()
}
func (this *studentview) find() {
     id := 0
     for {
        fmt.Printf("请输入查询的学生的编号(-1退出):")
        fmt.Scanf("%d",&id)
        if id == -1 {
           return
        } else {
          if this.studentservice.IndexFind(id) {
             student := this.studentservice.Find(id)
             fmt.Println(student.Show())
          } else {
             fmt.Println("查询的学生信息不存在")
          }
        }
     }
}
func (this *studentview) exit() {
     fmt.Println("是否退出程序请选择y/n")
     flag := ""
     for {
       fmt.Scanf("%s",&flag)
       if flag == "y" {
          fmt.Println("你已退出程序")
          service.Shutdb()
          this.studentservice.Exit()
       } else if flag == "n" {
          return
       } else {
          fmt.Println("输入字符有误，请重新输入y/n")
       }
     }
}
func (this *studentview) MainMenu() {
     this.studentservice = service.NewStudentService()
     for {
       fmt.Println("-------------学生信息管理软件------------")
       fmt.Println("-------------1.添加学生信息")
       fmt.Println("-------------2.修改学生信息")
       fmt.Println("-------------3.删除学生信息")
       fmt.Println("-------------4.学生信息列表")
       fmt.Println("-------------5.排        序")
       fmt.Println("-------------6.查        询")
       fmt.Println("-------------7.退        出")
       fmt.Printf("请选择(1-7):")
       fmt.Scanf("%d",&this.flag)
       switch this.flag {
           case 1 :
              this.add()
           case 2 :
              this.update()
           case 3 :
              this.deletes()
           case 4 :
              this.show()
           case 5 :
              this.sort()
           case 6 :
              this.find()
           case 7 :
              this.exit()
           default :
              fmt.Println("你输入的字符有误，请重新输入")
       }
     }
}
