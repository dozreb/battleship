package game

import (
	"math/rand"
	"time"
)

type Game struct {
	ID            string
	PlayerBoard   *Board
	ComputerBoard *Board
	Turn          string
	Over          bool
	Winner        string

	aiTargets []Point
	rng       *rand.Rand
}

func NewGame(id string) *Game {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Game{
		ID:            id,
		PlayerBoard:   NewRandomBoard(rng),
		ComputerBoard: NewRandomBoard(rng),
		Turn:          TurnPlayer,
		rng:           rng,
	}
}

func (g *Game) ToView(lastError string, lastPlayerShot *ShotResult, lastAIShots []ShotResult) GameView {
	enemyFleet := make([]EnemyShipStatus, 0, len(g.ComputerBoard.Ships))
	for _, ship := range g.ComputerBoard.Ships {
		enemyFleet = append(enemyFleet, EnemyShipStatus{
			ID:     ship.ID,
			Name:   ship.Name,
			Length: ship.Length,
			Sunk:   g.ComputerBoard.isShipSunk(ship),
		})
	}

	return GameView{
		ID:             g.ID,
		Turn:           g.Turn,
		Over:           g.Over,
		Winner:         g.Winner,
		PlayerBoard:    g.PlayerBoard.toView(true),
		ComputerBoard:  g.ComputerBoard.toView(false),
		EnemyFleet:     enemyFleet,
		LastError:      lastError,
		LastPlayerShot: lastPlayerShot,
		LastAIShots:    lastAIShots,
	}
}

func (g *Game) PlayerShoot(target Point) (ShotResult, []ShotResult, error) {
	if g.Over {
		return ShotResult{}, nil, ErrGameOver
	}
	if g.Turn != TurnPlayer {
		return ShotResult{}, nil, ErrNotPlayersTurn
	}

	result, err := g.ComputerBoard.Shoot(target)
	if err != nil {
		return ShotResult{}, nil, err
	}

	if g.ComputerBoard.allShipsSunk() {
		g.Over = true
		g.Winner = TurnPlayer
		return result, []ShotResult{}, nil
	}

	if result.Hit {
		return result, []ShotResult{}, nil
	}

	g.Turn = TurnComputer
	aiShots := g.playComputerTurn()
	return result, aiShots, nil
}

func (g *Game) playComputerTurn() []ShotResult {
	results := []ShotResult{}

	for {
		target := g.chooseAITarget()
		result, err := g.PlayerBoard.Shoot(target)
		if err != nil {
			continue
		}

		results = append(results, result)
		if result.Hit {
			if result.Sunk {
				g.aiTargets = nil
			} else {
				g.enqueueNeighbors(target)
			}
		}

		if g.PlayerBoard.allShipsSunk() {
			g.Over = true
			g.Winner = TurnComputer
			g.Turn = TurnComputer
			return results
		}

		if !result.Hit {
			g.Turn = TurnPlayer
			return results
		}
	}
}

func (g *Game) chooseAITarget() Point {
	for len(g.aiTargets) > 0 {
		candidate := g.aiTargets[0]
		g.aiTargets = g.aiTargets[1:]
		if !inBounds(candidate) {
			continue
		}
		if _, exists := g.PlayerBoard.Shots[key(candidate)]; exists {
			continue
		}
		return candidate
	}

	for {
		candidate := Point{X: g.rng.Intn(BoardSize), Y: g.rng.Intn(BoardSize)}
		if _, exists := g.PlayerBoard.Shots[key(candidate)]; exists {
			continue
		}
		return candidate
	}
}

func (g *Game) enqueueNeighbors(p Point) {
	neighbors := []Point{
		{X: p.X + 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y},
		{X: p.X, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
	}

	for i := len(neighbors) - 1; i >= 0; i-- {
		n := neighbors[i]
		if !inBounds(n) {
			continue
		}
		if _, exists := g.PlayerBoard.Shots[key(n)]; exists {
			continue
		}
		if containsPoint(g.aiTargets, n) {
			continue
		}
		g.aiTargets = append([]Point{n}, g.aiTargets...)
	}
}

func containsPoint(points []Point, target Point) bool {
	for _, p := range points {
		if p == target {
			return true
		}
	}
	return false
}
