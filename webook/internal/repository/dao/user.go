package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound   = gorm.ErrRecordNotFound
)

var _ UserDao = (*GormUserDao)(nil)

type UserDao interface {
	FindById(ctx context.Context, id int64) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
	Insert(ctx context.Context, u User) error
	FindByPhone(ctx context.Context, phone string) (User, error)
}

type GormUserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &GormUserDao{
		db: db,
	}
}

func (dao *GormUserDao) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).First(&u, "id = ?", id).Error
	return u, err
}

func (dao *GormUserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u,"email = ?").Error
	return u, err
}

func (dao *GormUserDao) Insert(ctx context.Context, u User) error {
	//存毫秒数
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			return ErrDuplicateEmail
		}
	}
	return err
}

func (dao *GormUserDao) FindByPhone(ctx context.Context, phone string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("phone = ?", phone).First(&u).Error
	return u, err
}

// User 在Dao中直接对应数据库表结构
// entity,model,po
type User struct {
	Id       int64          `gorm:"primaryKey,autoIncrement"`
	Email    sql.NullString `gorm:"unique"`
	Password string
	Nickname string
	Phone    sql.NullString `gorm:"unique"`
	Birthday string
	Profile  string
	//创建时间，毫秒数
	Ctime int64
	//更新时间，毫秒数
	Utime int64
}
