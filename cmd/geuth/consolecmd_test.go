// Copyright 2016 The eutherum Authors
// This file is part of eutherum.
//
// eutherum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// eutherum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with eutherum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"crypto/rand"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/eutherum/eutherum/params"
)

const (
	ipcAPIs  = "admin:1.0 debug:1.0 eth:1.0 ethash:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0"
	httpAPIs = "eth:1.0 net:1.0 rpc:1.0 web3:1.0"
)

// spawns geuth with the given command line args, using a set of flags to minimise
// memory and disk IO. If the args don't set --datadir, the
// child g gets a temporary data directory.
func runMinimalGeuth(t *testing.T, args ...string) *testgeuth {
	// --ropsten to make the 'writing genesis to disk' faster (no accounts)
	// --networkid=1337 to avoid cache bump
	// --syncmode=full to avoid allocating fast sync bloom
	allArgs := []string{"--ropsten", "--networkid", "1337", "--syncmode=full", "--port", "0",
		"--nat", "none", "--nodiscover", "--maxpeers", "0", "--cache", "64"}
	return runGeuth(t, append(allArgs, args...)...)
}

// Tests that a node embedded within a console can be started up properly and
// then terminated by closing the input stream.
func TestConsoleWelcome(t *testing.T) {
	coinbase := "0x8605cdbbdb6d264aa742e77020dcbc58fcdce182"

	// Start a geuth console, make sure it's cleaned up and terminate the console
	geuth := runMinimalGeuth(t, "--miner.etherbase", coinbase, "console")

	// Gather all the infos the welcome message needs to contain
	geuth.SetTemplateFunc("goos", func() string { return runtime.GOOS })
	geuth.SetTemplateFunc("goarch", func() string { return runtime.GOARCH })
	geuth.SetTemplateFunc("gover", runtime.Version)
	geuth.SetTemplateFunc("geuthver", func() string { return params.VersionWithCommit("", "") })
	geuth.SetTemplateFunc("niltime", func() string {
		return time.Unix(0, 0).Format("Mon Jan 02 2006 15:04:05 GMT-0700 (MST)")
	})
	geuth.SetTemplateFunc("apis", func() string { return ipcAPIs })

	// Verify the actual welcome message to the required template
	geuth.Expect(`
Welcome to the Geuth JavaScript console!

instance: Geuth/v{{geuthver}}/{{goos}}-{{goarch}}/{{gover}}
coinbase: {{.Etherbase}}
at block: 0 ({{niltime}})
 datadir: {{.Datadir}}
 modules: {{apis}}

To exit, press ctrl-d
> {{.InputLine "exit"}}
`)
	geuth.ExpectExit()
}

// Tests that a console can be attached to a running node via various means.
func TestAttachWelcome(t *testing.T) {
	var (
		ipc      string
		httpPort string
		wsPort   string
	)
	// Configure the instance for IPC attachment
	if runtime.GOOS == "windows" {
		ipc = `\\.\pipe\geuth` + strconv.Itoa(trulyRandInt(100000, 999999))
	} else {
		ws := tmpdir(t)
		defer os.RemoveAll(ws)
		ipc = filepath.Join(ws, "geuth.ipc")
	}
	// And HTTP + WS attachment
	p := trulyRandInt(1024, 65533) // Yeah, sometimes this will fail, sorry :P
	httpPort = strconv.Itoa(p)
	wsPort = strconv.Itoa(p + 1)
	geuth := runMinimalGeuth(t, "--miner.etherbase", "0x8605cdbbdb6d264aa742e77020dcbc58fcdce182",
		"--ipcpath", ipc,
		"--http", "--http.port", httpPort,
		"--ws", "--ws.port", wsPort)
	t.Run("ipc", func(t *testing.T) {
		waitForEndpoint(t, ipc, 3*time.Second)
		testAttachWelcome(t, geuth, "ipc:"+ipc, ipcAPIs)
	})
	t.Run("http", func(t *testing.T) {
		endpoint := "http://127.0.0.1:" + httpPort
		waitForEndpoint(t, endpoint, 3*time.Second)
		testAttachWelcome(t, geuth, endpoint, httpAPIs)
	})
	t.Run("ws", func(t *testing.T) {
		endpoint := "ws://127.0.0.1:" + wsPort
		waitForEndpoint(t, endpoint, 3*time.Second)
		testAttachWelcome(t, geuth, endpoint, httpAPIs)
	})
}

func testAttachWelcome(t *testing.T, geuth *testgeuth, endpoint, apis string) {
	// Attach to a running geuth note and terminate immediately
	attach := runGeuth(t, "attach", endpoint)
	defer attach.ExpectExit()
	attach.CloseStdin()

	// Gather all the infos the welcome message needs to contain
	attach.SetTemplateFunc("goos", func() string { return runtime.GOOS })
	attach.SetTemplateFunc("goarch", func() string { return runtime.GOARCH })
	attach.SetTemplateFunc("gover", runtime.Version)
	attach.SetTemplateFunc("geuthver", func() string { return params.VersionWithCommit("", "") })
	attach.SetTemplateFunc("etherbase", func() string { return geuth.Etherbase })
	attach.SetTemplateFunc("niltime", func() string {
		return time.Unix(0, 0).Format("Mon Jan 02 2006 15:04:05 GMT-0700 (MST)")
	})
	attach.SetTemplateFunc("ipc", func() bool { return strings.HasPrefix(endpoint, "ipc") })
	attach.SetTemplateFunc("datadir", func() string { return geuth.Datadir })
	attach.SetTemplateFunc("apis", func() string { return apis })

	// Verify the actual welcome message to the required template
	attach.Expect(`
Welcome to the Geuth JavaScript console!

instance: Geuth/v{{geuthver}}/{{goos}}-{{goarch}}/{{gover}}
coinbase: {{etherbase}}
at block: 0 ({{niltime}}){{if ipc}}
 datadir: {{datadir}}{{end}}
 modules: {{apis}}

To exit, press ctrl-d
> {{.InputLine "exit" }}
`)
	attach.ExpectExit()
}

// trulyRandInt generates a crypto random integer used by the console tests to
// not clash network ports with other tests running cocurrently.
func trulyRandInt(lo, hi int) int {
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(hi-lo)))
	return int(num.Int64()) + lo
}
