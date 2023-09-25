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

func Fprintf(msg string, err error) {
	if err != nil {
		//fmt.Printf("insert failed, err:%v\n", err)
		fmt.Printf(Errinfo()+msg, err)
	}
}

func Fprintfr(msg string, err error) {
	if err != nil {
		//fmt.Printf("insert failed, err:%v\n", err)
		fmt.Printf(Errinfo()+msg, err)
		return
	}

}
func ReturnErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func Fprintfunc(msg string, err error, add func()) {
	if err != nil {
		add()
		fmt.Printf(Errinfo()+msg, err)
	}
}

func Fprintfuncr(msg string, err error, add func()) {
	if err != nil {
		add()
		fmt.Printf(Errinfo()+msg, err)
		return
	}
}

func Fprintln(msg string, err error) {
	if err != nil {
		//fmt.Println("insert failed, err:", err)
		fmt.Println(Errinfo(), msg, err)
	}
}

func Fprintlnr(msg string, err error) {
	if err != nil {
		//fmt.Println("insert failed, err:", err)
		fmt.Println(Errinfo(), msg, err)
		return
	}
}
func Lprintf(msg string, err error) {
	if err != nil {
		//fmt.Printf("insert failed, err:%v\n", err)
		log.Printf(Errinfo()+msg, err)
	}
}

func Lprintfr(msg string, err error) {
	if err != nil {
		//fmt.Printf("insert failed, err:%v\n", err)
		log.Printf(Errinfo()+msg, err)
		return
	}

}

func Lprintfunc(msg string, err error, add func()) {
	if err != nil {
		add()
		log.Printf(Errinfo()+msg, err)
	}
}
func Lprintfuncr(msg string, err error, add func()) {
	if err != nil {
		add()
		log.Printf(Errinfo()+msg, err)
		return
	}
}
func Lprintln(msg string, err error) {
	if err != nil {
		//fmt.Println("insert failed, err:", err)
		log.Println(Errinfo(), msg, err)
	}
}

func Lprintlnr(msg string, err error) {
	if err != nil {
		//fmt.Println("insert failed, err:", err)
		log.Println(Errinfo(), msg, err)
		return
	}
}

// 