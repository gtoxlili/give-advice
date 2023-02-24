package pool

func Get(size int) []byte {
	return defaultAllocator.Get(size)
}

func Put(buf []byte) {
	defaultAllocator.Put(buf)
}
