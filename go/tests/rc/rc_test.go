package main

import (
	"sort"
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func TestFlags(t *testing.T) {
	argsAndSolution := map[string]string{
		"testdata/testingSimpleFunc.go": `Cheating:
	TYPE:             	NAME:             	LOCATION:
	illegal-import    	regexp            	testdata/testingSimpleFunc.go:4:2
	illegal-call      	len               	testdata/testingSimpleFunc.go:10:9
	illegal-access    	regexp.MustCompile	testdata/testingSimpleFunc.go:8:8
	illegal-definition	SimpleFunc        	testdata/testingSimpleFunc.go:7:1
`,

		"-no-for -no-lit=[a-z] testdata/printalphabet/printalphabet.go": `Cheating:
	TYPE:             	NAME:                	LOCATION:
	illegal-import    	fmt                  	testdata/printalphabet/printalphabet.go:4:2
	illegal-import    	github.com/01-edu/z01	testdata/printalphabet/printalphabet.go:6:2
	illegal-call      	append               	testdata/printalphabet/printalphabet.go:11:7
	illegal-definition	fillArray            	testdata/printalphabet/printalphabet.go:9:1
	illegal-call      	int                  	testdata/printalphabet/printalphabet.go:17:7
	illegal-access    	z01.PrintRune        	testdata/printalphabet/printalphabet.go:19:3
	illegal-access    	z01.PrintRune        	testdata/printalphabet/printalphabet.go:21:2
	illegal-definition	main                 	testdata/printalphabet/printalphabet.go:15:1
	illegal-access    	fmt.Println          	testdata/printalphabet/printalphabet.go:26:3
	illegal-definition	defFun               	testdata/printalphabet/printalphabet.go:25:2
	illegal-call      	defFun               	testdata/printalphabet/printalphabet.go:28:2
	illegal-definition	testingScope         	testdata/printalphabet/printalphabet.go:24:1
	illegal-loop      	for                  	testdata/printalphabet/printalphabet.go:10:2
	illegal-lit       	'a'                  	testdata/printalphabet/printalphabet.go:10:11
	illegal-lit       	'z'                  	testdata/printalphabet/printalphabet.go:10:21
	illegal-lit       	'a'                  	testdata/printalphabet/printalphabet.go:16:14
	illegal-lit       	'b'                  	testdata/printalphabet/printalphabet.go:16:19
	illegal-lit       	'c'                  	testdata/printalphabet/printalphabet.go:16:24
	illegal-lit       	'd'                  	testdata/printalphabet/printalphabet.go:16:29
	illegal-lit       	'e'                  	testdata/printalphabet/printalphabet.go:16:34
	illegal-lit       	'f'                  	testdata/printalphabet/printalphabet.go:16:39
	illegal-lit       	'a'                  	testdata/printalphabet/printalphabet.go:17:11
	illegal-lit       	'\n'                 	testdata/printalphabet/printalphabet.go:21:16
	illegal-lit       	"Hello"              	testdata/printalphabet/printalphabet.go:28:9
`,
		"-cast testdata/eightqueens.go": `Cheating:
	TYPE:             	NAME:                	LOCATION:
	illegal-import    	github.com/01-edu/z01	testdata/eightqueens.go:4:2
	illegal-access    	z01.PrintRune        	testdata/eightqueens.go:49:5
	illegal-access    	z01.PrintRune        	testdata/eightqueens.go:55:2
	illegal-definition	printQueens          	testdata/eightqueens.go:42:1
`,
		"-no-array testdata/printalphabet/printalphabet.go": `Cheating:
	TYPE:             	NAME:                	LOCATION:
	illegal-import    	fmt                  	testdata/printalphabet/printalphabet.go:4:2
	illegal-import    	github.com/01-edu/z01	testdata/printalphabet/printalphabet.go:6:2
	illegal-call      	append               	testdata/printalphabet/printalphabet.go:11:7
	illegal-definition	fillArray            	testdata/printalphabet/printalphabet.go:9:1
	illegal-call      	int                  	testdata/printalphabet/printalphabet.go:17:7
	illegal-access    	z01.PrintRune        	testdata/printalphabet/printalphabet.go:19:3
	illegal-access    	z01.PrintRune        	testdata/printalphabet/printalphabet.go:21:2
	illegal-definition	main                 	testdata/printalphabet/printalphabet.go:15:1
	illegal-access    	fmt.Println          	testdata/printalphabet/printalphabet.go:26:3
	illegal-definition	defFun               	testdata/printalphabet/printalphabet.go:25:2
	illegal-call      	defFun               	testdata/printalphabet/printalphabet.go:28:2
	illegal-definition	testingScope         	testdata/printalphabet/printalphabet.go:24:1
	illegal-slice     	rune                 	testdata/printalphabet/printalphabet.go:9:18
	illegal-slice     	rune                 	testdata/printalphabet/printalphabet.go:16:7
`,
		"-no-slices testdata/printalphabet/printalphabet.go": `Cheating:
	TYPE:             	NAME:                	LOCATION:
	illegal-import    	fmt                  	testdata/printalphabet/printalphabet.go:4:2
	illegal-import    	github.com/01-edu/z01	testdata/printalphabet/printalphabet.go:6:2
	illegal-call      	append               	testdata/printalphabet/printalphabet.go:11:7
	illegal-definition	fillArray            	testdata/printalphabet/printalphabet.go:9:1
	illegal-call      	int                  	testdata/printalphabet/printalphabet.go:17:7
	illegal-access    	z01.PrintRune        	testdata/printalphabet/printalphabet.go:19:3
	illegal-access    	z01.PrintRune        	testdata/printalphabet/printalphabet.go:21:2
	illegal-definition	main                 	testdata/printalphabet/printalphabet.go:15:1
	illegal-access    	fmt.Println          	testdata/printalphabet/printalphabet.go:26:3
	illegal-definition	defFun               	testdata/printalphabet/printalphabet.go:25:2
	illegal-call      	defFun               	testdata/printalphabet/printalphabet.go:28:2
	illegal-definition	testingScope         	testdata/printalphabet/printalphabet.go:24:1
	illegal-slice     	rune                 	testdata/printalphabet/printalphabet.go:9:18
	illegal-slice     	rune                 	testdata/printalphabet/printalphabet.go:16:7
`,
		"-no-these-slices=int,rune testdata/printalphabet/printalphabet.go": `Cheating:
	TYPE:             	NAME:                	LOCATION:
	illegal-import    	fmt                  	testdata/printalphabet/printalphabet.go:4:2
	illegal-import    	github.com/01-edu/z01	testdata/printalphabet/printalphabet.go:6:2
	illegal-call      	append               	testdata/printalphabet/printalphabet.go:11:7
	illegal-definition	fillArray            	testdata/printalphabet/printalphabet.go:9:1
	illegal-call      	int                  	testdata/printalphabet/printalphabet.go:17:7
	illegal-access    	z01.PrintRune        	testdata/printalphabet/printalphabet.go:19:3
	illegal-access    	z01.PrintRune        	testdata/printalphabet/printalphabet.go:21:2
	illegal-definition	main                 	testdata/printalphabet/printalphabet.go:15:1
	illegal-access    	fmt.Println          	testdata/printalphabet/printalphabet.go:26:3
	illegal-definition	defFun               	testdata/printalphabet/printalphabet.go:25:2
	illegal-call      	defFun               	testdata/printalphabet/printalphabet.go:28:2
	illegal-definition	testingScope         	testdata/printalphabet/printalphabet.go:24:1
	illegal-slice     	rune                 	testdata/printalphabet/printalphabet.go:9:18
	illegal-slice     	rune                 	testdata/printalphabet/printalphabet.go:16:7
`,
		`-allow-builtin testdata/doopprog/main.go fmt.Println strconv.Atoi os.Args`: ``,
		`-cast testdata/doopprog/main.go fmt.Println strconv.Atoi os.Args len`:      ``,
		`testdata/testingWrapping.go`: `Cheating:
	TYPE:             	NAME:           	LOCATION:
	illegal-call      	len             	testdata/utilDepth2/wrapper.go:4:9
	illegal-definition	LenWrapper      	testdata/utilDepth2/wrapper.go:3:1
	illegal-access    	util2.LenWrapper	testdata/util/util.go:10:9
	illegal-definition	LenWrapperU     	testdata/util/util.go:9:1
	illegal-access    	util.LenWrapperU	testdata/testingWrapping.go:8:9
	illegal-definition	Length          	testdata/testingWrapping.go:7:1
`,
		`testdata/testingWrapping.go len`: ``,
		`testdata/empty/empty len`: `	stat : no such file or directory
`,
		`testdata/empty/empty.go testdata/empty/empty`: `	testdata/empty/empty.go:1:1: expected ';', found 'EOF' (and 2 more errors)
`,
		`testdata/abc/main.go --cast github.com/01-edu/z01.PrintRune#2 --no-lit=[b-mB-Y]`: ``,
		`itoa.go len --cast`: `	stat itoa.go: no such file or directory
`,
	}
	Compare(t, argsAndSolution)
}

func Compare(t *testing.T, argsAndSol map[string]string) {
	for args, expected := range argsAndSol {
		a := strings.Split(args, " ")
		_, err := z01.MainOut("../rc", a...)
		if err != nil && !EqualResult(expected, err.Error()) {
			t.Errorf("./rc %s prints %q\n instead of %q\n", args, err.Error(), expected)
		}
	}
}

func EqualResult(sol, out string) bool {
	// split
	solSli := strings.Split(sol, "\n")
	outSli := strings.Split(out, "\n")
	// sort
	sort.Sort(sort.StringSlice(solSli))
	sort.Sort(sort.StringSlice(outSli))
	// join
	sol = strings.Join(solSli, " ")
	out = strings.Join(outSli, " ")
	// compare
	return sol == out
}
