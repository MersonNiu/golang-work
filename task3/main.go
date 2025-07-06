package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var globalDB *gorm.DB

func InitDB() *gorm.DB {
	if globalDB == nil {
		db, err := gorm.Open(sqlite.Open("gorm.db?_busy_timeout=5000"), &gorm.Config{})
		if err != nil {
			panic("连接失败：" + err.Error())
		}
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(1)
		sqlDB.SetMaxOpenConns(1)
		sqlDB.SetConnMaxLifetime(time.Hour)

		globalDB = db
	}
	return globalDB
}

type SchemaVersion struct {
	Version string `gorm:"primaryKey"`
}

func runMigrationOnce(db *gorm.DB) {
	currentVer := "v1.0"
	var lastVer SchemaVersion
	if db.Migrator().HasTable(&SchemaVersion{}) {
		db.First(&lastVer)
	} else {
		err := db.AutoMigrate(&SchemaVersion{})
		if err != nil {
			panic(err)
		}
	}
	if lastVer.Version != currentVer {
		err := db.AutoMigrate(&Student{})
		if err != nil {
			panic(err)
		}
		db.Save(&SchemaVersion{Version: currentVer})
	}
}

// SQL语句练习
// 题目1
type Student struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt *time.Time     `gorm:"column:created_at;autoCreateTime"` // 指针类型接收 NULL
	UpdatedAt *time.Time     `gorm:"column:updated_at;autoUpdateTime"` // 指针类型接收 NULL
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;softDelete:flag"`
	Name      string         `gorm:"column:name"`
	Age       int            `gorm:"column:age"`
	Grade     string         `gorm:"column:grade"`
}

func DataInit() {
	db := InitDB()
	students := []Student{
		{Name: "小明", Age: 12, Grade: "三年级"},
		{Name: "小美", Age: 13, Grade: "三年级"},
		{Name: "小芳", Age: 19, Grade: "五年级"},
		{Name: "小胖", Age: 12, Grade: "三年级"},
		{Name: "小宋", Age: 20, Grade: "三年级"},
		{Name: "小老弟", Age: 12, Grade: "二年级"},
	}
	result := db.Create(&students)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			log.Println("主键冲突")
		} else {
			log.Fatal("数据库错误:", result.Error)
		}
	}
}

// 题目2
type Account struct {
	ID          uint    `gorm:"primarykey"`
	AccountName string  `gorm:"column:accountname;default:'匿名账户'"`
	Balance     float64 `gorm:"column:balance;default:0.00"`
}
type Transaction struct {
	ID              uint    `gorm:"primarykey"`
	From_account_id int     `gorm:"column:from_account_id;default:0"`
	To_account_id   int     `gorm:"column:to_account_id;denfault:0"`
	Amount          float64 `gorm:"column:amount;default:0.00"`
}

func transacton(actname1, actname2 string) {
	db := InitDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		//查询账户
		var account1, account2 Account
		if err := tx.Where("accountname = ?", actname1).First(&account1).Error; err != nil {
			log.Fatalf("failed to find account1:%v", err)
		}
		if err := tx.Where("accountname = ?", actname2).First(&account2).Error; err != nil {
			log.Fatalf("falied to find account2:%v", err)
		}
		//账户余额与转账金额比较
		if account1.Balance > 100 {
			if err := tx.Model(&Account{}).Where("accountname=?", actname1).Update("balance", account1.Balance-100).Error; err != nil {
				return fmt.Errorf("更新账户A失败:%w", err)
			}
			if err := tx.Model(&Account{}).Where("accountname=?", actname2).Update("balance", account2.Balance+100).Error; err != nil {
				return fmt.Errorf("更新账户B失败:%w", err)
			}
			tranction := Transaction{From_account_id: int(account1.ID), To_account_id: int(account2.ID), Amount: 100}
			if err := tx.Create(&tranction).Error; err != nil {
				return fmt.Errorf("创建交易记录失败:%w", err)
			}
		} else {
			return fmt.Errorf("账户余额不足，主动回滚")
		}
		return nil
	})
	if err != nil {
		fmt.Println("事务失败：", err)
	}
}

