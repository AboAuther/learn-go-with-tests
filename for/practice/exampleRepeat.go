package practice

//修改测试代码，以便调用者可以指定字符重复的次数，然后修复代码
//写一个 ExampleRepeat 来完善你的函数文档

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
