package main

import "testing"

func TestMain(t *testing.T) {
	testcases := []struct {
		name  string
		input string
		want  string
		err   bool
	}{
		{
			name:  "Simple positive test case",
			input: "a4bc2d5e",
			want:  "aaaabccddddde",
			err:   false,
		},
		{
			name:  "Test without unpacking",
			input: "abcd",
			want:  "abcd",
			err:   false,
		},
		{
			name:  "Empty string",
			input: "",
			want:  "",
			err:   false,
		},
		{
			name:  "Negative test case",
			input: "45",
			want:  "",
			err:   true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := unpack(tc.input)
			if err != nil && !tc.err {
				t.Errorf("Error: %v", err)
				return
			}

			if got != tc.want {
				t.Errorf("got: %s, want: %s", got, tc.want)
			}
		})
	}
}
