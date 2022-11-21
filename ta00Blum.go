package blue

import "errors"
import "regexp"
import "strings"
import "time"

type Blue struct {
	idxx           string // ID
	mpxx            *Phnx // Mother's phone
	spxx            *Phnx // System's phone
	tixx         []string // Trck index
	tdxx map[string]*Trck // Trck details
	lsxx             bool // Life status
	usxx             bool // Use status
}
	func Blue_Estb (idxx string, mpcx, spcx uint16) (error, *Blue) {
		_ba00 := strings.ToLower (idxx)
		if regexp.MustCompile (`^[a-z0-9]+(\.[a-z0-9]+)*$`).MatchString(_ba00) == false {
			_ca00 := errors.New (
				"String can not be used as ID. [It contains one or more " +
				"illegal characters.]",
			)
			return _ca00, nil
		}
		_bb00 := &Blue {
			idxx:                   _ba00,
			mpxx:        Phnx_Estb (nil, nil, mpcx),
			spxx:        Phnx_Estb (nil, nil, spcx),
			tixx:             []string {},
			tdxx: make (map[string]*Trck),
			lsxx:                   false,
			usxx:                   false,
		}
		_bb00.mpxx.blue = _bb00
		_bb00.spxx.blue = _bb00
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
		for _, _bc00 := range i.tixx {
			if idxx == _bc00 {
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
			blum:                   i,
			idxx:                idxx,
			name:                name,
			phnx: Phnx_Estb (i, nil, pcxx),
			edxx:                edxx,
			code:                code,
			ssxx:                 "-",
		}
		_bd00.phnx.trck = _bd00
		/*--1--*/
		i.tixx = append (i.tixx, idxx)
		i.tdxx [idxx]   = _bd00
		/*--1--*/
		return nil
	}
	
	func (i *Blue) Runx () (error, *Phnx) {
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
		return nil, i.spxx
	}
	
	func (i *Blue) Halt () {
		if i.lsxx == false { return }
	}
	
	func tmxx (i *Blue) {}
	