// SQLX入门
func InitsqlxDB() *sqlx.DB {
	dsn := "file:test.db?cache=shared&mode=rwc"
	db, err := sqlx.Connect("sqlite3", dsn)
	if err != nil {
		log.Fatalf("数据库连接失败：%v", err)
		return nil
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(1 * time.Hour)
	return db
}

type Employee struct {
	Id         int     `db:"id" json:"id"`
	Name       string  `db:"name" json:"name"`
	Department string  `db:"department" json:"department"`
	Salary     float64 `db:"salary" json:"salary"`
}
type Book struct {
	Id     int     `db:"id" json:"id"`
	Title  string  `db:"title" json:"title"`
	Author string  `db:"author" json:"author"`
	Price  float64 `db:"price" json:"price"`
}

func Creattables() error {
	db := InitsqlxDB()
	tables := map[string]string{
		"employees": `
		CREATE TABLE IF NOT EXISTS employees(
			id INT PRIMARY KEY,
			name VARCHAR(50),
			department VARCHAR(10),
			salary DECIMAL(10,2)
		)`,
		"books": `
		CREATE TABLE IF NOT EXISTS books(
			id INT PRIMARY KEY,
			title VARCHAR(20),
			author VARCHAR(50),
			price DECIMAL(10,2)
		)`}
	for name, sql := range tables {
		if _, err := db.Exec(sql); err != nil {
			return fmt.Errorf("建表%s失败:%v", name, err)
		}
	}
	return nil
}
func Createlines() error {
	db := InitsqlxDB()
	// emps := []Employee{{
	// 	Id: 1, Name: "小美", Department: "人事部", Salary: 2000,
	// }, {
	// 	Id: 2, Name: "小明", Department: "技术部", Salary: 5000,
	// }, {
	// 	Id: 3, Name: "小军", Department: "组织部", Salary: 3000,
	// }}
	bks := []Book{{
		Id: 1, Title: "天空", Author: "人事部", Price: 30,
	}, {
		Id: 2, Title: "大地", Author: "人事部", Price: 100,
	}, {
		Id: 3, Title: "人群", Author: "人事部", Price: 80,
	}}
	// _, err := db.NamedExec(`INSERT INTO employees(id,name,department,salary)VALUES(:id,:name,:department,:salary)`, emps)
	// return err
	_, err := db.NamedExec(`INSERT INTO books(id,title,author,price)VALUES(:id,:title,:author,:price)`, bks)
	return err

}
func insertone() {
	db := InitsqlxDB()
	_, err := db.Exec(`INSERT INTO employees(id,name,department,salary)VALUES(?,?,?,?)`, 4, "阿龙", "技术部", 4500)
	if err != nil {
		fmt.Printf("插入数据失败: %v", err)
	}
}

// 题目1
func jishubu() {
	db := InitsqlxDB()
	var empss []Employee
	err := db.Select(&empss, `SELECT * FROM employees WHERE department = ? `, "技术部")
	if err != nil {
		log.Fatal(err)
	}
	for _, emp := range empss {
		fmt.Printf("%+v\n", emp)
	}
}
func zuigaogongzi() {
	db := InitsqlxDB()
	var empss []Employee
	err := db.Select(&empss, `SELECT * FROM employees WHERE salary = (SELECT MAX(salary) FROM employees)`)
	if err != nil {
		log.Fatal(err)
	}
	for _, emp := range empss {
		fmt.Printf("%+v\n", emp)
	}
}

// 题目2
func chashujv() {
	db := InitsqlxDB()
	var boks []Book
	err := db.Select(&boks, `SELECT * FROM books WHERE price > 50`)
	if err != nil {
		log.Fatal(err)
	}
	for _, emp := range boks {
		fmt.Printf("%+v\n", emp)
	}
}

// 进阶gorm
type User struct { //用户
	Id       int    `gorm:"primaryKey column:id"`
	Username string `gorm:"column:username"` //用户名字
	Postsum  int    `gorm:"column:postsum"`  //用户文章总数
	PostS    []Post `gorm:"foreignKey:UserID"`
}
type Post struct { //文章
	Id            int       `gorm:"primaryKey" column:"id"`
	UserID        int       `gorm:"column:user_id"`
	Author        string    `gorm:"column:author"`        //作者
	Title         string    `gorm:"column:title"`         //文章名字
	Commentstatus string    `gorm:"column:commentstatus"` //文章评论状态
	Comments      []Comment `gorm:"foreignKey:PostID"`
	User          User      `gorm:"foreignKey:UserID"` // 反向引用User
}
type Comment struct { //评论
	Id          int    `gorm:"primaryKey" column:"id"`
	Title       string `gorm:"column:title"`       //文章名字
	Commentsome string `gorm:"column:commentsome"` //文章评论
	PostID      int    `gorm:"column:post_id"`
	Post        Post   `gorm:"foreignKey:PostID"` // 反向引用Post
}

// 题目1模型定义
func modeldefine(structmodle interface{}, dataslice interface{}) {
	db := InitDB()
	err := db.AutoMigrate(structmodle)
	if err != nil {
		log.Fatalf("表结构迁移失败：%v", err)
	}
	result := db.Create(dataslice)
	if result.Error != nil {
		fmt.Printf("数据插入失败：%v", result.Error)
	}
}
func datainti() {
	users1 := []User{{Username: "老子"}, {Username: "悟空"}, {Username: "薛定谔"}}
	post1 := []Post{{Author: "老子", Title: "《道德经》"}, {Author: "老子", Title: "《骑牛》"},
		{Author: "悟空", Title: "《蟠桃种植法》"}, {Author: "悟空", Title: "《金箍棒108种打法》"}, {Author: "悟空", Title: "《西游回忆录》"}, {Author: "薛定谔", Title: "《猫》"}}
	title1 := []Comment{{Title: "《道德经》", Commentsome: "666"}, {Title: "《道德经》", Commentsome: "999"}, {Title: "《道德经》", Commentsome: "牛"},
		{Title: "《骑牛》", Commentsome: "好好好！"}, {Title: "《骑牛》", Commentsome: "腿已断！！！"}, {Title: "《蟠桃种植法》", Commentsome: "没有种子"}, {Title: "《蟠桃种植法》", Commentsome: "无法复现"},
		{Title: "《蟠桃种植法》", Commentsome: "没有天河水"}, {Title: "《金箍棒108种打法》", Commentsome: "大圣帅帅帅"}, {Title: "《西游回忆录》", Commentsome: "曾经有份真挚的感情"},
		{Title: "《西游回忆录》", Commentsome: "如果再给我一次机会"}, {Title: "《西游回忆录》", Commentsome: "我希望是一万年"}, {Title: "《猫》", Commentsome: "猫：我也不知道自己活着还是死了"}}
	modeldefine(&User{}, users1)
	modeldefine(&Post{}, post1)
	modeldefine(&Comment{}, title1)
}

// 题目2关联查询
func getuUserPostsWithComments(username string) {
	db := InitDB()
	var user User
	err := db.Preload("PostS.Comments").Where("username=?", username).First(&user).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("用户%s的文章:\n", user.Username)
	for _, post := range user.PostS {
		fmt.Printf("文章%s评论状态:%s\n", post.Title, post.Commentstatus)
		for _, comment := range post.Comments {
			fmt.Printf(" └ 评论(ID:%d):%s\n", comment.Id, comment.Commentsome)
		}
	}
}
func getMostCommentedPost() {
	db := InitDB()
	var topPosts []Post
	sql := `SELECT P.* 
			FROM posts p  
		JOIN (SELECT post_id,DENSE_RANK() OVER (ORDER BY comment_count DESC) AS rank
	    	FROM ( SELECT post_id,COUNT(*) AS comment_count 
	      			FROM comments GROUP BY post_id) AS comment_counts)AS ranked_posts
		ON p.id = ranked_posts.post_id 
		WHERE ranked_posts.rank = 1;`
	err := db.Raw(sql).Find(&topPosts).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("评论最多的文章(共%v篇):\n", len(topPosts))
	for _, post := range topPosts {
		fmt.Printf("文章ID:%v | 标题：%v | 作者：%v | 评论状态：%v\n",
			post.Id, post.Title, post.Author, post.Commentstatus)
	}
}

// 题目3钩子函数
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	sql := `UPDATE users
	SET postsum = (
		SELECT COUNT(*) FROM posts WHERE posts.user_id = users.id
	)
	WHERE users.id = ?`
	result := tx.Exec(sql, p.UserID)
	if result.Error != nil {
		log.Println("更新用户文章数失败：", result.Error)
		return result.Error
	}
	fmt.Printf("用户 ID %d 的 post 数量已更新。\n", p.UserID)
	return nil
}
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	sql := `SELECT COUNT (*) FROM comments WHERE post_id = ?`
	var comsum int
	tx.Raw(sql, c.PostID).Scan(&comsum)
	if comsum == 0 {
		sql2 := `UPDATE posts SET commentstatus = "无评论" WHERE id = ?`
		result := tx.Exec(sql2, c.PostID)
		if result.Error != nil {
			log.Println("更改评论状态失败：", result.Error)
			return result.Error
		} else {
			log.Println("更改评论状态成功！")
		}
	}
	return nil
}
func main() {

	db := InitDB()
	var come1 Comment
	db.First(&come1, 13)
	db.Delete(&come1)
	// post1 := Post{UserID: 2, Author: "悟空", Title: "《山区园林建设》"}
	// result := db.Create(&post1)
	// if result.Error != nil {
	// 	fmt.Println("插入数据失败：", result.Error)
	// }
	// getMostCommentedPost()
	// getuUserPostsWithComments("薛定谔")
	// db.AutoMigrate(&User{}, &Post{}, &Comment{})
	// datainti()
	// chashujv()
	// jishubu()
	// zuigaogongzi()
	// Createlines()
	// Creattables()
	// InitsqlxDB()
	// db := InitDB()
	// runMigrationOnce(db)
	// transacton("A", "B")
	// accountsA := []Account{{AccountName: "A", Balance: 150.00}}
	// accountsB := []Account{{AccountName: "B", Balance: 50.00}}
	// transactiono := []Transaction{{From_account_id: 0, To_account_id: 0, Amount: 0.00}}
	// db.Create(&accountsA)
	// db.Create(&accountsB)
	// db.Create(&transactiono)
	// sqlsentence := []string{
	// 	"INSERT INTO students (Name,Age,Grade) VALUES ('张三',20,'三年级')",
	// 	"SELECT * FROM students WHERE Age > 18",
	// 	"UPDATE students SET Grade='四年级' WHERE Name ='张三'",
	// 	"DELETE FROM students WHERE Age < 15",
	// db.Exec(sqlsentence[3])
	// db := InitDB()
	// db.Migrator().DropTable("users")
}
