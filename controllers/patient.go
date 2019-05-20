package controllers

import (
	"../config"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Syspar struct {
	Nopaskir    string
	Nopastng    string
	Nopaskan    string
	Nopastngmin string
	Nopasakhir  string
}

var Db = config.DbInit()

var resNopasakhir string

func CheckNoPatient() {

	var (
		noPatient  Syspar
		Nopasakhir string
	)

	tsql := fmt.Sprintf("SELECT NOPASKIR, NOPASTNG, NOPASKAN, NOPASTNGMIN, NOPASAKHIR FROM dbo.SYSPAR;")

	row := Db.QueryRow(tsql)
	err := row.Scan(&noPatient.Nopaskir, &noPatient.Nopastng, &noPatient.Nopaskan, &noPatient.Nopastngmin, &noPatient.Nopasakhir)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(noPatient)

		tmpNopaskir, error1 := strconv.ParseInt(noPatient.Nopaskir, 10, 8)
		tmpNopastng, error2 := strconv.ParseInt(noPatient.Nopastng, 10, 8)
		tmpNopaskan, error3 := strconv.ParseInt(noPatient.Nopaskan, 10, 8)
		tmpNopastngmin, error4 := strconv.ParseInt(noPatient.Nopastngmin, 10, 8)
		if error1 != nil && error2 != nil && error3 != nil && error4 != nil {
			log.Fatal("error parsing string to int")
		}

		var (
			resNopaskir    string
			resNopastngmin string
			resNopastng    string
			resNopaskan    string
		)

		if tmpNopaskir != 99 {

			if tmpNopaskir == 0 {
				resNopaskir = "01"
			} else {
				tmpNopaskir += 1
				resNopaskir = fmt.Sprintf("%v", tmpNopaskir)
				if (len(resNopaskir) == 1) {
					log.Println("this nopastngmin : %d", tmpNopaskir)
					resNopaskir = fmt.Sprintf("0%d", tmpNopaskir)
				}
			}

			resNopaskan = fmt.Sprintf("%v", tmpNopaskan)
			resNopastng = fmt.Sprintf("%v", tmpNopastng)
			resNopastngmin = fmt.Sprintf("%v", tmpNopastngmin)

			if (len(resNopastngmin) == 1) {
				resNopastngmin = fmt.Sprintf("0%v", tmpNopastngmin)
			}
			if (len(resNopastng) == 1) {
				resNopastng = fmt.Sprintf("0%v", tmpNopastng)
			}
			if (len(resNopaskan) == 1) {
				resNopaskan = fmt.Sprintf("0%v", tmpNopaskan)
				log.Println("NOPASKAN = %v", resNopaskan)
			}

		} else {
			resNopaskir = "00";
			if (tmpNopastngmin != 99) {

				if (tmpNopastngmin == 0) {
					resNopastngmin = "01";
				} else {
					tmpNopastngmin += 1;
					resNopastngmin = fmt.Sprintf("%v", tmpNopastngmin)
					if (len(resNopastngmin) == 1) {
						resNopastngmin = fmt.Sprintf("0%d", tmpNopastngmin)
					}
				}
				resNopaskan = fmt.Sprintf("%v", tmpNopaskan)
				resNopastng = fmt.Sprintf("%v", tmpNopastng)
				if (len(resNopastng) == 1) {
					resNopastng = fmt.Sprintf("0%v", tmpNopastng)
				}
				if (len(resNopaskan) == 1) {
					resNopaskan = fmt.Sprintf("0%v", tmpNopaskan)
					log.Println("NOPASKAN = %v", resNopaskan)
				}

			} else {
				resNopastngmin = "00";
				if (tmpNopastng != 99) {
					if (tmpNopastng == 0) {
						resNopastng = "01";
					} else {
						tmpNopastng += 1;
						resNopastng = fmt.Sprintf("%v", tmpNopastng)
						if (len(resNopastng) == 1) {
							resNopastng = fmt.Sprintf("0%d", tmpNopastng)
						}
					}
					resNopaskan = fmt.Sprintf("%v", tmpNopaskan)

					if (len(resNopaskan) == 1) {
						resNopaskan = fmt.Sprintf("0%v", tmpNopaskan)
					}
				} else {
					resNopastng = "00";
					if (tmpNopaskan != 99) {
						if (tmpNopaskan == 0) {
							resNopaskan = "01";
						} else {
							tmpNopaskan += 1;
							resNopaskan = fmt.Sprintf("%v", tmpNopaskan)
							if (len(resNopaskan) == 1) {
								resNopaskan = fmt.Sprintf("0%d", tmpNopaskan)
							}
						}
					} else {
						resNopaskan = "00";
					}
				}
			}
		}

		Nopasakhir = fmt.Sprintf("%s%s%s%s", resNopaskir, resNopastngmin, resNopastng, resNopaskan)

		updateSyspar := fmt.Sprintf("UPDATE dbo.SYSPAR SET NOPASKIR = '%s', NOPASTNGMIN = '%s', NOPASTNG = '%s', NOPASKAN = '%s', NOPASAKHIR = '%s' WHERE REGID= '%s'", resNopaskir, resNopastngmin, resNopastng, resNopaskan, Nopasakhir, "0101")

		nopatientUpdated, errUpdate := Db.Exec(updateSyspar)
		if errUpdate != nil {
			log.Println(nopatientUpdated)
			log.Println(errUpdate)
		}

		var patientNo string
		searchPatient := fmt.Sprintf("SELECT NOPASIEN FROM dbo.PASIEN WHERE NOPASIEN = '%s'", Nopasakhir)
		checkPatient := Db.QueryRow(searchPatient)
		errCheckPatient := checkPatient.Scan(&patientNo)
		if errCheckPatient != nil {
			log.Println(errCheckPatient.Error())
		}
		if patientNo == "" {
			fmt.Println("no patient belum dipakai")
			log.Println("NOPASAKHIR REAL = ", Nopasakhir)
			resNopasakhir = Nopasakhir
		} else {
			fmt.Println("no patient sudah dipakai")
			CheckNoPatient()
		}

	}
	log.Println("NOPASAKHIR REAL 222 = ", Nopasakhir)
	//return Nopasakhir
}

func GetNoPatient(c *gin.Context) {

	CheckNoPatient()

	Nopasakhir := resNopasakhir

	log.Println("NOPASAKHIR = ", Nopasakhir)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "",
		"data":    Nopasakhir,
	})

}

func store(c *gin.Context) {
	
}
