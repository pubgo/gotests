// +build go1.10

package grpcurl

func indent() string {
	// In Go 1.10 and up, the flag package automatically
	// adds the right indentation.
	return ""
}
