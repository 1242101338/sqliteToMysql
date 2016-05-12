package switcher

import (
	//"crypto/md5"
	//"crypto/rand"
	//"database/sql"
	"os"
	"strconv"
	"time"
	//"encoding/base64"
	//"encoding/hex"
	//"encoding/json"
	//"fmt"
	//"io"
	//"io/ioutil"
	"log"
	//"net/http"
	//"path"
	//sw "sqliteToMysql/switcher"
	"strings"
	//xupload "xinlanAdminTest/xinlanUpload"

	//"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

type HotClicks struct {
	Ip      string
	LogDate string
	hot_id  string
}
type HotComments struct {
	Id      int
	Name    string
	Img     string
	Content string
	LogDate string
	EventId int
}
type HotEvents struct {
	Id      int
	Title   string
	Status  string
	Content string
	LogDate string
	HotId   int
	UserId  int
}
type Hots struct {
	Id          int
	Title       string
	Description string
	LogDate     string
	TopImg      string
}
type UserInfo struct {
	Id        int
	UserName  string
	PassWord  string
	Privilege string
	LogAt     string
}
type HotZans struct {
	EventId int
	UserId  int
}

func Hot_zans() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v HotZans
	var str string

	rows, err := db.Query("select * from zans")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("hotzans.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.EventId, &v.UserId)
		str += strconv.Itoa(v.EventId) + "<>" + strconv.Itoa(v.UserId) + "\n"
	}
	f.WriteString(str)
	//读取txt数据，写入到mysql中
	db = ConnectMySql()

	defer func() {
		db.Close()
		err5 := recover()
		if err5 != nil {
			log.Println(err5)
		}
	}()
	f2, err3 := os.Open("hotzans.txt")
	if err3 != nil {
		panic(err3)
	}
	defer f2.Close()
	stat, err4 := f2.Stat()
	if err4 != nil {
		log.Println(err4)
	}
	size := stat.Size()
	a := make([]byte, size)
	f2.Read(a)
	//log.Println(string(a))
	arr := strings.Split(string(a), "\n")
	tx, err5 := db.Begin()
	perrorWithRollBack(err5, "插入失败", tx)
	for _, v := range arr {
		if v != "" {
			s1 := strings.Split(v, "<>")
			stmt, err6 := tx.Prepare("insert into zans values(?,?)")
			perrorWithRollBack(err6, "准备失败", tx)
			_, err = stmt.Exec(s1[0], s1[1])
			if err != nil {
				log.Println(err)
			}

		}
	}
	tx.Commit()
}

func User_info() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v UserInfo
	var str string
	var t time.Time
	rows, err := db.Query("select * from userinfo")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("userinfo.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Id, &v.UserName, &v.PassWord, &v.Privilege, &t)
		v.LogAt = t.Format("2006-01-02 15:04:05")
		str += strconv.Itoa(v.Id) + "<>" + v.UserName + "<>" + v.PassWord + "<>" + v.Privilege + "<>" + v.LogAt + "\n"
	}
	f.WriteString(str)
	//读取txt数据，写入到mysql中
	db = ConnectMySql()

	defer func() {
		db.Close()
		err5 := recover()
		if err5 != nil {
			log.Println(err5)
		}
	}()
	f2, err3 := os.Open("userinfo.txt")
	if err3 != nil {
		panic(err3)
	}
	defer f2.Close()
	stat, err4 := f2.Stat()
	if err4 != nil {
		log.Println(err4)
	}
	size := stat.Size()
	a := make([]byte, size)
	f2.Read(a)
	//log.Println(string(a))
	arr := strings.Split(string(a), "\n")
	tx, err5 := db.Begin()
	perrorWithRollBack(err5, "插入失败", tx)
	for _, v := range arr {
		if v != "" {
			s1 := strings.Split(v, "<>")
			stmt, err6 := tx.Prepare("insert into userinfo values(?,?,?,?,?)")
			perrorWithRollBack(err6, "准备失败", tx)
			_, err = stmt.Exec(s1[0], s1[1], s1[2], s1[3], s1[4])
			if err != nil {
				log.Println(err)
			}

		}
	}
	tx.Commit()
}
func Hotsall() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v Hots
	var str string
	var t time.Time
	rows, err := db.Query("select * from hots")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("hots.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Id, &v.Title, &v.Description, &t, &v.TopImg)
		v.LogDate = t.Format("2006-01-02 15:04:05")
		str += strconv.Itoa(v.Id) + "<>" + v.Title + "<>" + v.Description + "<>" + v.LogDate + "<>" + v.TopImg + "\n"
	}
	f.WriteString(str)
	//读取txt数据，写入到mysql中
	db = ConnectMySql()

	defer func() {
		db.Close()
		err5 := recover()
		if err5 != nil {
			log.Println(err5)
		}
	}()
	f2, err3 := os.Open("hots.txt")
	if err3 != nil {
		panic(err3)
	}
	defer f2.Close()
	stat, err4 := f2.Stat()
	if err4 != nil {
		log.Println(err4)
	}
	size := stat.Size()
	a := make([]byte, size)
	f2.Read(a)
	//log.Println(string(a))
	arr := strings.Split(string(a), "\n")
	tx, err5 := db.Begin()
	perrorWithRollBack(err5, "插入失败", tx)
	for _, v := range arr {
		if v != "" {
			s1 := strings.Split(v, "<>")
			stmt, err6 := tx.Prepare("insert into hots values(?,?,?,?,?)")
			perrorWithRollBack(err6, "准备失败", tx)
			_, err = stmt.Exec(s1[0], s1[1], s1[2], s1[3], s1[4])
			if err != nil {
				log.Println(err)
			}

		}
	}
	tx.Commit()
}

