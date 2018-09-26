package util

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Uint32Flag uint32

func (v *Uint32Flag) String() string {
	return strconv.FormatUint(uint64(*v), 10)
}

func (v *Uint32Flag) Set(s string) error {
	i, err := strconv.ParseUint(s, 0, 32)
	if err == nil {
		*v = Uint32Flag(i)
	}
	return err
}

func IsFlagDefault(s string) (ret bool) {
	ret = true
	flag.Visit(func(flag *flag.Flag) {
		if flag.Name == s {
			ret = false
		}
	})
	return
}

func PrintDefaults() {
	var flags, flagsr []*flag.Flag

	flag.VisitAll(func(flag *flag.Flag) {
		flags = append(flags, flag)
	})
	flagsr = make([]*flag.Flag, len(flags))

	for _, v := range flags {
		u := strings.SplitN(v.Usage, ":", 2)
		if len(u) < 2 {
			flagsr = append(flagsr, v)
			continue
		}
		v.Usage = u[1]
		i, err := strconv.Atoi(u[0])
		if err != nil || i >= len(flags) {
			flagsr = append(flagsr, v)
			continue
		}
		flagsr[i] = v
	}

	for _, v := range flagsr {
		if v == nil {
			continue
		}

		s := fmt.Sprintf("  -%s", v.Name) // Two spaces before -; see next two comments.
		name, usage := flag.UnquoteUsage(v)
		if len(name) > 0 {
			s += " " + name
		}

		// Boolean flags of one ASCII letter are so common we
		// treat them specially, putting their usage on the same line.
		if len(s) <= 4 { // space, space, '-', 'x'.
			s += "\t"
		} else {
			// Four spaces before the tab triggers good alignment
			// for both 4- and 8-space tab stops.
			s += "\n    \t"
		}
		s += strings.Replace(usage, "\n", "\n    \t", -1)

		if (func() bool {
			if _, ok := v.Value.(flag.Getter); ok {
				if _, ok := v.Value.(flag.Getter).Get().(string); ok {
					return true
				}
			}
			return false
		})() {
			// put quotes on the value
			s += fmt.Sprintf(" (default %q)", v.DefValue)
		} else {
			s += fmt.Sprintf(" (default %v)", v.DefValue)
		}
		fmt.Fprint(flag.CommandLine.Output(), s, "\n")
	}
}
