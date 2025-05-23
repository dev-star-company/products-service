package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const producthasproductreferenceNotFoundError = "producthasproductreference not found"

func ProductHasProductReferenceNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", producthasproductreferenceNotFoundError, id))
}

func IsProductHasProductReferenceNotFound(err error) bool {
	return strings.Contains(err.Error(), producthasproductreferenceNotFoundError)
}
