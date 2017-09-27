package brackets

import (
	"container/list"
	"fmt"
)

const testVersion = 5

type bracketStack struct {
	bracketList *list.List
}

func newStack() *bracketStack {
	return &bracketStack{
		bracketList: list.New(),
	}
}

func (b *bracketStack) Push(item rune) {
	b.bracketList.PushFront(item)
}

func (b *bracketStack) Pop() (rune, error) {
	if b.isEmpty() {
		return '\000', fmt.Errorf("stack is empty")
	}
	e := b.bracketList.Front()
	b.bracketList.Remove(e)
	return e.Value.(rune), nil
}

func (b *bracketStack) isEmpty() bool {
	fmt.Printf("stack len: %d\n", b.bracketList.Len())
	return b.bracketList.Len() == 0
}

func Bracket(input string) (bool, error) {
	stack := newStack()

	for _, bracket := range input {
		if bracket == '{' || bracket == '[' || bracket == '(' {
			stack.Push(bracket)
		}

		if bracket == '}' || bracket == ']' || bracket == ')' {
			v, err := stack.Pop()
			if err != nil {
				return false, nil
			}

			if (bracket == '}' && v != '{') || (bracket == ']' && v != '[') || (bracket == ')' && v != '(') {
				return false, nil
			}
		}
	}

	return stack.isEmpty(), nil
}
