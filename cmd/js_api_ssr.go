package cmd

const apiSSRContent = `

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
      }).catch( err => {
          console.error(` + "`" + `Error in url request ${url}` + "`" + `, err)
      });
    },
    find: (url, getParams, headerParams) => {
      return axios.get(this.getRouteUrl(url), {params: getParams, headers: headerParams })
        .then((response) => {
          return response.data;
      }).catch( err => {
          console.error(` + "`" + `Error in url request ${url}` + "`" + `, err)
      });
    },
    getServerUrl: () => {
      return this.serverUrl;
    },
    remove: (url, getParams, data, headerParams) => {
      return axios.delete(this.getRouteUrl(url), {data: data, params: getParams, headers: headerParams }).then((response) => {
        return response.data;
      }).catch( err => {
          console.error(` + "`" + `Error in url request ${url}` + "`" + `, err)
      });
    },
    setServerUrl: (url) => {
      this.serverUrl = url;
      return this;
    },
    update: (url, data, getParams, headerParams) => {
      return axios.put(this.getRouteUrl(url), data, {params: getParams, headers: headerParams }).then((response) => {
        return response.data;
      }).catch( err => {
          console.error(` + "`" + `Error in url request ${url}` + "`" + `, err)
      });
    },
  };
}

let apiSSR = new BackendApi();

export default apiSSR;
`