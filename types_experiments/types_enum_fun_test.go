package types_experiments

import "testing"

func TestAirfoil_GetManufacturerAbbreviation(t *testing.T) {
	type fields struct {
		wingspan     wingspan
		manufacturer ParachuteManufacturer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "PerformanceDesigns returns PD",
			fields: fields{
				wingspan:     wingspan{sqft: 120},
				manufacturer: PerformanceDesigns,
			},
			want: "PD",
		},
		{
			name: "NZAeroSports returns NZ",
			fields: fields{
				wingspan:     wingspan{sqft: 120},
				manufacturer: NZAeroSports,
			},
			want: "NZ",
		},
		{
			name: "FluidWings returns FW",
			fields: fields{
				wingspan:     wingspan{sqft: 120},
				manufacturer: FluidWings,
			},
			want: "FW",
		},
		{
			name: "Jyro returns JY",
			fields: fields{
				wingspan:     wingspan{sqft: 120},
				manufacturer: Jyro,
			},
			want: "JY",
		},
		{
			name: "Unknown returns what",
			fields: fields{
				wingspan:     wingspan{sqft: 120},
				manufacturer: Unknown,
			},
			want: "NoAbbreviation",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Airfoil{
				wingspan:     tt.fields.wingspan,
				manufacturer: tt.fields.manufacturer,
			}
			if got := w.GetManufacturerAbbreviation(); got != tt.want {
				t.Errorf("GetManufacturerAbbreviation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAirfoil_GetManufacturerPriority(t *testing.T) {
	type fields struct {
		wingspan     wingspan
		manufacturer ParachuteManufacturer
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "PerformanceDesigns returns 0",
			fields: fields{
				wingspan:     wingspan{sqft: 120},
				manufacturer: PerformanceDesigns,
			},
			want: 0,
		},
		{
			name: "NZAeroSports returns 1",
			fields: fields{
				wingspan:     wingspan{sqft: 120},
				manufacturer: NZAeroSports,
			},
			want: 1,
		},
		{
			name: "FluidWings returns 1",
			fields: fields{
				wingspan:     wingspan{sqft: 120},
				manufacturer: FluidWings,
			},
			want: 1,
		},
		{
			name: "Jyro returns 2",
			fields: fields{
				wingspan:     wingspan{sqft: 120},
				manufacturer: Jyro,
			},
			want: 2,
		},
		{
			name: "Anything else returns 3",
			fields: fields{
				wingspan:     wingspan{sqft: 120},
				manufacturer: Unknown,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Airfoil{
				wingspan:     tt.fields.wingspan,
				manufacturer: tt.fields.manufacturer,
			}
			if got := w.GetManufacturerPriority(); got != tt.want {
				t.Errorf("GetManufacturerPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}
