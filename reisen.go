package reisen

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-errors/errors"
)

type Err struct {
	Err *errors.Error
}

func (e *Err) Error() string {
	if os.Getenv("REISEN_VERBOSE") == "" {
		return e.Err.Error()
	}

	frames := e.Err.StackFrames()

	result := []string{
		fmt.Sprintf(
			"%s at %s line %d",
			e.Err.Error(),
			frames[0].File,
			frames[0].LineNumber,
		),
	}

	for _, frame := range frames[1:] {
		if frames[0].Package == frame.Package {
			result = append(result, fmt.Sprintf(
				"\t%s#%s() called at %s line %d",
				frame.Package,
				frame.Name,
				frame.File,
				frame.LineNumber,
			))
		}
	}

	return strings.Join(result, "\n")
}

func Wrap(e interface{}, skip int) error {
	switch t := e.(type) {
	case *Err:
		return t
	}
	return &Err{Err: errors.Wrap(e, skip)}
}

func Error(e interface{}) error {
	return Wrap(e, 2)
}

func Errorf(str string, args ...interface{}) error {
	return Wrap(fmt.Errorf(str, args...), 2)
}
