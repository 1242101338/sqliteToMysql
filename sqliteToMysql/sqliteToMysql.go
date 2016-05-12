// sqliteToMysql
package main

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
	sw "sqliteToMysql/switcher"
	"strings"
	//xupload "xinlanAdminTest/xinlanUpload"

	//"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

type Announcement struct {
	Id                   int
	Video_id             int
	Announcement         string
	Announcementdatatime string
}
type VideosAddclicks struct {
	Id         int
	Video_id   int
	Add_clicks int
}
type VideosClicks struct {
	Id       int
	Ip       string
	Video_id int
}
type Reply struct {
	Id         int
	Comment_id int
	Reply      string
}
type Video struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	VideoStream  string `json:"videostream"`
}
type VideoComment struct {
	Id        int    `json:"id"`
	Video_id  int    `json:"videoid"`
	Comment   string `json:"comment"`
	Logdate   string `json:"logdate"`
	User_id   int    `json:"userid"`
	User_img  string `json:"userimg"`
	User_name string `json:"username"`
	IsChecked int    `json:"ischecked"`
}

func main() {
	//直播
	//Videos()
	//Videos_Comments()
	//Videos_Addclicks()
	//Videos_Announcement()
	//Videos_Clicks()
	//Videos_Reply()
	//投票
	//sw.Vote()
	//sw.Votes_App_User()
	//sw.Votes_Candidate()
	sw.Votes_Clicks()
	//sw.Votes_Comments()
	//sw.Votes_Info()
	//sw.Votes_Seq()
	//热点
	//sw.Hot_clicks()
	//sw.Hot_comments()
	//sw.Hot_events()
	//sw.Hotsall()
	//sw.User_info()
	//sw.Hot_zans()
}
func Videos_Reply() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle_video.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v Reply
	var str string
	rows, err := db.Query("select * from videos_reply")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("videosReply.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Id, &v.Comment_id, &v.Reply)
		str += strconv.Itoa(v.Id) + "," + strconv.Itoa(v.Comment_id) + "," + v.Reply + "\n"
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
	f2, err3 := os.Open("videosReply.txt")
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
			_, err = db.Exec("insert into videos_reply(id,comment_id,reply) values(?,?,?)", s1[0], s1[1], s1[2])
			if err != nil {
				log.Println(err)
			}

		}
	}
}
func Videos_Clicks() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle_video.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v VideosClicks
	var str string
	rows, err := db.Query("select * from videos_clicks")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("videosClicks.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Id, &v.Ip, &v.Video_id)
		str += strconv.Itoa(v.Id) + "," + v.Ip + "," + strconv.Itoa(v.Video_id) + "\n"
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
	f2, err3 := os.Open("videosClicks.txt")
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
			_, err = db.Exec("insert into videos_clicks(id,ip,video_id) values(?,?,?)", s1[0], s1[1], s1[2])
			if err != nil {
				log.Println(err)
			}

		}
	}
}
func Videos_Announcement() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle_video.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v Announcement
	var str string
	var t time.Time
	rows, err := db.Query("select * from videos_announcement")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("videosAnnouncement.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Id, &v.Video_id, &v.Announcement, &t)
		v.Announcementdatatime = t.Format("2006-01-02 15:04:05")
		str += strconv.Itoa(v.Id) + "," + strconv.Itoa(v.Video_id) + "," + v.Announcement + "," + v.Announcementdatatime + "\n"
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
	f2, err3 := os.Open("videosAnnouncement.txt")
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
			_, err = db.Exec("insert into videos_announcement(id,video_id,announcement,announcement_datatime) values(?,?,?,?)", s1[0], s1[1], s1[2], s1[3])
			if err != nil {
				log.Println(err)
			}

		}
	}
}

func Videos_Addclicks() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle_video.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v VideosAddclicks
	var str string
	rows, err := db.Query("select * from videos_addclicks")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("videosAddclicks.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Id, &v.Video_id, &v.Add_clicks)
		str += strconv.Itoa(v.Id) + "," + strconv.Itoa(v.Video_id) + "," + strconv.Itoa(v.Add_clicks) + "\n"
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
	f2, err3 := os.Open("videosAddclicks.txt")
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
			_, err = db.Exec("insert into videos_addclicks(id,video_id,add_clicks) values(?,?,?)", s1[0], s1[1], s1[2])
			if err != nil {
				log.Println(err)
			}

		}
	}
}

func Videos_Comments() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle_video.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v VideoComment
	var str string
	rows, err := db.Query("select * from videos_comments")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("videosComments.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()
	var t time.Time
	for rows.Next() {
		rows.Scan(&v.Id, &v.Video_id, &v.Comment, &t, &v.User_id, &v.User_img, &v.User_name, &v.IsChecked)
		v.Logdate = t.Format("2006-01-02 15:04:05")
		str += strconv.Itoa(v.Id) + "," + strconv.Itoa(v.Video_id) + "," + v.Comment + "," + v.Logdate + "," + strconv.Itoa(v.User_id) + "," + v.User_img + "," + v.User_name + "," + strconv.Itoa(v.IsChecked) + "\n"
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
	f2, err3 := os.Open("videosComments.txt")
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
			//log.Println(s1[0], s1[1], s1[2], s1[3], s1[4], s1[5], s1[6], s1[7])
			_, err = db.Exec("insert into videos_comments(id,video_id,comment,logdate,user_id,user_img,user_name,is_checked) values(?,?,?,?,?,?,?,?)", s1[0], s1[1], s1[2], s1[3], s1[4], s1[5], s1[6], s1[7])
			if err != nil {
				log.Println(err)
			}

		}
	}
}

func Videos() {
	//读取sqlite数据，写入到txt
	db := ConnectDB("./middle_video.db")

	defer func() {
		db.Close()
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	var v Video
	var str string
	rows, err := db.Query("select * from videos")
	if err != nil {
		log.Println(err)
	}

	f, err2 := os.Create("videos.txt")
	if err2 != nil {
		panic(err2)
	}
	defer f.Close()

	for rows.Next() {
		rows.Scan(&v.Id, &v.Title, &v.Introduction, &v.VideoStream)
		str += strconv.Itoa(v.Id) + "," + v.Title + "," + v.Introduction + "," + v.VideoStream + "\n"
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
	f2, err3 := os.Open("videos.txt")
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
			_, err = db.Exec("insert into videos(id,title,introduction,videostream) values(?,?,?,?)", s1[0], s1[1], s1[2], s1[3])
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
