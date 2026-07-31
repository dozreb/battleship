export type Turn = "player" | "computer";

export interface BoardCellView {
  hasShip: boolean;
  hit: boolean;
  miss: boolean;
}

export interface EnemyShipStatus {
  id: string;
  name: string;
  length: number;
  sunk: boolean;
}

export interface ShotResult {
  point: { x: number; y: number };
  hit: boolean;
  sunk: boolean;
  shipId?: string;
  shipName?: string;
}

export interface GameView {
  id: string;
  turn: Turn;
  over: boolean;
  winner?: Turn;
  playerBoard: BoardCellView[][];
  computerBoard: BoardCellView[][];
  enemyFleet: EnemyShipStatus[];
  lastError?: string;
  lastPlayerShot?: ShotResult;
  lastAiShots?: ShotResult[];
}

export interface ShotResponse {
  game: GameView;
  playerShot: ShotResult;
  computerShots: ShotResult[];
}
