package services

import (
	"go-api-ligas/models"
	"go-api-ligas/repository"
)

func CreateGameStatistics(input struct {
	GameID        string  `json:"game_id"`
	PlayerID      string  `json:"player_id"`
	Points        int     `json:"points"`
	Rebounds      int     `json:"rebounds"`
	Assists       int     `json:"assists"`
	Steals        int     `json:"steals"`
	Blocks        int     `json:"blocks"`
	Fouls         int     `json:"fouls"`
	MinutesPlayed float32 `json:"minutes_played"`
}) (models.GameStatistics, error) {
	statistics := models.GameStatistics{
		GameID:        input.GameID,
		PlayerID:      input.PlayerID,
		Points:        input.Points,
		Rebounds:      input.Rebounds,
		Assists:       input.Assists,
		Steals:        input.Steals,
		Blocks:        input.Blocks,
		Fouls:         input.Fouls,
		MinutesPlayed: input.MinutesPlayed,
	}

	err := repository.CreateGameStatistics(&statistics)
	return statistics, err
}

func GetGameStatistics() ([]models.GameStatistics, error) {
	return repository.GetGameStatistics()
}

func GetGameStatisticsByID(id string) (models.GameStatistics, error) {
	return repository.GetGameStatisticsByID(id)
}

func UpdateGameStatistics(id string, updates map[string]interface{}) (models.GameStatistics, error) {
	return repository.UpdateGameStatistics(id, updates)
}

func DeleteGameStatistics(id string) error {
	return repository.DeleteGameStatistics(id)
}
