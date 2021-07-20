package main

import (
	"fmt"
	"math"
	"testing"
)

func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name    string
		c       Circle
		want    float64
		wantErr bool
	}{
		{
			name:    "if radius is 8 should return 201,06..., nil",
			c:       Circle{radius: 8},
			want:    math.Pi * math.Pow(8, 2),
			wantErr: false,
		},
		{
			name:    "if radius is less than 0 should return 0, error",
			c:       Circle{radius: 0},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circle.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	tests := []struct {
		name    string
		c       Circle
		want    float64
		wantErr bool
	}{
		{
			name:    "if radius is 8 should return 50,24..., nil",
			c:       Circle{radius: 8},
			want:    2 * math.Pi * 8,
			wantErr: false,
		},
		{
			name:    "if radius is 0 should return 0, error",
			c:       Circle{radius: 0},
			want:    0,
			wantErr: true,
		},
		{
			name:    "if radius is -5 should return 0, error",
			c:       Circle{radius: 0},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circle.Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Circle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_String(t *testing.T) {
	tests := []struct {
		name string
		c    Circle
		want string
	}{
		{
			name: "if radius is 8 should return \"Circle: radius %.2f\" ",
			c:    Circle{radius: 8},
			want: fmt.Sprintf("Circle: radius %.2f", 8.),
		},
		{
			name: "if radius is less than 0 should return \"Incorrect input; radius = %.2f\"",
			c:    Circle{radius: 0},
			want: fmt.Sprintf("Incorrect input; radius = %.2f", 0.),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Circle.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
