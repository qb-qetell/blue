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
	func phnx_Estb (blue *Blue, trck *Trck, cpct uint16) (*Phnx) {
		_ba00 := &Phnx {
			blue:                          blue,
			trck:                          trck,
			mtxx:                &sync.Mutex {},
			core: make (chan *Mssg, int (cpct)),
		}
		return _ba00
	}
	
	func (i *Phnx) Drop (sndr, rcpn string, mssg *Mssg, wndw ... time.Duration) (bool) {
		_ak00 := strings.Index (sndr, i.trck.idxx)
		if _ak00 !=   0 { return false }
		/*--2--*/
		if  mssg == nil { return false }
		/*--1--*/
		mssg.sndr = sndr
		/*--1--*/
		_ba00   := i.blue.mtxx
		if rcpn == i.blue.stxx.idxx  {
			_ba00 = i.blue.stxx
		} else {
			for _,  _cc00 := range  i.blue.tixx {
				_db00 := strings.Index (rcpn, _cc00)
				if _db00 == 0 {
					_ba00 = i.blue.tdxx [ _cc00]
					break
				}
			}
		}
		mssg.rcpn = rcpn
		/*--1--*/
		_bb00   := time.Millisecond   * 1
		if wndw != nil && len (wndw) > 0 {
			_bb00 = wndw [0]
		}
		/*--1--*/
		_bc00 := make (chan bool)
		go func (wndw time.Duration, chnl chan bool) {
			time.Sleep (wndw)
			chnl <- false
		} (_bb00, _bc00)
		/*--1--*/
		mssg.tmst = time.Now ().In (time.FixedZone ("+0000", 0)).Format (
			"2006-01-02 15:04:05 -0700",
		)
		select {
			case _ba00.phnx.core <-  mssg: {
				go func (chnl chan bool)  { _ = <- chnl } (_bc00)
				return true
			}
			case _ = <- _bc00: { return false }
		}
		/*--1--*/
		return false
	}
	
	func (i *Phnx) Ftch (wndw ... time.Duration) (*Mssg) {
		_bb00 := time.Millisecond * 1
		if wndw != nil && len (wndw) > 0 {
			_bb00 = wndw [0]
		}
		/*--1--*/
		_bc00 := make (chan bool)
		go func (wndw time.Duration, chnl chan bool) {
			time.Sleep (wndw)
			chnl <- false
		} (_bb00, _bc00)
		/*--1--*/
		select {
			case _ca00 := <- i.core: {
				go func (chnl chan bool)  { _ = <- chnl } (_bc00)
				return _ca00
			}
			case _ = <- _bc00: { return   nil }
		}
	}
func (i *Phnx) NtfySystAbtxFldxStrt (rsnx string, wndw ... time.Duration) (bool) {
	_ba00 := time.Millisecond * 4
	if wndw != nil && len (wndw) > 0 {
		_ba00 = wndw [0]
	}
	/*--1--*/
	_bb00 :=            Mssg_Estb ([]string {"txxx.ba00",   rsnx})
	_bc00 := _bb00.Send (i.trck.idxx, i.blue.stxx.idxx, i, _ba00 )
	return _bc00
}
func (i *Phnx) NtfySystAbtxScsfStrt (wndw ... time.Duration)  (bool) {
	_ba00 := time.Millisecond * 4
	if wndw != nil && len (wndw) > 0 {
		_ba00 = wndw [0]
	}
	/*--1--*/
	_bb00 :=                    Mssg_Estb ([]string {"txxx.ba10"})
	_bc00 := _bb00.Send (i.trck.idxx, i.blue.stxx.idxx, i, _ba00 )
	return _bc00
}
