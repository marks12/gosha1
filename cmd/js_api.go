package cmd

const apiContent = `function request(method, url, getParams, data, headerParams) {

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
                    switch (typeof params[f]) {
                        case "object":
                            let list = "";
                            for (let j = 0; j < params[f].length; j++) {
                                list += f + "[]=" + encodeURIComponent(params[f][j]);
                            }
                            uparams += list;
                            break;
                        default:
                            uparams += f + "=" + encodeURIComponent(params[f]);
                            break;
                    }
                }

                break;
        }

        return u + uparams;
    }

    function setHeader(req) {
        for (const f of Object.keys(headerParams)) {
            req.setRequestHeader(f, headerParams[f]);
        }
    }

    return new Promise(function(resolve, reject) {

        let xhr = new XMLHttpRequest();

        url = appendParams(url, getParams);

        xhr.open(method, url);
        xhr.onload = function() {

            if (this.status >= 200 && this.status < 300) {

                resolve(JSON.parse(xhr.response));

            } else {
                reject({
                    status: this.status,
                    statusText: xhr.statusText,
                });
            }
        };
        xhr.onerror = function() {
            reject({
                status: this.status,
                statusText: xhr.statusText,
            });
        };

        if (data) {

            if (headerParams) {
                setHeader(xhr);
            }

            xhr.setRequestHeader("Content-Type", "application/json");
            xhr.send(JSON.stringify(data));
        } else {
            xhr.send();
        }
    });
}

const api = {
    create(url, data, getParams, headerParams) {
        return request("POST", url, getParams, data, headerParams);
    },
    update(url, data, getParams, headerParams) {
        return request("PUT", url, getParams, data, headerParams);
    },
    find(url, getParams, headerParams) {
        return request("GET", url, getParams, null, headerParams);
    },
    remove(url, getParams, headerParams) {
        return request("DELETE", url, getParams, null, headerParams);
    },
};

export default api;
`