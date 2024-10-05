package main

type Move struct {
	position string
	nextMove *Move
}

func newMove(position string, nextMove *Move) *Move {
	move := &Move{position, nextMove}

	return move
}

type ChessMatch struct {
	head *Move
	tail *Move
}

func NewChessMatch() *ChessMatch {
	return &ChessMatch{}
}

func (c *ChessMatch) getMoveAtIndex(index int) *Move {
	if c.head == nil {
		return nil
	}

	if index == 0 {
		return c.head
	}

	move := c.head

	pos := 0

	for pos < index {
		move = move.nextMove
		pos++
	}

	return move
}

func (c *ChessMatch) getLastMove() *Move {
	return c.tail
}

func (c *ChessMatch) addMove(position string) {
	if c.head == nil {
		c.head, c.tail = newMove(position, nil), newMove(position, nil)
		return
	}

	lastMove := c.getLastMove()
	lastMove.nextMove = newMove(position, nil)

	if c.head != nil && c.head.nextMove == nil {
		c.head.nextMove = lastMove.nextMove
	}

	c.tail = lastMove.nextMove
}

func (c *ChessMatch) addAsHead(position string) {
	if c.head == nil {
		c.head, c.tail = newMove(position, nil), newMove(position, nil)
	}

	c.head = newMove(position, c.head)
}

func (c *ChessMatch) insertMoveAtIndex(index int, position string) {
	if c.head == nil {
		c.head, c.tail = newMove(position, nil), newMove(position, nil)
		return
	}

	if index == 0 {
		head := c.head
		c.head = newMove(position, head)
		return
	}

	move := c.getMoveAtIndex(index - 1)

	if move == nil {
		return
	}

	move.nextMove = newMove(position, move.nextMove)
}

func (c *ChessMatch) removeAtIndex(index int) error {
	if c.head == nil {
		return nil
	}

	if index == 0 {
		c.head = c.head.nextMove
		return nil
	}

	prevMove := c.getMoveAtIndex(index - 1)

	if prevMove != nil && prevMove.nextMove != nil {

		prevMove.nextMove = prevMove.nextMove.nextMove

		// Reassign tail if last element has been removed
		if prevMove.nextMove == nil {
			c.tail = prevMove
		}
	}

	return nil
}

func (c *ChessMatch) forEachMove(predicateFn func(*Move)) {
	pos := 0

	for {
		move := c.getMoveAtIndex(pos)
		if move == nil {
			break
		}
		predicateFn(move)
		pos++
	}
}
