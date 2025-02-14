package slices

import (
	"reflect"
	"testing"
)

func TestGetFirstElementInSlice(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Test slice with values",
			args:    args{slice: []string{"a", "b", "c"}},
			want:    "a",
			wantErr: false,
		},
		{
			name:    "Test with with no values",
			args:    args{slice: []string{}},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFirstElementInSlice(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFirstElementInSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFirstElementInSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLastElementInSlice(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test with slice with elements",
			args: args{
				slice: []string{"a", "b", "c"},
			},
			want:    "c",
			wantErr: false,
		},
		{
			name: "Test with slice with no elements",
			args: args{
				slice: []string{},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLastElementInSlice(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLastElementInSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLastElementInSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSubsetOfSlice(t *testing.T) {
	type args struct {
		slice    []string
		startIdx int
		count    int
	}
	tests := []struct {
		name         string
		args         args
		wantNewSlice []string
		wantErr      bool
	}{
		{
			name: "Test success",
			args: args{
				slice:    []string{"a", "b", "c", "d"},
				startIdx: 1,
				count:    2,
			},
			wantNewSlice: []string{"b", "c"},
			wantErr:      false,
		},
		{
			name: "Test slice with with no elements",
			args: args{
				slice:    []string{},
				startIdx: 0,
				count:    1,
			},
			wantNewSlice: nil,
			wantErr:      true,
		},
		{
			name: "Test slice with index above slice length",
			args: args{
				slice:    []string{"a", "b", "c", "d"},
				startIdx: 5,
				count:    2,
			},
			wantNewSlice: nil,
			wantErr:      true,
		},
		{
			name: "Test slice with with count resulting in higher than total",
			args: args{
				slice:    []string{"a", "b", "c", "d"},
				startIdx: 2,
				count:    2,
			},
			wantNewSlice: nil,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewSlice, err := GetSubsetOfSlice(tt.args.slice, tt.args.startIdx, tt.args.count)
			errIsNotNil := (err != nil)
			if errIsNotNil != tt.wantErr {
				t.Errorf("GetSubsetOfSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewSlice, tt.wantNewSlice) {
				t.Errorf("GetSubsetOfSlice() gotNewSlice = %v, want %v", gotNewSlice, tt.wantNewSlice)
			}
		})
	}
}

func TestAppendValueToSlice(t *testing.T) {
	type args struct {
		slice       []string
		appendValue string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test append with slice value",
			args: args{
				slice:       []string{"a", "b", "c"},
				appendValue: "d",
			},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "Test append with nil slice returns a slice with new value",
			args: args{
				slice:       nil,
				appendValue: "d",
			},
			want: []string{"d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AppendValueToSlice(tt.args.slice, tt.args.appendValue)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendValueToSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSliceSplitInHalf(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name           string
		args           args
		wantFirstHalf  []string
		wantSecondHalf []string
		wantErr        bool
	}{
		{
			name:           "Test with even number of elements",
			args:           args{slice: []string{"a", "b", "c", "d"}},
			wantFirstHalf:  []string{"a", "b"},
			wantSecondHalf: []string{"c", "d"},
			wantErr:        false,
		},
		{
			name:           "Test with odd number of elements",
			args:           args{slice: []string{"a", "b", "c", "d", "e"}},
			wantFirstHalf:  []string{"a", "b", "c"},
			wantSecondHalf: []string{"d", "e"},
			wantErr:        false,
		},
		{
			name:           "Test with no number of elements",
			args:           args{slice: []string{}},
			wantFirstHalf:  nil,
			wantSecondHalf: nil,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstHalf, gotSecondHalf, err := GetSliceSplitInHalf(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSliceSplitInHalf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFirstHalf, tt.wantFirstHalf) {
				t.Errorf("GetSliceSplitInHalf() gotFirstHalf = %v, want %v", gotFirstHalf, tt.wantFirstHalf)
			}
			if !reflect.DeepEqual(gotSecondHalf, tt.wantSecondHalf) {
				t.Errorf("GetSliceSplitInHalf() gotSecondHalf = %v, want %v", gotSecondHalf, tt.wantSecondHalf)
			}
		})
	}
}

func TestChuckSlice(t *testing.T) {
	type args struct {
		slice     []string
		chunkSize int
	}
	tests := []struct {
		name        string
		args        args
		wantChucked [][]string
		wantErr     bool
	}{
		{
			name: "Success chunk",
			args: args{
				slice:     []string{"a", "b", "c", "d", "e"},
				chunkSize: 2,
			},
			wantChucked: [][]string{
				{"a", "b"},
				{"c", "d"},
				{"e"},
			},
			wantErr: false,
		}, {
			name: "Negative or zero chuck size",
			args: args{
				slice:     []string{"a", "b", "c", "d", "e"},
				chunkSize: 0,
			},
			wantChucked: nil,
			wantErr:     true,
		}, {
			name: "Empty slice",
			args: args{
				slice:     []string{},
				chunkSize: 1,
			},
			wantChucked: nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChucked, err := ChuckSlice(tt.args.slice, tt.args.chunkSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChuckSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotChucked, tt.wantChucked) {
				t.Errorf("ChuckSlice() gotChucked = %v, want %v", gotChucked, tt.wantChucked)
			}
		})
	}
}

func TestAppendSliceToSlice(t *testing.T) {
	type args struct {
		slice        []string
		appendValues []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Appending slice with another slice",
			args: args{
				slice:        []string{"a", "b"},
				appendValues: []string{"c", "d"},
			},
			want: []string{"a", "b", "c", "d"},
		}, {
			name: "Appending slice with nil slice",
			args: args{
				slice:        []string{"a", "b"},
				appendValues: nil,
			},
			want: []string{"a", "b"},
		}, {
			name: "Appending nil slice with  slice",
			args: args{
				slice:        nil,
				appendValues: []string{"a", "b"},
			},
			want: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendSliceToSlice(tt.args.slice, tt.args.appendValues); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendSliceToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendVariadicValueToSlice(t *testing.T) {
	type args struct {
		slice        []string
		appendValues []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Appending slice with another slice",
			args: args{
				slice:        []string{"a", "b"},
				appendValues: []string{"c", "d"},
			},
			want: []string{"a", "b", "c", "d"},
		}, {
			name: "Appending slice with nil slice",
			args: args{
				slice:        []string{"a", "b"},
				appendValues: nil,
			},
			want: []string{"a", "b"},
		}, {
			name: "Appending nil slice with  slice",
			args: args{
				slice:        nil,
				appendValues: []string{"a", "b"},
			},
			want: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendVariadicValueToSlice(tt.args.slice, tt.args.appendValues...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppendVariadicValueToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func RemoveIndexFromSlice(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
