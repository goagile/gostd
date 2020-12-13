package bit

// Bit - representation of bit
type Bit uint

// Get - get "i" bit (true | false)
func (b Bit) Get(i uint) bool {
	return (b & (1 << i)) != 0
}

// Set - set "i" bit to "1"
func (b Bit) Set(i uint) Bit {
	return Bit(b | (1 << i))
}

// Clear - clear "i" bit
func (b Bit) Clear(i uint) Bit {
	return Bit(b &^ (1 << i))
}

// Update - update "i" bit to v
func (b Bit) Update(i uint, s bool) Bit {
	v := 0
	if s == true {
		v = 1
	}
	m := uint(^(1 << i))
	return Bit((uint(b) & m) | uint(v<<i))
}
