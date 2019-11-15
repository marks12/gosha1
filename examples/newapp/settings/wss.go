package settings

const WssPortDev = "5600"

const WssPortProd = "5050"

func GetWssPort() string {

    if IsDev() {
        return WssPortDev
    }

    return WssPortProd
}
