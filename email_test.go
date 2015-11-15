package input

import (
	"testing"
)

func TestParseEmail(t *testing.T) {
	// Test invalid inputs.
	for _, in := range []string{
		"",
		"abc",
		"abc@",
		"@bar.com",
		"abc@bar",
		"a b@x.cz",
		"ab@x.c z",
		"abc@.com",
		"something@@somewhere.com",
		"email@127.0.0.1",
		"email@[127.0.0.1]",
		"email@[2001:dB8::1]",
		"email@[2001:dB8:0:0:0:0:0:1]",
		"email@[::fffF:127.0.0.1]",
		"email@localhost",
		"\"test@test\"@example.com",
		"email@a.b",
		"example@atm.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"example@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.atm",
	} {
		out, err := ParseEmail(in)
		if err == nil {
			t.Errorf("Expected error parsing %s but no error occured, instead %s was returned", in, out)
		}
	}

	// Test valid inputs.
	for in, expOut := range map[string]string{
		"a @ x.cz ":                                                                              "a@x.cz",
		"A @ X.CZ":                                                                               "a@x.cz",
		"email@here.com":                                                                         "email@here.com",
		"email@a.bc":                                                                             "email@a.bc",
		"weirder-email@here.and.there.com":                                                       "weirder-email@here.and.there.com",
		"example@valid-----hyphens.com":                                                          "example@valid-----hyphens.com",
		"example@valid-with-hyphens.com":                                                         "example@valid-with-hyphens.com",
		"test@domain.with.idn.tld.उदाहरण.परीक्षा":                                                "test@domain.with.idn.tld.xn--p1b6ci4b4b3a.xn--11b5bs3a9aj6g",
		"example@atm.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa":            "example@atm.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"example@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.atm":            "example@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.atm",
		"example@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.bbbbbbbbbb.atm": "example@aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.bbbbbbbbbb.atm",
	} {
		actOut, err := ParseEmail(in)
		if err != nil {
			t.Errorf("Unexpected error parsing %s: %v", in, err)
		} else if actOut != expOut {
			t.Errorf("Expected parsed e-mail for %s to be %s, but it is %s", in, expOut, actOut)
		}
	}
}
