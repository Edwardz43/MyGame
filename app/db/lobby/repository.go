package lobby

// Repository ...
type Repository interface {
	GetLatest(gameID int) (int64, int, int8, int8, error)
	Update(gameID int, run int64, inn int, status int) error
	Countdown(gameID int, countdown int8) error
}
