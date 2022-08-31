package main

import (
	"fmt"
	builtinLog "log"

	"github.com/pkg/errors"
)

func errorsDemo() {
	builtinLog.Printf("error of task: %v", taskWithoutErrWrap())
	builtinLog.Printf("error of task: %v", taskWithErrWrap())
}

func taskWithErrWrap() error {
	if err := step1(); err != nil {
		return errors.Wrap(err, "task fail")
	}
	if err := step2(); err != nil {
		return errors.Wrap(err, "task fail")
	}
	return nil
}

func taskWithoutErrWrap() error {
	if err := step1(); err != nil {
		return fmt.Errorf("task fail 1")
	}
	if err := step2(); err != nil {
		return fmt.Errorf("task fail 2")
	}
	return nil
}

// ---

func step1() error {
	// return nil
	return fmt.Errorf("step1 fail")
}

func step2() error {
	return fmt.Errorf("step2 fail")
}
