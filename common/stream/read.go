package stream

import (
	"context"
	"io"

	"github.com/gtoxlili/give-advice/common/pool"
)

func ReadFlow(r io.Reader, end []byte, callback func([]byte) bool) error {
	defer io.Copy(io.Discard, r)

	// 1. 从复用池中获取一个 []byte 用于存储读取到的总数据
	content := pool.Get(256)
	defer pool.Put(content)
	// 2. 从复用池中获取一个 []byte 用于存储每次读取到的数据
	buf := pool.Get(1)
	defer pool.Put(buf)
	for {
		content = content[:0]
		for {
			// 3. 读取数据
			_, err := r.Read(buf)
			if err != nil {
				if err == context.Canceled {
					return nil
				}
				return err
			}
			// 4. 将读取到的数据追加到总数据中
			content = append(content, buf...)
			// 5. 判断是否读取到结束标识
			if len(content) >= len(end) {
				if bytesEqual(content[len(content)-len(end):], end) {
					if callback(content) {
						return nil
					}
					break
				}
			}
		}
	}
}

func bytesEqual(a, b []byte) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
