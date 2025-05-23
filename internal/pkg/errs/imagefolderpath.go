package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const imagefolderpathNotFoundError = "imagefolderpath not found"

func ImageFolderPathNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", imagefolderpathNotFoundError, id))
}

func IsImageFolderPathNotFound(err error) bool {
	return strings.Contains(err.Error(), imagefolderpathNotFoundError)
}
