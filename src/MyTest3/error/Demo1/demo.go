package main

import "fmt"

func main() {
	runError()

	fmt.Println("---------------------------")
	runPanicError()
}

type Student struct {
	Chinese int
	Math    int
	English int
}

var ss = []Student{
	{100, 90, 89},
	{80, 80, 80},
	{190, 40, 59},
	{80, 75, 66},
}

func runError() {
	i := 0

	for ; i < len(ss); i++ {
		flag, err := checkStudent(&ss[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		// 遇到异常数据就会立即返回，不能处理剩余的数据
		// 而且，正常逻辑中参杂异常处理，使得程序并不是那么优雅
		fmt.Printf("student %#v,及格? ：%t \n", ss[i], flag)
	}

}

func runPanicError() {
	i := 0
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		i++ // 跳过异常的数据，继续处理剩余的数据
		for ; i < len(ss); i++ {
			fmt.Printf("student %#v,及格? ：%t \n", ss[i], checkStudentS(&ss[i]))
		}
	}()

	for ; i < len(ss); i++ {
		fmt.Printf("student %#v,及格? ：%t \n", ss[i], checkStudentS(&ss[i]))
	}

}

// ----------------------------------------
func checkStudent(s *Student) (bool, error) {
	if s.Chinese > 100 || s.Math > 100 || s.English > 100 {
		return false, fmt.Errorf("student %#v, something error", s)
	}

	if s.Chinese > 60 && s.Math > 60 && s.English > 60 {
		return true, nil
	}

	return false, nil
}
func checkStudentS(s *Student) bool {
	if s.Chinese > 100 || s.Math > 100 || s.English > 100 {
		panic(fmt.Errorf("student %#v, something error", s))
	}

	if s.Chinese > 60 && s.Math > 60 && s.English > 60 {
		return true
	}

	return false
}
