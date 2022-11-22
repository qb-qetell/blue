package blue

type Trck struct {
	blum        *Blue // Blue
	idxx       string // Identity
	name       string // Identity
	phnx        *Phnx // Phone
	edxx interface {} // Execution data
	code         func (mixx, idxx, name string, phnx *Phnx, edxx interface {})
	ssxx       string // Start-up status: -, f, s
	snxx       string // Start-up note
}
	func (i *Trck) runx () {
		go i.code (i.blum.idxx, i.idxx, i.name, i.phnx, i.edxx)
	}
