package game

import (
	"fmt"
	"math/rand"
)

type shipSpec struct {
	Name   string
	Length int
	Count  int
}

var fleetSpec = []shipSpec{
	{Name: "Schlachtschiff", Length: 5, Count: 1},
	{Name: "Kreuzer", Length: 4, Count: 2},
	{Name: "Zerstoerer", Length: 3, Count: 2},
	{Name: "U-Boot", Length: 2, Count: 3},
}

type Board struct {
	Ships []Ship
	Shots map[string]ShotResult
}

func NewRandomBoard(rng *rand.Rand) *Board {
	ships := make([]Ship, 0, 8)
	occupied := make(map[string]bool)

	for _, spec := range fleetSpec {
		for i := 0; i < spec.Count; i++ {
			ship := placeShip(rng, spec, i, occupied)
			ships = append(ships, ship)
			for _, cell := range ship.Cells {
				occupied[key(cell)] = true
			}
		}
	}

	return &Board{
		Ships: ships,
		Shots: map[string]ShotResult{},
	}
}

func placeShip(rng *rand.Rand, spec shipSpec, idx int, occupied map[string]bool) Ship {
	for {
		horizontal := rng.Intn(2) == 0
		maxX := BoardSize - 1
		maxY := BoardSize - 1
		if horizontal {
			maxX = BoardSize - spec.Length
		} else {
			maxY = BoardSize - spec.Length
		}

		start := Point{X: rng.Intn(maxX + 1), Y: rng.Intn(maxY + 1)}
		cells := make([]Point, 0, spec.Length)
		valid := true

		for offset := 0; offset < spec.Length; offset++ {
			p := start
			if horizontal {
				p.X += offset
			} else {
				p.Y += offset
			}
			if !isValidPlacementCell(p, occupied) {
				valid = false
				break
			}
			cells = append(cells, p)
		}

		if !valid {
			continue
		}

		return Ship{
			ID:     fmt.Sprintf("%s-%d", spec.Name, idx+1),
			Name:   spec.Name,
			Length: spec.Length,
			Cells:  cells,
		}
	}
}

func isValidPlacementCell(p Point, occupied map[string]bool) bool {
	if !inBounds(p) {
		return false
	}
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			neighbor := Point{X: p.X + dx, Y: p.Y + dy}
			if !inBounds(neighbor) {
				continue
			}
			if occupied[key(neighbor)] {
				return false
			}
		}
	}
	return true
}

func (b *Board) Shoot(p Point) (ShotResult, error) {
	if !inBounds(p) {
		return ShotResult{}, ErrInvalidCoordinate
	}
	k := key(p)
	if _, exists := b.Shots[k]; exists {
		return ShotResult{}, ErrAlreadyShot
	}

	result := ShotResult{Point: p}
	for _, ship := range b.Ships {
		if shipContains(ship, p) {
			result.Hit = true
			result.ShipID = ship.ID
			result.ShipName = ship.Name
			break
		}
	}

	b.Shots[k] = result

	if result.Hit {
		ship := b.shipByID(result.ShipID)
		result.Sunk = b.isShipSunk(ship)
		b.Shots[k] = result
	}

	return result, nil
}

func (b *Board) shipByID(id string) Ship {
	for _, ship := range b.Ships {
		if ship.ID == id {
			return ship
		}
	}
	return Ship{}
}

func (b *Board) isShipSunk(ship Ship) bool {
	if ship.ID == "" {
		return false
	}
	for _, cell := range ship.Cells {
		shot, ok := b.Shots[key(cell)]
		if !ok || !shot.Hit {
			return false
		}
	}
	return true
}

func (b *Board) allShipsSunk() bool {
	for _, ship := range b.Ships {
		if !b.isShipSunk(ship) {
			return false
		}
	}
	return true
}

func (b *Board) toView(showShips bool) [][]BoardCellView {
	grid := make([][]BoardCellView, BoardSize)
	for y := 0; y < BoardSize; y++ {
		row := make([]BoardCellView, BoardSize)
		for x := 0; x < BoardSize; x++ {
			p := Point{X: x, Y: y}
			cell := BoardCellView{}
			if showShips && b.hasShipAt(p) {
				cell.HasShip = true
			}
			if shot, ok := b.Shots[key(p)]; ok {
				if shot.Hit {
					cell.Hit = true
				} else {
					cell.Miss = true
				}
			}
			row[x] = cell
		}
		grid[y] = row
	}
	return grid
}

func (b *Board) hasShipAt(p Point) bool {
	for _, ship := range b.Ships {
		if shipContains(ship, p) {
			return true
		}
	}
	return false
}

func shipContains(ship Ship, p Point) bool {
	for _, cell := range ship.Cells {
		if cell == p {
			return true
		}
	}
	return false
}

func key(p Point) string {
	return fmt.Sprintf("%d:%d", p.X, p.Y)
}

func inBounds(p Point) bool {
	return p.X >= 0 && p.X < BoardSize && p.Y >= 0 && p.Y < BoardSize
}
