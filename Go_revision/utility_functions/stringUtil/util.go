package stringutil



func Reverse(s string) string {
	reverse := ""
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}


func IsPalindrome(s string)bool{
	for i:=0;i<len(s)/2;i++{
		if s[i]!=s[len(s)-1-i]{
			return false
		}
	}
	return true
}