package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const featuresvaluestypesNotFoundError = "featuresvaluestypes not found"

func FeaturesValuesTypesNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", featuresvaluestypesNotFoundError, id))
}

func IsFeaturesValuesTypesNotFound(err error) bool {
	return strings.Contains(err.Error(), featuresvaluestypesNotFoundError)
}
