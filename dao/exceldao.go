package dao

import (
	"domain/model"

	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func ReadExcel() []*model.Domain {

	path := Inputstr("请输入文件的路径：")

	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)

	}

	domainForms := []*model.Domain{}
	rows := f.GetRows("Sheet1")
	for _, row := range rows {
		if row[0] == "域名" {
			continue
		}
		domainForm := &model.Domain{
			DomainName:        row[0],
			Project:           row[1],
			Service:           row[2],
			CDN:               row[3],
			HTTPS:             row[4],
			Backend:           row[5],
			WhiteList:         row[6],
			WhiteListLocation: row[7],
			Notes:             row[8],
		}

		domainForms = append(domainForms, domainForm)

	}

	return domainForms
}
