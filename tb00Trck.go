package blue

import "fmt"
import "time"

type Trck struct {
	blue        *Blue // Blue
	idxx       string // Identity
	name       string // Identity
	phnx        *Phnx // Phone
	edxx interface {} // Execution data
	code         func (mixx, idxx, name string, phnx *Phnx, edxx interface {}) (error)
	ssxx       string // Start-up status: -, f, s
	snxx       string // Start-up note
	lsxx       string // Life status: -, l, d
}
	func (i *Trck) runx () {
		go func (i *Trck) {
			defer func () {
				_ba00 := recover ()
				if _ba00 != nil {
					_ca00 := fmt.Sprintf ("Track paniced. [%v]",   _ba00 )
					_cd00 := Mssg_Estb ([]string {"txxx.bb00",     _ca00})
			                _cd00.Send (
			                	i.idxx, i.blue.stxx.idxx, i.phnx,
			                	(time.Hour * 240000),
			                )
				}
			} ()
			/*--1--*/
			_ba00 := i.code (i.blue.idxx, i.idxx, i.name, i.phnx, i.edxx)
			/*--1--*/
			if _ba00 != nil {
				_ca00 :=    Mssg_Estb ([]string {"txxx.bb00", _ba00.Error ()})
				_ca00.Send (
					i.idxx, i.blue.stxx.idxx, i.phnx, (time.Hour * 240000),
				)
				return
			}
			/*--1--*/
			_bb00 :=                            Mssg_Estb ([]string {"txxx.bb10"})
			_bb00.Send    (i.idxx, i.blue.stxx.idxx, i.phnx, (time.Hour * 240000))
			return
		} (i)
	}
