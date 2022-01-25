package main

import (
	"github.com/r-c-correa/nr-to-excel/pkg/config"
	"github.com/r-c-correa/nr-to-excel/pkg/errr"
	"github.com/r-c-correa/nr-to-excel/pkg/excel"
	"github.com/r-c-correa/nr-to-excel/pkg/nrql"
)

func main() {
	cfg, err := config.Load("./config.json")
	errr.PanicIfIsNotNull(err)

	response := nrql.GET(cfg)

	excel.SaveDataInExcel("./result.xlsx", response)
}
