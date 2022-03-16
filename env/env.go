package env

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	Built   = "0"
	Version = "0.0.0"
	Githash = "no githash set"
	OSArch  = "no arch set"
)
var e *RuntimeEnv
var m sync.Mutex

type RuntimeEnv struct {
	Built     string `json:"built"`
	Version   string `json:"version"`
	GoVersion string `json:"goVersion"`
	Githash   string `json:"githash"`
	OSArch    string `json:"osarch"`
}

func CompileInfo() *RuntimeEnv {
	m.Lock()
	defer m.Unlock()

	if e == nil {
		e = &RuntimeEnv{
			Built:     Built,
			Version:   Version,
			GoVersion: runtime.Version(),
			Githash:   Githash,
			OSArch:    OSArch,
		}
	}
	return e
}

func (e *RuntimeEnv) Print(named string) {
	t, _ := strconv.ParseInt(e.Built, 0, 64)
	fmt.Printf(`%s (release)
  Version: %s
  Go version: %s
  Git commit: %s
  OS/Arch: %s
  Built: %s
`, named,
		e.Version,
		e.GoVersion,
		e.Githash,
		e.OSArch,
		time.Unix(t, 0),
	)
}
