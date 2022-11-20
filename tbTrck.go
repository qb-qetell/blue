package blum

type Trck struct {
	blum        *Blum // Blum
	idxx       string // Identity
	name       string // Identity
	phnx        *Phnx // Phone
	edxx interface {} // Execution data
	code         func (mixx, idxx, name string, phnx *Phnx, edxx interface {})
	ssxx       string // Start-up status: -, f, s
	snxx       string // Start-up note
}
	func (i *Trck) runx () {
	}
