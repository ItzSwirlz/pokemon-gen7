package nex

import (
	"fmt"
	"os"

	"github.com/ItzSwirlz/pokemon-gen7/globals"
	nex "github.com/PretendoNetwork/nex-go"
)

var serverBuildString string

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewServer()
	globals.AuthenticationServer.SetPRUDPVersion(1)
	globals.AuthenticationServer.SetPRUDPProtocolMinorVersion(3)
	globals.AuthenticationServer.SetDefaultNEXVersion(&nex.NEXVersion{
		Major: 3,
		Minor: 3,
		Patch: 0,
	})
	globals.AuthenticationServer.SetKerberosPassword(globals.KerberosPassword)

	// FIXME: What is the access key, we must have it somewhere
	// globals.AuthenticationServer.SetAccessKey("876138df")

	globals.AuthenticationServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==Pokemon S/M/US/UM (Gen 7) - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	registerCommonAuthenticationServerProtocols()

	globals.AuthenticationServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_POKEGEN7_AUTHENTICATION_SERVER_PORT")))
}
