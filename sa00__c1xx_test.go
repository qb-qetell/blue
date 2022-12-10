package blue

import "fmt"
import "time"

func c1xx (mixx, idxx, name string, phnx *Phnx, edxx interface {})     (error) {
	fmt.Println ("c1xx:", "mixx:!", mixx)
	fmt.Println ("c1xx:", "idxx:!", idxx)
	fmt.Println ("c1xx:", "name:!", name)
	fmt.Println ("c1xx:", "edxx:!", edxx)
	/*--1--*/
	_cb00 := phnx.NtfySystAbtxScsfStrt ()
	if _cb00 != nil {
		_ca00 := fmt.Sprintf (
			"c1xx: %s: Could not send message. [%s]", idxx, _cb00.Error (),
		)
		fmt.Println ( _ca00  )
	}
	/*--1--*/
	for i := 1; i <= 10; i ++ {
		_ca00 :=  Mssg_Estb ([]string {"sxxx.bc00"})
		_cb00 := _ca00.Send (phnx.trck.idxx, "1234.i4xxy", phnx, time.Hour * 1)
		if _cb00 != nil {
			fmt.Println (_cb00)
		}
	}
	/*--1--*/
	return nil
}
