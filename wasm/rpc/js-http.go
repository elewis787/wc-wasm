package rpc

import (
	"io"
	"net/http"
	"syscall/js"
)

type jshttp struct {
	client *http.Client
}

func NewHTTP() *jshttp {
	return &jshttp{
		client: http.DefaultClient,
	}
}

func (j *jshttp) Get() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		requestUrl := args[0].String()
		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			resolve := args[0]
			reject := args[1]
			go func() {
				res, err := j.client.Get(requestUrl)
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}
				defer res.Body.Close()

				data, err := io.ReadAll(res.Body)
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					errorObject := errorConstructor.New(err.Error())
					reject.Invoke(errorObject)
					return
				}

				// Create a new JavaScript object
				responseObj := js.Global().Get("Object").New()

				// Set the fields of the JavaScript object
				responseObj.Set("body", string(data))
				responseObj.Set("status", res.StatusCode)
				responseObj.Set("statusText", res.Status)

				resolve.Invoke(responseObj)
			}()
			return nil
		})

		promise := js.Global().Get("Promise")
		return promise.New(handler)
	})
}
