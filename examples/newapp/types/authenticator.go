package types

import (
    "newapp/core"
    "newapp/dbmodels"
    "newapp/flags"
    "newapp/settings"
    "net/http"
    "strings"
)

type Authenticator struct {
    Token        string
    FunctionType string
    UrlPath      string
    validator
}

func (auth *Authenticator) IsAuthorized() bool {
    if *flags.Auth {
        return true
    }

    if len(auth.Token) < 1 {
        return false
    }

    dbAuth := dbmodels.Auth{}
    core.Db.Where(dbmodels.Auth{Token: auth.Token}).First(&dbAuth)
    if dbAuth.IsActive {
        usedResources := []dbmodels.Resource{}
        core.Db.Where(dbmodels.Resource{
            Code:   clearPath(auth.UrlPath),
            TypeId: 1,
        }).Find(&usedResources)
        ids := []int{}
        for _, r := range usedResources {
            ids = append(ids, r.ID)
        }
        roleResources := []dbmodels.RoleResource{}
        if dbAuth.UserId > 0 {
            core.Db.Model(dbmodels.RoleResource{}).Where("role_id in (select role_id from user_roles where user_id = ?) and resource_id in (?)", dbAuth.UserId, ids).Find(&roleResources)
        } else {
            core.Db.Model(dbmodels.RoleResource{}).Where("role_id = 4 and resource_id in (?)", ids).Find(&roleResources)
        }
        switch auth.FunctionType {
        case settings.FunctionTypeFind:
            for _, rr := range roleResources {
                if rr.Find {
                    return true
                }
            }
            return false

        case settings.FunctionTypeRead:
            for _, rr := range roleResources {
                if rr.Read {
                    return true
                }
            }
            return false

        case settings.FunctionTypeCreate:
            for _, rr := range roleResources {
                if rr.Create {
                    return true
                }
            }
            return false

        case settings.FunctionTypeUpdate:
            for _, rr := range roleResources {
                if rr.Update {
                    return true
                }
            }
            return false

        case settings.FunctionTypeDelete:
            for _, rr := range roleResources {
                if rr.Delete {
                    return true
                }
            }
            return false

        case settings.FunctionTypeFindOrCreate:
            for _, rr := range roleResources {
                if rr.FindOrCreate {
                    return true
                }
            }
            return false
        }
    }

    return false
}

func clearPath(s string) string {
    if strings.Count(s, "/") > 3 {
        return s[0:strings.LastIndex(s, "/")]
    }

    return s
}

func (auth *Authenticator) SetToken(r *http.Request) error {

    auth.Token = r.Header.Get("Token")

    return nil
}

func (authenticator *Authenticator) Validate(functionType string) {

    switch functionType {

    case settings.FunctionTypeFind:

        break;
    case settings.FunctionTypeCreate:

        break;
    case settings.FunctionTypeRead:

        break;
    case settings.FunctionTypeUpdate:

        break;
    case settings.FunctionTypeDelete:

        break;

    default:
        authenticator.validator.validationErrors = append(authenticator.validator.validationErrors, "Usupported function type: "+functionType)
        break;
    }
}
