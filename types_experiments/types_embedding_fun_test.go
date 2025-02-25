package types_experiments

import "testing"

func Test_wingspan_CalculateWingload(t *testing.T) {
	type fields struct {
		sqft int
	}
	type args struct {
		exitWeightLbs int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name:   "Calculate",
			fields: fields{sqft: 120},
			args:   args{exitWeightLbs: 220},
			want:   1.83,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wingspan{
				sqft: tt.fields.sqft,
			}
			if got := w.CalculateWingload(tt.args.exitWeightLbs); got != tt.want {
				t.Errorf("CalculateWingload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateWingload(t *testing.T) {
	type args struct {
		sqft          int
		exitWeightLbs int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Calculate base",
			args: args{exitWeightLbs: 220, sqft: 120},
			want: 1.83,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateWingload(tt.args.sqft, tt.args.exitWeightLbs); got != tt.want {
				t.Errorf("CalculateWingload() = %v, want %v", got, tt.want)
			}
		})
	}
}
