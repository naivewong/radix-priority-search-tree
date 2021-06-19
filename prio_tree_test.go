package priotree

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// Ok fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// NotOk fails the test if an err is nil.
func NotOk(tb testing.TB, err error) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: expected error, got nothing \033[39m\n\n", filepath.Base(file), line)
		tb.FailNow()
	}
}

// Equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}, msgAndArgs ...interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:%s\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, formatMessage(msgAndArgs), exp, act)
		tb.FailNow()
	}
}

// NotEquals fails the test if exp is equal to act.
func NotEquals(tb testing.TB, exp, act interface{}) {
	if reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: Expected different exp and got\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func formatMessage(msgAndArgs []interface{}) string {
	if len(msgAndArgs) == 0 {
		return ""
	}

	if msg, ok := msgAndArgs[0].(string); ok {
		return fmt.Sprintf("\n\nmsg: "+msg, msgAndArgs[1:]...)
	}
	return ""
}

// Assert fails the test if the condition is false.
func Assert2(condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		panic("")
	}
}

// Ok fails the test if an err is not nil.
func Ok2(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		panic("")
	}
}

// NotOk fails the test if an err is nil.
func NotOk2(err error) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: expected error, got nothing \033[39m\n\n", filepath.Base(file), line)
		panic("")
	}
}

// Equals fails the test if exp is not equal to act.
func Equals2(exp, act interface{}, msgAndArgs ...interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:%s\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, formatMessage(msgAndArgs), exp, act)
		panic("")
	}
}

// NotEquals fails the test if exp is equal to act.
func NotEquals2(exp, act interface{}) {
	if reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: Expected different exp and got\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		panic("")
	}
}

func TestPrioTree1(t *testing.T) {
	tree := NewPriorityTree(3)
	tree.Insert(2, 4)
	tree.Insert(0, 5)
	tree.Insert(3, 6)
	tree.Insert(4, 5)
	tree.Insert(0, 7)

	{
		n := tree.FirstOverlap(1, 5)
		NotEquals(t, (*PriorityTreeNode)(nil), n)
		Equals(t, 0, n.start)
		Equals(t, 7, n.last)

		n = tree.NextOverlap(1, 5, n)
		NotEquals(t, (*PriorityTreeNode)(nil), n)
		Equals(t, 3, n.start)
		Equals(t, 6, n.last)

		n = tree.NextOverlap(1, 5, n)
		NotEquals(t, (*PriorityTreeNode)(nil), n)
		Equals(t, 0, n.start)
		Equals(t, 5, n.last)

		n = tree.NextOverlap(1, 5, n)
		NotEquals(t, (*PriorityTreeNode)(nil), n)
		Equals(t, 2, n.start)
		Equals(t, 4, n.last)

		n = tree.NextOverlap(1, 5, n)
		NotEquals(t, (*PriorityTreeNode)(nil), n)
		Equals(t, 4, n.start)
		Equals(t, 5, n.last)

		n = tree.NextOverlap(1, 5, n)
		Equals(t, (*PriorityTreeNode)(nil), n)
	}

	Equals(t, false, tree.Delete(0, 3))

	{
		Equals(t, true, tree.Delete(3, 6))
		n := tree.FirstOverlap(1, 5)
		NotEquals(t, (*PriorityTreeNode)(nil), n)
		Equals(t, 0, n.start)
		Equals(t, 7, n.last)

		n = tree.NextOverlap(1, 5, n)
		NotEquals(t, (*PriorityTreeNode)(nil), n)
		Equals(t, 0, n.start)
		Equals(t, 5, n.last)

		n = tree.NextOverlap(1, 5, n)
		NotEquals(t, (*PriorityTreeNode)(nil), n)
		Equals(t, 2, n.start)
		Equals(t, 4, n.last)

		n = tree.NextOverlap(1, 5, n)
		NotEquals(t, (*PriorityTreeNode)(nil), n)
		Equals(t, 4, n.start)
		Equals(t, 5, n.last)

		n = tree.NextOverlap(1, 5, n)
		Equals(t, (*PriorityTreeNode)(nil), n)
	}

	{
		Equals(t, true, tree.Delete(0, 7))
		n := tree.FirstOverlap(1, 5)
		NotEquals(t, (*PriorityTreeNode)(nil), n)
		Equals(t, 0, n.start)
		Equals(t, 5, n.last)

		n = tree.NextOverlap(1, 5, n)
		NotEquals(t, (*PriorityTreeNode)(nil), n)
		Equals(t, 2, n.start)
		Equals(t, 4, n.last)

		n = tree.NextOverlap(1, 5, n)
		NotEquals(t, (*PriorityTreeNode)(nil), n)
		Equals(t, 4, n.start)
		Equals(t, 5, n.last)

		n = tree.NextOverlap(1, 5, n)
		Equals(t, (*PriorityTreeNode)(nil), n)
	}
}