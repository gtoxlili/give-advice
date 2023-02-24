package validate

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gtoxlili/give-advice/common/pool"
)

var (
	validate = validator.New()
	trans    ut.Translator
)

func init() {
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	translations.RegisterDefaultTranslations(validate, trans)
}

func Struct(s any) error {
	r := validate.Struct(s)
	if r != nil {
		err := r.(validator.ValidationErrors)
		sb := pool.GetBuffer()
		defer pool.PutBuffer(sb)
		for _, e := range err {
			sb.WriteString(e.Translate(trans))
			sb.WriteString(" | ")
		}
		sb.Truncate(sb.Len() - 3)
		return errors.New(sb.String())
	}
	return nil
}
