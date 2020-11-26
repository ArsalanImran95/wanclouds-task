package main

import (
	"database/sql"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
func main() {

	// this section is for database connection
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// this section is for database creation
	_,err = db.Exec("CREATE DATABASE "+ "testcloud")
	if err != nil {
		panic(err)
	}

	// this section is use database
	_,err = db.Exec("USE "+"testcloud")
	if err != nil {
		panic(err)
	}
	// this section is to create a table, this table will be used to store data.
	_,err = db.Exec("CREATE TABLE blood ( first_name varchar(45) , last_name varchar(45), age varchar(45) , blood_type varchar(45))")
	if err != nil {
		panic(err)
	}
	// this section is for reading excel sheet.
	xlsx, err := excelize.OpenFile("./bloodgroup.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	var col [4] string
	// this section is for testing excel sheet
	cell, _ := xlsx.GetCellValue("Sheet1", "A2")
	fmt.Println(cell)
	//this section is to iterate through values from excel sheet
	rows, _ := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		i := 0
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
			col[i] = colCell
			i++
		}
		// this section is to insert values into blood table.
		stmt, err := db.Prepare("INSERT INTO blood(first_name,last_name,age,blood_type) VALUES(?,?,?,?)")
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(col[0],col[1],(col[2]),col[3])
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
		fmt.Println()
	}


}