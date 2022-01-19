class Cloud {

    server = "https://cloud.gosha.info"
    token = localStorage.getItem("GoshaToken") || ""
    isUserAuthorized = false;

    urls = {
        auth: "/api/v1/auth",
        me: "/api/v1/currentUser/me",
    };

    authStore = [];

    setServerUrl(url) {
        this.server = url;
    }

    getServerUrl() {
        return this.server;
    }

    setToken(token) {
        this.token = token;
        localStorage.setItem("GoshaToken", this.token);
    }

    isAuthorized() {
        return this.isUserAuthorized;
    }

    setAuthorized(isAuth) {
        this.isUserAuthorized = isAuth;

        if (this.authStore && this.authStore.length) {
            for (let i=0;i<this.authStore.length;i++) {
                if (typeof this.authStore[i] === 'function') {
                    this.authStore[i](isAuth);
                }
            }
        }

    }

    getToken() {
        return this.token;
    }

    getHeaders(additional) {

        let def = {
            'Content-Type': 'application/json',
            'X-Agent': 'gosha',
            'Token': this.getToken(),
        }

        if (additional) {
            def = Object.assign(def, additional);
        }

        return def;
    }

    post(route, data, headers) {
        let url = this.server + route;
        return fetch(url, {
            method: 'POST',
            body: JSON.stringify(data),
            headers: this.getHeaders(headers),
        }).then((res)=>{
            return res.json();
        })
    }

    get(route, headers) {
        let url = this.server + route;
        return fetch(url, {
            method: 'GET',
            headers: this.getHeaders(headers),
        }).then((res)=>{
            return res.json();
        })
    }

    auth(data, callbackSuccess, callbackError) {

        callbackSuccess = this.getDefaultCallback(callbackSuccess)
        callbackError = this.getDefaultCallback(callbackError)

        return this.post(this.urls.auth, data).then((result)=>{

            if (result && result.Model && result.Model.Token) {
                this.setToken(result.Model.Token);
                callbackSuccess(this.getToken());
            } else {
                callbackError("Неверный логин или пароль");
            }

        }).catch((error)=>{
            callbackError("Сервер недоступен");
        });
    }

    saveDocument(data, callbackSuccess, callbackError) {

        callbackSuccess = this.getDefaultCallback(callbackSuccess)
        callbackError = this.getDefaultCallback(callbackError)

        return this.post(this.urls.document, data).then((result)=>{

            if (result && result.Model && result.Model.Token) {
                this.setToken(result.Model.Token);
                callbackSuccess(this.getToken());
            } else {
                callbackError("Неверный логин или пароль");
            }

        }).catch((error)=>{
            callbackError("Сервер недоступен");
        });
    }

    getDefaultCallback(callback) {
        let def = (data, err)=>{console.error("error. Data: ", data, err)};
        if (!callback) {
            console.error("callback not set")
            return def
        }
        return callback
    }

    checkAuthorized(callbackSuccess, callbackError) {

        this.isUserAuthorized = false;

        callbackSuccess = this.getDefaultCallback(callbackSuccess)
        callbackError = this.getDefaultCallback(callbackError)

        return this.get(this.urls.me).then((result)=>{

            if (result && result.Model && result.Model.Id) {
                callbackSuccess(true);
                this.setAuthorized(true);
            } else {
                callbackError(false, "Ошибка авторизации");
                this.setAuthorized(false);
            }

        }).catch((error)=>{
            this.setAuthorized(false);
            callbackError(false, "Сервер недоступен");
        });
    }

    onAuthorize(method) {
        this.authStore.push(method);
    }
}

let myCloud = new Cloud();

export default myCloud
