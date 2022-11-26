package blue

import "fmt"
import "time"

func c4xx (mixx, idxx, name string, phnx *Phnx, edxx interface {}) (error)     {
	fmt.Println ("c4xx:", "mixx:!", mixx)
	fmt.Println ("c4xx:", "idxx:!", idxx)
	fmt.Println ("c4xx:", "name:!", name)
	fmt.Println ("c4xx:", "edxx:!", edxx)
	/*--1--*/
	_cb00 := phnx.NtfySystAbtxScsfStrt ()
	if _cb00 == false {
		_ca00 := fmt.Sprintf ("c1xx: %s: Could not send message.", idxx)
		fmt.Println ( _ca00 )
	}
	/*--1--*/
	for {
		_ca00 := phnx.Ftch (time.Second * 2)
		if _ca00 == nil {
			//_ba00.Halt ()
			continue
		}
		fmt.Println ("c4xx:", _ca00.FSXX (), _ca00.FRXX (), _ca00.FCXX (), _ca00.FTXX ())
	}
	/*--1--*/
	return nil
}
