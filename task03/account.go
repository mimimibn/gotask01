package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
*
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
建表语句
CREATE TABLE `accounts` (

	`id` int(11) NOT NULL,
	`balance` int(11) DEFAULT NULL,
	PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
CREATE TABLE `transactions` (

	`id` int(11) NOT NULL,
	`from_account_id` int(11) DEFAULT NULL,
	`to_account_id` int(11) DEFAULT NULL,
	`amount` int(11) DEFAULT NULL,
	PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/
type Accounts struct {
	Id      int32 `gorm:"column:id;type:int(11);primaryKey;not null;" json:"id"`
	Balance int32 `gorm:"column:balance;type:int(11);default:NULL;" json:"balance"`
}
type Transactions struct {
	Id            int32 `gorm:"column:id;type:int(11);primaryKey;not null;" json:"id"`
	FromAccountId int32 `gorm:"column:from_account_id;type:int(11);default:NULL;" json:"from_account_id"`
	ToAccountId   int32 `gorm:"column:to_account_id;type:int(11);default:NULL;" json:"to_account_id"`
	Amount        int32 `gorm:"column:amount;type:int(11);default:NULL;" json:"amount"`
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		fmt.Println("链接数据库失败", err)
		return
	}
	//账户1 设置为101元
	//账户2 设置为10元
	accountss := []Accounts{{Balance: 101}, {Balance: 10}}
	db.CreateInBatches(accountss, len(accountss))
	accounts1 := accountss[0]
	accounts2 := accountss[1]
	//设置转账金额
	var num int32 = 100
	txerr := db.Transaction(func(tx *gorm.DB) error {
		findAccount1Err := tx.Where(" id = ?", accounts1.Id).First(&accounts1).Error
		if findAccount1Err != nil {
			return findAccount1Err
		}
		if accounts1.Balance < num {
			return errors.New("A账户金额不足")
		}
		updateAccount2Err := tx.Model(&Accounts{}).Where("id = ?", accounts2.Id).UpdateColumn("Balance", gorm.Expr("Balance + ?", num)).Error
		if updateAccount2Err != nil {
			return updateAccount2Err
		}
		updateAccount1Err := tx.Model(&Accounts{}).Where("id = ?", accounts1.Id).UpdateColumn("Balance", gorm.Expr("Balance - ?", num)).Error
		if updateAccount1Err != nil {
			return updateAccount1Err
		}
		transactions := Transactions{FromAccountId: accounts1.Id, ToAccountId: accounts2.Id, Amount: num}
		saveTransactionsErr := tx.Save(&transactions).Error
		if saveTransactionsErr != nil {
			return saveTransactionsErr
		}
		return nil
	})
	if txerr != nil {
		fmt.Println("事务回滚", txerr)
	} else {
		fmt.Println("转账成功")
	}
}
