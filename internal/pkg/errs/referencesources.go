package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const referencesourcesNotFoundError = "referencesources not found"

func ReferenceSourcesNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", referencesourcesNotFoundError, id))
}

func IsReferenceSourcesNotFound(err error) bool {
	return strings.Contains(err.Error(), referencesourcesNotFoundError)
}
