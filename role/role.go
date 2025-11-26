package role

type Good int

type Personer interface {
	New()
	Story()
	Talk(string)
	Tellcg()
	Getgood() Good
	Getcg() bool
}
type Now_everyone struct {
	One_good map[string]Good
	One_cg   map[string]bool
}
