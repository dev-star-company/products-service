package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const productpricesNotFoundError = "productprices not found"

func ProductPricesNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", productpricesNotFoundError, id))
}

func IsProductPricesNotFound(err error) bool {
	return strings.Contains(err.Error(), productpricesNotFoundError)
}
