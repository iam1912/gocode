package service
//该包是功能的方面的代码，只处理功能不包含界面上的显示
import (
   "github.com/XIE_3/Student"
   "os"
   "fmt"
   "github.com/jinzhu/gorm"
  _"github.com/jinzhu/gorm/dialects/mysql"
)
var (
   DB *gorm.DB
   Err error
)

type StudentService struct {
     Students []model.Student
}

func NewStudentService() *StudentService {//初始化学生变量并打开数据库
     studentService := &StudentService{}
     Opendb()
     return studentService
}
func Opendb()  {
     DB, Err = gorm.Open("mysql",
     "root:15219331409@/stu?charset=utf8&parseTime=True&loc=Local")
     if Err != nil {
         fmt.Println(Err)
     }
     DB.AutoMigrate(&model.Student{})
}
func Shutdb() {
     DB.Close()
}
func (this *StudentService) List() []model.Student {//返回存放学习信息的切片
     //Opendb()
     DB.Find(&this.Students)
     return  this.Students
}
func (this *StudentService) Add(student model.Student) bool {//添加学生
     //Opendb()
     DB.Create(&student)
     return true
}
func (this *StudentService) IndexFind(id int) bool {//查找时先查找学生编号是否存在
     index := -1
     for i:=0;i<len(this.Students);i++ {
         if this.Students[i].ID == id {
              index = i
         }
     }
     if index != -1 {
        return true
     } else {
        return false
     }
}
func (this *StudentService) Deletes(id int) bool {//删除学生信息
     //Opendb()
     DB.Where("ID = ?",id).Delete(model.Student{})
     return true
}
func (this *StudentService) Exit() {
     os.Exit(0)
}
func (this *StudentService) Modify(id int,args...interface{}) bool {
//修改成绩，专业，备注
     //Opendb()
     DB.Table("students").Where("ID = ?",id).Update("socre",args[0].(int))
     DB.Table("students").Where("ID = ?",id).Updates(model.Student{Major:args[1].(string),Note:args[2].(string)})
     return true
}
func (this *StudentService) Sorts(i int) []model.Student {//排序
     //Opendb()
     if i == 0 {
       DB.Order("socre").Find(&this.Students)
     } else {
       DB.Order("socre desc").Find(&this.Students)
     }
     return this.Students
}
func (this *StudentService) Find(id int) model.Student {//查找
     //Opendb()
     var student model.Student
     DB.Where("ID = ?",id).Find(&student)
     return student
}
