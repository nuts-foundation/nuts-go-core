package docs

import (
	"github.com/spf13/pflag"
	"os"
	"sort"
	"strings"
)

// GeneratePartitionedConfigOptionsDocs is like GenerateConfigOptionsDocs but renders a table with subsections.
func GeneratePartitionedConfigOptionsDocs(fileName string, flags map[string]*pflag.FlagSet) {
	sortedKeys := make([]string, 0)
	for key, _ := range flags {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	values := make([][]rstValue, 0)
	for _, key := range sortedKeys {
		// Make partition captions bold
		values = append(values, []rstValue{{
			value: key,
			bold:  true,
		}})
		values = append(values, flagsToSortedValues(flags[key])...)
	}
	generateRstTable(fileName, values)
}


// GenerateConfigOptionsDocs generates an RST-style (ReStructuredText) table for the given flag set, which can serve
// as documentation for the application. If the target fileName exists, it is overwritten.
func GenerateConfigOptionsDocs(fileName string, flags *pflag.FlagSet) {
	values := flagsToSortedValues(flags)
	generateRstTable(fileName, values)
}

func flagsToSortedValues(flags *pflag.FlagSet) [][]rstValue {
	values := make([][]rstValue, 0)
	flags.VisitAll(func(f *pflag.Flag) {
		if f.Hidden {
			return
		}
		values = append(values, vals(f.Name, f.DefValue, f.Usage))
	})
	// We want global properties (the ones without dots) to appear at the top, so we need some custom sorting
	sort.Slice(values, func(i, j int) bool {
		s1 := values[i][0].value
		s2 := values[j][0].value
		if strings.Contains(s1, ".") {
			if strings.Contains(s2, ".") {
				return s1 < s2
			}
			return false
		} else {
			if strings.Contains(s2, ".") {
				return true
			}
			return s1 < s2
		}
	})
	return values
}

func generateRstTable(fileName string, values [][]rstValue) {
	optionsFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer optionsFile.Close()
	printRstTable(vals("Key", "Default", "Description"), values, optionsFile)
	if err := optionsFile.Sync(); err != nil {
		panic(err)
	}
}
