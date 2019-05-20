package controllers

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type PoliSturct struct {
	PoliCode string
	PoliName string
}

type PoliSchedule struct {
	Doctor_id string
	Specialist_id string
	Poli_id string
	Day_id string
	tmpTime string
	Time_id int
	Start_time time.Time
	End_time time.Time
}

func GetPoli(c *gin.Context) {

	var (
		poli    PoliSturct
		arrPoli []PoliSturct
	)

	tsql := fmt.Sprintf("SELECT KODEBAGIAN, NAMABAGIAN FROM dbo.BAGIAN;")

	rows, err := Db.Query(tsql)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		if err := rows.Scan(&poli.PoliCode, &poli.PoliName); err != nil {
			log.Fatal(err.Error())

		} else {
			arrPoli = append(arrPoli, poli)
		}
	}

	defer rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    arrPoli,
		"count":   len(arrPoli),
	})

}

func GetSchedule(c *gin.Context)  {
	var (
		poliSchedule PoliSchedule
		arrSchedule []PoliSchedule
	)

	tsql := fmt.Sprintf("SELECT KODEDOKTER, KDSPESIAL, KODEBAGIAN, KODEHARI, JAMMULAI, JAMSELESAI, KODEWAKTU FROM dbo.JWDOKTER;")
	rows, err := Db.Query(tsql)
	if err != nil {
		log.Fatalln(err.Error())
	}


	log.Println("ini rows next : ",rows.Next())
	i := 0
	for rows.Next() {
		if err := rows.Scan(&poliSchedule.Doctor_id, &poliSchedule.Specialist_id, &poliSchedule.Poli_id, &poliSchedule.Day_id, &poliSchedule.Start_time, &poliSchedule.End_time, &poliSchedule.tmpTime); err != nil {
			log.Fatalln(err.Error())
		} else {
			arrSchedule = append(arrSchedule, poliSchedule)
			if arrSchedule[i].tmpTime == "P" {
				arrSchedule[i].Time_id = 1
			} else if arrSchedule[i].tmpTime == "S" {
				arrSchedule[i].Time_id = 2
			} else if arrSchedule[i].tmpTime == "L" {
				arrSchedule[i].Time_id = 3
			} else {
				arrSchedule[i].Time_id = 99
			}
		}
		i++
	}

	//for i:=1; i < len(arrSchedule); i++ {
	//	if arrSchedule[i].tmpTime == "P" {
	//		arrSchedule[i].Time_id = 1
	//	} else if arrSchedule[i].tmpTime == "S" {
	//		arrSchedule[i].Time_id = 2
	//	} else if arrSchedule[i].tmpTime == "L" {
	//		arrSchedule[i].Time_id = 3
	//	} else {
	//		arrSchedule[i].Time_id = 99
	//	}
	//}

	defer rows.Close()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code": 200,
		"message": "",
		"data": arrSchedule,
	})
}
