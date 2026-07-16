# battleship

## Goal
Battleship wird ein browserbasiertes Schiffeversenken. Der Spieler kann gegen den Computer oder im Multiplayer gegen einen anderen Spieler antreten. Zudem soll man sich registrieren können, um für gewonnene Spiele Punkte zu erhalten. Schließlich ist später noch eine Crossplattform-Möglichkeit vorgesehen, sodass Spieler aus dem Browser heraus auch gegen Spieler, die eine mobile App (Unity) nutzen, spielen können. 

Die App soll stufenweise in Versionen entwickelt werden.

## Requirements v1
Als Spieler möchte ich gegen den Computer spielen.
- Der Computergegner soll serverseitig realisiert werden (besserer Schutz gegen Cheaten, etc.).
- Verhalten des Computergegner: Der Computer feuert zufällig. Bei einem Treffer, versucht er das Schiff zu versenken.
- Modus: Man hat immer nur einen Schuss, danach ist der andere Spieler dran. Bei einem Treffer hat man jedoch einen weiteren Schuss und zwar solange bis man wieder daneben schießt.

# Requirements v2
TBD

## Frontend
### Responsives Design
- Mobile Ansicht (unter 800px): Nur ein Spielfeld. Der Spieler kann seinen Treffer platzieren und bekommt visuelles Feedback, ob es ein Treffer war oder nicht. Danach wechselt das Spielfeld und man sieht, wohin der Comupter zielt. Der Spieler kann in dieser Ansicht seine eigenen Schiffe sehen und erkennt ob diese bereits getroffen bzw. versenkt wurden.
- Desktop / Tablet Ansicht (ab 800px): Beide Spielfelder sind nebeneinander: Spieler und Computeransicht  

# Requirements v2
TBD

## Tech-Stack
- C# mit .NET 8 Framework
- MVVM Architektur
- MainWindow für alle Views außer Login / Register
- MySQL Datenbank
- Datenbankanbindung über MySQL.Data
- BCrypt.Net-Next für Login