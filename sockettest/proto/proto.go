// proto
//
package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode消息编码
func Encode(message string) ([]byte, error) {
	// int32 占用四个字节
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)

	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err 
	}

	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

// Decode解码
func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuffer := bytes.NewBuffer(lengthByte) // bytes放入Buffer中

	var length int32 // 数据长度
	err := binary.Read(lengthBuffer, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	// 检查缓冲中的数据长度是否够
	if int32(reader.Buffered()) < length + 4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, 4 + length)
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	
	return string(pack[4:]), nil
}