package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const productsNotFoundError = "products not found"

func ProductsNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", productsNotFoundError, id))
}

func IsProductsNotFound(err error) bool {
	return strings.Contains(err.Error(), productsNotFoundError)
}
