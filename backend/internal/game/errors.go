package game

import "errors"

var (
	ErrNotFound         = errors.New("game not found")
	ErrGameOver         = errors.New("game is over")
	ErrNotPlayersTurn   = errors.New("it is not the player's turn")
	ErrInvalidCoordinate = errors.New("coordinate out of bounds")
	ErrAlreadyShot      = errors.New("field was already targeted")
)
