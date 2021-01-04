package cmd

const usualErrorCode = `
package errors

const DefaultErrorLanguageId = 1

// Error codes
const (
	ErrorCodeUndefined                          ErrorCode = 0
	ErrorCodeNotFound                           ErrorCode = 100
	ErrorCodeInvalidAuthorize                   ErrorCode = 200
	ErrorCodeUnsupportedFunctionType            ErrorCode = 300
	ErrorCodeParseId                            ErrorCode = 400
	ErrorCodeInvalidPerPage                     ErrorCode = 500
	ErrorCodeRabbitQueueNameNotSet              ErrorCode = 600
	ErrorCodeNotValid                           ErrorCode = 700
	ErrorCodeSqlError                           ErrorCode = 750
	ErrorCodeInvalidCurrentPage                 ErrorCode = 800
	ErrorCodeNotEmpty                           ErrorCode = 900
	ErrorCodeAgeToSmall                         ErrorCode = 1000
	ErrorCodeFieldLengthTooShort                ErrorCode = 1100
)

// Error messages EN
const (
	ErrorEnMessageUndefined                          = "Undefined error"
	ErrorEnMessageNotFound                           = "Not found"
	ErrorEnMessageInvalidAuthorize                   = "Invalid authorize"
	ErrorEnMessageUnsupportedFunctionType            = "Unsupported function type"
	ErrorEnMessageParseId                            = "Error parse Id"
	ErrorEnMessageInvalidPerPage                     = "Invalid PerPage"
	ErrorEnMessageRabbitQueueNameNotSet              = "RabbitQueueName not set for application"
	ErrorEnMessageNotValid                           = "Field not valid"
	ErrorEnMessageSqlError                           = "Sql error"
	ErrorEnMessageInvalidCurrentPage                 = "Invalid CurrentPage"
	ErrorEnMessageNotEmpty                           = "Field in not empty"
	ErrorEnMessageAgeToSmall                         = "You must be over 14 years old"
	ErrorEnMessageFieldLengthTooShort                = "Field value has too short length"
)
`

var usualTemplateErrorCodes = template{
    Path:    "./errors/codes.go",
    Content: usualErrorCode,
}

const usualError = `// Пакет кастомных ошибок
// Полностью совместим со стандартным интерфейсом error
// Ошибки из пакета должны создаваться только
// с использованием конструкторов: NewErrorWithCode и NewValidationError
package errors

import "strings"

// Тип ошибки содержащий в себе код ошибки
// Включается в ValidatorError
type ErrorWithCode interface {
    Error() string
    ErrorCode() int
    GetField() string
    SetErrorCode(code ErrorCode)
}

type ErrorCode int

func (e ErrorCode) GetInt() int {
    return int(e)
}

type Error struct {
    Field string
    Text string
    Code int
}

func (e Error) Error() string {
    return e.Text
}

func (e Error) ErrorCode() int {
    return e.Code
}

func (e *Error) SetErrorCode(code ErrorCode) {
    e.Code = int(code)
}


func (e Error) GetField() string {
    return e.Field
}

func NewErrorWithCode(err string, code ErrorCode, field string) ErrorWithCode{
    return &Error{Text: err, Code: int(code), Field: field}
}

func New(err string) ErrorWithCode{
    return &Error{Text: err, Code: int(ErrorCodeNotValid), Field: ""}
}

// Тип ошибки для валидатора.
// Это единственное место где мы можем возвращать множественную ошибку
type ValidatorErrorInterface interface {
    Error() string
    Errors() []ErrorWithCode
    IsEmpty() bool
    AddError(ErrorWithCode)
}

type ValidatorError struct {
    errors []ErrorWithCode
}

func (v ValidatorError) Error() string {
    strErrors := []string{}
    for _, e := range v.errors {
        strErrors = append(strErrors, e.Error())
    }
    return strings.Join(strErrors, ",")
}

func (v ValidatorError) Errors() []ErrorWithCode {
    return v.errors
}

func (v *ValidatorError) AddError(e ErrorWithCode) {
    v.errors = append(v.errors, e)
}

func (v *ValidatorError) IsEmpty() bool{
    return len(v.errors) == 0
}

func JoinValidatorError(args... ValidatorErrorInterface) ValidatorError {
    err := ValidatorError{}
    for _, validatorError := range args {
        for _, e := range validatorError.Errors() {
            err.AddError(e)
        }
    }

    return err
}
`

var usualTemplateError = template{
    Path:    "./errors/error.go",
    Content: usualError,
}
