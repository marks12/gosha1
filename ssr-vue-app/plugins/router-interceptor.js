export default ({ app }) => {

  app.router.afterEach((to, from) => {
    app.socket.send({
      type: "Event",
      data: {
        Event: "UrlChange",
        Url: to.path,
        ElementName: ""
      }
    });
  });

}
