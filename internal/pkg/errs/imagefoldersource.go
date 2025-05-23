package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const imagefoldersourceNotFoundError = "imagefoldersource not found"

func ImageFolderSourceNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", imagefoldersourceNotFoundError, id))
}

func IsImageFolderSourceNotFound(err error) bool {
	return strings.Contains(err.Error(), imagefoldersourceNotFoundError)
}