func Hot_events() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v HotEvents
	var str string
	var t time.Time
	rows, err := db.Query("select * from events")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("hotevents.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Id, &v.Title, &v.Status, &v.Content, &t, &v.HotId, &v.UserId)
		v.LogDate = t.Format("2006-01-02 15:04:05")
		str += strconv.Itoa(v.Id) + " {} " + v.Title + " {} " + v.Status + " {} " + v.Content + " {} " + v.LogDate + " {} " + strconv.Itoa(v.HotId) + " {} " + strconv.Itoa(v.UserId) + "\\"
	}
	f.WriteString(str)
	//读取txt数据，写入到mysql中
	db = ConnectMySql()

	defer func() {
		db.Close()
		err5 := recover()
		if err5 != nil {
			log.Println(err5)
		}
	}()
	f2, err3 := os.Open("hotevents.txt")
	if err3 != nil {
		panic(err3)
	}
	defer f2.Close()
	stat, err4 := f2.Stat()
	if err4 != nil {
		log.Println(err4)
	}
	size := stat.Size()
	a := make([]byte, size)
	f2.Read(a)
	//log.Println(string(a))
	arr := strings.Split(string(a), "\\")
	tx, err5 := db.Begin()
	perrorWithRollBack(err5, "插入失败", tx)
	for _, v := range arr {
		if v != "" {
			s1 := strings.Split(v, " {} ")
			stmt, err6 := tx.Prepare("insert into events values(?,?,?,?,?,?,?)")
			if err6 != nil {
				log.Println(err)
			}
			perrorWithRollBack(err6, "准备失败", tx)
			log.Println(s1[3])
			_, err = stmt.Exec(s1[0], s1[1], s1[2], s1[3], s1[4], s1[5], s1[6])
			if err != nil {
				log.Println(err)
			}

		}
	}
	tx.Commit()
}
func Hot_comments() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v HotComments
	var str string
	var t time.Time
	rows, err := db.Query("select * from comments")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("hotcomments.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Id, &v.Name, &v.Img, &v.Content, &t, &v.EventId)
		v.LogDate = t.Format("2006-01-02 15:04:05")
		str += strconv.Itoa(v.Id) + "</>" + v.Name + "</>" + v.Img + "</>" + v.Content + "</>" + v.LogDate + "</>" + strconv.Itoa(v.EventId) + "\n"
	}
	f.WriteString(str)
	//读取txt数据，写入到mysql中
	db = ConnectMySql()

	defer func() {
		db.Close()
		err5 := recover()
		if err5 != nil {
			log.Println(err5)
		}
	}()
	f2, err3 := os.Open("hotcomments.txt")
	if err3 != nil {
		panic(err3)
	}
	defer f2.Close()
	stat, err4 := f2.Stat()
	if err4 != nil {
		log.Println(err4)
	}
	size := stat.Size()
	a := make([]byte, size)
	f2.Read(a)
	//log.Println(string(a))
	arr := strings.Split(string(a), "\n")
	tx, err5 := db.Begin()
	perrorWithRollBack(err5, "插入失败", tx)
	for _, v := range arr {
		if v != "" {
			s1 := strings.Split(v, "</>")
			stmt, err6 := tx.Prepare("insert into comments values(?,?,?,?,?,?)")
			perrorWithRollBack(err6, "准备失败", tx)
			_, err = stmt.Exec(s1[0], s1[1], s1[2], s1[3], s1[4], s1[5])
			if err != nil {
				log.Println(err)
			}

		}
	}
	tx.Commit()
}
func Hot_clicks() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle_bak.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v HotClicks
	var str string
	var t time.Time
	rows, err := db.Query("select * from clicks")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("hotclicks.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Ip, &t, &v.hot_id)
		v.LogDate = t.Format("2006-01-02 15:04:05")
		str += v.Ip + "<>" + v.LogDate + "<>" + v.hot_id + "\n"
	}
	f.WriteString(str)
	//读取txt数据，写入到mysql中
	db = ConnectMySql()

	defer func() {
		db.Close()
		err5 := recover()
		if err5 != nil {
			log.Println(err5)
		}
	}()
	f2, err3 := os.Open("hotclicks.txt")
	if err3 != nil {
		panic(err3)
	}
	defer f2.Close()
	stat, err4 := f2.Stat()
	if err4 != nil {
		log.Println(err4)
	}
	size := stat.Size()
	a := make([]byte, size)
	f2.Read(a)
	//log.Println(string(a))
	arr := strings.Split(string(a), "\n")
	tx, err5 := db.Begin()
	perrorWithRollBack(err5, "插入失败", tx)
	for _, v := range arr {
		if v != "" {
			s1 := strings.Split(v, "<>")
			stmt, err6 := tx.Prepare("insert into clicks values(?,?,?)")
			perrorWithRollBack(err6, "准备失败", tx)
			_, err = stmt.Exec(s1[0], s1[1], s1[2])
			if err != nil {
				log.Println(err)
			}

		}
	}
	tx.Commit()
}

/*func ConnectDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	return db
}

func ConnectMySql() *sql.DB {
	dbinfo := "root" + ":" + "123" + "@/" + "xinlanAdmin" + "?charset=utf8"
	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
		panic(err)
	}
	return db
}
func perrorWithRollBack(e error, errMsg string, tx *sql.Tx) {
	if e != nil {
		tx.Rollback()
		log.Println(e)
		panic(errMsg)
	}
}*/
