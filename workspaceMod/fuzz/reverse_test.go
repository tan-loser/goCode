package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

// 单元测试
func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev, revErr := Reverse(tc.in)
		if revErr == nil && rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
	fmt.Println("demo test: success!")
}

// 模糊测试：模糊测试的一个好处是它可以为你生成输入，并可能识别出你编写测试用例时没有触及的边缘情况。
// go test -run=FuzzReverse ，如果你那文件里有其他测试，并且你只想运行 fuzz 测试。
// 运行 FuzzReverse 并启用 fuzzing，以查看任何随机生成的字符串输入是否会引发失败。
// 这通过使用 go test 并设置新标志 -fuzz 为参数 Fuzz 来执行。
// 另一个有用的标志是 -fuzztime ，它可以限制 fuzzing 所花费的时间。
// 例如，在下面的测试中指定 -fuzztime 10s 意味着，只要之前没有发生失败，测试将在 10 秒后默认退出。
// go test -fuzz=Fuzz -fuzztime 30s
func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345", "dd"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	// f.Fuzz(func(t *testing.T, orig string) {
	// 	rev := Reverse(orig)
	// 	doubleRev := Reverse(rev)
	// 	if orig != doubleRev {
	// 		t.Errorf("Before: %q, after: %q", orig, doubleRev)
	// 	}
	// 	if utf8.ValidString(orig) && !utf8.ValidString(rev) {
	// 		t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
	// 	}
	// 	fmt.Println("fuzz test:" + orig)
	// })

	// f.Fuzz(func(t *testing.T, orig string) {
	// 	rev := Reverse(orig)
	// 	doubleRev := Reverse(rev)
	// 	t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
	// 	if orig != doubleRev {
	// 		t.Errorf("Before: %q, after: %q", orig, doubleRev)
	// 	}
	// 	if utf8.ValidString(orig) && !utf8.ValidString(rev) {
	// 		t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
	// 	}
	// })

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
