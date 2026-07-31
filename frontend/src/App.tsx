import { useEffect, useMemo, useRef, useState } from "react";
import { createGame, getGame, shoot } from "./api";
import type { BoardCellView, GameView, ShotResult } from "./types";

function statusText(game: GameView | null): string {
  if (!game) {
    return "Spiel wird gestartet...";
  }
  if (game.over) {
    return game.winner === "player" ? "Du hast gewonnen" : "Computer hat gewonnen";
  }
  return game.turn === "player" ? "Du bist am Zug" : "Computer ist am Zug";
}

function formatShot(result: ShotResult): string {
  const pos = `(${result.point.x}, ${result.point.y})`;
  if (result.hit) {
    if (result.sunk) {
      return `Treffer auf ${pos} - ${result.shipName ?? "Schiff"} versenkt`;
    }
    return `Treffer auf ${pos}`;
  }
  return `Fehlschuss auf ${pos}`;
}

interface BoardProps {
  title: string;
  grid: BoardCellView[][];
  clickable: boolean;
  onShoot?: (x: number, y: number) => void;
}

function Board({ title, grid, clickable, onShoot }: BoardProps) {
  return (
    <section className="board-card">
      <h2>{title}</h2>
      <div className="board-grid" role="grid" aria-label={title}>
        {grid.map((row, y) =>
          row.map((cell, x) => {
            const className = [
              "board-cell",
              cell.hasShip ? "ship" : "",
              cell.hit ? "hit" : "",
              cell.miss ? "miss" : ""
            ]
              .filter(Boolean)
              .join(" ");
            const disabled = !clickable || cell.hit || cell.miss;
            return (
              <button
                type="button"
                key={`${x}-${y}`}
                className={className}
                onClick={() => onShoot?.(x, y)}
                disabled={disabled}
                aria-label={`Feld ${x}, ${y}`}
              >
                {cell.hit ? "X" : cell.miss ? "o" : ""}
              </button>
            );
          })
        )}
      </div>
    </section>
  );
}

export default function App() {
  const [game, setGame] = useState<GameView | null>(null);
  const [busy, setBusy] = useState<boolean>(true);
  const [error, setError] = useState<string>("");
  const [logLines, setLogLines] = useState<string[]>([]);
  const [isMobile, setIsMobile] = useState<boolean>(window.innerWidth < 800);
  const shotLockRef = useRef<boolean>(false);

  useEffect(() => {
    const onResize = () => setIsMobile(window.innerWidth < 800);
    window.addEventListener("resize", onResize);
    return () => window.removeEventListener("resize", onResize);
  }, []);

  useEffect(() => {
    void startNewGame();
  }, []);

  const visibleBoard = useMemo(() => {
    if (!game) {
      return "computer" as const;
    }
    if (!isMobile) {
      return "both" as const;
    }
    return game.turn === "player" ? "computer" : "player";
  }, [game, isMobile]);

  async function startNewGame() {
    try {
      shotLockRef.current = true;
      setBusy(true);
      setError("");
      setLogLines([]);
      const created = await createGame();
      setGame(created);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Spielstart fehlgeschlagen");
    } finally {
      shotLockRef.current = false;
      setBusy(false);
    }
  }

  async function handleShoot(x: number, y: number) {
    if (!game || game.turn !== "player" || game.over || busy || shotLockRef.current) {
      return;
    }

    shotLockRef.current = true;
    try {
      setBusy(true);
      setError("");
      const response = await shoot(game.id, x, y);
      const computerShots = response.computerShots ?? [];
      const newLog: string[] = [
        `Spieler: ${formatShot(response.playerShot)}`,
        ...computerShots.map((s) => `Computer: ${formatShot(s)}`)
      ];
      setLogLines((prev) => [...newLog, ...prev].slice(0, 12));
      setGame(response.game);
    } catch (err) {
      const message = err instanceof Error ? err.message : "Schuss fehlgeschlagen";

      if (message.toLowerCase().includes("already targeted")) {
        try {
          const refreshed = await getGame(game.id);
          setGame(refreshed);
        } catch {
          // Ignore refresh failures and keep original error feedback.
        }
      }

      setError(message);
    } finally {
      shotLockRef.current = false;
      setBusy(false);
    }
  }

  return (
    <main className="app-shell">
      <header className="topbar">
        <h1>Battleship V1</h1>
        <button type="button" onClick={() => void startNewGame()} disabled={busy}>
          Neues Spiel
        </button>
      </header>

      <p className="status">{statusText(game)}</p>
      {error ? <p className="error">{error}</p> : null}

      {game ? (
        <>
          <section className="fleet-strip" aria-label="Feindschiffe">
            {game.enemyFleet.map((ship) => (
              <div key={ship.id} className={`fleet-ship ${ship.sunk ? "sunk" : ""}`}>
                <span>{ship.name}</span>
                <span>{"[]".repeat(ship.length)}</span>
              </div>
            ))}
          </section>

          <section className="boards-layout">
            {(visibleBoard === "both" || visibleBoard === "computer") && (
              <Board
                title="Computerfeld"
                grid={game.computerBoard}
                clickable={game.turn === "player" && !game.over && !busy}
                onShoot={handleShoot}
              />
            )}
            {(visibleBoard === "both" || visibleBoard === "player") && (
              <Board title="Spielerfeld" grid={game.playerBoard} clickable={false} />
            )}
          </section>
        </>
      ) : null}

      <section className="log-card">
        <h2>Letzte Aktionen</h2>
        <ul>
          {logLines.length === 0 ? <li>Noch keine Schuesse</li> : null}
          {logLines.map((line, index) => (
            <li key={`${line}-${index}`}>{line}</li>
          ))}
        </ul>
      </section>
    </main>
  );
}
