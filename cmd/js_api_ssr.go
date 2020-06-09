package cmd

const apiSSRContent = `

import axios from "axios";

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
            return axios.get(this.getRouteUrl(url), {params: getParams, headers: this.appendToken(headerParams)})
                .then((response) => {
                    return response.data;
                }).catch(err => {
                    console.error("Error in url request " + url, err)
                });
        },
        create: (url, data, getParams, headerParams) => {
            getParams = appendGlobalFilter(getParams)
            return axios.post(this.getRouteUrl(url), data, {
                params: getParams,
                headers: this.appendToken(headerParams)
            }).then((response) => {
                return response.data;
            }).catch(err => {
                console.error("Error in url request " + url, err)
            });
        },
        update: (url, data, getParams, headerParams) => {
            getParams = appendGlobalFilter(getParams)
            return axios.put(this.getRouteUrl(url), data, {
                params: getParams,
                headers: this.appendToken(headerParams)
            }).then((response) => {
                return response.data;
            }).catch(err => {
                console.error("Error in url request " + url, err)
            });
        },
        remove: (url, getParams, data, headerParams) => {
            getParams = appendGlobalFilter(getParams)
            return axios.delete(this.getRouteUrl(url), {
                data: data,
                params: getParams,
                headers: this.appendToken(headerParams)
            }).then((response) => {
                return response.data;
            }).catch(err => {
                console.error("Error in url request " + url, err)
            });
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

let apiSSR = new BackendApi();

export default apiSSR;
`
