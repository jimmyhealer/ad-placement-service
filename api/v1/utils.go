package v1

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func errorMessageHandler(err error, requestErrorMessages map[string]string) string {
	if err == nil {
		return ""
	}

	errs := err.(validator.ValidationErrors)
	err_msgs := make([]string, 0)
	for _, e := range errs {
		field := e.Field()
		// if field is array, need remove "[*]" from field name
		if strings.Contains(field, "[") {
			field = field[:strings.Index(field, "[")]
		}
		key := field + "." + e.Tag()
		if msg, ok := requestErrorMessages[key]; ok {
			err_msgs = append(err_msgs, msg)
		} else {
			err_msgs = append(err_msgs, e.Error())
		}
	}
	return strings.Join(err_msgs, "; ")
}
