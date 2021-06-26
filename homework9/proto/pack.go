package proto

func Packet(message []byte) []byte {
	protocolVersion := 1
	operation := 2
	sequence := 312312
	pvLength := len(Int16ToBytes(protocolVersion))
	opLength := len(Int32ToBytes(operation))
	sqLength := len(Int32ToBytes(sequence))
	headerLength := pvLength + opLength + sqLength
	h := Int16ToBytes(headerLength)
	packageLength := PackageLength + HeaderLength + headerLength + len(message)
	messageHeader := append(append(append(append(Int32ToBytes(packageLength), h...), Int16ToBytes(protocolVersion)...), Int32ToBytes(operation)...), Int32ToBytes(sequence)...)
	return append(messageHeader, message...)
}
