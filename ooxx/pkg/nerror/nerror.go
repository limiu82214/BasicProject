package nerror

import (
	"fmt"

	"github.com/pkg/errors"
)

func PrettyError(err error) string {
	errMsg := ""
	errCause := errors.Cause(err)
	errMsg += fmt.Sprintf("%v\n", err)
	errMsg += fmt.Sprintf("%+v\n", errCause)

	return errMsg
}
