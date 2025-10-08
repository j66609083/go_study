package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 用户
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Password  string
	Email     string
	PostCount uint // 拥有的文章数量
	Posts     []Post
}

// 文章
type Post struct {
	ID            uint
	Title         string // 文章标题
	Content       string // 文章内容
	UserId        uint   // 作者id
	CommentStatus string // 评论状态：有评论，无评论
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Comments      []Comment
}

// 评论
type Comment struct {
	ID        uint
	Content   string
	UserId    uint // 评论人id
	PostId    uint // 评论的文章id
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	var postCount int64
	tx.Model(&Post{UserId: p.UserId}).Count(&postCount)
	tx.Model(&User{}).Where("id = ?", p.UserId).Update("post_count", postCount)
	// tx.Model(&User{}).Where("id = ?", p.UserId).Update("post_count", gorm.Expr("post_count + 1"))
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("AfterDelete Comment: ", c)
	var commentCount int64
	tx.Model(&Comment{PostId: c.PostId}).Count(&commentCount)
	if commentCount == 0 {
		tx.Model(&Post{}).Where("id = ?", c.PostId).Update("comment_status", "无评论")
	} else {
		tx.Model(&Post{}).Where("id = ?", c.PostId).Update("comment_status", "有评论")
	}
	return nil
}
func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&Post{}).Where("id = ?", c.PostId).Update("comment_status", "有评论")
	return nil
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	// 题目1：模型定义
	/*
		使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
		编写Go代码，使用Gorm创建这些模型对应的数据库表。
	*/
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Post{}, &Comment{})

	// 题目2：关联查询
	// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	var someUser User
	db.Model(User{ID: 1}).Preload("Posts").Preload("Posts.Comments").First(&someUser)
	log.Println(someUser)

	// 编写Go代码，使用Gorm查询评论数量最多的文章信息。
	type CommentCount struct {
		PostId uint
		Total  int
	}
	var countResult CommentCount
	db.Model(&Comment{}).
		Select("post_id, count(1) as total").
		Group("post_id").
		Order("total desc").
		Find(&countResult)
	log.Println(countResult)

	var maxCommentPost Post
	db.Model(&Post{ID: countResult.PostId}).First(&maxCommentPost)
	log.Println("评论数量最多的文章信息: ", maxCommentPost)

	// 测试钩子函数
	post1 := Post{Title: "标题测试1", Content: "内容测试1", UserId: 1}
	result := db.Create(&post1)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	db.Delete(&Comment{ID: 2, PostId: 1})
}
