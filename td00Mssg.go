package blue

import "time"

type Mssg struct {
	sndr       string
	rcpn       string
	core interface {}
	tmst       string
}
	func Mssg_Estb (core interface {}) (*Mssg) {
		return &Mssg {core: core}
	}
	func (i *Mssg) Send (sndr, rcpn string, phnx *Phnx, wndw ... time.Duration) (bool) {
		if phnx == nil { return false }
		/*--1--*/
		_bb00 := time.Nanosecond * 1
		if wndw != nil && len (wndw) > 0 {
			_bb00 = wndw [0]
		}
		/*--1--*/
		_bc00 := phnx.Drop (sndr, rcpn, i, _bb00)
		/*--1--*/
		return _bc00
	}
	func (i *Mssg) FSXX ()   (   string   ) { return i.sndr }
	func (i *Mssg) FRXX ()   (   string   ) { return i.rcpn }
	func (i *Mssg) FCXX ()   (interface {}) { // Fetch message core
		if i.core == nil { return  "" }
		return i.core
	}
	func (i *Mssg) FTXX ()   (   string   ) { return i.tmst }
