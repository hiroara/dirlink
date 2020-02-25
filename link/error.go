package link

import (
	"fmt"
)

type linkFailure struct {
	link  *Link
	cause error
}

func (err *linkFailure) Error() string {
	return fmt.Sprintf("Link failure (%s -> %s): %s", err.link.Src, err.link.Dest, err.cause)
}
