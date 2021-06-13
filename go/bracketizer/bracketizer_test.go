package bracketizer_test

import (
	"code_test/go/bracketizer"
	"testing"
)

var (
	bracketizeUs = []string{
		"{}",    // 0
		"}{",    // 1
		"{{}",   // 2
		"",      // 3
		"     ", // 4
		"F{{DF}SDF{SDFSD}SDF{SDF}SDFSD{FSD}FSD{FS{DFSD}SD{FSDF}SDF}SD}FSD{FSDF}",                  // 5
		"F{{DF}SDF{SDFSD}SDF{SDF}SD{F{S{D{FSD}F{SD{FS{DFSD}SD{FSDF}SDF}SD}FSD{FS}D}F}}",           // 6
		"  F{{ DF}SDF{SD FSD} SDF{SDF}SDFSD{    FSD}FSD{FS{DFSD} SD{FSDF}SD F}SD}FSD{FSD F} ",     // 7
		"  F{{DF}SDF{SDFSD}SDF{SDF}S   D{F{S{D{FSD} F{SD{FS{DFSD}SD{FSDF }SDF}SD}FSD{ FS}D}F}}  ", // 8
	}
)

func TestBracketize(t *testing.T) {

	// 0
	clean := bracketizer.SweepForBrackets(bracketizeUs[0])

	if !clean {
		t.Errorf("SweepForBrackets should be clean. Expected: %v || Got: %v", true, clean)
	}

	// 1
	clean = bracketizer.SweepForBrackets(bracketizeUs[1])

	if clean {
		t.Errorf("SweepForBrackets should be dirty. Expected: %v || Got: %v", false, clean)
	}

	// 2
	clean = bracketizer.SweepForBrackets(bracketizeUs[2])

	if clean {
		t.Errorf("SweepForBrackets should be dirty. Expected: %v || Got: %v", false, clean)
	}

	// 3
	clean = bracketizer.SweepForBrackets(bracketizeUs[3])

	if !clean {
		t.Errorf("SweepForBrackets should be clean. Expected: %v || Got: %v", true, clean)
	}

	// 4
	clean = bracketizer.SweepForBrackets(bracketizeUs[4])

	if !clean {
		t.Errorf("SweepForBrackets should be clean. Expected: %v || Got: %v", true, clean)
	}

	// 5
	clean = bracketizer.SweepForBrackets(bracketizeUs[5])

	if !clean {
		t.Errorf("SweepForBrackets should be clean. Expected: %v || Got: %v", true, clean)
	}

	// 6
	clean = bracketizer.SweepForBrackets(bracketizeUs[6])

	if clean {
		t.Errorf("SweepForBrackets should be dirty. Expected: %v || Got: %v", false, clean)
	}

	// 7
	clean = bracketizer.SweepForBrackets(bracketizeUs[7])

	if !clean {
		t.Errorf("SweepForBrackets should be clean. Expected: %v || Got: %v", true, clean)
	}

	// 8
	clean = bracketizer.SweepForBrackets(bracketizeUs[8])

	if clean {
		t.Errorf("SweepForBrackets should be dirty. Expected: %v || Got: %v", false, clean)
	}
}
