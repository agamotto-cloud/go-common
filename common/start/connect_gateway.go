package start

import (
	"github.com/rs/zerolog/log"
	"net"
	"net/http"
)

// 和线上的gateway建立一个TCP连接,获取gateway单线发来的请求,用于调试只能线上调试的接口
// 连接远程的gateway,获取gateway单线发来的请求
func connectGateway(gatewayUrl string, srv *http.Server) {
	log.Info().Str("gateUrl", gatewayUrl).Msg("connect gateway")
	//和gateway建立一个TPC连接
	conn, err := net.Dial("tcp", gatewayUrl)
	if err != nil {
		log.Log().Err(err)
		return
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Log().Err(err)
			return
		}
	}(conn)
	go func() {
		//handleRequest(conn)
		srv.Serve(&TestListener{conn})

	}()

}

func handleRequest(client net.Conn) {

}

// 实现一个 net.Listener
type TestListener struct {
	conn net.Conn
}

func (t *TestListener) Accept() (net.Conn, error) {
	return t.conn, nil
}

func (t *TestListener) Close() error {
	return t.conn.Close()
}
func (t *TestListener) Addr() net.Addr {
	return t.conn.LocalAddr()
}
