package cmd

const usualAuthLogic = `package logic

import (
    "{ms-name}/types"
    "{ms-name}/dbmodels"
    "{ms-name}/core"
    "{ms-name}/settings"
    "{ms-name}/common"
    "errors"
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "log"
)

func AuthFind(filter types.AuthFilter)  (result []types.Auth, totalRecords int, err error) {
	return
}

func AuthCreate(filter types.AuthFilter) (data types.Auth, err error) {

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
        hashErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(typeModel.Password+settings.PASSWORD_SALT))

        if hashErr == nil {

            token := common.RandomString(32)
            dbAuth.Token = token

            core.Db.Model(dbmodels.Auth{}).Where(dbmodels.Auth{Token: token}).First(&dbAuth)
            dbAuth.IsActive = true
            dbAuth.UserId = dbUser.ID
            dbAuth.Email = dbUser.Email
            passLen := "empty"
            if len(dbUser.Password) > 0 {
                passLen = "used"
            }
            dbAuth.Password = passLen
            q := core.Db.Save(&dbAuth)

            if q.Error != nil {

                log.Println("GetNewToken > Ошибка сохранения данных:", q.Error)

            } else {

                typeModel.Token = dbAuth.Token
                typeModel.Password = "******"

                return typeModel, nil
            }
        }

        fmt.Println("AuthCreate > Create Auth error:", hashErr)
        return types.Auth{}, errors.New("cant create Auth")

    } else {

        fmt.Println("AuthCreate > Invalid data:", dbAuth)
        return types.Auth{}, errors.New(dbAuth.GetValidationErrors())
    }
}

func AuthRead(filter types.AuthFilter)  (data types.Auth, err error) {
	return
}

func AuthUpdate(filter types.AuthFilter)  (data types.Auth, err error) {
	return 
}

func AuthDelete(filter types.AuthFilter) (isOk bool, err error) {

    dbAuth := dbmodels.Auth{}

    query := core.Db

    q := query.Model(dbmodels.Auth{}).Where(dbmodels.Auth{Token: filter.Token}).Find(&dbAuth)

    dbAuth.IsActive = false
    q = core.Db.Model(dbmodels.Auth{}).Save(&dbAuth)

    if q.Error != nil {
        err = q.Error
        return
    }

    isOk = true
    return
}

func AuthFindOrCreate(filter types.AuthFilter)  (data types.Auth, err error) {
	return 
}

`
var usualTemplateAuthLogic = template{
	Path:    "./logic/auth.go",
	Content: assignMsName(usualAuthLogic),
}
