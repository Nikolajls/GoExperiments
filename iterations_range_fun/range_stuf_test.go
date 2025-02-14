package iterations_range_fun

import "testing"

func TestSumOverSliceInt(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ADditions",
			args: args{
				[]int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOverSliceInt(tt.args.numbers); got != tt.want {
				t.Errorf("SumOverSliceInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddNumbersFromStartToEnd(t *testing.T) {
	type args struct {
		startNo int
		endNo   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				startNo: 1,
				endNo:   5,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddNumbersFromStartToEnd(tt.args.startNo, tt.args.endNo); got != tt.want {
				t.Errorf("AddNumbersFromStartToEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMapDictionaryContainKeyValueLookup(t *testing.T) {
	type args struct {
		dict map[string]string
		key  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				dict: map[string]string{"A": "VAL", "B": "VALUE"},
				key:  "B",
			},
			want:    "VALUE",
			wantErr: false,
		}, {
			name: "Failure due to not containing key",
			args: args{
				dict: map[string]string{"A": "VAL", "B": "VALUE"},
				key:  "C",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMapDictionaryContainKeyValueLookup(tt.args.dict, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMapDictionaryContainKeyValueLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMapDictionaryContainKeyValueLookup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMapDictionaryContainKeyValueIteration(t *testing.T) {
	type args struct {
		dict map[string]string
		key  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				dict: map[string]string{"A": "VAL", "B": "VALUE"},
				key:  "B",
			},
			want:    "VALUE",
			wantErr: false,
		}, {
			name: "Failure due to not containing key",
			args: args{
				dict: map[string]string{"A": "VAL", "B": "VALUE"},
				key:  "C",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMapDictionaryContainKeyValueIteration(tt.args.dict, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMapDictionaryContainKeyValueIteration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMapDictionaryContainKeyValueIteration() got = %v, want %v", got, tt.want)
			}
		})
	}
}
