package damy

import (
	"database/sql"
	"log"
	"myapp/app/models"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

// sqlのDBを指すポインタとしてDbを宣言
var Db *sql.DB



type Councilor struct{
	Id int `json:"id" form:"id" query:"id"`
	Name string `json:"name" form:"name" query:"name"`
	Address string `json:"address" form:"address" query:"address"`
}






	// ----------------------------------------------------------
	// login機能実装
	// ----------------------------------------------------------
	func login(c echo.Context) error{
		u:= new(models.User)		
		if err:= c.Bind(u); err != nil{
			log.Fatal(err)
		}
		// Db,_:= sql.Open("sqlite3","./test.sql")
		// defer Db.Close()

		user,err:=models.GetUser(u)
		if err != nil{
			log.Fatalln(err)
		}

		if models.Encrypt(u.PassWord) ==models.Encrypt(user.PassWord){
			return c.JSON(http.StatusCreated, "OK")
		}else{
			return c.JSON(http.StatusCreated,"NotFound")
		}	
	}
	
	// ----------------------------------------------------------
	// login機能実装
	// ----------------------------------------------------------

	// ----------------------------------------------------------
	// サインアップ機能実装
	// ----------------------------------------------------------
		func createAccount(c echo.Context)error{
			u:= new(models.User)
			if err:= c.Bind(u); err != nil{
				log.Fatal(err)
			}
			err := u.CreateUser()
			if err != nil{
				log.Fatalln(err)
			}
			return c.JSON(http.StatusCreated, "OK")
		}
	// ----------------------------------------------------------
	// サインアップ機能実装
	// ----------------------------------------------------------

	// ----------------------------------------------------------
	// 議員情報を取得
	// ----------------------------------------------------------

func getCouncilor(c echo.Context)error{
	var councilor Councilor
	id:= c.Param("id")
	log.Print(id)
	// Db,_:= sql.Open("sqlite3","././coucils.sql")
	// defer Db.Close()
	cmd:= "SELECT * FROM councils WHERE id = ?"
	err := Db.QueryRow(cmd,id).Scan(
		&councilor.Id,
		&councilor.Name,
		&councilor.Address,
	)
	if err != nil{
		log.Fatal(err)
	}
	
	return c.JSON(http.StatusCreated, councilor)
}

func getCouncilors(c echo.Context)error{
	// Db, _ := sql.Open("sqlite3", "./coucils.sql")
	// defer Db.Close()
	var councilors []models.Councilor
	councilors,err:=models.GetCouncilorList()
	if err != nil{
		log.Fatalln(err)
	}
	return c.JSON(http.StatusCreated, councilors)
}

	// ----------------------------------------------------------
	// 議員情報を取得
	// ----------------------------------------------------------

	