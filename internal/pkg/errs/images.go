package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const imagesNotFoundError = "images not found"

func ImagesNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", imagesNotFoundError, id))
}

func IsImagesNotFound(err error) bool {
	return strings.Contains(err.Error(), imagesNotFoundError)
}
