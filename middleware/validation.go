package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type ValidationError struct {
	engine *validator.Validate
}
type ValidationInterface interface {
	BindAndValidate(c *gin.Context, payload any) map[string]string
}
type EnumBool interface {
	IsValidEnum() bool
}
type EnumApproval interface {
	IsValid() bool
}

type Status string

const (
	Approve Status = "approve" // add + 1 otherwise validation won't work for 0
	Pending Status = "pending"
	Reject  Status = "reject"
)

func (s *Status) IsValid() bool {
	switch *s {
	case Approve, Pending, Reject:
		return true
	}
	return false
}

type BoolType string

const (
	True  BoolType = "True"
	False BoolType = "False"
)

func (b *BoolType) IsValidEnum() bool {
	switch *b {
	case True, False:
		return true
	}
	return false
}
func ValidateEnumApproval(fl validator.FieldLevel) bool {
	status := Status(fl.Field().Interface().(string))
	return status.IsValid()
}
func ValidateEnumFlag(fl validator.FieldLevel) bool {
	boolType := BoolType(fl.Field().Interface().(string))
	return boolType.IsValidEnum()
}

func NewValidation() ValidationInterface {
	engine := validator.New()
	engine.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	err := engine.RegisterValidation("enum-approval", ValidateEnumApproval)
	if err != nil {
		return nil
	}
	err = engine.RegisterValidation("enum-flag", ValidateEnumFlag)
	fmt.Println("oke")
	if err != nil {
		return nil
	}
	return &ValidationError{engine: engine}
}

func (v *ValidationError) BindAndValidate(c *gin.Context, payload any) map[string]string {
	errs := make(map[string]string)
	err := c.Bind(payload)
	if err != nil {
		var errJSON *json.UnmarshalTypeError
		if errors.As(err, &errJSON) {
			field := errJSON.Field
			errVal := errJSON.Error()
			return map[string]string{
				field: errVal,
			}
		}
		return map[string]string{
			"error": err.Error(),
		}
	}
	err = v.engine.Struct(payload)
	if err != nil {
		var errVals validator.ValidationErrors
		if errors.As(err, &errVals) {
			for i, _ := range errVals {
				errs[errVals[i].Field()] = getErrorMsg(errVals[i])
			}
			return errs
		}
		return errs
	}

	return errs
}
func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "This field need email format " + fe.Param()
	case "url":
		return "This field need url format " + fe.Param()
	case "enum-flag":
		return "This field just need input True or False in type string " + fe.Param()
	case "enum-approval":
		return "This field just need input approve, pending or reject in type string " + fe.Param()
	}
	return "Unknown error"
}
