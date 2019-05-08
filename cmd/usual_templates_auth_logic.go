package cmd

const usualAuthLogic = `
package logic

import (
    "{ms-name}/types"
    "{ms-name}/dbmodels"
    "{ms-name}/core"
    "{ms-name}/common"
    "errors"
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "log"
)

func AuthCreate(filter types.AuthFilter)  (data types.Auth, err error) {

    query := core.Db

    typeModel := filter.GetAuthModel()
    dbAuth := AssignAuthDbFromType(typeModel)
    dbAuth.Validate()

    if dbAuth.IsValid() {

        dbUser := dbmodels.User{}

        query = query.Where(dbmodels.User{Email: dbAuth.Email}).Find(&dbUser)

        if dbUser.ID < 1 {
            return types.Auth{}, errors.New("cant create Auth")
        }

        hashErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(dbAuth.Password))

        fmt.Println("hashErr", hashErr)

        if hashErr == nil {

            token := common.RandomString(32)

            dbUser.Token = token
            q := core.Db.Save(&dbUser)

            if q.Error != nil {

                log.Println("GetNewToken > Ошибка сохранения данных:", q.Error)

            } else {

                typeModel.Token = token

                return typeModel, nil
            }
        }

    } else {

        fmt.Println("AuthCreate > Invalid data:", dbAuth)
        return types.Auth{}, errors.New(dbAuth.GetValidationErrors())
    }

    fmt.Println("AuthCreate > Create Auth error:", query.Error)
    return types.Auth{}, errors.New("cant create Auth")
}

func AuthDelete(filter types.AuthFilter)  (isOk bool, err error) {

    dbUser := dbmodels.User{}

    query := core.Db

    q := query.Where(dbmodels.User{Token: filter.Token}).Find(&dbUser)

    dbUser.Token = ""
    q = core.Db.Save(&dbUser)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}
`
var usualTemplateAuthLogic = template{
	Path:    "./logic/auth.go",
	Content: assignMsName(usualAuthLogic),
}
