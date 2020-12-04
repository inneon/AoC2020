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

	t.Run("valid birth year", func(t *testing.T) {
		got := validBirthYear("2002")
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	for _, testCase := range []string {
		"",
		"2003",
		"1800",
		"1999AD",
		"123",
	} {
		t.Run("invalid birth year", func(t *testing.T) {
			got := validBirthYear(testCase)
			want := false
			if got != want {
				t.Errorf("got %v want %v for case %q", got, want, testCase)
			}
		})
	}

	t.Run("valid issue year", func(t *testing.T) {
		got := validIssueYear("2015")
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	for _, testCase := range []string {
		"",
		"2025",
		"1800",
		"1999AD",
		"123",
	} {
		t.Run("invalid issue year", func(t *testing.T) {
			got := validIssueYear(testCase)
			want := false
			if got != want {
				t.Errorf("got %v want %v for case %q", got, want, testCase)
			}
		})
	}

	t.Run("valid expiration year", func(t *testing.T) {
		got := validExpirationYear("2025")
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	for _, testCase := range []string {
		"",
		"2015",
		"1800",
		"1999AD",
		"123",
	} {
		t.Run("invalid expiration year", func(t *testing.T) {
			got := validExpirationYear(testCase)
			want := false
			if got != want {
				t.Errorf("got %v want %v for case %q", got, want, testCase)
			}
		})
	}

	for _, testCase := range []string {
		"60in",
		"190cm",
	} {
		t.Run("valid height", func(t *testing.T) {
			got := validHeight(testCase)
			want := true
			if got != want {
				t.Errorf("got %v want %v for case %q", got, want, testCase)
			}
		})
	}

	for _, testCase := range []string {
		"",
		"190",
		"190in",
		"77in",
	} {
		t.Run("invalid height", func(t *testing.T) {
			got := validHeight(testCase)
			want := false
			if got != want {
				t.Errorf("got %v want %v for case %q", got, want, testCase)
			}
		})
	}

	for _, testCase := range []string {
		"#ff00ff",
		"#123456",
	} {
		t.Run("valid hair colour", func(t *testing.T) {
			got := validHairColour(testCase)
			want := true
			if got != want {
				t.Errorf("got %v want %v for case %q", got, want, testCase)
			}
		})
	}

	for _, testCase := range []string {
		"",
		"#123abz",
		"123abc",
		"#ff00fff",
		"f#ff00ff",
	} {
		t.Run("invalid hair colour", func(t *testing.T) {
			got := validHairColour(testCase)
			want := false
			if got != want {
				t.Errorf("got %v want %v for case %q", got, want, testCase)
			}
		})
	}

	t.Run("valid eye colour", func(t *testing.T) {
		got := validEyeColour("brn")
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("invalid eye colour", func(t *testing.T) {
		got := validEyeColour("wat")
		want := false
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	for _, testCase := range []string {
		"000000001",
		"123456789",
	} {
		t.Run("valid passport number", func(t *testing.T) {
			got := validPassportNumber(testCase)
			want := true
			if got != want {
				t.Errorf("got %v want %v for case %q", got, want, testCase)
			}
		})
	}

	for _, testCase := range []string {
		"",
		"asdf",
		"1234567890",
	} {
		t.Run("invalid passport number", func(t *testing.T) {
			got := validPassportNumber(testCase)
			want := false
			if got != want {
				t.Errorf("got %v want %v for case %q", got, want, testCase)
			}
		})
	}

	t.Run("Example 2", func(t *testing.T) {
		got := ValidPassportsStrict([]string{
			"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
			"hcl:#623a2f",
			"",
			"eyr:2029 ecl:blu cid:129 byr:1989",
			"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
			"",
			"hcl:#888785",
			"hgt:164cm byr:2001 iyr:2015 cid:88",
			"pid:545766238 ecl:hzl",
			"eyr:2022",
			"",
			"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
		})
		want := 4
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		got := ValidPassportsStrict([]string{
			"eyr:1972 cid:100",
			"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
			"",
			"iyr:2019",
			"hcl:#602927 eyr:1967 hgt:170cm",
			"ecl:grn pid:012533040 byr:1946",
			"",
			"hcl:dab227 iyr:2012",
			"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
			"",
			"hgt:59cm ecl:zzz",
			"eyr:2038 hcl:74454a iyr:2023",
			"pid:3556412378 byr:2007",
		})
		want := 0
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}