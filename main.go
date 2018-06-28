package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Tournament struct {
	id        string
	startDate string
	endDate   string
}

type Leagues struct {
	id   int
	name string
	logo string
	guid string
}

func getLeagues() {
	var dat map[string]interface{}
	// We can use GET form to get result.
	resp, err := http.Get("https://api.lolesports.com/api/v1/navItems")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &dat); err != nil {
		log.Fatal(err)
	}

	var dat2 []interface{}
	dat2 = dat["leagues"].([]interface{})
	for _, v := range dat2 {
		var leagueTmp Leagues
		for k2, v2 := range v.(map[string]interface{}) {
			if k2 == "id" {
				leagueTmp.id = int(v2.(float64))
			} else if k2 == "guid" {
				leagueTmp.guid = v2.(string)
			} else if k2 == "name" {
				leagueTmp.name = v2.(string)
			} else if k2 == "logoUrl" {
				leagueTmp.logo = v2.(string)
			}
		}
		getTournament(leagueTmp)
		fmt.Println(leagueTmp)
	}
}

func getTournament(leagueTmp Leagues) {
	resp, err := http.Get("https://api.lolesports.com/api/v1/leagues/" + strconv.Itoa(leagueTmp.id))
	fmt.Println("https://api.lolesports.com/api/v1/leagues/" + strconv.Itoa(leagueTmp.id))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var datLeague map[string]interface{}
	if err := json.Unmarshal(body, &datLeague); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(datLeague)
	var dat2 []interface{}

	layout := "2006-01-02"

	dat2 = datLeague["highlanderTournaments"].([]interface{})
	for _, v := range dat2 {
		var tournamentTmp Tournament
		for k2, v2 := range v.(map[string]interface{}) {
			if k2 == "id" {
				tournamentTmp.id = v2.(string)
			} else if k2 == "startDate" {
				tournamentTmp.startDate = v2.(string)
			} else if k2 == "endDate" {
				tournamentTmp.endDate = v2.(string)
			}
		}
		tEnd, _ := time.Parse(layout, tournamentTmp.endDate)
		if tEnd.After(time.Now()) {
			fmt.Println(tournamentTmp)
		}
	}
	time.Now()
}

func main() {
	getLeagues()
}
