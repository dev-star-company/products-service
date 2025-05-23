package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const varianttypeNotFoundError = "varianttype not found"

func VariantTypeNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", varianttypeNotFoundError, id))
}

func IsVariantTypeNotFound(err error) bool {
	return strings.Contains(err.Error(), varianttypeNotFoundError)
}
