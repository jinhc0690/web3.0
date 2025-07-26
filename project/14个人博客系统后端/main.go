package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null;" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"unique;not null" json:"email"`
	Post     []Post
	Comment  []Comment
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	UserId  uint   `json:"userid"`
	Comment []Comment
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	UserId  uint
	PostId  uint ` json:"postid"`
}

type Token struct {
	Token string `header:"token"`
}

var secretKey = "grbkxthd"

// 后续的需要认证的接口需要验证该 JWT 的有效性
func CheckToken(c *gin.Context) uint {
	fmt.Println("CheckToken start")
	var tokenInfo Token
	if err := c.ShouldBindHeader(&tokenInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}
	fmt.Println("获取token", tokenInfo)
	tokenstring := tokenInfo.Token
	claim := jwt.MapClaims{}
	// 解析 Token
	_, err := jwt.ParseWithClaims(
		tokenstring,
		&claim, // 使用你的自定义 Claims 结构
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil // 返回签名密钥
		},
	)
	fmt.Println("token还原", claim)
	exp := claim["exp"]

	var expUnix int64
	switch v := exp.(type) {
	case float64:
		expUnix = int64(v) // 直接转换 float64 -> int64
	case int64:
		expUnix = v // 已经是目标类型
	default:
		// 处理无效类型
	}

	// 现在 expUnix 是 int64 类型，可直接与 time.Now().Unix() 比较
	currentTime := time.Now().Unix()
	if expUnix < currentTime {
		fmt.Println("Token 已过期")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token 已过期"})
		return 0
	}

	if err != nil {
		// 错误处理
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 0
	}

	id := claim["id"]
	var result uint
	switch v := id.(type) {
	case float64:
		result = uint(v) // 直接转换 float64 -> uint
	case uint:
		result = v // 已经是目标类型
	default:
		// 处理无效类型
	}

	fmt.Println("CheckToken end", result)
	return result
}

// func CheckToken() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("CheckToken before")
// 		var tokenInfo Token
// 		if err := c.ShouldBindHeader(&tokenInfo); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		fmt.Println("获取token", tokenInfo)
// 		tokenstring := tokenInfo.Token
// 		claim := jwt.MapClaims{}
// 		// 解析 Token
// 		_, err := jwt.ParseWithClaims(
// 			tokenstring,
// 			&claim, // 使用你的自定义 Claims 结构
// 			func(token *jwt.Token) (interface{}, error) {
// 				return []byte(secretKey), nil // 返回签名密钥
// 			},
// 		)
// 		fmt.Println("token还原", claim)
// 		exp := claim["exp"]

// 		var expUnix int64
// 		switch v := exp.(type) {
// 		case float64:
// 			expUnix = int64(v) // 直接转换 float64 -> int64
// 		case int64:
// 			expUnix = v // 已经是目标类型
// 		default:
// 			// 处理无效类型
// 			return
// 		}

