package clock

const testVersion = 4

// You can find more details and hints in the test file.

type Clock struct {
	h int
	m int
}

func New(hour, minute int) Clock {
	return Clock{h: hour, m: minute}
}

func (Clock) String() string {
}

func (Clock) Add(minutes int) Clock {
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
