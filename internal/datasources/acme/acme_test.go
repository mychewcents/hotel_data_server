package acme

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHotels(t *testing.T) {
	type scenario struct {
		desc           string
		hotelIDsFilter map[string]bool
		destIDsFilter  map[int]bool
		resultLen      int
		expectedErr    error
	}

	scenarios := []scenario{
		{
			desc:      "testing the flow",
			resultLen: 3,
		},
		{
			desc: "add hotel id filter -> no result to be returned",
			hotelIDsFilter: map[string]bool{
				"ABCD": true,
			},
			resultLen: 0,
		},
		{
			desc: "add destination id filter",
			destIDsFilter: map[int]bool{
				5432: true,
			},
			resultLen: 2,
		},
		{
			desc: "add destination + hotel id filter -> no result",
			hotelIDsFilter: map[string]bool{
				"ABCD": true,
			},
			destIDsFilter: map[int]bool{
				5432: true,
			},
			resultLen: 0,
		},
		{
			desc: "add destination + hotel id filter -> 1 result",
			hotelIDsFilter: map[string]bool{
				"iJhz": true,
			},
			destIDsFilter: map[int]bool{
				5432: true,
			},
			resultLen: 1,
		},
	}

	for _, tc := range scenarios {
		testCase := tc
		t.Run(testCase.desc, func(t *testing.T) {
			obj := GetHandler()
			hotels, err := obj.GetHotels(testCase.hotelIDsFilter, testCase.destIDsFilter)

			if tc.expectedErr != nil {
				assert.Equal(t, err, tc.expectedErr)
			} else {
				assert.Equal(t, tc.resultLen, len(hotels))
			}
		})
	}
}
