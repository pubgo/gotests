package netdx

var (
	l3, l4, l7 bool
)

const (
	L3 = "l3"
	L4 = "l4"
	L7 = "l7"
)

// OSI Layer 3 - network layer
func turnOnL3()       { l3 = true }
func turnOffL3()      { l3 = false }
func L3Enabled() bool { return l3 }

// OSI Layer 4 - transport layer
func turnOnL4()       { l4 = true }
func turnOffL4()      { l4 = false }
func L4Enabled() bool { return l4 }

// OSI Layer 7 - application layer
func turnOnL7()       { l7 = true }
func turnOffL7()      { l7 = false }
func L7Enabled() bool { return l7 }
