package errs

import (
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const toolsNotFoundError = "tools not found"

func ToolsNotFound(id int) error {
	return status.Error(codes.NotFound, fmt.Sprintf("%s with id: %d", toolsNotFoundError, id))
}

func IsToolsNotFound(err error) bool {
	return strings.Contains(err.Error(), toolsNotFoundError)
}
