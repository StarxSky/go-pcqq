package utils

// 解包
type PackDecrypt struct {
	src []byte
	position int
}

// 置数据
func (pack *PackDecrypt) SetData(data []byte)  {
	pack.src = data
	pack.position = 0
}
// 取字节集
func (pack *PackDecrypt) GetBin(length int) []byte {
	temp := pack.src[0:length]
	pack.src = pack.src[length:]
	pack.position = pack.position + length
	return temp
}
// 取短整数
func (pack *PackDecrypt) GetShort() int16 {
	num := BytesToInt16(pack.src[0:2])
	pack.src = pack.src[2:]
	pack.position = pack.position + 2
	return num
}
// 取长整数
func (pack *PackDecrypt) GetLong() int64 {
	num := BytesToInt64(BytesMerge([]byte{0,0,0,0},pack.src[0:4]))
	pack.src = pack.src[4:]
	pack.position = pack.position + 4
	return num
}
// 取字节
func (pack *PackDecrypt) GetByte() byte {
	temp := pack.src[0]
	pack.src = pack.src[1:]
	pack.position = pack.position + 1
	return temp
}
// 取所有数据
func (pack *PackDecrypt) GetAll() []byte {
	return pack.src
}
// 取位置
func (pack *PackDecrypt) GetPosition() int {
	return pack.position
}
