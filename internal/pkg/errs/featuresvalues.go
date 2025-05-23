package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const featuresvaluesNotFoundError = "featuresvalues not found"

func FeaturesValuesNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", featuresvaluesNotFoundError, id))
}

func IsFeaturesValuesNotFound(err error) bool {
	return strings.Contains(err.Error(), featuresvaluesNotFoundError)
}
