package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const productreferencesNotFoundError = "productreferences not found"

func ProductReferencesNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", productreferencesNotFoundError, id))
}

func IsProductReferencesNotFound(err error) bool {
	return strings.Contains(err.Error(), productreferencesNotFoundError)
}
