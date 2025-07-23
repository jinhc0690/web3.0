package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SQL语句练习 题目1：基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

type Student struct {
	ID    int
	Name  string
	Age   int
	Grade string
}

// SQL语句练习 题目2：事务语句
// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
type Account struct {
	ID      int
	Balance float64
}

type Transaction struct {
	ID            int
	FromAccountId int
	ToAccountId   int
	Amount        float64
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	account := &Account{ID: t.FromAccountId}
	tx.Debug().Select("Balance").First(account)
	if account.Balance-t.Amount < 0 {
		return errors.New("余额不足")
	}
	return
}

func main() {
	// 创建db连接
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("===============================", db, err)

	// SQL语句练习 题目1：基本CRUD操作
	// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
	// 要求 ：
	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

	// 创建表结构
	db.AutoMigrate(&Student{})

	// 向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	db.Debug().Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})

	// 查询 students 表中所有年龄大于 18 岁的学生信息。
	db.Debug().Where("age > ?", 18).Find(&Student{})

	// 将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	db.Debug().Model(&Student{}).Where("Name", "张三").Update("Grade", "四年级")

	// 删除 students 表中年龄小于 15 岁的学生记录。
	db.Debug().Where("age < ?", 15).Delete(&Student{})

	// SQL语句练习 题目2：事务语句
	// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
	// 要求 ：
	// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

	// 创建表结构
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})

	accountA := Account{Balance: 50}
	accountB := Account{Balance: 200}
	db.Create(&accountA)
	db.Create(&accountB)

	db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		transaction := Transaction{FromAccountId: accountA.ID, ToAccountId: accountB.ID, Amount: 100}
		if err := tx.Debug().Create(&transaction).Error; err != nil {
			fmt.Println(err)
			// 返回任何错误都会回滚事务
			return err
		}

		db.Debug().Model(accountA).Update("Balance", gorm.Expr("Balance - ?", transaction.Amount))
		db.Debug().Model(accountB).Update("Balance", gorm.Expr("Balance + ?", transaction.Amount))

		// 返回 nil 提交事务
		return nil
	})

}
