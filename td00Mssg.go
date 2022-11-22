package blue

import "time"

type Mssg struct {
	sndr        *Trck
	rcpn        *Trck
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
	func (i *Mssg) FSXX ()   (   string   ) { // Fetch message sender
		if i.sndr == nil { return  "" }
		return i.sndr.idxx
	}
	func (i *Mssg) FRXX ()   (   string   ) { // Fetch message recipient
		if i.rcpn == nil { return  "" }
		return i.rcpn.idxx
	}
	func (i *Mssg) FCXX ()   (interface {}) { // Fetch message core
		if i.core == nil { return  "" }
		return i.core
	}
	func (i *Mssg) FTXX ()   (   string   ) { // Fetch message timestamp
		return i.tmst 
	}
