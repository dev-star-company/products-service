package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const featuresNotFoundError = "features not found"

func FeaturesNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", featuresNotFoundError, id))
}

func IsFeaturesNotFound(err error) bool {
	return strings.Contains(err.Error(), featuresNotFoundError)
}
