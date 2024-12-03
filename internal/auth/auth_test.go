package auth

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBearerToken(t *testing.T) {
	want := "hello-world"
	headers := http.Header{
		"Authorization": []string{fmt.Sprintf("Bearer %s", want)},
	}

	got, err := GetBearerToken(headers)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
