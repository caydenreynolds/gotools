package unique

//Gives a unique string
type Str struct {
	bytes []byte
}

//Returns the next unique string
func (u *Str) Next() string {
	u.incrementByte(len(u.bytes) - 1)
	return string(u.bytes)
}

//Recursively increments the value of the bytes by 1
func (u *Str) incrementByte(byteIndex int) {
	//If value is beyond the end of the bytes, increase the size. Also handles initialisation
	if byteIndex == -1 {
		u.increaseByteWidth()
		byteIndex = 0
	}
	println(byteIndex, len(u.bytes))
	//increment byte
	u.bytes[byteIndex] = u.bytes[byteIndex] + 1

	//If byte has overflowed, increment next byte
	if u.bytes[byteIndex] == 0 {
		u.incrementByte(byteIndex - 1)
	}
}

//Adds one byte the the front of the bytes
func (u *Str) increaseByteWidth() {
	newBytes := make([]byte, len(u.bytes)+1)

	for i := 0; i < len(u.bytes); i++ {
		newBytes[i+1] = u.bytes[i]
	}

	u.bytes = newBytes
	println(len(u.bytes))
}
