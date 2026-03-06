package oppositesattract

func Lovefunc(flower1, flower2 int) bool {
	return (flower1+flower2)%2 == 1
}