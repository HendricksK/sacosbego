package models

import (
	"strings"
	"fmt"
	"log"
)

func BuildSelectQueryWithAggregate(fields []string, model string, model_aggregate string) string {
	fieldList := fmt.Sprint(fields,",")
	fieldList = strings.ReplaceAll(fieldList,"[","")
	fieldList = strings.ReplaceAll(fieldList,"]","")
	fieldList = strings.ReplaceAll(fieldList," ",",")
	fieldList = strings.TrimRight(fieldList,",")

	log.Println("SELECT "  + fieldList + " FROM " + model + " " + model + " LEFT JOIN " + model_aggregate + " " + model_aggregate + " ON " + model_aggregate + "." + model + "_id = " + model + ".id")

	return "SELECT "  + fieldList + " FROM " + model + " " + model + " LEFT JOIN " + model_aggregate + " " + model_aggregate + " ON " + model_aggregate + "." + model + "_id = " + model + ".id"
} 


func BuildSelectQuery(fields []string, model string, id int) string {
	fieldList := fmt.Sprint(fields,",")
	fieldList = strings.ReplaceAll(fieldList,"[","")
	fieldList = strings.ReplaceAll(fieldList,"]","")
	fieldList = strings.ReplaceAll(fieldList," ",",")
	fieldList = strings.TrimRight(fieldList,",")

	return "SELECT "  + fieldList + " FROM " + model
}

