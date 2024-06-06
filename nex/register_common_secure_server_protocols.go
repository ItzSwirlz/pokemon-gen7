package nex

import (
	"github.com/ItzSwirlz/pokemon-gen7/globals"
	matchmakingext "github.com/PretendoNetwork/nex-protocols-common-go/matchmaking-ext"
	secureconnection "github.com/PretendoNetwork/nex-protocols-common-go/secure-connection"
)

func registerCommonSecureServerProtocols() {
	secureconnection.NewCommonSecureConnectionProtocol(globals.SecureServer)
	matchmakingext.NewCommonMatchMakingExtProtocol(globals.SecureServer)
}
