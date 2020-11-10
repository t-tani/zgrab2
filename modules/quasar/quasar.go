package quasar

import(
	"bytes"
	"errors"
	"net"
)

func init(){

}

func checkInitResponse(logStruct *QuasarLog, connection net.Conn) error {
	var data []byte = make([]byte, 4 * 1024)
	size, err := connection.Read(data)
	if err != nil {
		return err
	}
	connection.Close()

	// check tcp payload size
	if size != 0x44 {
		return errors.New("Invalid payload size")
	}

	// check tcp payload header(4bytes)
	if bytes.Compare(data[:4],[]byte{0x40,0x00,0x00,0x00}) != 0{
		return errors.New("Invalid header")
	}

	logStruct.IsQuasar = true
	return nil
}