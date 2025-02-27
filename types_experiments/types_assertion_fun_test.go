package types_experiments

import "testing"

type mySqftStruct struct {
	sqft int
}

func (s *mySqftStruct) Sqft() int {
	return s.sqft
}

func TestCalculateWingLoadSafeTypeAssertion(t *testing.T) {

	type args struct {
		arg          interface{}
		exitWeightKg int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Success with int",
			args: args{
				arg:          120,
				exitWeightKg: 100,
			},
			want:    1.84,
			wantErr: false,
		},

		{
			name: "Success with wingspan",
			args: args{
				arg:          wingspan{sqft: 120},
				exitWeightKg: 100,
			},
			want:    1.84,
			wantErr: false,
		},
		{
			name: "Success with parachute ",
			args: args{
				arg: Parachute{
					wingspan{sqft: 120},
					PerformanceDesigns,
					Competition,
					false,
				},
				exitWeightKg: 100,
			},
			want:    1.84,
			wantErr: false,
		},
		{
			name: "Success with struct with an Sqft() method ",
			args: args{
				arg:          &mySqftStruct{sqft: 120},
				exitWeightKg: 100,
			},
			want:    1.84,
			wantErr: false,
		},
		{
			name: "Fails with string",
			args: args{
				arg:          "120",
				exitWeightKg: 100,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateWingLoadSafeTypeAssertion(tt.args.arg, tt.args.exitWeightKg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateWingLoadWithTypeAssertionOfWingspan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateWingLoadWithTypeAssertionOfWingspan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateWingLoadSwitchTypeAssertion(t *testing.T) {
	type args struct {
		arg          interface{}
		exitWeightKg int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Success with int",
			args: args{
				arg:          120,
				exitWeightKg: 100,
			},
			want:    1.84,
			wantErr: false,
		},

		{
			name: "Success with wingspan",
			args: args{
				arg:          wingspan{sqft: 120},
				exitWeightKg: 100,
			},
			want:    1.84,
			wantErr: false,
		},
		{
			name: "Success with airfoil ",
			args: args{
				arg: Parachute{
					wingspan{sqft: 120},
					FluidWings,
					HighPerformance,
					false,
				},
				exitWeightKg: 100,
			},
			want:    1.84,
			wantErr: false,
		}, {
			name: "Success with struct with an Sqft() method ",
			args: args{
				arg:          &mySqftStruct{sqft: 120},
				exitWeightKg: 100,
			},
			want:    1.84,
			wantErr: false,
		},
		{
			name: "Fails with string",
			args: args{
				arg:          "120",
				exitWeightKg: 220,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateWingLoadSwitchTypeAssertion(tt.args.arg, tt.args.exitWeightKg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateWingLoadWithTypeAssertionOfWingspan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalculateWingLoadWithTypeAssertionOfWingspan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateWingLoadNotSafeTypeAssertion(t *testing.T) {
	type args struct {
		arg          interface{}
		exitWeightKg int
	}
	tests := []struct {
		name      string
		args      args
		want      float64
		wantErr   bool
		wantPanic bool
	}{
		{
			name: "Success with wingspan",
			args: args{
				arg:          wingspan{sqft: 120},
				exitWeightKg: 100,
			},
			want:      1.84,
			wantPanic: true,
		},
		{
			name: "Failure with int",
			args: args{
				arg:          120,
				exitWeightKg: 100,
			},
			want:      0,
			wantPanic: true,
		},
		{
			name: "Failure with Parachute",
			args: args{
				arg: Parachute{
					wingspan{sqft: 120},
					FluidWings,
					Intermediary,
					false,
				},
				exitWeightKg: 100,
			},
			want:      0,
			wantPanic: true,
		},
		{
			name: "Fails with string",
			args: args{
				arg:          "120",
				exitWeightKg: 100,
			},
			want:      0,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil && !tt.wantPanic {
					t.Errorf("The code did not panic")
				}
			}()

			got := CalculateWingLoadNotSafeTypeAssertion(tt.args.arg, tt.args.exitWeightKg)

			if got != tt.want {
				t.Errorf("CalculateWingLoadWithTypeAssertionOfWingspan() got = %v, want %v", got, tt.want)
			}
		})
	}
}
