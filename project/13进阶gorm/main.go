package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 进阶gorm 题目1：模型定义
// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
// 要求 ：
// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	ID   int
	Name string
	Num  int
	Post []Post
}

type Post struct {
	ID      int
	Name    string
	UserId  int
	State   string
	Comment []Comment
}

type Comment struct {
	ID     int
	Name   string
	PostID int
}

// 进阶gorm 题目2：关联查询
// 基于上述博客系统的模型定义。
// 要求 ：
// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

// 进阶gorm 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	result := tx.Debug().Model(&User{}).Where(p.UserId).Update("num", gorm.Expr("num + ?", 1))

	return result.Error
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	result := tx.Debug().Where("post_id = ?", c.PostID).First(&Comment{})
	if result.RowsAffected == 0 {
		tx.Debug().Model(&Post{}).Where(c.PostID).Update("state", "无评论")
	}
	return
}

func main() {
	// 创建db连接
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("===============================", db, err)

	// 进阶gorm 题目1：模型定义
	// 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
	// 要求 ：
	// 使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

	// 创建表结构
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})

	user := User{Name: "jhc", Num: 100, Post: []Post{{Name: "每天锻炼", Comment: []Comment{{Name: "写的不错"}}}}}
	db.Debug().Create(&user)

	user = User{Name: "yqq", Num: 99, Post: []Post{{Name: "如何月入百万", Comment: []Comment{{Name: "我真的赚到了"}}}}}
	db.Debug().Create(&user)

	// 进阶gorm 题目2：关联查询
	// 基于上述博客系统的模型定义。
	// 要求 ：
	// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

	// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	db.Debug().Preload("Post").Preload("Post.Comment").Find(&User{}, 1)

	// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	comment := Comment{}
	db.Debug().Select("post_id").Group("post_id").Order("count(*) desc").Limit(1).Find(&comment)
	db.Debug().Where("id = ?", comment.PostID).Find(&Post{})

	// 进阶gorm 题目3：钩子函数
	// 继续使用博客系统的模型。
	// 要求 ：
	// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

	// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	post := Post{Name: "开发工程师手册", UserId: 1}
	db.Debug().Create(&post)

	// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	com := Comment{ID: 1, PostID: 1}
	db.Debug().Delete(&com)
}
