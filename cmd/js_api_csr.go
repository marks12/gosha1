package cmd

const apiCSRContent = `

function request(method, url, getParams, data, headerParams) {

    function appendParams(u, params) {
        let uparams = "";
        switch (typeof params) {
            case "object":
                if (Object.keys(params).length < 1) {
                    return u;
                }
                u = u + (u.includes("?") ? "" : "?");
                for (const f of Object.keys(params)) {

                    if (uparams !== "") {
                        uparams += "&";
                    }

                    let list = "";

                    switch (typeof params[f]) {
                        case "object":

                            if (params && params[f]) {
                                for (let j = 0; j < params[f].length; j++) {
                                    if (list.length) {
                                        list += "&";
                                    }
                                    list += f + "[]=" + encodeURIComponent(params[f][j]);
                                }

                                uparams += list;
                            }
                            break;
                        default:
                            uparams += f + "=" + encodeURIComponent(params[f]);
                            break;
                    }
                }
                break;
        }
        return (u + uparams).replace(/[&]+/, '&');
    }

    function setHeader(req) {
        for (const f of Object.keys(headerParams)) {
            req.setRequestHeader(f, headerParams[f]);
        }
    }

    return new Promise(function (resolve, reject) {

        let xhr = new XMLHttpRequest();

        url = appendParams(url, getParams);

        xhr.open(method, url);
        xhr.setRequestHeader("Content-Type", "application/json");
        if (headerParams) {
            setHeader(xhr);
        }

        xhr.onload = function () {

            if (this.status >= 200 && this.status < 300) {

				if(xhr.response) {
				  try {
					data = JSON.parse(xhr.response)
					resolve(data);
				  } catch(e) {
					reject({
					  status: this.status,
					  statusText: "json parse error",
					});
				  }
				}

            } else {
                reject({
                    status: this.status,
                    statusText: xhr.statusText,
		            errorResponse: JSON.parse(xhr.response),
                });
            }
        };
        xhr.onerror = function () {
            reject({
                status: this.status,
                statusText: xhr.statusText,
		        errorResponse: JSON.parse(xhr.response),
            });
        };

        if (data) {
            xhr.send(JSON.stringify(data));
        } else {
            xhr.send();
        }
    });
}

function BackendApi() {

    this.serverUrl = "";

    this.token = "";

    this.defaultParameters = {};

    this.getRouteUrl = (url) => {
        return this.serverUrl + url;
    };

    this.appendToken = (params) => {
        if (!params) {
            params = {};
        }
        params.Token = this.token;
        return params;
    };

    let appendGlobalFilter = (filter) => {

        for (let param in filter) {
            if (filter.hasOwnProperty(param) && !filter[param] && this.defaultParameters.hasOwnProperty(param)) {
                filter[param] = this.defaultParameters[param];
            }
        }

        return filter;
    }

    return {
        find: (url, getParams, headerParams) => {
            getParams = appendGlobalFilter(getParams)
            return request("GET", this.getRouteUrl(url), getParams, null, this.appendToken(headerParams));
        },
        create: (url, data, getParams, headerParams) => {
            getParams = appendGlobalFilter(getParams)
            return request("POST", this.getRouteUrl(url), getParams, data, this.appendToken(headerParams));
        },
        update: (url, data, getParams, headerParams) => {
            getParams = appendGlobalFilter(getParams)
            return request("PUT", this.getRouteUrl(url), getParams, data, this.appendToken(headerParams));
        },
        remove: (url, getParams, data, headerParams) => {
            getParams = appendGlobalFilter(getParams)
            return request("DELETE", this.getRouteUrl(url), getParams, data, this.appendToken(headerParams));
        },
        getServerUrl: () => {
            return this.serverUrl;
        },
        setServerUrl: (url) => {
            this.serverUrl = url;
            return this;
        },
        getToken: () => {
            return this.token;
        },
        setToken: (token) => {
            this.token = token;
            return this;
        },
        clearGlobalFilters: () => {
            this.defaultParameters = {};
            return this;
        },
        clearGlobalFilter: (name) => {
            delete this.defaultParameters[name];
            return this;
        },
        setGlobalFilterValue: (parameter, value) => {
            this.defaultParameters[parameter] = value;
            return this;
        },
        getGlobalFilters: () => {
            return this.defaultParameters;
        },
        getGlobalFilter: (name) => {
            return this.defaultParameters[name];
        },
    };
}

let apiCSR = new BackendApi();

export default apiCSR;
`