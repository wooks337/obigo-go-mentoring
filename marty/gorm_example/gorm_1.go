package gorm_example

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	ID             uuid.UUID `gorm:"column:id;type:char(36);primary_key"` // column 은 name 규칙과 다르게 db column 을 구성 하면 명시해 주기도 한다.
	UID            uuid.UUID `gorm:"column:uid;not null;index"`
	Email          string    `gorm:"column:email;not null;index"`
	ApproveMessage string    `gorm:"column:approve_message;type(varchar(4000))"`
}

const loyalUser string = "marty@obigo.com"

func main() {
	dsn := "root:root@(10.28.3.180:3307)/SchoolDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// create
	err = Create(db)
	if err != nil {
		fmt.Errorf("Create error :  %v", err)
	}

	// get user
	err = GetUser(db)
	if err != nil {
		fmt.Errorf("GetUser error :  %v", err)
	}

}

func Create(db *gorm.DB) (err error) {

	user := User{
		ID:             uuid.New(),
		UID:            uuid.New(),
		Email:          "marty@obigo.com",
		ApproveMessage: "good",
	}

	/* gorm 에서 hook 처리 => springboot 에서 @Transactional 에서 해주던 기능
	create / update / delete 는 이러한 hook 처리를 하지 않으면 DB transaction error 발생시
	DB  에 Lock 걸리는 현상이 발생하곤 한다.
	*/
	tx := db.Begin()
	result := db.Create(&user)
	if result.Error != nil {
		tx.Rollback()
		return errors.New("Invalid data !!")
	}
	tx.Commit()

	return nil
}

func GetUser(db *gorm.DB) (err error) {
	query := db.Table("user").Select("id, email")

	// 검색 필터 같은 기능 이용시 아래 와 같이 subQuery를 이용
	if loyalUser != "" {
		query = query.Where("email like ?", "%"+loyalUser+"%")
	}

	if query.Error != nil {
		return query.Error
	}

	return nil
}
