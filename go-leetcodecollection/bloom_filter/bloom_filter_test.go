package bloom_filter

import "testing"

func TestAddAndGet(t *testing.T) {
	filter := New(1024, 3, false)
	// chaining operation
	filter.Add([]byte("Hello")).
		AddString("World").
		AddUInt16(uint16(16)).
		AddUInt32(uint32(32)).
		AddUInt64(uint64(64)).
		AddUint16Batch([]uint16{1, 2, 3})

	t.Logf("Hello exist:%t", filter.Test([]byte("Hello")))
	t.Logf("World exist:%t", filter.TestString("World"))
	t.Logf("uint 16 exist:%t", filter.TestUInt16(uint16(16)))
	t.Logf("uint 32 exist:%t", filter.TestUInt32(uint32(32)))
	t.Logf("uint 64 exist:%t", filter.TestUInt64(uint64(64)))

	t.Logf("key exist:%t", filter.Test([]byte("key")))
	t.Logf("exist exist:%t", filter.TestString("exist"))
	t.Logf("uint 128 exist:%t", filter.TestUInt16(uint16(128)))
	t.Logf("uint 33 exist:%t", filter.TestUInt32(uint32(33)))
	t.Logf("uint 65 exist:%t", filter.TestUInt64(uint64(65)))

}
