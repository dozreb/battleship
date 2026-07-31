package game

const BoardSize = 10

const (
	TurnPlayer   = "player"
	TurnComputer = "computer"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Ship struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Length int     `json:"length"`
	Cells  []Point `json:"cells"`
}

type ShotResult struct {
	Point    Point  `json:"point"`
	Hit      bool   `json:"hit"`
	Sunk     bool   `json:"sunk"`
	ShipID   string `json:"shipId,omitempty"`
	ShipName string `json:"shipName,omitempty"`
}

type BoardCellView struct {
	HasShip bool `json:"hasShip"`
	Hit     bool `json:"hit"`
	Miss    bool `json:"miss"`
}

type EnemyShipStatus struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Length int    `json:"length"`
	Sunk   bool   `json:"sunk"`
}

type GameView struct {
	ID             string            `json:"id"`
	Turn           string            `json:"turn"`
	Over           bool              `json:"over"`
	Winner         string            `json:"winner,omitempty"`
	PlayerBoard    [][]BoardCellView `json:"playerBoard"`
	ComputerBoard  [][]BoardCellView `json:"computerBoard"`
	EnemyFleet     []EnemyShipStatus `json:"enemyFleet"`
	LastError      string            `json:"lastError,omitempty"`
	LastPlayerShot *ShotResult       `json:"lastPlayerShot,omitempty"`
	LastAIShots    []ShotResult      `json:"lastAiShots,omitempty"`
}

type ShotRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type ShotResponse struct {
	Game         GameView    `json:"game"`
	PlayerShot   ShotResult  `json:"playerShot"`
	ComputerShots []ShotResult `json:"computerShots"`
}
