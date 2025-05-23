package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const promotiohasproductNotFoundError = "promotiohasproduct not found"

func PromotionHasProductNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", promotiohasproductNotFoundError, id))
}

func IsPromotionHasProductNotFound(err error) bool {
	return strings.Contains(err.Error(), promotiohasproductNotFoundError)
}
