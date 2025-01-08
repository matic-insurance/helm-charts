package golden_testing

import "flag"

type Flags struct {
	UpdateGolden bool
}

var flags = &Flags{}

func ReadFlags() *Flags {
	if flag.Lookup("update-golden") == nil {
		flag.BoolVar(&flags.UpdateGolden, "update-golden", false, "update golden test output files")
	}
	return flags
}
