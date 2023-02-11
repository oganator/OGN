package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

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

// Get - Provides static data to Retool
func (c *TestController) Get() {
	// temp := make(map[string]interface{})

	// c.Data["json"] = temp
	// c.ServeJSON()
}

// Post - SQL
// func (c *TestController) Post() {
// 	clientQuery := GetTestString(c, "query")
// 	fmt.Println(clientQuery)
// 	db, err := sql.Open("sqlite", "./models/ogndata.db")
// 	defer db.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	rows, _ := db.Query(clientQuery)
// 	finaldata := MarshalQuery(rows)
// 	c.Data["json"] = finaldata
// 	c.ServeJSON()
// }

// Post - html response, can't get json ;(
// func (c *TestController) Post() {
// query, _ := json.Marshal(GetTestString(c, "query"))
// _ = os.WriteFile("views/test.tpl", query, 0644)
// c.TplName = "test.tpl"
// temp := make(map[interface{}]interface{})
// temp["EntityModelsMap"] = EntityModelsMap
// temp["ModelsList"] = ModelsList
// temp["FundsList"] = FundsList
// temp["Units"] = Units
// c.Data = temp
// }

// Post - switch with separate clauses for each struct field. returns JSON
func (c *TestController) Post() {
	c.TplName = "test.tpl"
	object := GetTestString(c, "object")
	switch object {
	case "table":
		asset := GetTestString(c, "asset")
		assetint := EntityModelsList[asset]
		c.Data["json"] = EntityModelsMap[assetint].EntityModel.Table
		c.ServeJSON()
	case "fundslist":
		response := make([]KeyValue, len(FundsList))
		index := 0
		for i, v := range FundsList {
			response[index].Key = i
			response[index].Value = v
			index++
		}
		c.Data["json"] = response
		c.ServeJSON()
	case "assetslist":
		c.Data["json"] = AssetsList
		c.ServeJSON()
	case "fundmodelslist":
		fund := strings.TrimSpace(GetTestString(c, "fund"))
		fundint := FundsList[fund]
		response := make([][]KeyValue, 0)
		for _, v := range EntityMap[fundint].Models {
			response = append(response, v.ModelDetails(DetailsInput{Name: true, StartDate: true, SalesDate: true}))
			c.Data["json"] = response
		}
		c.ServeJSON()
	case "json":
		query, _ := json.Marshal(GetTestString(c, "query"))
		_ = os.WriteFile("views/test.tpl", query, 0644)
		temp := make(map[interface{}]interface{})
		temp["EntityModelsMap"] = EntityModelsMap
		temp["EntityMap"] = EntityMap
		temp["Units"] = Units
		c.Data = temp
	case "table2":
		asset := GetTestString(c, "asset")
		timeframe := GetTestInt(c, "timeframe")
		assetint := EntityModelsList[asset]
		test := EntityModelsMap[assetint].EntityModel.MakeTable2(CFTableCOA, timeframe)
		c.Data["json"] = test
		c.ServeJSON()
	case "assetmodelslist":
		fundmodel := GetTestString(c, "fundmodel")
		fundint := FundModelsList[fundmodel]
		response := make([][]KeyValue, 0)
		for _, v := range EntityModelsMap[fundint].EntityModel.ChildEntityModels {
			response = append(response, v.ModelDetails(DetailsInput{Name: true}))
			c.Data["json"] = response
		}
		c.ServeJSON()
	}
}

func GetModels(entityInt int) (modelsTable [][]KeyValue) {
	modelsTable = make([][]KeyValue, 0)
	for _, v := range EntityModelsMap {
		if v.EntityModel.Entity.MasterID == entityInt {
			fmt.Println(v.EntityModel.Name)
			modelsTable = append(modelsTable, v.EntityModel.ModelDetails(DetailsInput{Name: true}))
		}
	}
	return modelsTable
}

// used to marshall a sql query into a json response
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
	return finalRows
}
