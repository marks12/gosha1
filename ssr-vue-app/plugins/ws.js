import Vue from 'vue'

export default ({ app }) => {
  function Ws() {
    let wsSocket,
      types = {},
      sessions = {},
      reconnectTimeout = 5000,
      serverUrl = process.env.NODE_ENV === 'development'
        ? 'ws://127.0.0.1:5600/ws'
        : `ws://127.0.0.1:5600/ws`
        // : `wss://${window.location.hostname}/ws`
    ;

    function connect() {
      wsSocket = new WebSocket(serverUrl);

      wsSocket.onmessage = messageHandler;

      wsSocket.onclose = function (event) {

        if (!event.wasClean) {
          console.log('Код: ' + event.code + ' причина: ' + event.reason);

          setTimeout(function () {
            reconnect();
          }, reconnectTimeout);
        }

      };

      wsSocket.onopen = function () {
        send({
          type: 'SetToken',
          // TODO: Set really token
          data: 'super_long_token_not_really_token'
        });

        let url = new URL(window.location);
        let sessionId = url.searchParams.get('sessionId');
        if (sessionId) {
          send({
            type: 'AdminSession',
            data: sessionId
          });
        }

      }
    }

    function messageHandler(event) {
      let incomingMessage = event.data;

      if (!incomingMessage) return;

      incomingMessage = JSON.parse(incomingMessage);
      let response = new Response(incomingMessage);

      if (response.type) {
        for (let id in types) {
          if (types.hasOwnProperty(id) && types[id].getTypeName() === response.type) {
            types[id].getTypeFunc()(response);
          }
        }
      }

      if (response.sessionId && sessions.hasOwnProperty(response.sessionId)) {
        sessions[response.sessionId](event);
      }
    }

    function reconnect() {
      closeConnect();
      connect();
    }

    function closeConnect() {
      wsSocket.close();
    }

    function setUrl(url) {
      serverUrl = url;
    }

    function bindType(typeName, func) {
      let newType = new Type();
      newType.setFunction(func);
      newType.setType(typeName);

      let id = Date.now() + Math.random();

      types[id] = newType;

      return id;
    }

    function bindSession(sessionId, func) {
      sessions[sessionId] = func;
    }

    function removeType(typeName) {
      types = types.find(r => r.getTypeName() !== typeName);
    }

    function Type() {
      let typeAddr = '';
      let func = function () {
      };

      function setType(a) {
        typeAddr = a;
      }

      function setFunction(f) {
        func = f;
      }

      function getTypeName() {
        return typeAddr;
      }

      function getTypeFunc() {
        return func;
      }

      return {
        setType: setType,
        setFunction: setFunction,
        getTypeName: getTypeName,
        getTypeFunc: getTypeFunc
      }
    }

    function Request({type = '', data = data}) {

      return {
        type: type,
        data: data
      }
    }

    function Response({Type = '', Result = '', Data = '', Success = false}) {

      return {
        type: Type,
        result: Result,
        data: Data,
        success: Success
      }
    }

    function send(requestData) {
      let request = new Request(requestData);
      let data = JSON.stringify(request);

      wsSocket.send(data);
    }

    function setReconnectTimeout(time) {
      reconnectTimeout = time;
    }

    function isNeedAutoConnect() {
      return window.location.pathname.startsWith('/im3');
    }

    return {
      connect: connect,
      isNeedAutoConnect: isNeedAutoConnect,
      reconnect: reconnect,
      closeConnect: closeConnect,
      setUrl: setUrl,
      removeType: removeType,
      send: send,
      setReconnectTimeout: setReconnectTimeout,
      bindType: bindType,
      bindSession: bindSession,
    }
  }

  let socket = new Ws();

  socket.connect();

  app.socket = socket;
  Vue.prototype.$socket = socket;
}
