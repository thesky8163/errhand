package errhand

import (
	"fmt"
	"log"
	"runtime"
)

//var err error

func Errinfo() (errmsg string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("Failed to get caller information")
	}
	funcName := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("Error in function %s at %s:%d: ", funcName, file, line)
}

func Fprintf(format string, a ...interface{}) {
	err, ok := a[len(a)-1].(error)
	if ok {
		if err != nil {
			fmt.Printf(Errinfo()+format, a...)
		}
	} else {
		fmt.Println("errhand err: The last parameter type for error handling must be errors")
	}
}

func FprintfR(format string, a ...interface{}) {
	err, ok := a[len(a)-1].(error)
	if ok {
		if err != nil {
			fmt.Printf(Errinfo()+format, a...)
			return
		}
	} else {
		fmt.Println("errhand err: The last parameter type for error handling must be errors")
	}
}

func ReturnErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func Fprintfunc(format string, add func(), a ...interface{}) {
	err, ok := a[len(a)-1].(error)
	if ok {
		if err != nil {
			add()
			fmt.Printf(Errinfo()+format, a...)
		}
	} else {
		fmt.Println("errhand err: The last parameter type for error handling must be errors")
	}
}

func FprintfuncR(format string, add func(), a ...interface{}) {
	err, ok := a[len(a)-1].(error)
	if ok {
		if err != nil {
			add()
			fmt.Printf(Errinfo()+format, a...)
			return
		}
	} else {
		fmt.Println("errhand err: The last parameter type for error handling must be errors")
	}
}

func Fprintln(a ...interface{}) {
	err, ok := a[len(a)-1].(error)
	if ok {
		if err != nil {
			//fmt.Println("insert failed, err:", err)
			fmt.Print(Errinfo())
			fmt.Println(a...)
		}
	} else {
		fmt.Println("errhand err: The last parameter type for error handling must be errors")
	}
}

func FprintlnR(a ...interface{}) {
	err, ok := a[len(a)-1].(error)
	if ok {
		if err != nil {
			//fmt.Println("insert failed, err:", err)
			fmt.Print(Errinfo())
			fmt.Println(a...)
			return
		}
	} else {
		fmt.Println("errhand err: The last parameter type for error handling must be errors")
	}
}

func Lprintf(format string, a ...interface{}) {
	err, ok := a[len(a)-1].(error)
	if ok {
		if err != nil {
			log.Printf(Errinfo()+format, a...)
		}
	} else {
		fmt.Println("errhand err: The last parameter type for error handling must be errors")
	}
}

func LprintfR(format string, a ...interface{}) {
	err, ok := a[len(a)-1].(error)
	if ok {
		if err != nil {
			log.Printf(Errinfo()+format, a...)
			return
		}
	} else {
		fmt.Println("errhand err: The last parameter type for error handling must be errors")
	}
}

func Lprintfunc(format string, add func(), a ...interface{}) {
	err, ok := a[len(a)-1].(error)
	if ok {
		if err != nil {
			add()
			log.Printf(Errinfo()+format, a...)
		}
	} else {
		fmt.Println("errhand err: The last parameter type for error handling must be errors")
	}
}

func LprintfuncR(format string, add func(), a ...interface{}) {
	err, ok := a[len(a)-1].(error)
	if ok {
		if err != nil {
			add()
			log.Printf(Errinfo()+format, a...)
			return
		}
	} else {
		fmt.Println("errhand err: The last parameter type for error handling must be errors")
	}
}

func Lprintln(a ...interface{}) {
	err, ok := a[len(a)-1].(error)
	if ok {
		if err != nil {
			log.Println(Errinfo(), a)
		}
	} else {
		fmt.Println("errhand err: The last parameter type for error handling must be errors")
	}
}

func LprintlnR(a ...interface{}) {
	err, ok := a[len(a)-1].(error)
	if ok {
		if err != nil {
			log.Println(Errinfo(), a)
			return
		}
	} else {
		fmt.Println("errhand err: The last parameter type for error handling must be errors")
	}
}


