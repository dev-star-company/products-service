package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const brandNotFoundError = "brand not found"

func BrandNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", brandNotFoundError, id))
}

func IsBrandNotFound(err error) bool {
	return strings.Contains(err.Error(), brandNotFoundError)
}
