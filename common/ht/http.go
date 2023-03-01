package ht

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	json "github.com/bytedance/sonic"
	"github.com/gtoxlili/give-advice/common/pool"
)

var (
	Client = &http.Client{}
)

func WithTimeout(timeout time.Duration) {
	Client.Timeout = timeout
}

type Header map[string]string

type Response[R any] struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Data       *R
}

func Request[R any](ctx context.Context, method string, url string, body any, headers Header) (*Response[R], error) {
	fmt.Println("url:", url)
	bodyReader := pool.GetBuffer()
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
		// 传入 map[string]interface{} 或者 struct
	} else if body != nil {
		bys, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader.Write(bys)
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
		Data:       new(R),
	}
	if resp.StatusCode != http.StatusOK {
		return res, nil
	}

	resReader := pool.GetBuffer()
	defer pool.PutBuffer(resReader)
	buf := pool.Get(calcSize(resReader))
	defer pool.Put(buf)
	_, err = io.CopyBuffer(resReader, resp.Body, buf)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(resReader.Bytes(), res.Data); err != nil {
		return nil, err
	}
	return res, nil
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
