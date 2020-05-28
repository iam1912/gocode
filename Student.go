package model
import "fmt"

type Student struct {
     ID int //`gorm:"type:char(6);not null"`
     Name string //`gorm:"type:char(8);not null"`
     Major  string //`gorm:"type:char(10);not null"`
     Sex int //`gorm:"type:tinyint(1);not null;default:1"`
     Birthday string //`gorm:"not null"`
     Socre int //`gorm:"type:tinyint(1),not null"`
     Note string //`gorm:"type:tinytext;null"`
}
func Newstudent(id int, name string, major string,sex int,
                 birthday string, socre int, note string) Student {
     return Student{
        ID : id,
        Name : name,
        Major : major,
        Sex : sex,
        Birthday : birthday,
        Socre : socre,
        Note : note,
     }
}
func (this *Student) Show() string {
     info := fmt.Sprintf("%d\t%s\t%s\t%d\t%s\t%d\t%s",this.ID, this.Name,
             this.Major, this.Sex, this.Birthday, this.Socre, this.Note)
     return info
}
