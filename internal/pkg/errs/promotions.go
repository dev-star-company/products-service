package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const promotionsNotFoundError = "promotions not found"

func PromotionsNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", promotionsNotFoundError, id))
}

func IsPromotionsNotFound(err error) bool {
	return strings.Contains(err.Error(), promotionsNotFoundError)
}
