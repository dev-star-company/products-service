package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const infotypesNotFoundError = "infotypes not found"

func InfoTypesNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", infotypesNotFoundError, id))
}

func IsInfoTypesNotFound(err error) bool {
	return strings.Contains(err.Error(), infotypesNotFoundError)
}
