package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const producthasfeatureNotFoundError = "producthasfeature not found"

func ProductHasFeatureNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", producthasfeatureNotFoundError, id))
}

func IsProductHasFeatureNotFound(err error) bool {
	return strings.Contains(err.Error(), producthasfeatureNotFoundError)
}
