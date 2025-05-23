package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const categoryNotFoundError = "category not found"

func CategoryNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", categoryNotFoundError, id))
}

func IsCategoryNotFound(err error) bool {
	return strings.Contains(err.Error(), categoryNotFoundError)
}
