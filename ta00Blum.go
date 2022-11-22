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
	lsxx             bool // Life status
	usxx             bool // Use status
}
	func Blue_Estb (idxx string, mpcx, spcx uint16) (error, *Blue) {
		_ba00 := strings.ToLower (idxx)
		if regexp.MustCompile (`^[a-z0-9]+(\.[a-z0-9]+)*$`).MatchString (_ba00) == false{
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
			lsxx:                        false,
			usxx:                        false,
		}
		_bc00 := &Trck {
			blum:                        _bb00,
			idxx:                        _ba00,
			name:               "Mother Track",
			phnx: phnx_Estb (_bb00, nil, mpcx),
			edxx:                          nil,
			code:                          nil,
			ssxx:                          "-",
		}
		_bb00.mtxx           = _bc00
		_bb00.mtxx.phnx.trck = _bc00
		_bd00 := &Trck {
			blum:                        _bb00,
			idxx:                 _ba00 + ".!",
			name:               "System Track",
			phnx: phnx_Estb (_bb00, nil, spcx),
			edxx:                          nil,
			code:                          nil,
			ssxx:                          "-",
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
		code         func (mixx, idxx, name string, phnx *Phnx, edxx interface {}),
	 ) (error) {
		_ba00 := strings.ToLower (idxx)
		if regexp.MustCompile (`^[a-z0-9]+(\.[a-z0-9]+)*$`).MatchString(_ba00) == false {
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
			blum:                        i,
			idxx:                    _bb00,
			name:                     name,
			phnx: phnx_Estb (i, nil, pcxx),
			edxx:                     edxx,
			code:                     code,
			ssxx:                      "-",
		}
		_bd00.phnx.trck = _bd00
		/*--1--*/
		i.tixx = append (i.tixx, _bb00)
		i.tdxx [idxx]   = _bd00
		/*--1--*/
		return nil
	}
	
	func (i *Blue) Strt () (error, *Phnx) {
		if i.usxx == true {
			_ca00 := errors.New ("Blum has already been run.")
			return _ca00, nil
		}
		i.usxx = true
		/*--1--*/
		go func (i *Blue) {
			go tmxx (i)
			for _, _da00 := range i.tixx {
				go i.tdxx [_da00].code (
					i.idxx,
					i.tdxx [_da00].idxx,
					i.tdxx [_da00].name,
					i.tdxx [_da00].phnx,
					i.tdxx [_da00].edxx,
				)
				for {
					if        i.tdxx [_da00].ssxx == "-" {
						time.Sleep (time.Millisecond * 1)
						continue
					} else if i.tdxx [_da00].ssxx == "f" {
						return
					} else if i.tdxx [_da00].ssxx == "s" {
						break
					} else {
						return
					}
				}
			}
		} (i)
		/*--1--*/
		return nil, i.stxx.phnx
	}
	
	func (i *Blue) Halt () {
		_ba00 := Mssg_Estb ([]string {"ca00"})
		_ba00.Send (i.mtxx.idxx, i.stxx.idxx, i.stxx.phnx, (time.Hour * 24))
		return
	}
	
	func tmxx (i *Blue) {}
	