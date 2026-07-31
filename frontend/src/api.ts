import type { GameView, ShotResponse } from "./types";

const API_BASE = import.meta.env.VITE_API_BASE ?? "http://localhost:8080";

async function parseJson<T>(response: Response): Promise<T> {
  if (!response.ok) {
    const payload = (await response.json().catch(() => ({}))) as { error?: string };
    throw new Error(payload.error ?? "Unbekannter API-Fehler");
  }
  return response.json() as Promise<T>;
}

export async function createGame(): Promise<GameView> {
  const response = await fetch(`${API_BASE}/api/game/new`, {
    method: "POST"
  });
  return parseJson<GameView>(response);
}

export async function getGame(id: string): Promise<GameView> {
  const response = await fetch(`${API_BASE}/api/game/${id}`);
  return parseJson<GameView>(response);
}

export async function shoot(id: string, x: number, y: number): Promise<ShotResponse> {
  const response = await fetch(`${API_BASE}/api/game/${id}/shot`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ x, y })
  });
  return parseJson<ShotResponse>(response);
}
