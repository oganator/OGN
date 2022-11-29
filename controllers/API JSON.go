package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"database/sql"

	_ "modernc.org/sqlite"

	beego "github.com/astaxie/beego"
)

// ViewEntityController - default route - "/". Executed from submit button
type TestController struct {
	beego.Controller
}

// GetString -
func GetTestString(c *TestController, field string) string {
	c.Data[field] = c.GetString(field)
	return c.Data[field].(string)
}

// GetFloat -
func GetTestFloat(c *TestController, field string) (result float64) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.ParseFloat(temp, 64)
	return result
}

// GetInt -
func GetTestInt(c *TestController, field string) (result int) {
	c.Data[field] = c.GetString(field)
	temp := c.Data[field].(string)
	result, _ = strconv.Atoi(temp)
	return result
}

// Get - Provides a list of funds
func (c *TestController) Get() {
	c.Data["json"] = FundsList
	c.ServeJSON()
}

// Post - Provides a list of funds
func (c *TestController) Post() {
	clientQuery := GetTestString(c, "query")
	db, err := sql.Open("sqlite", "./models/ogndata.db")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	rows, _ := db.Query(clientQuery)

	finaldata := MarshalQuery(rows)

	c.Data["json"] = finaldata
	c.ServeJSON()
}

func MarshalQuery(rows *sql.Rows) (finalRows []interface{}) {
	columnTypes, _ := rows.ColumnTypes()
	count := len(columnTypes)
	finalRows = []interface{}{}

	for rows.Next() {
		scanArgs := make([]interface{}, count)
		for i, v := range columnTypes {
			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
			case "INT4":
				scanArgs[i] = new(sql.NullInt64)
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		rows.Scan(scanArgs...)
		masterData := map[string]interface{}{}
		for i, v := range columnTypes {
			if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
				masterData[v.Name()] = z.Bool
				continue
			}
			if z, ok := (scanArgs[i]).(*sql.NullString); ok {
				masterData[v.Name()] = z.String
				continue
			}
			if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
				masterData[v.Name()] = z.Int64
				continue
			}
			if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
				masterData[v.Name()] = z.Float64
				continue
			}
			if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
				masterData[v.Name()] = z.Int32
				continue
			}
			masterData[v.Name()] = scanArgs[i]
		}
		finalRows = append(finalRows, masterData)

	}
	json.Marshal(finalRows)
	fmt.Println(finalRows)
	return finalRows
}
