package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const pricetypeNotFoundError = "pricetype not found"

func PriceTypeNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", pricetypeNotFoundError, id))
}

func IsPriceTypeNotFound(err error) bool {
	return strings.Contains(err.Error(), pricetypeNotFoundError)
}
