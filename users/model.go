package users

import (
	//"database/sql"
	"crypto/rand"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//Users user info
type Users struct {
	ID       int    `json:"user-id"         xorm:"id"`
	UID      string `json:"user-uid"        xorm:"uid"`
	Name     string `json:"user-name"       xorm:"name"`
	Email    string `json:"user-email"      xorm:"email"`
	Password string `json:"user-password"   xorm:"password"`
}

func createUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func getUser(e string, p string) {
	//engine, _ := xorm.NewEngine("mysql", "root:root@tcp(serversideforrxpractice_db_1:3306)/sample_db")
	//engine, _ := xorm.NewEngine("mysql", "root:root@tcp(db:3306)/sample_db")
	engine, _ := xorm.NewEngine("mysql", "root:@/sample_db")
	user := Users{Email: e, Password: p}
	user.UID = createUID()
	_, err := engine.Insert(&user)
	if err != nil {
		fmt.Printf("can not insert")
		log.Fatal(err)
		return
	}
	return
}
