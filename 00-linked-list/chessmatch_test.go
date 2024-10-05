package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestInsertMoveAt(t *testing.T) {
	cMatch := NewChessMatch()

	cMatch.insertMoveAtIndex(0, "e4")
	cMatch.insertMoveAtIndex(1, "e5")

	firstMove := cMatch.getMoveAtIndex(0)

	if firstMove == nil || firstMove.position != "e4" {
		t.Error("GetMoveAt returns nil")

	}

	if firstMove.nextMove == nil || firstMove.nextMove.position != "e5" {
		t.Error("GetMoveAt next position is invalid")
	}

	secondMove := cMatch.getMoveAtIndex(1)

	if secondMove == nil || secondMove.position != "e5" {
		t.Error("GetMoveAt incorrect data at index")
	}

}

func TestRemoveAtIndex(t *testing.T) {
	cMatch := NewChessMatch()

	err := cMatch.removeAtIndex(1)

	if err != nil {
		t.Errorf("Move at %d not found", 1)
	}

	cMatch.addMove("e1")
	cMatch.addMove("e2")
	cMatch.addMove("e3")

	err = cMatch.removeAtIndex(1)

	if err != nil {
		t.Errorf("Move at %d not found", 1)
	}

	if cMatch.head.position != "e1" {
		t.Error("Head position incorrect")
	}
	if cMatch.head.nextMove.position != "e3" {
		t.Error("Head next move incorrect")
	}

	cMatch.addMove("e4")
	cMatch.addMove("e5")

	err = cMatch.removeAtIndex(3)

	if cMatch.head.position != "e1" {
		t.Error("Head position incorrect")
	}
	fmt.Println(cMatch.tail)
	if cMatch.tail.position != "e4" {
		t.Error("Tail position is incorrect after deleting last element")
	}
}

func TestGetMoveAtIndex(t *testing.T) {
	cMatch := NewChessMatch()

	cMatch.addMove("e1")
	cMatch.addMove("e2")
	cMatch.addMove("e3")
	cMatch.addMove("e4")
	cMatch.addMove("e5")

	firstMove := cMatch.getMoveAtIndex(0)

	if firstMove == nil {
		t.Error("incorrect move. Returned nil")
	}

	if firstMove.position != "e1" {
		t.Errorf("incorrect move. Returned %s", firstMove.position)
	}

	thirdMove := cMatch.getMoveAtIndex(2)

	if thirdMove == nil {
		t.Error("incorrect move. Returned nil")
	}

	if thirdMove.position != "e3" {
		t.Errorf("incorrect move. Returned %s", thirdMove.position)
	}

	lastMove := cMatch.getMoveAtIndex(4)

	if lastMove == nil {
		t.Error("incorrect move. Returned nil")
	}

	if lastMove.position != "e5" {
		t.Errorf("incorrect move. Returned %s", lastMove.position)
	}
}

func TestGetLastMove(t *testing.T) {
	cMatch := NewChessMatch()

	cMatch.addMove("e1")
	cMatch.addMove("e2")
	cMatch.addMove("e3")
	cMatch.addMove("e4")
	cMatch.addMove("e5")

	cMatch.removeAtIndex(1)

	lastMove := cMatch.getLastMove()

	if lastMove == nil {
		t.Error("Last move is nil")
	}

	if lastMove.position != "e5" {
		t.Errorf("Last move incorrect. Returned %s", lastMove.position)
	}

	cMatch.removeAtIndex(3)

	lastMove = cMatch.getLastMove()

	if lastMove == nil {
		t.Error("Last move is nil")
	}

	if lastMove.position != "e4" {
		t.Errorf("Last move incorrect. Returned %s", lastMove.position)
	}
}

func TestAddMove(t *testing.T) {
	cMatch := NewChessMatch()

	cMatch.addMove("e1")

	if cMatch.head == nil {
		t.Error("Head set incorrectly after 1st addMove. Returned nil")
	}
	if cMatch.head.position != "e1" {
		t.Errorf("Head set incorrectly after 1st addMove. Returned %s", cMatch.head.position)
	}
	if cMatch.tail == nil {
		t.Error("Tail set incorrectly after 1st addMove. Returned nil")
	}
	if cMatch.tail.position != "e1" {
		t.Errorf("Tail set incorrectly after 1st addMove. Returned %s", cMatch.head.position)
	}

	cMatch.addMove("e2")

	if cMatch.tail == nil {
		t.Error("Tail set incorrectly after 1st addMove. Returned nil")
	}
	if cMatch.tail.position != "e2" {
		t.Errorf("Tail set incorrectly after 1st addMove. Returned %s", cMatch.head.position)
	}

	cMatch.addMove("e3")

	if cMatch.tail == nil {
		t.Error("Tail set incorrectly after 1st addMove. Returned nil")
	}
	if cMatch.tail.position != "e3" {
		t.Errorf("Tail set incorrectly after 1st addMove. Returned %s", cMatch.head.position)
	}

	if cMatch.head == nil {
		t.Error("Tail set incorrectly after 1st addMove. Returned nil. Expected e1")
	}
	if cMatch.head.position != "e1" {
		t.Errorf("Head has been incorrectly modified. Returned %s. Expected e1", cMatch.head.position)
	}
}

func TestAddAsHead(t *testing.T) {
	cMatch := NewChessMatch()

	cMatch.addAsHead("e1")

	if cMatch.head == nil {
		t.Error("Head has returned nil")
	}

	if cMatch.head.position != "e1" {
		t.Errorf("Head has returned %s. Expected e1", cMatch.head.position)
	}
	if cMatch.tail == nil {
		t.Error("Tail has returned nil")
	}

	if cMatch.tail.position != "e1" {
		t.Errorf("Tail has returned %s. Expected e1", cMatch.head.position)
	}

	cMatch.addAsHead("e2")

	if cMatch.head.position != "e2" {
		t.Errorf("Head has returned %s. Expected e2", cMatch.head.position)
	}

	if cMatch.head.nextMove.position != "e1" {
		t.Errorf("Head nextMove position has returned %s. Expected e1", cMatch.head.position)
	}

	if cMatch.tail.position != "e1" {
		t.Errorf("Tail has returned %s. Expected e1", cMatch.head.position)
	}

	cMatch.addAsHead("e3")

	if cMatch.head.position != "e3" {
		t.Errorf("Head has returned %s. Expected e2", cMatch.head.position)
	}

	if cMatch.head.nextMove.position != "e2" {
		t.Errorf("Head nextMove position has returned %s. Expected e1", cMatch.head.position)
	}

	if cMatch.tail.position != "e1" {
		t.Errorf("Tail has returned %s. Expected e1", cMatch.head.position)
	}
}

func TestForEachMove(t *testing.T) {
	cMatch := NewChessMatch()

	cMatch.addMove("e1")
	cMatch.addMove("e2")
	cMatch.addMove("e3")
	cMatch.addMove("e4")

	cMatch.forEachMove(func(move *Move) {
		move.position = move.position + "-test"
	})

	pos := 0

	for {
		if pos == 3 {
			break
		}

		targetMovePosition := cMatch.getMoveAtIndex(pos).position
		expectedPosition := targetMovePosition[:2] + "-test"

		if !strings.EqualFold(targetMovePosition, targetMovePosition[:2]+"-test") {

			t.Errorf("For each didn't work. Received %s. Expected %s-test", targetMovePosition, expectedPosition)
		}

		pos++
	}
}
