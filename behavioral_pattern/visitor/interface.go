package visitor

type JK string

const (
	Laki      JK = "male"
	Perempuan JK = "female"
)

type pelanggan interface {
	Setuju(m Montir)
}

type Montir interface {
	DoForGantiOli(p pelanggan)
	DoForServisRingan(p pelanggan)
	GetJenisKelamin() JK
}
