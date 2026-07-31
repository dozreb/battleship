# battleship

## Goal
Battleship wird ein browserbasiertes Schiffeversenken. Der Spieler kann gegen den Computer oder im Multiplayer gegen einen anderen Spieler antreten. Zudem soll man sich registrieren können, um für gewonnene Spiele Punkte zu erhalten. Schließlich ist später noch eine Crossplattform-Möglichkeit vorgesehen, sodass Spieler aus dem Browser heraus auch gegen Spieler, die eine mobile App (Unity) nutzen, spielen können. 

Die App soll stufenweise in Versionen entwickelt werden.

## Requirements v1
Als Spieler möchte ich gegen den Computer spielen.
- Der Computergegner soll serverseitig realisiert werden (besserer Schutz gegen Cheaten, etc.).
- Verhalten des Computergegner: Der Computer feuert zufällig. Bei einem Treffer, versucht er das Schiff zu versenken.
- Modus: Man hat immer nur einen Schuss, danach ist der andere Spieler dran. Bei einem Treffer hat man jedoch einen weiteren Schuss und zwar solange bis man wieder daneben schießt.
- Es gibt folgende Schiffe: ein Schlachtschiff (5 Felder), zwei Kreuzer (je 4 Felder), zwei Zerstörer (je 3 Felder) und drei U-Boote (je 2 Felder)
- Folgende Regeln gelten für das setzen der Schiffe: 
    - Die Schiffe dürfen nicht aneinander stoßen.
    - Die Schiffe dürfen nicht über Eck gebaut sein oder Ausbuchtungen besitzen.
    - Die Schiffe dürfen auch am Rand liegen.
    - Die Schiffe dürfen nicht diagonal aufgestellt werden.
- Die Schiffe des Computers und (vorerst) auch des Spielers werden zufällig platziert.
- Der Spieler nutzt vorerst nur die responsive Ansicht.
- In v1 ist noch keine DB-Anbindung vorgesehen und auch kein Login.

# Requirements v2
TBD

## Frontend
### Responsives Design
- Mockup der Spielfelder (Spieler ist am Zug, Computer ist am Zug) ist in der Battleship-Mockup.png
- Mobile Ansicht (unter 800px): Nur ein Spielfeld. Der Spieler kann seinen Treffer platzieren und bekommt visuelles Feedback, ob es ein Treffer war oder nicht. Danach wechselt das Spielfeld und man sieht, wohin der Comupter zielt. Der Spieler kann in dieser Ansicht seine eigenen Schiffe sehen und erkennt ob diese bereits getroffen bzw. versenkt wurden.
- Desktop / Tablet Ansicht (ab 800px): Beide Spielfelder sind nebeneinander: Spieler und Computeransicht
- Übersicht der feindlichen Schiffe: Der Spieler sieht immer eine Übersicht der feindlichen Schiffe und ob diese bereits versenkt wurden.
  

# Requirements v2
TBD

## Tech-Stack
- Backend: 
    - Go und Gin
    - Rest-API
    - Websockets
- Frontend: React
- Datenbank: PostgreSQL

## V1 Implementation Status
V1 ist als lauffaehiger Startstand umgesetzt.

- Singleplayer gegen serverseitige KI
- 10x10 Board, zufaellige Schiffsplatzierung fuer beide Seiten
- Schiffsregeln: nur horizontal/vertikal, keine Beruehrung (auch diagonal)
- Turn-Logik: Treffer gibt Zusatzschuss, Fehlschuss wechselt den Zug
- Spielende sofort beim Versenken des letzten gegnerischen Schiffs
- Doppelschuss auf dasselbe Feld wird als ungueltig abgelehnt
- Responsive UI:
    - unter 800px nur ein Board mit automatischem Wechsel
    - ab 800px beide Boards nebeneinander
- Treffer/Fehlschuss-Feedback ueber Markierungen im Feld
- Feindschiff-Uebersicht markiert nur vollstaendig versenkte Schiffe

Hinweis: Fuer V1 wird REST-only verwendet. WebSocket, Login und DB sind fuer spaetere Versionen vorgesehen.

## Projektstruktur

- backend: Go + Gin REST-API und Spiel-Engine
- frontend: React + TypeScript UI (Vite)

## Lokal starten

### Backend

```bash
cd backend
go mod tidy
go run .
```

Backend laeuft dann auf `http://localhost:8080`.

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Frontend laeuft standardmaessig auf `http://localhost:5173` und spricht per Default mit `http://localhost:8080`.

Optional kann die API-URL ueber `VITE_API_BASE` gesetzt werden.

## REST API (V1)

- `POST /api/game/new` -> neues Spiel erstellen
- `GET /api/game/:id` -> aktuellen Spielzustand abrufen
- `POST /api/game/:id/shot` -> Spieler-Schuss ausfuehren (`{ "x": 0, "y": 0 }`)

Typische Fehlercodes:
- `400` fuer ungueltige Eingaben (z. B. Koordinaten ausserhalb des Boards)
- `404` falls Spiel-ID nicht existiert
- `409` fuer unzulaessige Aktionen (z. B. Feld bereits beschossen, Spiel bereits beendet)