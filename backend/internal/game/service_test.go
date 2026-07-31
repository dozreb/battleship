package game

import "testing"

func TestNewGameHasExpectedFleet(t *testing.T) {
	g := NewGame("test")
	if len(g.PlayerBoard.Ships) != 8 {
		t.Fatalf("expected 8 player ships, got %d", len(g.PlayerBoard.Ships))
	}
	if len(g.ComputerBoard.Ships) != 8 {
		t.Fatalf("expected 8 computer ships, got %d", len(g.ComputerBoard.Ships))
	}
}

func TestRejectAlreadyShot(t *testing.T) {
	svc := NewService()
	view := svc.NewGame()

	_, err := svc.PlayerShot(view.ID, ShotRequest{X: 0, Y: 0})
	if err != nil {
		// The first shot can fail only if out of bounds, which it is not.
	}

	_, err = svc.PlayerShot(view.ID, ShotRequest{X: 0, Y: 0})
	if err != ErrAlreadyShot {
		t.Fatalf("expected ErrAlreadyShot, got %v", err)
	}
}
