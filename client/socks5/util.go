package socks5

import (
	"bufio"
	"errors"
	"fmt"
)

// AssertProtocolVersion checks if socks protocol version is supported.
// If can't read version or the version is not support, return an error.
// Otherwise return nil.
func AssertProtocolVersion(reader *bufio.Reader) error {
	version, err := reader.ReadByte()
	if err != nil {
		err = errors.New("Unable to read socks version")
		return err
	}
	if version != ProtocolVersion {
		err = fmt.Errorf("Support socks version %v but found %v", ProtocolVersion, version)
		return err
	}
	return nil
}
