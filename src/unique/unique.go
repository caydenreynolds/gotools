package unique

//Gives a unique string
type Str struct {
	Bytes []byte
}

//Returns the next unique string
func (u *Str) Next() string {
	u.incrementByte(len(u.Bytes) - 1)
	return string(u.Bytes)
}

//Recursively increments the value of the Bytes by 1
func (u *Str) incrementByte(byteIndex int) {
	//If value is beyond the end of the Bytes, increase the size. Also handles initialisation
	if byteIndex == -1 {
		u.increaseByteWidth()
		byteIndex = 0
	}
	println(byteIndex, len(u.Bytes))
	//increment byte
	u.Bytes[byteIndex] = u.Bytes[byteIndex] + 1

	//If byte has overflowed, increment next byte
	if u.Bytes[byteIndex] == 0 {
		u.incrementByte(byteIndex - 1)
	}
}

//Adds one byte the the front of the Bytes
func (u *Str) increaseByteWidth() {
	newBytes := make([]byte, len(u.Bytes)+1)

	for i := 0; i < len(u.Bytes); i++ {
		newBytes[i+1] = u.Bytes[i]
	}

	u.Bytes = newBytes
	println(len(u.Bytes))
}
