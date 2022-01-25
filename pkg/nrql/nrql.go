package nrql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/r-c-correa/nr-to-excel/pkg/config"
	"github.com/r-c-correa/nr-to-excel/pkg/errr"
)

var client = &http.Client{}

func GET(cfg *config.Config) []map[string]interface{} {
	method := "GET"
	result := []map[string]interface{}{}
	layout := "2006-01-02T15:04:05Z"
	tStart, err := time.Parse(layout, cfg.Start)
	errr.PanicIfIsNotNull(err)

	tEnd, err := time.Parse(layout, cfg.End)
	errr.PanicIfIsNotNull(err)

	keys := []string{}

	for tStart.Before(tEnd) {
		query := cfg.NRQL
		tEndLocal := tStart.Add(time.Minute * time.Duration(cfg.WindowInMinutes))

		query += " SINCE '" + tStart.Format("2006-01-02 15:04:05") + " +0000'"
		query += " UNTIL '" + tEndLocal.Format("2006-01-02 15:04:05") + " +0000'"

		tEndLocal = tEndLocal.Add(time.Minute * time.Duration(-1))

		bbody := []byte(fmt.Sprintf(`
			{
				actor {
					account(id: %d) {
						nrql(query: "%v") {
							results
						}
					}
				}
			}
		`, cfg.AccountID, query))

		req, err := http.NewRequest(method, "https://api.newrelic.com/graphql", bytes.NewBuffer(bbody))
		errr.PanicIfIsNotNull(err)

		req.Header.Add("API-Key", cfg.ApiKey)
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		errr.PanicIfIsNotNull(err)

		b, err := ioutil.ReadAll(res.Body)
		errr.PanicIfIsNotNull(err)

		var response Response
		err = json.Unmarshal(b, &response)
		errr.PanicIfIsNotNull(err)

		body := string(b)
		if strings.Contains(body, `{"errorClass":"TIMEOUT"}`) {
			fmt.Println(body)

			return result
		}

		err = res.Body.Close()
		errr.PanicIfIsNotNull(err)

		if cfg.PrimaryKey == "" {
			result = append(result, response.Data.Actor.Account.NRQL.Results...)
		} else {
			for _, row := range response.Data.Actor.Account.NRQL.Results {
				pk := row[cfg.PrimaryKey].(string)
				if contains(keys, pk) {
					continue
				}

				keys = append(keys, pk)
				result = append(result, row)
			}
		}

		tStart = tEndLocal

		fmt.Printf("%v Query: %v (Total Rows: %v)\n", time.Now().Format("2006-01-02T15:04:05Z"), query, len(response.Data.Actor.Account.NRQL.Results))

		time.Sleep(time.Second * 1)
	}

	return result
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
