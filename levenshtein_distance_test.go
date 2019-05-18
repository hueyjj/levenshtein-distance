package levenshtein

import "testing"

func TestDistance(t *testing.T) {
	s1, s2 := "kitten", "sitting"
	distance := Distance(s1, s2)
	if distance != 3 {
		t.Errorf("Expected Distance(%s, %s)=%d, but got %d", s1, s2, 3, distance)
	}

	s1, s2 = "kitten", "sunday"
	distance = Distance(s1, s2)
	if distance != 6 {
		t.Errorf("Expected Distance(%s, %s)=%d, but got %d", s1, s2, 6, distance)
	}

	s1, s2 = "Sunday", "Saturday"
	distance = Distance(s1, s2)
	if distance != 3 {
		t.Errorf("Expected Distance(%s, %s)=%d, but got %d", s1, s2, 3, distance)
	}
}
