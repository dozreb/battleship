package game

import (
	"strconv"
	"sync"
	"sync/atomic"
)

type Service struct {
	mu      sync.RWMutex
	games   map[string]*Game
	nextID  uint64
}

func NewService() *Service {
	return &Service{
		games: map[string]*Game{},
	}
}

func (s *Service) NewGame() GameView {
	id := strconv.FormatUint(atomic.AddUint64(&s.nextID, 1), 10)
	game := NewGame(id)

	s.mu.Lock()
	s.games[id] = game
	s.mu.Unlock()

	return game.ToView("", nil, nil)
}

func (s *Service) GetGameView(id string) (GameView, error) {
	s.mu.RLock()
	game, ok := s.games[id]
	s.mu.RUnlock()
	if !ok {
		return GameView{}, ErrNotFound
	}
	return game.ToView("", nil, nil), nil
}

func (s *Service) PlayerShot(id string, req ShotRequest) (ShotResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	game, ok := s.games[id]
	if !ok {
		return ShotResponse{}, ErrNotFound
	}

	playerShot, aiShots, err := game.PlayerShoot(Point{X: req.X, Y: req.Y})
	if err != nil {
		return ShotResponse{}, err
	}

	view := game.ToView("", &playerShot, aiShots)
	return ShotResponse{
		Game:          view,
		PlayerShot:    playerShot,
		ComputerShots: aiShots,
	}, nil
}
