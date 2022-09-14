package main

type arbre struct {
	Val    map[string]string
	Left   *arbre
	Right  *arbre
	centre *arbre
}

// number of key in the val : ("name", "mob_nb", "mob_type", "boss_type", "feu", "forge", "secret_destination", "secret_unlock")
func (a *arbre) Insert(val map[string]string) {
	if a.Val == nil {
		a.Val = val
		a.Left = &arbre{}
		a.Right = &arbre{}
		a.centre = &arbre{}
	} else {
		if a.Val["name"] > val["name"] {
			a.Left.Insert(val)
		} else if a.Val["name"] < val["name"] {
			a.Right.Insert(val)
		} else {
			a.centre.Insert(val)
		}
	}
}

func main() {

}
