package parser

/*
* Functions needed to read a block from a file
*/

import (
  "bytes"
  "encoding/binary"
)

// Close file
func (f *File) Close() {
	f.file.Close()
}

// Reset sets cursor's position to 0
func (f *File) Reset() {
  // TODO: To implement
}

// Seek moves cursor's position to offset
func (f *File) Seek(offset int64, whence int) (int64, error) {
	return f.file.Seek(offset, whence)
}

// Size returns the size of a file
func (f *File) Size() (int64, error) {
	fInfo, err := f.file.Stat()
	if err != nil {
		return 0, err
	}
	return fInfo.Size(), err
}

// Peek read length bytes without moving cursor
func (f *File) Peek(length int) ([]byte, error) {
	pos, err := f.Seek(0, 1)
	if err != nil {
		return nil, err
	}
	val := make([]byte, length)
	f.file.Read(val)
	_, err = f.Seek(pos, 0)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// ReadByte reads next one byte of data
func (f *File) ReadByte() byte {
	val := make([]byte, 1)
	f.file.Read(val)
	return val[0]
}

// ReadBytes reads next length bytes of data
func (f *File) ReadBytes(length uint64) []byte {
	val := make([]byte, length)
	f.file.Read(val)
	return val
}

// ReadUint16 reads next 4 bytes of data as uint16, LE
func (f *File) ReadUint16() uint16 {
	val := make([]byte, 2)
	f.file.Read(val)
	return binary.LittleEndian.Uint16(val)
}

// ReadInt32 reads next 8 bytes of data as int32, LE
func (f *File) ReadInt32() int32 {
	raw := make([]byte, 4)
	f.file.Read(raw)
	var val int32
	binary.Read(bytes.NewReader(raw), binary.LittleEndian, &val)
	return val
}

// ReadUint32 reads next 8 bytes of data as uint32, LE
func (f *File) ReadUint32() uint32 {
	val := make([]byte, 4)
	f.file.Read(val)
	return binary.LittleEndian.Uint32(val)
}

// ReadInt64 reads next 16 bytes of data as int64, LE
func (f *File) ReadInt64() int64 {
	raw := make([]byte, 8)
	f.file.Read(raw)
	var val int64
	binary.Read(bytes.NewReader(raw), binary.LittleEndian, &val)
	return val
}

// ReadUint64 reads next 16 bytes of data as uint64, LE
func (f *File) ReadUint64() uint64 {
	val := make([]byte, 8)
	f.file.Read(val)
	return binary.LittleEndian.Uint64(val)
}

//func (f *File) ReadVarint() uint64 {
//	chSize := f.ReadByte()
//	if chSize < 253 {
//		return uint64(chSize)
//	} else if chSize == 253 {
//		return uint64(f.ReadUint16())
//	} else if chSize == 254 {
//		return uint64(f.ReadUint32())
//	} else {
//		return f.ReadUint64()
//	}
//}

// ReadVarint reads N byte of data as uint64, LE.
// N depends on the first byte
func (f *File) ReadVarint() uint64 {
	intType := f.ReadByte()
	if intType == 0xFF {
		return f.ReadUint64()
	} else if intType == 0xFE {
		return uint64(f.ReadUint32())
	} else if intType == 0xFD {
		return uint64(f.ReadUint16())
	}

	return uint64(intType)
}

// ShiftVarint reads N byte of data as uint64, LE.
// N depends on the first byte
func (f *File) ShiftVarint() uint64 {
  return 10
}