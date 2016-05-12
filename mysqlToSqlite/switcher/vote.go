package switcher

import (
	//"crypto/md5"
	//"crypto/rand"
	"database/sql"
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

type Votes struct {
	Id         int
	Title      string
	TopImg     string
	ProfileImg string
	LogDate    string
	IsOnline   int
	VoteType   int
}
type VotesAppUser struct {
	DeviceToken string
	Vote_id     int
	If_voted    int
}
type VotesCandidate struct {
	Id       int
	Vote_id  int
	Name     string
	Work     string
	Img      string
	Thumb    string
	IsOnline int
}
type VotesCicks struct {
	Ip      string
	LogDate string
	Vote_id int
}
type VotesComments struct {
	Vote_id int
	Comment string
	LogDate string
}
type VotesInfo struct {
	Vote_id       int
	Vote_for      int
	Vote_from     string
	Vote_datetime string
}
type VotesSeq struct {
	Vote_id int
	Seq     int
}

func Votes_Seq() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v VotesSeq
	var str string
	rows, err := db.Query("select * from votes_seq")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("votesseq.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Vote_id, &v.Seq)
		str += strconv.Itoa(v.Vote_id) + "<>" + strconv.Itoa(v.Seq) + "\n"
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
	f2, err3 := os.Open("votesseq.txt")
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
			stmt, err6 := tx.Prepare("insert into votes_seq values(?,?)")
			perrorWithRollBack(err6, "准备失败", tx)
			_, err = stmt.Exec(s1[0], s1[1])
			if err != nil {
				log.Println(err)
			}

		}
	}
	tx.Commit()
}
func Votes_Info() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v VotesInfo
	var str string
	var t time.Time
	rows, err := db.Query("select * from votes_info")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("voteinfo.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Vote_id, &v.Vote_for, &v.Vote_from, &t)
		v.Vote_datetime = t.Format("2006-01-02 15:04:05")
		str += strconv.Itoa(v.Vote_id) + "<>" + strconv.Itoa(v.Vote_for) + "<>" + v.Vote_from + "<>" + v.Vote_datetime + "\n"
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
	f2, err3 := os.Open("voteinfo.txt")
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
			stmt, err6 := tx.Prepare("insert into votes_info values(?,?,?,?)")
			perrorWithRollBack(err6, "准备失败", tx)
			_, err = stmt.Exec(s1[0], s1[1], s1[2], s1[3])
			if err != nil {
				log.Println(err)
			}

		}
	}
	tx.Commit()
}
func Votes_Comments() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v VotesComments
	var str string
	var t time.Time
	rows, err := db.Query("select * from votes_comments")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("votecomments.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Vote_id, &v.Comment, &t)
		v.LogDate = t.Format("2006-01-02 15:04:05")
		str += strconv.Itoa(v.Vote_id) + "<>" + v.Comment + "<>" + v.LogDate + "\n"
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
	f2, err3 := os.Open("votecomments.txt")
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
			stmt, err6 := tx.Prepare("insert into votes_comments values(?,?,?)")
			perrorWithRollBack(err6, "准备失败", tx)
			_, err = stmt.Exec(s1[0], s1[1], s1[2])
			if err != nil {
				log.Println(err)
			}

		}
	}
	tx.Commit()
}
func Votes_Clicks() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v VotesCicks
	var str string
	var t time.Time
	rows, err := db.Query("select * from votes_clicks")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("voteclicks.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Ip, &t, &v.Vote_id)
		v.LogDate = t.Format("2006-01-02 15:04:05")
		log.Println(v)
		str += v.Ip + "<>" + v.LogDate + "<>" + strconv.Itoa(v.Vote_id) + "\n"
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
	f2, err3 := os.Open("voteclicks.txt")
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
			log.Println(s1)
			stmt, err6 := tx.Prepare("insert into votes_clicks values(?,?,?)")
			perrorWithRollBack(err6, "准备失败", tx)
			_, err = stmt.Exec(s1[0], s1[1], s1[2])
			if err != nil {
				log.Println(err)
			}

		}
	}
	tx.Commit()
}
func Votes_Candidate() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v VotesCandidate
	var str string

	rows, err := db.Query("select * from votes_candidate")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("candidate.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Id, &v.Vote_id, &v.Name, &v.Work, &v.Img, &v.Thumb, &v.IsOnline)
		str += strconv.Itoa(v.Id) + "<>" + strconv.Itoa(v.Vote_id) + "<>" + v.Name + "<>" + v.Work + "<>" + v.Img + "<>" + v.Thumb + "<>" + strconv.Itoa(v.IsOnline) + "\n"
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
	f2, err3 := os.Open("candidate.txt")
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
			stmt, err6 := tx.Prepare("insert into votes_candidate values(?,?,?,?,?,?,?)")
			perrorWithRollBack(err6, "准备失败", tx)
			_, err = stmt.Exec(s1[0], s1[1], s1[2], s1[3], s1[4], s1[5], s1[6])
			if err != nil {
				log.Println(err)
			}

		}
	}
	tx.Commit()
}
func Votes_App_User() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v VotesAppUser
	var str string

	rows, err := db.Query("select * from votes_app_user")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("appuser.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.DeviceToken, &v.Vote_id, &v.If_voted)
		str += v.DeviceToken + "," + strconv.Itoa(v.Vote_id) + "," + strconv.Itoa(v.If_voted) + "\n"
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
	f2, err3 := os.Open("appuser.txt")
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
			s1 := strings.Split(v, ",")
			stmt, err6 := tx.Prepare("insert into votes_app_user values(?,?,?)")
			perrorWithRollBack(err6, "准备失败", tx)
			_, err = stmt.Exec(s1[0], s1[1], s1[2])
			if err != nil {
				log.Println(err)
			}

		}
	}
	tx.Commit()
}
func Vote() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v Votes
	var str string
	var t time.Time
	rows, err := db.Query("select * from votes")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("votes.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Id, &v.Title, &v.TopImg, &v.ProfileImg, &t, &v.IsOnline, &v.VoteType)
		v.LogDate = t.Format("2006-01-02 15:04:05")
		str += strconv.Itoa(v.Id) + "," + v.Title + "," + v.TopImg + "," + v.ProfileImg + "," + v.LogDate + "," + strconv.Itoa(v.IsOnline) + "," + strconv.Itoa(v.VoteType) + "\n"
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
	f2, err3 := os.Open("votes.txt")
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
	for _, v := range arr {
		if v != "" {
			s1 := strings.Split(v, ",")
			_, err = db.Exec("insert into votes values(?,?,?,?,?,?,?)", s1[0], s1[1], s1[2], s1[3], s1[4], s1[5], s1[6])
			if err != nil {
				log.Println(err)
			}

		}
	}
}
func ConnectDB(dbPath string) *sql.DB {
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
}
