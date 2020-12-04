package day04

import "testing"

func TestPassportScanner(t *testing.T) {
	t.Run("An invalid passport", func(t *testing.T) {
		got := ValidPassports([]string{"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884","hcl:#cfa07d byr:1929"})
		want := 0
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("A valid passport", func(t *testing.T) {
		got := ValidPassports([]string{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd","byr:1937 iyr:2017 cid:147 hgt:183cm"})
		want := 1
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("2 valid passports", func(t *testing.T) {
		got := ValidPassports([]string{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd","byr:1937 iyr:2017 cid:147 hgt:183cm", "", "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd","byr:1937 iyr:2017 cid:147 hgt:183cm"})
		want := 2
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("example 1", func(t *testing.T) {
		got := ValidPassports([]string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
			"byr:1937 iyr:2017 cid:147 hgt:183cm",
			"",
			"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
			"hcl:#cfa07d byr:1929",
			"",
			"hcl:#ae17e1 iyr:2013",
			"eyr:2024",
			"ecl:brn pid:760753108 byr:1931",
			"hgt:179cm",
			"",
			"hcl:#cfa07d eyr:2025 pid:166559648",
			"iyr:2011 ecl:brn hgt:59in",
		})
		want := 2
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}