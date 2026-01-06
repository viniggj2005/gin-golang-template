package routes

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var Trans ut.Translator

// InitValidator inicializa o tradutor para pt_BR
func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		pt := pt_BR.New()
		uni := ut.New(pt, pt)
		Trans, _ = uni.GetTranslator("pt_BR")
		pt_translations.RegisterDefaultTranslations(v, Trans)
	}
}
