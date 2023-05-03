package goattribute

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Sample struct{}

func TestCopyInt(t *testing.T) {
	// Test cases for int
	t.Run("int", func(t *testing.T) {
		var target int

		err := CopyInt(&target, 42)
		assert.NoError(t, err)
		assert.Equal(t, 42, target)

		err = CopyInt(&target, "123")
		assert.NoError(t, err)
		assert.Equal(t, 123, target)

		err = CopyInt(&target, 123.456)
		assert.NoError(t, err)
		assert.Equal(t, 123, target)

		err = CopyInt(&target, "1.2")
		assert.NoError(t, err)
		assert.Equal(t, 1, target)

		err = CopyInt(&target, nil)
		assert.NotNil(t, err)

		err = CopyInt(&target, "")
		assert.NotNil(t, err)

		err = CopyInt(&target, &Sample{})
		assert.NotNil(t, err)
	})

	// Test cases for int8
	t.Run("int8", func(t *testing.T) {
		var target int8

		err := CopyInt(&target, int8(42))
		assert.NoError(t, err)
		assert.Equal(t, int8(42), target)

		err = CopyInt(&target, "123")
		assert.NoError(t, err)
		assert.Equal(t, int8(123), target)

		err = CopyInt(&target, 123.456)
		assert.NoError(t, err)
		assert.Equal(t, int8(123), target)

		err = CopyInt(&target, "1.2")
		assert.NoError(t, err)
		assert.Equal(t, int8(1), target)

		err = CopyInt(&target, nil)
		assert.NotNil(t, err)

		err = CopyInt(&target, "")
		assert.NotNil(t, err)

		err = CopyInt(&target, &Sample{})
		assert.NotNil(t, err)

		// overflow, but expected
		err = CopyInt(&target, 260)
		assert.Equal(t, int8(4), target)
	})

	// Test cases for uint8
	t.Run("uint8", func(t *testing.T) {
		var target uint8

		err := CopyInt(&target, uint8(42))
		assert.NoError(t, err)
		assert.Equal(t, uint8(42), target)

		err = CopyInt(&target, "123")
		assert.NoError(t, err)
		assert.Equal(t, uint8(123), target)

		err = CopyInt(&target, 123.456)
		assert.NoError(t, err)
		assert.Equal(t, uint8(123), target)

		err = CopyInt(&target, "1.2")
		assert.NoError(t, err)
		assert.Equal(t, uint8(1), target)

		err = CopyInt(&target, nil)
		assert.NotNil(t, err)

		err = CopyInt(&target, "")
		assert.NotNil(t, err)

		err = CopyInt(&target, &Sample{})
		assert.NotNil(t, err)
	})

	// Test cases for int32
	t.Run("int32", func(t *testing.T) {
		var target int32

		err := CopyInt(&target, int32(42))
		assert.NoError(t, err)
		assert.Equal(t, int32(42), target)

		err = CopyInt(&target, "123")
		assert.NoError(t, err)
		assert.Equal(t, int32(123), target)

		err = CopyInt(&target, 123.456)
		assert.NoError(t, err)
		assert.Equal(t, int32(123), target)

		err = CopyInt(&target, "1.2")
		assert.NoError(t, err)
		assert.Equal(t, int32(1), target)

		err = CopyInt(&target, nil)
		assert.NotNil(t, err)

		err = CopyInt(&target, "")
		assert.NotNil(t, err)

		err = CopyInt(&target, &Sample{})
		assert.NotNil(t, err)
	})

	// Test cases for int64
	t.Run("int64", func(t *testing.T) {
		var target int64

		err := CopyInt(&target, int64(42))
		assert.NoError(t, err)
		assert.Equal(t, int64(42), target)

		err = CopyInt(&target, "123")
		assert.NoError(t, err)
		assert.Equal(t, int64(123), target)

		err = CopyInt(&target, 123.456)
		assert.NoError(t, err)
		assert.Equal(t, int64(123), target)

		err = CopyInt(&target, "1.2")
		assert.NoError(t, err)
		assert.Equal(t, int64(1), target)

		err = CopyInt(&target, nil)
		assert.NotNil(t, err)

		err = CopyInt(&target, "")
		assert.NotNil(t, err)

		err = CopyInt(&target, &Sample{})
		assert.NotNil(t, err)
	})

	// Test cases for uint
	t.Run("uint", func(t *testing.T) {
		var target uint

		err := CopyInt(&target, uint(11))
		assert.NoError(t, err)
		assert.Equal(t, uint(11), target)

		err = CopyInt(&target, int8(22))
		assert.NoError(t, err)
		assert.Equal(t, uint(22), target)

		err = CopyInt(&target, uint8(33))
		assert.NoError(t, err)
		assert.Equal(t, uint(33), target)

		err = CopyInt(&target, uint8(44))
		assert.NoError(t, err)
		assert.Equal(t, uint(44), target)
	})

	// Test cases for int16
	t.Run("int16", func(t *testing.T) {
		var target int16

		err := CopyInt(&target, int16(42))
		assert.NoError(t, err)
		assert.Equal(t, int16(42), target)

		err = CopyInt(&target, int8(43))
		assert.NoError(t, err)
		assert.Equal(t, int16(43), target)

		err = CopyInt(&target, uint8(44))
		assert.NoError(t, err)
		assert.Equal(t, int16(44), target)
	})

	// Test cases for uint16
	t.Run("uint16", func(t *testing.T) {
		var target uint16

		err := CopyInt(&target, uint16(12))
		assert.NoError(t, err)
		assert.Equal(t, uint16(12), target)

		err = CopyInt(&target, int8(23))
		assert.NoError(t, err)
		assert.Equal(t, uint16(23), target)

		err = CopyInt(&target, uint8(34))
		assert.NoError(t, err)
		assert.Equal(t, uint16(34), target)
	})

	t.Run("mix", func(t *testing.T) {
		var i1 int
		assert.NoError(t, CopyInt(&i1, int64(0)))
		assert.Equal(t, 0, i1)

		assert.NoError(t, CopyInt(&i1, int8(-12)))
		assert.Equal(t, -12, i1)
	})
}
