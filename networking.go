package networking

import "net";

// Right now, the error message (if any) is discarded from the calls. However, the reason these are being abstracted into another function, is so that later we CAN add our own error handling nice and easy like.

// localhost:port
func StartServer(port int) (server *TCPListener) {
	server, _ := net.ListenTCP("tcp", "localhost:" + String(int))
	return(*server)
}

// Shutdown the server.
func (self *TCPListener) CloseServer() {
	_ := self.Close()
}

// Accept a TCP connection.
func (self *TCPListener) AcceptUser() (conn *TCPConn) {
	con, _ := self.AcceptTCP()
	return(*con)
}