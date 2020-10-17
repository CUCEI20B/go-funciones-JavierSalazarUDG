package main

func Burbuja(s []int64) {
	var temp int64
	for j := 0; j < len(s); j++ {
		for i := 0; i < len(s)-1; i++ {
			if i < len(s) && s[i] > s[i+1] {
				temp = s[i]
				s[i] = s[i+1]
				s[i+1] = temp
			}
		}
	}
}

func main() {

}
