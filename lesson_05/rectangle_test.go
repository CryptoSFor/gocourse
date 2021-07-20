package main

import (
	"fmt"
	"testing"
)

func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name    string
		r       Rectangle
		want    float64
		wantErr bool
	}{
		{
			name:    "if height is 2 and width is 8 should return 16",
			r:       Rectangle{height: 2, width: 8},
			want:    16.,
			wantErr: false,
		},
		{
			name:    "if height is -2 and width is 8 should return 0",
			r:       Rectangle{height: -2, width: 8},
			want:    0.,
			wantErr: true,
		},
		{
			name:    "if height is 2 and width is -8 should return 0",
			r:       Rectangle{height: 2, width: -8},
			want:    0.,
			wantErr: true,
		},
		{
			name:    "if height is -2 and width is -8 should return 0",
			r:       Rectangle{height: -2, width: -8},
			want:    0.,
			wantErr: true,
		},
		{
			name:    "if height is 0 and width is 8 should return 0",
			r:       Rectangle{height: 0, width: 8},
			want:    0.,
			wantErr: true,
		},
		{
			name:    "if height is -2 and width is 0 should return 0",
			r:       Rectangle{height: -2, width: 0},
			want:    0.,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rectangle.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	tests := []struct {
		name    string
		r       Rectangle
		want    float64
		wantErr bool
	}{
		{
			name:    "if height is 2 and width is 8 should return 16",
			r:       Rectangle{height: 2, width: 8},
			want:    20.,
			wantErr: false,
		},
		{
			name:    "if height is -2 and width is 8 should return 0",
			r:       Rectangle{height: -2, width: 8},
			want:    0.,
			wantErr: true,
		},
		{
			name:    "if height is 2 and width is -8 should return 0",
			r:       Rectangle{height: 2, width: -8},
			want:    0.,
			wantErr: true,
		},
		{
			name:    "if height is -2 and width is -8 should return 0",
			r:       Rectangle{height: -2, width: -8},
			want:    0.,
			wantErr: true,
		},
		{
			name:    "if height is 0 and width is 8 should return 0",
			r:       Rectangle{height: 0, width: 8},
			want:    0.,
			wantErr: true,
		},
		{
			name:    "if height is -2 and width is 0 should return 0",
			r:       Rectangle{height: -2, width: 0},
			want:    0.,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rectangle.Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Rectangle.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_String(t *testing.T) {
	tests := []struct {
		name string
		r    Rectangle
		want string
	}{
		{
			name: "if height = 2 and width = 8 should return \"Rectangle with height 2 and width 8\"",
			r:    Rectangle{height: 2, width: 8},
			want: fmt.Sprintf("Rectangle with height %.2f and width %.2f", 2., 8.),
		},
		{
			name: "if height = -2 and width = 8 should return \"Incorrect input; height = -2 and width = 8\"",
			r:    Rectangle{height: -2, width: 8},
			want: fmt.Sprintf("Incorrect input; height = %.2f and width = %.2f", -2., 8.),
		},
		{
			name: "if height = 2 and width = -8 should return \"Incorrect input; height = 2 and width = -8\"",
			r:    Rectangle{height: 2, width: -8},
			want: fmt.Sprintf("Incorrect input; height = %.2f and width = %.2f", 2., -8.),
		},
		{
			name: "if height = -2 and width = -8 should return \"Incorrect input; height = -2 and width = -8\"",
			r:    Rectangle{height: -2, width: -8},
			want: fmt.Sprintf("Incorrect input; height = %.2f and width = %.2f", -2., -8.),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("Rectangle.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
