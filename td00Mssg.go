package blue

import "time"

type Mssg struct {
	sndr  *Trck
	rcpn  *Trck
	core interface {}
	tmst string
}
	func Mssg_Estb () (*Mssg) { return nil }
	func (*Mssg) Fill ([]byte) {}
	func (*Mssg) Send (rcpn  string, phnx *Phnx, wndw ... time.Duration) (error) { return nil}
	func (*Mssg) FSXX () (   string   ) { return ""  }
	func (*Mssg) FRXX () (   string   ) { return ""  }
	func (*Mssg) FCXX () (interface {}) { return nil }
	func (*Mssg) FTXX () (   string   ) { return ""  }

