package testutil

import (
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/fatih/color"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)

// assert fails the test now if exp is not equal to act.
func AssertDeepEqual(t *testing.T, exp, act interface{}, msg interface{}) {
	deepEqual(t, exp, act, true, msg)
}

func RequireDeepEqual(t *testing.T, exp, act interface{}, msg interface{}) {
	deepEqual(t, exp, act, false, msg)
}

// equals fails the test if exp is not equal to act.
func DeepEqual(t *testing.T, exp, act interface{}) {
	deepEqual(t, exp, act, false, "")
}

func deepEqual(t *testing.T, exp, act interface{}, assert bool, msg interface{}) {

	expt := reflect.TypeOf(exp)
	if expt.Kind() == reflect.Ptr {
		expv := reflect.ValueOf(exp)
		exp = expv.Elem().Interface()
	}

	actt := reflect.TypeOf(act)
	if actt.Kind() == reflect.Ptr {
		actv := reflect.ValueOf(act)
		act = actv.Elem().Interface()
	}

	if !reflect.DeepEqual(exp, act) {

		_, file, line, _ := runtime.Caller(2)

		red.Printf("\n\n==============================================================\n")
		if msg != "" {
			red.Printf("%s:%d: FAIL: Not Equals (%v/%v) [%v]\n",
				filepath.Base(file), line, reflect.TypeOf(exp), reflect.TypeOf(act), msg)
		} else {
			red.Printf("%s:%d: FAIL: Not Equals (%v/%v) \n",
				filepath.Base(file), line, reflect.TypeOf(exp), reflect.TypeOf(act))
		}
		red.Printf("==============================================================\n\n")

		//dump diff data
		dmp := diffmatchpatch.New()
		expText := fmt.Sprintf("%s", exp)

		actText := fmt.Sprintf("%s", act)
		diffs := dmp.DiffMain(expText, actText, false)

		var top bytes.Buffer
		var bottom bytes.Buffer

		for _, diff := range diffs {
			switch diff.Type {
			case diffmatchpatch.DiffDelete:
				if top.Len() > 0 {
					fmt.Printf(top.String())
					top.Reset()
				}
				if bottom.Len() > 0 {
					fmt.Printf("\n\n...\n\n")
					fmt.Printf(bottom.String())
					bottom.Reset()
				}

				red.Printf("%s", diff.Text)
			case diffmatchpatch.DiffInsert:
				if top.Len() > 0 {
					fmt.Printf(top.String())
					top.Reset()
				}
				if bottom.Len() > 0 {
					fmt.Printf("\n\n...\n\n")
					fmt.Printf(bottom.String())
					bottom.Reset()
				}

				green.Printf("%s", diff.Text)
			case diffmatchpatch.DiffEqual:
				grab := 10
				lines := strings.Split(diff.Text, "\n")

				if len(lines) <= 2*grab {
					top.WriteString(diff.Text)
				} else {
					top.WriteString(strings.Join(lines[:grab], "\n"))
					bottom.WriteString(strings.Join(lines[len(lines)-grab:], "\n"))
				}
			default:
				fmt.Printf("Unknown diff type: %v", diff.Type)
				break
			}
		}

		if top.Len() > 0 {
			fmt.Println(top.String())
		}

		if assert {
			t.FailNow()
		} else {
			t.Fail()
		}
	}
}
