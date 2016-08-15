package phpfunc

import "testing"

func TestStrReplace(t *testing.T) {
	from := "1"
	to := "2"
	str := "111111"
	newstr := StrReplace(from, to, str, -1)
	t.Error("newstr", newstr)

	from2 := "1"
	to2 := []string{"2", "3"}
	str2 := "111111"
	newstr = StrReplace(from2, to2, str2, -1)
	t.Error("newstr", newstr)

	from3 := []string{"2", "3"}
	to3 := "4"
	str3 := "1111112222233333"
	newstr = StrReplace(from3, to3, str3, -1)
	t.Error("newstr", newstr)

	from4 := []string{"1", "2", "3"}
	to4 := []string{"4", "5", "6"}
	str4 := "111111222222233333333"
	newstr = StrReplace(from4, to4, str4, -1)
	t.Error("newstr", newstr)
}
