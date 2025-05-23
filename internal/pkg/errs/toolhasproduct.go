package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const toolhasproductNotFoundError = "toolhasproduct not found"

func ToolHasProductNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", toolhasproductNotFoundError, id))
}

func IsToolHasProductNotFound(err error) bool {
	return strings.Contains(err.Error(), toolhasproductNotFoundError)
}
