package serializer

import (
	"github.com/stretchr/testify/require"
	"grpc_test/sample"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../temp/laptop.bin"
	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)
}
