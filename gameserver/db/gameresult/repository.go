package gameresult

import (
	"github.com/Edwardz43/mygame/gameserver/db/models"
)

// Repository represent the author's repository contract
type Repository interface {
	AddNewOne(gameType int8, run int64, detail string, modID int) (int64, error)
	GetOne(gameType int8, run int64) (*models.GameResult, error)
	GetByRun(gameType int8, runStart int64, runEnd int64) ([]*models.GameResult, error)
}
