package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)

type Value struct {
	typ   string
	str   string
	num   int
	bulk  string
	array []Value
}

type Resp struct {
	reader *bufio.Reader
}

func NewResp(reader io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(reader)}
}

func (r *Resp) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		n += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], n, nil
}

func (r *Resp) readInteger() (x int, err error) {
	line, _, err := r.readLine()
	if err != nil {
		return 0, err
	}

	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i64), nil
}

func (r *Resp) readArray() (Value, error) {
	v := Value{}
	v.typ = "array"

	size, err := r.readInteger()
	// fmt.Println("Size is ", size)

	if err != nil {
		fmt.Printf("Error reading size of array", err)
		return Value{}, err
	}

	v.array = make([]Value, size)

	for i := 0; i < size; i++ {
		val, err := r.Read()
		if err != nil {
			return v, err
		}

		// add parsed value to array
		v.array = append(v.array, val)
	}
	// fmt.Println("Returning from readArray", v)
	return v, nil
}

func (r *Resp) readBulk() (Value, error) {
	v := Value{}

	v.typ = "bulk"

	len, err := r.readInteger()
	if err != nil {
		return v, err
	}

	// fmt.Println("length of bulk", len)

	bulk := make([]byte, len)

	r.reader.Read(bulk)

	v.bulk = string(bulk)

	// fmt.Println("bulk string is ", string(bulk))

	// Read the trailing CRLF
	r.readLine()
	// fmt.Println("returning from readBulk", v)

	return v, nil
}

func (r *Resp) Read() (Value, error) {
	_type, error := r.reader.ReadByte()

	if error != nil {
		return Value{}, error
	}

	switch _type {
	case ARRAY:
		return r.readArray()
	case BULK:
		return r.readBulk()
	default:
		fmt.Printf("Unknown value observed: %v", _type)
		return Value{}, nil
	}
}
