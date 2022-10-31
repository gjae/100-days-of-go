package main
import (
	"fmt"
	s "strings"
	"unicode"
)


var f = fmt.Printf

func main() {
	f("EqualsFold: %v\n", s.EqualFold("Mihalis", "MIHALis"))
	f("EqualsFold: %v\n", s.EqualFold("Mihalis", "MIHALIS"))

	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "Ha"))

	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Prefix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Prefix: %v\n", s.HasSuffix("Mihalis", "IS"))

	t := s.Fields("This is a string!")
	f("Fields: %v\n", t)

	f("%s\n", s.Split("abcd efg", ""))
	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
}