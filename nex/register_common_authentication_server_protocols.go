package nex

import (
	"os"

	"github.com/ItzSwirlz/pokemon-gen7/globals"
	nex "github.com/PretendoNetwork/nex-go"
	ticket_granting "github.com/PretendoNetwork/nex-protocols-common-go/ticket-granting"
)

func registerCommonAuthenticationServerProtocols() {
	ticketGrantingProtocol := ticket_granting.NewCommonTicketGrantingProtocol(globals.AuthenticationServer)

	secureStationURL := nex.NewStationURL("")
	secureStationURL.SetScheme("prudps")
	secureStationURL.SetAddress(os.Getenv("PN_POKEGEN7_SECURE_SERVER_HOST"))
	secureStationURL.SetPort(61005) // TODO: this may need to be changed
	secureStationURL.SetCID(1)
	secureStationURL.SetPID(2)
	secureStationURL.SetSID(1)
	secureStationURL.SetStream(10)
	secureStationURL.SetType(2)

	ticketGrantingProtocol.SetSecureStationURL(secureStationURL)
	ticketGrantingProtocol.SetBuildName(serverBuildString)

	globals.AuthenticationServer.SetPasswordFromPIDFunction(globals.PasswordFromPID)
}
