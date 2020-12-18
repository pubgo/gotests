package netdx

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	_ "unsafe"
)

//go:linkname RoundTrip http.(*Transport).RoundTrip
func RoundTrip(transport *http.Transport, req *http.Request) (reps *http.Response, err error) {
	// Feel free to diagnosis
	var (
		url = fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.URL.Host, req.URL.Path)
		id  = generateID()
		now = time.Now()
	)

	L7Printf("[HTTP-DX] id: %d, send a request to url: %s, method: %s \n", id, url, req.Method)

	if VerboseEnabled() {
		L7Printf("[HTTP-DX] id: %d, header:%+v, form: %+v, postForm: %+v, multipartForm %+v \n",
			id, req.Header, req.Form, req.PostForm, req.MultipartForm)
	}

	timer := time.AfterFunc(time.Second, func() {
		L7Errorf("[HTTP-DX] id: %d, url: %s, err: connection is stuck \n", id, url)
	})
	defer timer.Stop()

	defer func() {

		defer func() {
			if !VerboseEnabled() || err != nil {
				return
			}

			pFn := L7Printf
			if reps.StatusCode != http.StatusOK {
				pFn = L7Errorf
			}

			if reps.Body != nil {
				body, _ := ioutil.ReadAll(reps.Body)
				reps.Body = ioutil.NopCloser(bytes.NewBuffer(body))

				pFn("[HTTP-DX] id: %d, url: %s, body: %s \n", id, url, string(body))
			}
		}()

		if err != nil {
			L7Errorf("[HTTP-DX] id: %d, url: %s, err: %v \n", id, url, err)
			return
		}

		if reps.StatusCode != http.StatusOK {
			L7Errorf("[HTTP-DX] id: %d, url: %s, statusCode: %v \n", id, url, reps.StatusCode)
			return
		}

		L7Printf("[HTTP-DX] id: %d, url: %s, cost: %v, OK \n", id, url, time.Since(now))
	}()

	return roundTrip(transport, req)
}

func httpDx() {
	checkGCFlags()
}

//go:linkname roundTrip net/http.(*Transport).roundTrip
func roundTrip(*http.Transport, *http.Request) (*http.Response, error)
