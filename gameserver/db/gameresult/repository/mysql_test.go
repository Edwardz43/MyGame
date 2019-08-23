package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	gameserver "github.com/Edwardz43/mygame/gameserver/app"
	"github.com/Edwardz43/mygame/gameserver/db/gameresult/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetByBetNo(t *testing.T) {

}

func TestAddNewOne(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO GameResult").ExpectExec().WithArgs(1, "", time.Now(), 0).WillReturnResult(sqlmock.NewResult(1, 1))

	a := repository.NewMysqlGameResultRepository(db)

	var n int64

	gr := gameserver.GameResult{
		Run:        1,
		Timestamp:  time.Now(),
		GameType:   gameserver.Dice,
		GameDetail: "{d1:1, d2:2, d3:3}",
	}

	if n, err = a.AddNewOne(&gr); err != nil {
		t.Errorf("an error '%s' was not expected when add a new game result", err)
	}
	assert.NotZero(t, n)
	assert.NotEqual(t, -1, n)
}
