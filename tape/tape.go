package tape

// Tape implements an "infinite" tape for the turing machine with a head control.
type Tape struct {
	head int64
	tape []byte
}

// RTape is a interface implemented by the Tape struct. Only supports Read operations on the tape.
type RTape interface {
	MoveRight()
	MoveLeft()
	ReadSymbol() byte
}

const (
	BlankSymbol     = byte('$')
	initialTapeSize = 1000
)

func (t *Tape) isHeadLocationOnTape() bool {
	return t.head >= 0 && t.head < int64(len(t.tape))
}

/* increaseTapeSize increases the size of the tape to make head valid.
if the head points to a location not in the tape, the tape size has to be increased in case of write. */
func (t *Tape) increaseTapeSize() {
	if t.isHeadLocationOnTape() {
		return
	}

	if t.head < 0 {
		newTapeSize := Max(2*Abs(t.head), int64(len(t.tape)))

		newTape := make([]byte, newTapeSize)
		for newTapeIndex := range newTape {
			newTape[newTapeIndex] = BlankSymbol
		}

		t.tape = append(newTape, t.tape...)
		t.head = newTapeSize - Abs(t.head)
	} else {
		tapeLen := int64(len(t.tape))
		newTapeSize := Max(2*(t.head-tapeLen+1), tapeLen)

		newTape := make([]byte, newTapeSize)
		for newTapeIndex := range newTape {
			newTape[newTapeIndex] = BlankSymbol
		}

		t.tape = append(t.tape, newTape...)
	}
}

func (t *Tape) MoveRight() {
	t.head++
}

func (t *Tape) MoveLeft() {
	t.head--
}

func (t *Tape) ReadSymbol() byte {
	if !t.isHeadLocationOnTape() {
		return BlankSymbol
	}

	return t.tape[t.head]
}

// WriteSymbol writes the given symbol on the tape in the place pointed by the head.
func (t *Tape) WriteSymbol(symbol byte) {
	if !t.isHeadLocationOnTape() {
		t.increaseTapeSize()
	}

	t.tape[t.head] = symbol
}

func NewTape() *Tape {
	tape := make([]byte, initialTapeSize)

	for tapeIndex := range tape {
		tape[tapeIndex] = BlankSymbol
	}

	return &Tape{
		head: int64(initialTapeSize / 2),
		tape: tape,
	}
}
