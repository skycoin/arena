package pkg

type Addend struct{
	AddnA int
	AddnB int
	AddnS int
}

func (a *Addend)Adds() {
	a.AddnS = a.AddnA + a.AddnB
}

