package cmd

const apiContent = `
import apiFront from "./apiCSR";
import apiSSR from "./apiSSR";

// for Client Side Rendering
let api = apiFront;
if(process && !process.client) {
  // for Server Side Rendering
  let api = apiSSR;
}

export default api;
`