package testu

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"
)

// assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func OK(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

type Comparator func(exp, act interface{}) bool

func EqualsSlice[T any](tb testing.TB, exp, act []T) {
	for exp_element := range exp {
		found := false
		for act_element := range act {
			if reflect.DeepEqual(exp_element, act_element) {
				found = true
				break
			}
		}
		if !found {
			_, file, line, _ := runtime.Caller(1)
			fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
			tb.FailNow()
		}
	}
}

type AlwaysValidValidator struct{}

func (validator AlwaysValidValidator) Validate(interface{}) ([]string, error) {
	return nil, nil
}

type FakePresenter struct {
	Data    interface{}
	DataErr error
}

func (presenter *FakePresenter) Present(data interface{}) {
	presenter.Data = data
}

func (presenter *FakePresenter) PresentError(data error) {
	presenter.DataErr = data
}

func TimeParseOrPanic(value string) time.Time {
	time, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}
	return time
}

type FakeClock struct {
	NowReturn time.Time
}

func (c FakeClock) Now() time.Time {
	return c.NowReturn
}