// 		// 现在 expUnix 是 int64 类型，可直接与 time.Now().Unix() 比较
// 		currentTime := time.Now().Add(time.Hour * 24).Unix()
// 		if expUnix < currentTime {
// 			fmt.Println("Token 已过期")
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Token 已过期"})
// 			return
// 		} else if err != nil {
// 			// 错误处理
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		} else {
// 			c.Next()
// 			fmt.Println("CheckToken after")
// 		}
// 	}
// }

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("===============================", db, err)
	if err != nil {
		panic("failed to connect database")
	}

	// 创建表结构
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})

	router := gin.Default()

	// Register
	router.POST("/register", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("注册", user)
		Password := user.Password
		// 密码为空检查
		if Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "密码不可为空！"})
			return
		}
		// 加密密码
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = string(hashedPassword)

		if err := db.Debug().Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	})

	// Login
	router.POST("/login", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("登录", user)

		var storedUser User
		if err := db.Debug().Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
			return
		}

		// 验证密码
		if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		// 生成 JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       storedUser.ID,
			"username": storedUser.Username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	})

	// 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
	router.POST("/addPost", func(c *gin.Context) {
		UserId := CheckToken(c)
		if UserId <= 0 {
			return
		}
		var post Post = Post{UserId: UserId}
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("文章的创建", post)

		Title := post.Title
		// 标题为空检查
		if Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "标题不可为空！"})
			return
		}
		Content := post.Content
		// 内容为空检查
		if Content == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "内容不可为空！"})
			return
		}
		if err := db.Debug().Create(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "文章创建成功"})
	})

	// 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
	router.GET("/queryPostAll", func(c *gin.Context) {
		UserId := CheckToken(c)
		if UserId <= 0 {
			return
		}
		fmt.Println("文章的读取所有列表")
		postRows, err := db.Debug().Model(&Post{}).Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to queryall post"})
			return
		}

		defer postRows.Close()

		var post Post
		posts := []Post{}
		var flag bool
		for postRows.Next() {
			// ScanRows 将一行扫描至 post
			db.ScanRows(postRows, &post)

			posts = append(posts, post)

			flag = true
		}
		if !flag {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "文章暂无数据"})
			return
		}
		fmt.Println(posts)

		c.JSON(http.StatusOK, gin.H{"success": "文章所有读取成功", "post": posts})
	})

	router.GET("/queryPostSingle/:id", func(c *gin.Context) {
		UserId := CheckToken(c)
		if UserId <= 0 {
			return
		}
		id := c.Param("id")
		fmt.Println("文章的读取单个", id)
		postRows, err := db.Debug().Model(&Post{}).Where(id).Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to querysingle post"})
			return
		}

		defer postRows.Close()

		var post Post
		var flag bool
		for postRows.Next() {
			// ScanRows 将一行扫描至 post
			db.ScanRows(postRows, &post)
			flag = true
		}
		if !flag {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询不到对应文章的详细信息"})
			return
		}
		fmt.Println(post)

		c.JSON(http.StatusOK, gin.H{"success": "文章单个读取成功", "post": post})
	})

	// 实现文章的更新功能，只有文章的作者才能更新自己的文章。
	router.PUT("/updatePost", func(c *gin.Context) {
		UserId := CheckToken(c)
		if UserId <= 0 {
			return
		}
		var post Post = Post{UserId: UserId}
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("文章的修改", post)

		Title := post.Title
		// 标题为空检查
		if Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "标题不可为空！"})
			return
		}
		Content := post.Content
		// 内容为空检查
		if Content == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "内容不可为空！"})
			return
		}
		result := db.Debug().Where(&post, "ID", "UserId").Find(&Post{})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query post"})
			return
		}
		if result.RowsAffected == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "不是对应文章的作者"})
			return
		}
		if err := db.Debug().Updates(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "文章修改成功"})
	})

	// 实现文章的删除功能，只有文章的作者才能删除自己的文章。
	router.DELETE("/deletePost/:id", func(c *gin.Context) {
		UserId := CheckToken(c)
		if UserId <= 0 {
			return
		}
		id := c.Param("id")
		fmt.Println("文章的删除", id)
		result := db.Debug().Where("ID = ? and user_id = ?", id, UserId).Find(&Post{})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query post"})
			return
		}
		if result.RowsAffected == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "不是对应文章的作者"})
			return
		}
		if err := db.Debug().Delete(&Post{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "文章删除成功"})
	})

	// 实现评论的创建功能，已认证的用户可以对文章发表评论。
	router.POST("/addComment", func(c *gin.Context) {
		UserId := CheckToken(c)
		if UserId <= 0 {
			return
		}
		var comment Comment = Comment{UserId: UserId}
		if err := c.ShouldBindJSON(&comment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("评论的创建", comment)

		Content := comment.Content
		// 内容为空检查
		if Content == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "内容不可为空！"})
			return
		}
		if err := db.Debug().Create(&comment).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": "评论创建成功"})

	})

	// 实现评论的读取功能，支持获取某篇文章的所有评论列表。
	router.GET("/queryComment/:id", func(c *gin.Context) {
		UserId := CheckToken(c)
		if UserId <= 0 {
			return
		}
		id := c.Param("id")
		fmt.Println("评论的读取", id)
		commentRows, err := db.Debug().Model(&Comment{}).Where("post_id = ?", id).Rows()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query comment"})
			return
		}

		defer commentRows.Close()

		var comment Comment
		comments := []Comment{}
		var flag bool
		for commentRows.Next() {
			// ScanRows 将一行扫描至 comment
			db.ScanRows(commentRows, &comment)

			comments = append(comments, comment)

			flag = true
		}
		if !flag {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询不到对应文章的评论信息"})
			return
		}

		fmt.Println(comments)

		c.JSON(http.StatusOK, gin.H{"success": "评论读取成功", "comment": comments})
	})

	erro := router.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if erro != nil {
		panic(erro)
	}

}
