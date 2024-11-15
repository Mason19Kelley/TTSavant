package db

import (
	"TTSavant/structs"
	"context"

	"github.com/jackc/pgx/v5"
)

type IttfPlayerRanking = structs.IttfPlayerRanking

func InsertPlayerRankings(players []IttfPlayerRanking, conn *pgx.Conn) error {
    batch := &pgx.Batch{}
    for _, player := range players {
        batch.Queue(
            `INSERT INTO ittf_player_rankings (
                IttfId, PlayerName, CountryCode, CountryName, 
                AssociationCountryCode, AssociationCountryName, 
                CategoryCode, AgeCategoryCode, SubEventCode, 
                RankingYear, RankingMonth, RankingWeek, 
                RankingPointsCareer, RankingPointsYTD, 
                RankingPosition, CurrentRank, PreviousRank, 
                RankingDifference, PublishDate
            ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`,
            player.IttfId, player.PlayerName, player.CountryCode, player.CountryName,
            player.AssociationCountryCode, player.AssociationCountryName,
            player.CategoryCode, player.AgeCategoryCode, player.SubEventCode,
            player.RankingYear, player.RankingMonth, player.RankingWeek,
            player.RankingPointsCareer, player.RankingPointsYTD,
            player.RankingPosition, player.CurrentRank, player.PreviousRank,
            player.RankingDifference, player.PublishDate,
        )
    }
    br := conn.SendBatch(context.Background(), batch)
    _, err := br.Exec()
    if err != nil {
        return err
    }
    return br.Close()
}