package main

import (
	"TTSavant/db"
	"TTSavant/http"
	"TTSavant/structs"
	"encoding/json"
	"fmt"
)


func get_players(startRank int, endRank int) []structs.IttfPlayer {
	var result []structs.IttfPlayer;
	err := http.Get(fmt.Sprintf("internalttu/RankingsCurrentWeek/CurrentWeek/GetRankingIndividuals?CategoryCode=SEN&SubEventCode=MS&StartRank=%d&EndRank=%d", startRank, endRank), "Result", &result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func main() {
    players := get_players(1, 100)
    playersJSON, err := json.MarshalIndent(players, "", "  ")
    if err != nil {
        fmt.Println(err)
		fmt.Println(playersJSON)
        return
    }
    conn, err := db.GetDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	db.InsertPlayers(players, conn)
	
}
