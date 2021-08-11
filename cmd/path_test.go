package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathIsValid(t *testing.T) {

	type test struct {
		path string
		want bool
	}

	tests := []test{
		{path: "/this/path/doesnt/exist/test/test/", want: false},
		{path: "/tmp/cd871346dfjyasdfjgazg2t22187", want: false},
		{path: "/tmp", want: true},
		{path: "~/", want: true},
	}

	for _, tc := range tests {
		got, gotErr := pathIsValid(tc.path)
		assert.NoError(t, gotErr)
		assert.Equal(t, got, tc.want)
	}
}
