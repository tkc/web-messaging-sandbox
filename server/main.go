package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var html = `<!DOCTYPE html>
<html>
<head>
    <title>Authorization Response</title>
</head>
<body>
<script type="text/javascript">
    (function(window, document) {
        // send web message
        const targetOrigin = "http://localhost:4000";
        const webMessageRequest = {};
        const authorizationResponse = {
            type: "authorization_response",
            response: {
                "token": "sample_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            }
        };

        const mainWin = (window.opener) ? window.opener : window.parent;
        if (webMessageRequest["web_message_uri"] && webMessageRequest["web_message_target"]) {
            window.addEventListener("message", function(evt) {
                if (evt.origin != targetOrigin) return;
                switch (evt.data.type) {
                    case "relay_response":
                        const messageTargetWindow = evt.source.frames[webMessageRequest["web_message_target"]];
                        if (messageTargetWindow) {
                            messageTargetWindow.postMessage(authorizationResponse, webMessageRequest["web_message_uri"]);
                            window.close();
                        }
                        break;
                }
            });
            mainWin.postMessage({
                type: "relay_request"
            }, targetOrigin);
        } else {
            mainWin.postMessage(authorizationResponse, targetOrigin);
        }
    })(this, this.document);
</script>
</body>
</html>`

func main() {
	e := echo.New()

	// Routes
	e.GET("/authorize", func(c echo.Context) error {
		// TODO : check user_id in session
		return c.HTML(http.StatusOK, html)
	})

	// serve
	e.Logger.Fatal(e.Start(":8080"))
}
