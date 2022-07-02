package duplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicates_GetAll(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	structTests := []struct {
		name string
		args []testStruct
		want []testStruct
	}{
		{
			name: "Struct tests: No duplicates",
			args: []testStruct{
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Tom",
					Age:  9,
				},
				{
					Name: "Andy",
					Age:  12,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
				{
					Name: "Sarah",
					Age:  14,
				},
			},
			want: nil,
		},
		{
			name: "Struct tests: Has 2 duplicates",
			args: []testStruct{
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
				{
					Name: "Sarah",
					Age:  14,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
			},
			want: []testStruct{
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
			},
		},
	}
	for _, tt := range structTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetAll(tt.args))
		})
	}

	intTests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Int tests: No duplications",
			args: []int{1, 2, 3, 4, 5},
			want: nil,
		},
		{
			name: "Int tests: Has 2 duplicates",
			args: []int{1, 1, 3, 4, 4},
			want: []int{1, 4},
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetAll(tt.args))
		})
	}

	stringTests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "Strings test: No duplicates",
			args: []string{"one", "two", "three"},
			want: nil,
		},
		{
			name: "Strings test: Has 2 duplicates",
			args: []string{"one", "two", "two", "two", "three"},
			want: []string{"two", "two"},
		},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetAll(tt.args))
		})
	}

	float64Tests := []struct {
		name string
		args []float64
		want []float64
	}{
		{
			name: "Float64 tests: No duplicates",
			args: []float64{12.3, 10.5, 145.6},
			want: nil,
		},
		{
			name: "Float64 tests: Has 2 duplicates",
			args: []float64{9.5, 10.5, 15.9, 10.1, 10.5, 15.9},
			want: []float64{10.5, 15.9},
		},
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetAll(tt.args))
		})
	}
}

func TestDuplicates_Detect(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	structTests := []struct {
		name string
		args []testStruct
		want bool
	}{
		{
			name: "Struct tests: No duplicates",
			args: []testStruct{
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Tom",
					Age:  9,
				},
				{
					Name: "Andy",
					Age:  12,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
				{
					Name: "Sarah",
					Age:  14,
				},
			},
			want: false,
		},
		{
			name: "Struct tests: Has 2 duplicates",
			args: []testStruct{
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
				{
					Name: "Sarah",
					Age:  14,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
			},
			want: true,
		},
	}
	for _, tt := range structTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Detect(tt.args))
		})
	}

	intTests := []struct {
		name string
		args []int
		want bool
	}{
		{
			name: "Int tests: No duplications",
			args: []int{1, 2, 3, 4, 5},
			want: false,
		},
		{
			name: "Int tests: Has 2 duplicates",
			args: []int{1, 1, 3, 4, 4},
			want: true,
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Detect(tt.args))
		})
	}

	stringTests := []struct {
		name string
		args []string
		want bool
	}{
		{
			name: "Strings test: No duplicates",
			args: []string{"one", "two", "three"},
			want: false,
		},
		{
			name: "Strings test: Has 2 duplicates",
			args: []string{"one", "two", "two", "two", "three"},
			want: true,
		},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Detect(tt.args))
		})
	}

	float64Tests := []struct {
		name string
		args []float64
		want bool
	}{
		{
			name: "Float64 tests: No duplicates",
			args: []float64{12.3, 10.5, 145.6},
			want: false,
		},
		{
			name: "Float64 tests: Has 2 duplicates",
			args: []float64{9.5, 10.5, 15.9, 10.1, 10.5, 15.9},
			want: true,
		},
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Detect(tt.args))
		})
	}
}

func TestDuplicates_Count(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	structTests := []struct {
		name string
		args []testStruct
		want int
	}{
		{
			name: "Struct tests: No duplicates",
			args: []testStruct{
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Tom",
					Age:  9,
				},
				{
					Name: "Andy",
					Age:  12,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
				{
					Name: "Sarah",
					Age:  14,
				},
			},
			want: 0,
		},
		{
			name: "Struct tests: Has 2 duplicates",
			args: []testStruct{
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
				{
					Name: "Sarah",
					Age:  14,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
			},
			want: 2,
		},
	}
	for _, tt := range structTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Count(tt.args))
		})
	}

	intTests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "Int tests: No duplications",
			args: []int{1, 2, 3, 4, 5},
			want: 0,
		},
		{
			name: "Int tests: Has 2 duplicates",
			args: []int{1, 1, 3, 4, 4},
			want: 2,
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Count(tt.args))
		})
	}

	stringTests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "Strings test: No duplicates",
			args: []string{"one", "two", "three"},
			want: 0,
		},
		{
			name: "Strings test: Has 2 duplicates",
			args: []string{"one", "two", "two", "two", "three"},
			want: 2,
		},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Count(tt.args))
		})
	}

	float64Tests := []struct {
		name string
		args []float64
		want int
	}{
		{
			name: "Float64 tests: No duplicates",
			args: []float64{12.3, 10.5, 145.6},
			want: 0,
		},
		{
			name: "Float64 tests: Has 2 duplicates",
			args: []float64{9.5, 10.5, 15.9, 10.1, 10.5, 15.9},
			want: 2,
		},
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Count(tt.args))
		})
	}
}

func TestDuplicates_Remove(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	structTests := []struct {
		name string
		args []testStruct
		want []testStruct
	}{
		{
			name: "Struct tests: No duplicates",
			args: []testStruct{
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Tom",
					Age:  9,
				},
				{
					Name: "Andy",
					Age:  12,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
				{
					Name: "Sarah",
					Age:  14,
				},
			},
			want: []testStruct{
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Tom",
					Age:  9,
				},
				{
					Name: "Andy",
					Age:  12,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
				{
					Name: "Sarah",
					Age:  14,
				},
			},
		},
		{
			name: "Struct tests: Has 2 duplicates",
			args: []testStruct{
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
				{
					Name: "Sarah",
					Age:  14,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
			},
			want: []testStruct{
				{
					Name: "Josh",
					Age:  12,
				},
				{
					Name: "Ashley",
					Age:  11,
				},
				{
					Name: "Sarah",
					Age:  14,
				},
			},
		},
	}
	for _, tt := range structTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Remove(tt.args))
		})
	}

	intTests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "Int tests: No duplications",
			args: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Int tests: Has 2 duplicates",
			args: []int{1, 1, 3, 4, 4},
			want: []int{1, 3, 4},
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Remove(tt.args))
		})
	}

	stringTests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "Strings test: No duplicates",
			args: []string{"one", "two", "three"},
			want: []string{"one", "two", "three"},
		},
		{
			name: "Strings test: Has 2 duplicates",
			args: []string{"one", "two", "two", "two", "three"},
			want: []string{"one", "two", "three"},
		},
	}
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Remove(tt.args))
		})
	}

	float64Tests := []struct {
		name string
		args []float64
		want []float64
	}{
		{
			name: "Float64 tests: No duplicates",
			args: []float64{12.3, 10.5, 145.6},
			want: []float64{12.3, 10.5, 145.6},
		},
		{
			name: "Float64 tests: Has 2 duplicates",
			args: []float64{9.5, 10.5, 15.9, 10.1, 10.5, 15.9},
			want: []float64{9.5, 10.5, 15.9, 10.1},
		},
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Remove(tt.args))
		})
	}
}
