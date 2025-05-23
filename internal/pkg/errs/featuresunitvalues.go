package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const featuresunitvaluesNotFoundError = "featuresunitvalues not found"

func FeaturesUnitValuesNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", featuresunitvaluesNotFoundError, id))
}

func IsFeaturesUnitValuesNotFound(err error) bool {
	return strings.Contains(err.Error(), featuresunitvaluesNotFoundError)
}
