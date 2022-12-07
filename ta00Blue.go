package blue

import "errors"
import "regexp"
import "strings"
import "time"

type Blue struct {
	idxx           string // ID
	mtxx            *Trck // Mother's phone
	stxx            *Trck // System's phone
	tixx         []string // Trck index
	tdxx map[string]*Trck // Trck details
	sisx             bool // Shutdown in-progress status
	scsx             bool // Shutdown complete status
	lsxx             bool // Life status
	usxx             bool // Use status
}
	func Blue_Estb (idxx string, mpcx, spcx uint16) (error, *Blue) {
		_ba00 := strings.ToLower (idxx)
		if regexp.MustCompile(`^[a-z0-9]+(\.[a-z0-9]+)*$`).MatchString (_ba00) == false {
			_ca00 := errors.New (
				"String can not be used as ID. [It contains one or more " +
				"illegal characters.]",
			)
			return _ca00, nil
		}
		_bb00 := &Blue {
			idxx:                        _ba00,
			mtxx:                          nil,
			stxx:                          nil,
			tixx:                  []string {},
			tdxx:      make (map[string]*Trck),
			sisx:                        false,
			scsx:                        false,
			lsxx:                        false,
			usxx:                        false,
		}
		_bc00 := &Trck {
			blue:                        _bb00,
			idxx:                        _ba00,
			name:               "Mother Track",
			phnx: phnx_Estb (_bb00, nil, mpcx),
			edxx:                          nil,
			code:                          nil,
			ssxx:                          "-",
			lsxx:                          "-",
		}
		_bb00.mtxx           = _bc00
		_bb00.mtxx.phnx.trck = _bc00
		_bd00 := &Trck {
			blue:                        _bb00,
			idxx:                 _ba00 + ".!",
			name:               "System Track",
			phnx: phnx_Estb (_bb00, nil, spcx),
			edxx:                          nil,
			code:                          nil,
			ssxx:                          "-",
			lsxx:                          "-",
		}
		_bb00.stxx           = _bd00
		_bb00.stxx.phnx.trck = _bd00
		/*--1--*/
		return nil, _bb00
	}
	
	func (i *Blue) Elbr (
		idxx       string, // ID
		name       string, // Name
		pcxx       uint16, // Phone capacity
		edxx interface {}, // Execution data
		code         func(mixx, idxx, name string, phnx *Phnx, edxx interface {})(error),
	 ) (error) {
		_ba00 := strings.ToLower (idxx)
		if regexp.MustCompile(`^[a-z0-9]+(\.[a-z0-9]+)*$`).MatchString (_ba00) == false {
			_ca00 := errors.New (
				"String can not be used as ID. [It contains one or more " +
				"illegal characters.]",
			)
			return _ca00
		}
		_bb00 := i.idxx + "." + _ba00
		for _, _bc00 := range i.tixx {
			if _bc00 == _bb00 {
				_ca00 := errors.New ("A track is already using this ID.")
				return _ca00
			}
		}
		/*--2--*/
		if code == nil {
			_ca00 := errors.New ("Track's code was not provided.")
			return _ca00
		}
		/*--1--*/
		_bd00 := &Trck {
			blue:                        i,
			idxx:                    _bb00,
			name:                     name,
			phnx: phnx_Estb (i, nil, pcxx),
			edxx:                     edxx,
			code:                     code,
			ssxx:                      "-",
			lsxx:                      "-",
		}
		_bd00.phnx.trck = _bd00
		/*--1--*/
		i.tixx = append (i.tixx, _bb00)
		i.tdxx [_bb00] = _bd00
		/*--1--*/
		return nil
	}
	
	func (i *Blue) Strt () (error, *Phnx) {
		if i.usxx == true {
			_ca00 := errors.New ("Blum has already been run.")
			return _ca00, nil
		}
		/*--1--*/
		i.usxx = true
		/*--1--*/
		
		/*--1--*/
		go func (i *Blue) {
			go tmxx (i)
			for _, _da00   := range i.tixx {
				go i.tdxx [_da00].runx ()
				for {
					if        i.tdxx [_da00].ssxx == "-" {
						time.Sleep (time.Millisecond * 1)
						continue
					} else if i.tdxx [_da00].ssxx == "f" {
						return
					} else if i.tdxx [_da00].ssxx == "s" {
						break
					}
				}
			}
			_cb00 :=       Mssg_Estb ([]string {"sxxx.ba15"}) // Start up successful
			_cb00.Send (i.stxx.idxx, i.mtxx.idxx, i.stxx.phnx, (time.Hour * 240000))
		} (i)
		/*--1--*/
		return nil, i.mtxx.phnx
	}
	
	func (i *Blue) Halt () {
		_ba00 := Mssg_Estb ([]string {"txxx.bc00"})
		_ba00.Send (i.mtxx.idxx, i.stxx.idxx, i.mtxx.phnx, (time.Hour * 240000))
		return
	}
	
	func tmxx (i *Blue) { for {
		if i.scsx == true {
			_ba00 := Mssg_Estb ([]string {"sxxx.bc00"})
			_ba00.Send (i.stxx.idxx, i.mtxx.idxx, i.stxx.phnx, (time.Hour * 240000))
			return
		}
		/*--1--*/
		_ba00 := i.stxx.phnx.Ftch (time.Millisecond * 1)
		if _ba00 ==  nil {
			continue
		}
		/*--1--*/
		_bb00, _bc00 := _ba00.FCXX ().([]string)
		if _bc00 == false  || len (_bb00) < 1 {
			continue
		}
		/*--1--*/
		if /*--*/ _bb00 [0]   ==  "txxx.ba00" { // Track could not start up
			if len (_bb00) < 2 { continue }
			_ca00  :=  ""
			for _,  _cc00 := range i.tixx {
				_ca00  = _cc00
				_db00 := strings.Index (_ba00.sndr, _cc00)
				if _db00 == 0 {
					i.tdxx [_cc00].ssxx = "f"
					i.tdxx [_cc00].snxx = _bb00 [1]
					break
				}
			}
			if _ca00 == "" { continue }
			/*--2--*/
			_cc50 := _bb00 [1]
			if len (_bb00) < 2 { _cc50 = "" }
			/*--2--*/
			_cd00 :=                Mssg_Estb ([]string {"sxxx.ba00", _ca00, _cc50})
			_cd00.Send (i.stxx.idxx, i.mtxx.idxx, i.stxx.phnx, (time.Hour * 240000))
			go func (i *Blue) {
				_da00 := Mssg_Estb ([]string {"txxx.bc00"})
				_da00.Send (
					i.stxx.idxx, i.stxx.idxx, i.stxx.phnx,
					(time.Hour * 240000),
				)
			} (i)
		} else if _bb00 [0] == "txxx.ba10" { // Track has started up
			_ca00  :=  ""
			for _,  _cc00 := range  i.tixx {
				_ca00  = _cc00
				_db00 := strings.Index (_ba00.sndr, _cc00)
				if _db00 == 0 {
					i.tdxx [_cc00].ssxx = "s"
					i.tdxx [_cc00].lsxx = "l"
					break
				}
			}
			if _ca00 == "" { continue }
			/*--2--*/
			_cd00 :=                       Mssg_Estb ([]string {"sxxx.ba10", _ca00})
			_cd00.Send (i.stxx.idxx, i.mtxx.idxx, i.stxx.phnx, (time.Hour * 240000))
		} else if _bb00 [0] == "txxx.bb00" { // Track has shutdown due to an error
			_ca00  :=  ""
			for _,  _cc00 := range  i.tixx {
				_ca00  = _cc00
				_db00 := strings.Index (_ba00.sndr, _cc00)
				if _db00 == 0  {
					i.tdxx [_cc00].lsxx = "d"
					break
				}
			}
			if _ca00 == "" { continue }
			/*--2--*/
			_cc50 := _bb00 [1]
			if len (_bb00) < 2 { _cc50 = "" }
			/*--2--*/
			_cd00 :=                Mssg_Estb ([]string {"sxxx.bb00", _ca00, _cc50})
			_cd00.Send (i.stxx.idxx, i.mtxx.idxx, i.stxx.phnx, (time.Hour * 240000))
		} else if _bb00 [0] == "txxx.bb10" { // Track has shutdown gracefully
			_ca00  :=  ""
			for _,  _cc00 := range  i.tixx {
				_ca00  = _cc00
				_db00 := strings.Index (_ba00.sndr, _cc00)
				if _db00 == 0  {
					i.tdxx [_cc00].lsxx = "d"
					break
				}
			}
			if _ca00 == "" { continue }
			/*--2--*/
			_cd00 :=                       Mssg_Estb ([]string {"sxxx.bb10", _ca00})
			_cd00.Send (i.stxx.idxx, i.mtxx.idxx, i.stxx.phnx, (time.Hour * 240000))
		} else if _bb00 [0] == "txxx.bc00" { // Mother/System: Shutdown blue
			if _ba00.sndr != i.mtxx.idxx && _ba00.sndr != i.stxx.idxx {
				continue
			}
			/*--2--*/
			if i.sisx == true { continue }
			i.sisx = true
			/*--2--*/
			for iCA00 := len (i.tixx); iCA00 >= 1; iCA00 -- {
				_cc00 := i.tixx [iCA00 - 1]
				_cd00 := Mssg_Estb ([]string {"sxxx.bc00"})
				/*--3--*/
				_ce00 := false
				for i.tdxx [_cc00].lsxx == "l" {
					_ce00 = _cd00.Send (
						i.stxx.idxx             ,
						i.tdxx [_cc00].idxx     ,
						i.stxx.phnx             ,
					)
					if _ce00 ==true { break }
					time.Sleep (time.Millisecond * 1)
				}
				/*--3--*/
			}
			/*--2--*/
			go func (i *Blue) () {
				for iEA00 := len (i.tixx); iEA00 >= 1; iEA00 -- {
					_fc00 :=  i.tixx  [iEA00  - 1]
					for i.tdxx [_fc00].lsxx  == "l" {
						time.Sleep (time.Millisecond * 1)
					}
				}
				i.scsx = true
			} (i)
		}
	}}
