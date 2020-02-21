

import axios from "axios";

function BackendApi() {

  this.serverUrl = "";

  this.getRouteUrl = (url) => {
    return this.serverUrl + url;
  };

  return {
    create: (url, data, getParams, headerParams) => {
      return axios.post(this.getRouteUrl(url),data, {params: getParams, headers: headerParams }).then((response) => {
        return response.data;
      });
    },
    find: (url, getParams, headerParams) => {
      return axios.get(this.getRouteUrl(url), {params: getParams, headers: headerParams })
        .then((response) => {
          return response.data;
        });
    },
    getServerUrl: () => {
      return this.serverUrl;
    },
    remove: (url, getParams, data, headerParams) => {
      return axios.delete(this.getRouteUrl(url), {data: data, params: getParams, headers: headerParams }).then((response) => {
        return response.data;
      });
    },
    setServerUrl: (url) => {
      this.serverUrl = url;
      return this;
    },
    update: (url, data, getParams, headerParams) => {
      return axios.put(this.getRouteUrl(url), data, {params: getParams, headers: headerParams }).then((response) => {
        return response.data;
      });
    },
  };
}

let apiSSR = new BackendApi();

export default apiSSR;
