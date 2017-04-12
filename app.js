/*global app */
(function () {
    'use strict';

    var msg = "this is app data";
    app.send = function() {
        app.ws.sock.send(msg);
        console.log("sent data")
    }

})();
