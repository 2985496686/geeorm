package test

import (
	"GeeORM"
	"GeeORM/session"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

type User2 struct {
	Id       int32 `geeorm:"primary key AUTO_INCREMENT"`
	Username string
	Password string
	Gender   rune
	Weight   int32
	Height   int32
}

func Test(t *testing.T) {
	engine, _ := GeeORM.NewEngine("mysql", "root:111111@tcp(localhost:3306)/study2")
	session := engine.NewSession()
	var id int
	var username, password, gender, email string
	rows := session.Raw("select *from t_user where id = ?", 3).QueryRow()
	rows.Scan(&id, &username, &password, &gender, &email)
	fmt.Println(id, username, password, gender, email)
	//session.Model(&User2{}).CreateTable()
	session.Model(&User2{}).DropTable()
	fmt.Println(session.Model(&User2{}).HasTable())

}

func (u *User2) BeforeQuery(s *session.Session) error {
	fmt.Println("开始查找！")
	return nil
}

func (u *User2) AfterQuery(s *session.Session) error {
	u.Password = "********"
	return nil
}

func TestHook(t *testing.T) {
	engine, _ := GeeORM.NewEngine("mysql", "root:111111@tcp(localhost:3306)/study2")
	session := engine.NewSession()
	var users []User2
	session.Find(&users)
	for _, i := range users {
		fmt.Println(i)
	}
}
