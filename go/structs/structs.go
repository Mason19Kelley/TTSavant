package structs

type IttfPlayerRanking struct {
	IttfId                 string `json:"IttfId"`
	PlayerName             string `json:"PlayerName"`
	CountryCode            string `json:"CountryCode"`
	CountryName            string `json:"CountryName"`
	AssociationCountryCode string `json:"AssociationCountryCode"`
	AssociationCountryName string `json:"AssociationCountryName"`
	CategoryCode           string `json:"CategoryCode"`
	AgeCategoryCode        string `json:"AgeCategoryCode"`
	SubEventCode           string `json:"SubEventCode"`
	RankingYear            string `json:"RankingYear"`
	RankingMonth           string `json:"RankingMonth"`
	RankingWeek            string `json:"RankingWeek"`
	RankingPointsCareer    *int   `json:"RankingPointsCareer"` // Use pointer to handle null values
	RankingPointsYTD       string `json:"RankingPointsYTD"`
	RankingPosition        string `json:"RankingPosition"`
	CurrentRank            string `json:"CurrentRank"`
	PreviousRank           string `json:"PreviousRank"`
	RankingDifference      string `json:"RankingDifference"`
	PublishDate            string `json:"PublishDate"`
}