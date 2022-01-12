class Cloud {

    server = "https://gosha.info"
    token = localStorage.getItem("GoshaToken") || ""

    urls = {
        auth: "/api/v1/auth",
        me: "/api/v1/me",
    };

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

    getToken() {
        return this.token;
    }

    getHeaders(additional) {

        let def = {
            'Content-Type': 'application/json',
            'X-Agent': 'gosha',
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

    auth(data, callbackSuccess, callbackError) {

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
}

let myCloud = new Cloud();

export default myCloud
