package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stratosnet/stratos-chain/app"
	stratos "github.com/stratosnet/stratos-chain/types"
)

func init() {
	go func() {
		pprofPort := 1234
		fmt.Println("pprof registers handlers and listen port:", pprofPort)
		http.ListenAndServe(":"+strconv.Itoa(pprofPort), nil)
	}()
}

func main() {
	registerDenoms()

	rootCmd, _ := NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}

// RegisterDenoms registers the base and display denominations to the SDK.
func registerDenoms() {
	if err := sdk.RegisterDenom(stratos.DisplayDenom, sdk.OneDec()); err != nil {
		panic(err)
	}

	if err := sdk.RegisterDenom(stratos.USTOS, sdk.NewDecWithPrec(1, stratos.BaseDenomUnit)); err != nil {
		panic(err)
	}
}
