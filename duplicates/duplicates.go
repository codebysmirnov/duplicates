// Duplicates is feature set for work with duplicates.
// For work with this set you need use go1.18 version
// because all features use generics.

package duplicates

// GetAll return slice of all duplicates from
// fom transferred slice
func GetAll[T comparable](s []T) (d []T) {
	if len(s) <= 1 {
		return d
	}

	seen := make(map[T]struct{}, len(s))
	for i := range s {
		if _, ok := seen[s[i]]; !ok {
			seen[s[i]] = struct{}{}
			continue
		}
		d = append(d, s[i])
	}

	return d
}

// Detect returns result of have any duplicates check
// if function detect duplicate - function return true
func Detect[T comparable](s []T) bool {
	if len(s) <= 1 {
		return false
	}
	seen := make(map[T]struct{}, len(s))
	for i := range s {
		if _, ok := seen[s[i]]; !ok {
			seen[s[i]] = struct{}{}
			continue
		}
		return true
	}
	return false
}

// Count returns count of all duplicates
func Count[T comparable](s []T) int {
	return len(GetAll(s))
}

// Remove returns new slice without duplicates
func Remove[T comparable](s []T) (ns []T) {
	if len(s) <= 1 {
		return s
	}
	seen := make(map[T]struct{}, len(s))
	for i := range s {
		if _, ok := seen[s[i]]; !ok {
			ns = append(ns, s[i])
			seen[s[i]] = struct{}{}
			continue
		}
	}
	return ns
}
