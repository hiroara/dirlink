package controller

import (
	"fmt"
	"strings"
)

type aggregatedError struct {
	errors []error
}

func aggregateErrors(errors []error) *aggregatedError {
	if len(errors) == 0 {
		return nil
	}
	return &aggregatedError{errors}
}

func (err *aggregatedError) Error() string {
	if len(err.errors) == 1 {
		return err.errors[0].Error()
	}
	msgs := make([]string, 0, len(err.errors))
	for _, e := range err.errors {
		msgs = append(msgs, e.Error())
	}
	return fmt.Sprintf("Multiple errors occur: %s", strings.Join(msgs, ", "))
}
