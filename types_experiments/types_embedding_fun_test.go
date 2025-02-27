package types_experiments

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateWingload(t *testing.T) {
	type args struct {
		sqft         int
		exitWeightKg int
	}
	tests := []struct {
		name                    string
		args                    args
		wantWingload            float64
		wantSqftLoadedWithGrams int
	}{
		{
			name:                    "Calculate base",
			args:                    args{exitWeightKg: 100, sqft: 120},
			wantWingload:            1.84,
			wantSqftLoadedWithGrams: 833,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWingload, gotSqftLoadedWithGrams := CalculateWingload(tt.args.sqft, tt.args.exitWeightKg)
			assert.Equalf(t, tt.wantWingload, gotWingload, "CalculateWingload(%v, %v)", tt.args.sqft, tt.args.exitWeightKg)
			assert.Equalf(t, tt.wantSqftLoadedWithGrams, gotSqftLoadedWithGrams, "CalculateWingload(%v, %v)", tt.args.sqft, tt.args.exitWeightKg)
		})
	}
}

func Test_wingspan_CalculateWingload(t *testing.T) {
	type fields struct {
		sqft int
	}
	type args struct {
		exitWeightKg int
	}
	tests := []struct {
		name                    string
		fields                  fields
		args                    args
		wantWingload            float64
		wantSqftLoadedWithGrams int
	}{
		{
			name:                    "Calculate",
			fields:                  fields{sqft: 120},
			args:                    args{exitWeightKg: 100},
			wantWingload:            1.84,
			wantSqftLoadedWithGrams: 833,
		},
		{
			name:                    "Calculate 190sqft 95kg",
			fields:                  fields{sqft: 190},
			args:                    args{exitWeightKg: 95},
			wantWingload:            1.1,
			wantSqftLoadedWithGrams: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wingspan{
				sqft: tt.fields.sqft,
			}
			gotWingload, gotSqftLoadedWithGrams := w.CalculateWingload(tt.args.exitWeightKg)
			assert.Equalf(t, tt.wantWingload, gotWingload, "CalculateWingload(%v)", tt.args.exitWeightKg)
			assert.Equalf(t, tt.wantSqftLoadedWithGrams, gotSqftLoadedWithGrams, "CalculateWingload(%v)", tt.args.exitWeightKg)
		})
	}
}
