(function (window) {
    'use strict';

    var ws = {};
    window.onload = function() {
        ws.sock = null;
        ws.uri = "ws://127.0.0.1:8080/ws";
        ws.sock = new WebSocket(ws.uri);
        ws.sock.onopen = function() {
            console.log("connected to " + ws.uri);
        }
        ws.sock.onclose = function(e) {
            console.log("connection closed (" + e.code + ")");
        }
        ws.sock.onmessage = function(e) {
            console.log("message received: " + e.data);
        }
        // ws.sock.send(msg)
    };


    window.app = window.app || {};
    window.app.ws = ws;
})(window);
