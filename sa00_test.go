package blue

import "fmt"
import "testing"
import "time"

func TestBa00_ (t *testing.T) {
	_ak00, _ba00 :=  Blue_Estb ("1234", 0, 0)
	if _ak00 != nil {
		_ca00 := fmt.Sprintf ("Could not create blue. [%s]", _ak00.Error ())
		fmt.Println (_ca00)
		return
	}
	/*--1--*/
	_bb10 := _ba00.Elbr ("i1xx", "Instance 1", 0, "ED 1", c1xx)
	if _bb10 != nil {
		_ca00 := fmt.Sprintf ("Could not addd track. [%s]", _bb10.Error ())
		fmt.Println (_ca00)
		return
	}
	_bb20 := _ba00.Elbr ("i2xx", "Instance 2", 0, "ED 2", c1xx)
	if _bb20 != nil {
		_ca00 := fmt.Sprintf ("Could not addd track. [%s]", _bb20.Error ())
		fmt.Println (_ca00)
		return
	}
	_bb30 := _ba00.Elbr ("i3xx", "Instance 3", 0, "ED 3", c1xx)
	if _bb30 != nil {
		_ca00 := fmt.Sprintf ("Could not addd track. [%s]", _bb30.Error ())
		fmt.Println (_ca00)
		return
	}
	_bb40 := _ba00.Elbr ("i4xx", "Instance 4", 0, "ED 4", c4xx)
	if _bb40 != nil {
		_ca00 := fmt.Sprintf ("Could not addd track. [%s]", _bb40.Error ())
		fmt.Println (_ca00)
		return
	}
	/*--1--*/
	_bc10, _bc20 := _ba00.Strt ()
	if _bc10 != nil {
		_ca00 := fmt.Sprintf ("Could not start blue. [%s]", _bc10.Error ())
		fmt.Println (_ca00)
		return
	}
	/*--1--*/
	go func () {
		time.Sleep (time.Second * 2)
		_ba00.Halt  ()
	} ()
	for {
		_ca00 := _bc20.Ftch (time.Second * 2)
		if _ca00 == nil {
			//_ba00.Halt ()
			continue
		}
		fmt.Println ("mthr:", _ca00.FSXX (), _ca00.FRXX (), _ca00.FCXX (), _ca00.FTXX ())
	}
}
