package notation

import (
	"testing"
)

func Test_ShiftEquality(t *testing.T) {
	type args struct {
		base int
		m    []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "decimal number system",
			args: args{
				base: 10,
				m:    []string{"1", "2", "3", "3", "2", "1"},
			},
			want: true,
		},
		{
			name: "decimal number system non strict equality",
			args: args{
				base: 10,
				m:    []string{"2", "2", "3", "3", "3", "1"},
			},
			want: true,
		},
		{
			name: "decimal number system wrong sequence",
			args: args{
				base: 10,
				m:    []string{"3", "2", "1", "1", "2", "3"},
			},
			want: false,
		},
		{
			name: "binary number system",
			args: args{
				base: 2,
				m:    []string{"0", "0", "1", "1", "0", "0"},
			},
			want: true,
		},
		{
			name: "binary number system non strict equality",
			args: args{
				base: 2,
				m:    []string{"0", "1", "1", "1", "1", "0"},
			},
			want: true,
		},
		{
			name: "binary number system wrong sequence",
			args: args{
				base: 2,
				m:    []string{"1", "1", "1", "0", "0", "0"},
			},
			want: false,
		},
		{
			name: "hexadecimal number system",
			args: args{
				base: 16,
				m:    []string{"3", "8", "9", "F", "C", "B", "A", "2"},
			},
			want: true,
		},
		{
			name: "hexadecimal number system wrong sequence",
			args: args{
				base: 16,
				m:    []string{"A", "8", "F", "F", "C", "B", "A", "2"},
			},
			want: false,
		},
		{
			name: "pentecostal system sequence",
			args: args{
				base: 50,
				m: []string{
					"0",
					"1",
					"2",
					"3",
					"4",
					"5",
					"6",
					"7",
					"8",
					"9",
					"A",
					"B",
					"C",
					"D",
					"E",
					"F",
					"G",
					"H",
					"I",
					"J",
					"K",
					"L",
					"L",
					"K",
					"J",
					"I",
					"H",
					"G",
					"F",
					"E",
					"D",
					"C",
					"B",
					"A",
					"9",
					"8",
					"7",
					"6",
					"5",
					"4",
					"3",
					"2",
					"1",
					"0",
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eq, err := ShiftEquality(tt.args.base, tt.args.m...)
			if err != nil {
				t.Error(err)
			}

			if eq != tt.want {
				t.Errorf("expected value and actual value do not match arg: %v", tt.args)
			}
		})
	}
}
