package cmd

const apiContent = `
import apiFront from "./apiCSR";
import apiSSR from "./apiSSR";

let api;
if(process && !process.client) {
  // for Server Side Rendering
  api = apiSSR;
} else {
  // for Client Side Rendering
  api = apiFront;
}

export default api;
`