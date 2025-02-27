package types_experiments

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyJumperCanJump(t *testing.T) {
	type args struct {
		ruleset   IsParachuteAllowedForJumper
		jumper    Skydiver
		parachute Parachute
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Success with AllowedToJump",
			args: args{
				ruleset: AllowedToJump,
				jumper: Skydiver{
					nakedWeightKg:        85,
					gearWeightKg:         10,
					jumpCount:            1,
					highPerformanceJumps: 0,
				},
				parachute: Parachute{
					wingspan:     wingspan{sqft: 99},
					manufacturer: PerformanceDesigns,
					Level:        Competition,
					Elliptical:   false,
				},
			},
			want: true,
		},
		{
			name: "Not allowed to jump with NotAllowedToJump",
			args: args{
				ruleset: NotAllowedToJump,
				jumper: Skydiver{
					nakedWeightKg:        65,
					gearWeightKg:         10,
					jumpCount:            1,
					highPerformanceJumps: 0,
				},
				parachute: Parachute{
					Level:        Beginner,
					Elliptical:   false,
					manufacturer: PerformanceDesigns,
					wingspan:     wingspan{sqft: 280},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, VerifyJumperCanJump(tt.args.ruleset, tt.args.jumper, tt.args.parachute), "VerifyJumperCanJump(%v, %v, %v)", tt.args.ruleset, tt.args.jumper, tt.args.parachute)
		})
	}
}

func TestDFURuleset(t *testing.T) {
	type args struct {
		jumper    Skydiver
		parachute Parachute
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "DFU Jumper beginner jumper exceed load limit for 190 ",
			args: args{
				jumper: Skydiver{
					nakedWeightKg:        90,
					gearWeightKg:         10,
					jumpCount:            180,
					highPerformanceJumps: 0,
				},
				parachute: Parachute{
					Level:      Intermediary,
					Elliptical: false,
					wingspan:   wingspan{sqft: 190},
				},
			},
			want: false,
		},
		{
			name: "DFU Jumper beginner jumper within load limit for 190 ",
			args: args{
				jumper: Skydiver{
					nakedWeightKg:        80,
					gearWeightKg:         10,
					jumpCount:            180,
					highPerformanceJumps: 0,
				},
				parachute: Parachute{
					Level:      Intermediary,
					Elliptical: false,
					wingspan:   wingspan{sqft: 190},
				},
			},
			want: true,
		},
		{
			name: "DFU Jumper more experienced jumper has higher load limit for 190 ",
			args: args{
				jumper: Skydiver{
					nakedWeightKg:        90,
					gearWeightKg:         10,
					jumpCount:            200,
					highPerformanceJumps: 0,
				},
				parachute: Parachute{
					Level:      Intermediary,
					Elliptical: false,
					wingspan:   wingspan{sqft: 190},
				},
			},
			want: true,
		},
		{
			name: "DFU Jumper more experienced jumper has higher load limit for 190, but not elliptical ",
			args: args{
				jumper: Skydiver{
					nakedWeightKg:        90,
					gearWeightKg:         10,
					jumpCount:            200,
					highPerformanceJumps: 0,
				},
				parachute: Parachute{
					Level:      Intermediary,
					Elliptical: true,
					wingspan:   wingspan{sqft: 190},
				},
			},
			want: false,
		},
		{
			name: "DFU Jumper more experienced jumper has higher load limit but still not allowed to jump smaller parachute",
			args: args{
				jumper: Skydiver{
					nakedWeightKg:        50,
					gearWeightKg:         10,
					jumpCount:            399,
					highPerformanceJumps: 0,
				},
				parachute: Parachute{
					Level:      Intermediary,
					Elliptical: true,
					wingspan:   wingspan{sqft: 107},
				},
			},
			want: false,
		},
		{
			name: "DFU Jumper more experienced jumper cannot jump high performance before at least 600 jumps",
			args: args{
				jumper: Skydiver{
					nakedWeightKg:        70,
					gearWeightKg:         10,
					jumpCount:            505,
					highPerformanceJumps: 0,
				},
				parachute: Parachute{
					Level:      HighPerformance,
					Elliptical: false,
					wingspan:   wingspan{sqft: 107},
				},
			},
			want: false,
		},
		{
			name: "DFU Jumper more experienced jumper can jump high performance after at least 600 jumps",
			args: args{
				jumper: Skydiver{
					nakedWeightKg:        70,
					gearWeightKg:         10,
					jumpCount:            600,
					highPerformanceJumps: 0,
				},
				parachute: Parachute{
					Level:      HighPerformance,
					Elliptical: false,
					wingspan:   wingspan{sqft: 107},
				},
			},
			want: true,
		},
		{
			name: "DFU Jumper very experienced jumper cannot jump competition before at least 800 jumps with 200 HighPerformance",
			args: args{
				jumper: Skydiver{
					nakedWeightKg:        70,
					gearWeightKg:         10,
					jumpCount:            900,
					highPerformanceJumps: 190,
				},
				parachute: Parachute{
					Level:      Competition,
					Elliptical: false,
					wingspan:   wingspan{sqft: 80},
				},
			},
			want: false,
		},
		{
			name: "DFU Jumper very experienced jumper can jump competition before at least 800 jumps with 200 HighPerformance",
			args: args{
				jumper: Skydiver{
					nakedWeightKg:        70,
					gearWeightKg:         10,
					jumpCount:            900,
					highPerformanceJumps: 200,
				},
				parachute: Parachute{
					Level:      Competition,
					Elliptical: false,
					wingspan:   wingspan{sqft: 80},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DFURuleset(tt.args.jumper, tt.args.parachute), "DFURuleset(%v, %v)", tt.args.jumper, tt.args.parachute)
		})
	}
}
