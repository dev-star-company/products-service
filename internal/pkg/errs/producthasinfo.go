package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const producthasinfoNotFoundError = "producthasinfo not found"

func ProductHasInfoNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", producthasinfoNotFoundError, id))
}

func IsProductHasInfoNotFound(err error) bool {
	return strings.Contains(err.Error(), producthasinfoNotFoundError)
}
