package loaderservice

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teadove/teasutils/utils/test_utils"
)

func TestGetImageDate(t *testing.T) {
	t.Parallel()

	r := NewService(nil, nil, "", "", nil)

	date, err := r.getImageDate("test.jpg")
	require.NoError(t, err)

	test_utils.Pprint(date)
}
