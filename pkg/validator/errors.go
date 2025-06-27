package validator

import (
	"strings"

	"github.com/gobuffalo/validate"
)

func FormatErrors(errors *validate.Errors) string {
	var res string
	for _, v := range errors.Errors {
		res = res + strings.Join(v, ", ") + ", "
		if len(v) > 1 {
		}
	}
	res = strings.TrimRight(res, ", ")
	return res
}
