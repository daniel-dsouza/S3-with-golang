package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

// Controller handle all base methods
type Controller struct {
}

// SendJSON marshals v to a json struct and sends appropriate
func (c *Controller) SendJSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")

	b, err := json.Marshal(v)

	if err != nil {
		log.Print(fmt.Sprintf("Error while encoding JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Internal server error"}`)
	} else {
		w.WriteHeader(code)
		io.WriteString(w, string(b))
	}
}

// SendJPEG returns a jpeg image
func (c *Controller) SendJPEG(w http.ResponseWriter, r *http.Request, buffer *bytes.Buffer, code int) {
	w.Header().Add("Content-Type", "image/jpeg")
	w.Header().Add("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	w.WriteHeader(code)
	_, err := w.Write(buffer.Bytes())

	if err != nil { // this is technicall wrong, the header should be written before content
		log.Println("unable to write image")
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		//w.WriteHeader(code)
	}
}

// SendWebSocket upgrades to a websocket connection, and returns the connection
func (c *Controller) SendWebSocket(w http.ResponseWriter, r *http.Request) *websocket.Conn {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil
	}

	return conn
}
