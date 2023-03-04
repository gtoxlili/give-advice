package ht

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	json "github.com/bytedance/sonic"
	"github.com/gtoxlili/advice-hub/common/pool"
	"github.com/gtoxlili/advice-hub/constants"
)

var (
	Client = &http.Client{}
)

func init() {
	if constants.ProxyAddr != "" {
		WithProxy(constants.ProxyAddr)
	}
}

func WithTimeout(timeout time.Duration) {
	Client.Timeout = timeout
}

func WithProxy(proxy string) {
	proxyUrl, _ := url.Parse(proxy)
	Client.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
}

type H map[string]string

type UnmarshalHandler[R any] func(*bytes.Buffer) R

type Response[R any] struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Data       R
}

func Get[R any](ctx context.Context, url string, headers H) (*Response[R], error) {
	return request[R](ctx, http.MethodGet, url, nil, headers, nil)
}

func Post[R any](ctx context.Context, url string, body any, headers H) (*Response[R], error) {
	return request[R](ctx, http.MethodPost, url, body, headers, nil)
}

func RequestWithUnmarshalHandler[R any](ctx context.Context, method string, url string, body any, headers H, unmarshalHandler UnmarshalHandler[R]) (*Response[R], error) {
	return request[R](ctx, method, url, body, headers, unmarshalHandler)
}

func request[R any](ctx context.Context, method string, url string, body any, headers H, unmarshalHandler UnmarshalHandler[R]) (*Response[R], error) {
	bodyReader := &bytes.Buffer{}

	if body != nil {
		bodyReader = pool.GetBuffer()
		defer pool.PutBuffer(bodyReader)
		// 直接传入 io.Reader
		if r, ok := body.(io.Reader); ok {
			_, err := io.Copy(bodyReader, r)
			if err != nil {
				return nil, err
			}
			// 传入 url.Values
		} else if v, ok := body.(interface{ Encode() string }); ok {
			bodyReader.WriteString(v.Encode())
			if _, ok = headers["Content-Type"]; !ok {
				headers["Content-Type"] = "application/x-www-form-urlencoded"
			}
			// 传入 map[string]interface{} 或者 struct
		} else {
			bys, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			bodyReader.Write(bys)
			if _, ok = headers["Content-Type"]; !ok {
				headers["Content-Type"] = "application/json"
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := &Response[R]{
		Headers:    resp.Header,
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
	}
	if resp.StatusCode != http.StatusOK {
		return res, err
	}

	resReader := pool.GetBuffer()
	defer pool.PutBuffer(resReader)
	buf := pool.Get(calcSize(resReader))
	defer pool.Put(buf)
	_, err = io.CopyBuffer(resReader, resp.Body, buf)
	if err != nil {
		return nil, err
	}
	res.Data, err = unmarshal[R](resReader, unmarshalHandler)
	return res, err
}

func unmarshal[R any](src *bytes.Buffer, unmarshalHandler UnmarshalHandler[R]) (R, error) {
	if unmarshalHandler != nil {
		return unmarshalHandler(src), nil
	}
	res := new(R)
	if err := json.Unmarshal(src.Bytes(), res); err != nil {
		return *new(R), err
	}
	return *res, nil
}

func calcSize(src io.Reader) int {
	size := 32 * 1024
	if l, ok := src.(*io.LimitedReader); ok && int64(size) > l.N {
		if l.N < 1 {
			size = 1
		} else {
			size = int(l.N)
		}
	}
	return size
}
