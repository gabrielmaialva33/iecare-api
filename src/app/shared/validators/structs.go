package validators

import (
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptbr2 "github.com/go-playground/validator/v10/translations/pt_BR"
	"iecare-api/src/app/shared/utils"
	"iecare-api/src/database"
	"reflect"
	"regexp"
)

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Field       string `json:"field"`
	Value       string `json:"value"`
	Param       string `json:"param"`
	Message     string `json:"message"`
}

var validate *validator.Validate

// ValidateStruct validates a struct (all the fields)
func ValidateStruct(model interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	//var messages []string

	validate = validator.New()

	// Register custom validators
	_ = validate.RegisterValidation("unique", Unique)

	// Register translations
	pt := pt_BR.New()
	uni := ut.New(pt, pt)
	trans, _ := uni.GetTranslator("pt_BR")
	if err := ptbr2.RegisterDefaultTranslations(validate, trans); err != nil {
		panic(err)
	}

	// Custom translations for custom tags
	_ = validate.RegisterTranslation("unique", trans, func(ut ut.Translator) error {
		return ut.Add("unique", "{0} já está em uso", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("unique", fe.Field())
		return t
	})

	if err := validate.Struct(model); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var response ErrorResponse
			response.FailedField = err.StructNamespace()
			response.Tag = err.Tag()
			response.Field = utils.Underscore(err.Field())
			response.Value = err.Value().(string)
			response.Param = err.Param()
			response.Message = err.Translate(trans)
			errors = append(errors, &response)
			//messages = append(messages, response.Message)
		}
	}

	return errors //, messages
}

// ValidatePartialStruct validates a partial struct (only the fields that are present)
func ValidatePartialStruct(model interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate = validator.New()
	_ = validate.RegisterValidation("unique", Unique)

	var fields []string
	val := reflect.Indirect(reflect.ValueOf(model))
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Interface() != "" {
			fields = append(fields, val.Type().Field(i).Name)
		}
	}

	if err := validate.StructPartial(model, fields...); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var response ErrorResponse
			response.FailedField = err.StructNamespace()
			response.Tag = err.Tag()
			response.Field = utils.Underscore(err.Field())
			response.Value = err.Value().(string)
			response.Param = err.Param()
			errors = append(errors, &response)
		}
	}

	return errors
}

// Unique checks if a field is unique in the database
func Unique(fl validator.FieldLevel) bool {
	var count int64

	model := fl.Top().Interface()
	field := utils.Underscore(fl.StructFieldName())
	value := fl.Field().String()

	database.DB.Model(model).Where(field+" = ?", value).Count(&count)

	return count == 0
}

// UUID checks if a field is a valid UUID
func UUID(fl validator.FieldLevel) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(fl.Field().String())
}
