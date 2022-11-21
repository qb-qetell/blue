package blue

import "strings"
import "sync"
import "time"

type Phnx struct {
	blue       *Blue
	trck       *Trck
	mtxx *sync.Mutex
	core  chan *Mssg
}
	func Phnx_Estb (blue *Blue, trck *Trck, cpct uint16) (*Phnx) {
		_ba00 := &Phnx {
			blue:                          blue,
			trck:                          trck,
			mtxx:                &sync.Mutex {},
			core: make (chan *Mssg, int (cpct)),
		}
		return _ba00
	}
	
	func (i *Phnx) Drop (sndr, rcpn string, mssg *Mssg, wndw ... time.Duration) (bool) {
		_ak00 := strings.Index (sndr, i.blue.idxx + "." + i.trck.idxx)
		if _ak00 !=   0 { return false }
		/*--2--*/
		if  mssg == nil { return false }
		/*--1--*/
		_ba00 := i.blue.mpxx
		if i.blue.idxx + ".!"  ==  rcpn {
			_ba00 = i.blue.spxx
		} else {
			for _,  _cc00 := range i.blue.tixx {
				_da00 := i.blue.idxx +  "." + _cc00
				_db00 := strings.Index (rcpn, _da00)
				if _db00 == 0 {
					_ba00 = i.blue.tdxx [_da00].phnx
					break
				}
			}
		}
		/*--1--*/
		_bb00 := time.Nanosecond * 1
		if wndw != nil && len (wndw) > 0 {
			_bb00 = wndw [0]
		}
		/*--1--*/
		_bc00 := make (chan bool)
		go func (wndw time.Duration, chnl chan bool) {
			time.Sleep (wndw)
			chnl <- true
		} (_bb00, _bc00)
		/*--1--*/
		select {
			case _ba00.core <-  mssg: { return  true }
			case _        = <- _bc00: { return false }
		}
		/*--1--*/
		return false
	}
	
	func (i *Phnx) Ftch (wndw ... time.Duration) (*Mssg) {
		_bb00 := time.Nanosecond * 1
		if wndw != nil && len (wndw) > 0 {
			_bb00 = wndw [0]
		}
		/*--1--*/
		_bc00 := make (chan bool)
		go func (wndw time.Duration, chnl chan bool) {
			time.Sleep (wndw)
			chnl <- true
		} (_bb00, _bc00)
		/*--1--*/
		select {
			case _ca00 := <- i.core: { return _ca00 }
			case _      = <-  _bc00: { return   nil }
		}
	}
