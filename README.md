# Typogenerator: a typosquatting generator in Golang

A Golang typosquat generator utilizing various strategies to generate potential variants of a string. Some strategies are similar to those utilized by [dnstwist](https://github.com/elceef/dnstwist). This library is not intended to be a complete port of dnstwist and may include additional strategies.

## Usage

See files under `cmd/` for example usage.

### Fuzz

```go
all := []strategy.Strategy{
	strategy.Omission,
	strategy.Repetition,
}

results, err := typogenerator.Fuzz("zenithar", all...)
if err != nil {
	fmt.Println(err)
}

for _, r := range results {
	for _, p := range r.Permutations {
		fmt.Println(p)
	}
}

// enithar
// znithar
// zeithar
// zenthar
// ...
```

### FuzzDomain

```go
all := []strategy.Strategy{
	strategy.Omission,
	strategy.Repetition,
}

results, err := typogenerator.FuzzDomain("example.com", all...)
if err != nil {
	fmt.Println(err)
}

for _, r := range results {
	for _, p := range r.Permutations {
		fmt.Println(p)
	}
}

// xample.com
// eample.com
// exmple.com
// exaple.com
// ...
```

## Fuzzing Algorithms (strategies)

1. **Addition** - Addition of a single character to the end of a string
1. **BitSquatting** - Generates a string with one bit difference from the input
1. **DoubleHit** - Addition of a single character that is adjacent to each character (double hitting of a key)
1. **Homoglyph** - Substituiton of a single character with another that looks similar
1. **Hyphenation** - Addition of a hypen `-` between the first and last character in a string
1. **Omission** - Removal of a single character in a string
1. **Prefix** - Addition of predefined prefixes to the start of a string
1. **Repetition** - Repetition of characters in a string (pressing a key twice)
1. **Replace** - Replacement of a single character that is adjacent to each character (pressing wrong key)
1. **Similar** - Replacement of a single character that looks the same. This is a subset of Homoglyph but is language specific.
1. **SubDomain** - Addition of a period `.` between the first and last character in a string
1. **Transposition** - Swapping of adjacent characters in a string
1. **VowelSwap** - Swapping of vowels in string with all other vowels
1. **TLDReplace** - Replaces the TLD with a list of commonly used TLDs. Only works with `FuzzDomain`.
1. **TLDRepeat** - Repeats the TLD after the domain name. Only works with `FuzzDomain`.

## Languages

The following strategies are language dependent:

1. `DoubleHit`
1. `Replace`
1. `Similar`

Supported languages include:

1. English
1. French
1. German
1. Spanish
