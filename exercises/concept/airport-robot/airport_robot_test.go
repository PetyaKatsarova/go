// package airportrobot
package main

import "testing"

// testRunnerTaskID=2
func TestSayHello_Italien(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		want     string
	}{
		{
			testName: "name without spaces",
			name:     "Petka",
			want:     "I can speak Italian: Ciao Petka!",
		},
		{
			testName: "name without spaces",
			name:     "Luiggi",
			want:     "I can speak Italian: Ciao Luiggi!",
		},
		{
			testName: "full name",
			name:     "Tomaso Giulio Micheli",
			want:     "I can speak Italian: Ciao Tomaso Giulio Micheli!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := SayHello(tt.name, Italian{}); got != tt.want {
				t.Errorf("SayHello(%q, \"Italian{}\") = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}

// testRunnerTaskID=3
func TestSayHello_Bulgarian(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		want     string
	}{
		{
			testName: "name without spaces",
			name:     "Petka",
			want:     "I can speak Bulgarian: Zdravei, Petka!",
		},
		{
			testName: "full name",
			name:     "Petka Motoretka",
			want:     "I can speak Bulgarian: Zdravei, Petka Motoretka!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := SayHello(tt.name, Bulgarian{}); got != tt.want {
				t.Errorf("SayHello(%q, \"Bulgarian{}\") = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
