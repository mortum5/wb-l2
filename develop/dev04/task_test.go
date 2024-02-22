package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAnagram(t *testing.T) {
	testCases := []struct {
		desc string
		data []string
		want map[string][]string
	}{
		{
			desc: "normal",
			data: []string{
				"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "столик", "листок",
			},
			want: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			desc: "none",
			data: []string{
				"123f", "ghskjas", "sfasdfk",
			},
			want: map[string][]string{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			anagram(tC.data)
			got := ans
			assert.Equal(t, tC.want, got)
		})
	}
}
