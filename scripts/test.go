package main

import (
	"encoding/json"
	"fmt"
)


func get_players(startRank int, endRank int) []IttfPlayer {
	var result []IttfPlayer;
	err := Get(fmt.Sprintf("internalttu/RankingsCurrentWeek/CurrentWeek/GetRankingIndividuals?CategoryCode=SEN&SubEventCode=MS&StartRank=%d&EndRank=%d", startRank, endRank), "Result", &result)
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
    conn, err := getDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	insertPlayers(players, conn)
	
}
