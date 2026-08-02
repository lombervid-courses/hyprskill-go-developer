package main

import "errors"

type Queue struct {
	input  Stack
	output Stack
}

func (q *Queue) Push(value int) {
	q.input.Push(value)
}

func (q *Queue) Pop() (int, error) {
	outputVal, outputErr := q.output.Pop()
	if outputErr == nil {
		return outputVal, nil
	}

	inputVal, inputErr := q.input.Pop()
	if inputErr != nil {
		return 0, errors.New("Queue is empty") // if input stack is empty
	}

	// if the output stack is empty but the input is not empty
	for inputErr == nil { // while input stack not empty...
		q.output.Push(inputVal)            // rearrange input to output
		inputVal, inputErr = q.input.Pop() // and read again
	}

	return q.output.Pop()
}
