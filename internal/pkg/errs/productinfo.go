package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const productinfoNotFoundError = "productinfo not found"

func ProductInfoNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", productinfoNotFoundError, id))
}

func IsProductInfoNotFound(err error) bool {
	return strings.Contains(err.Error(), productinfoNotFoundError)
}
