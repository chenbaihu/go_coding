package fun

//#include "fun.h"
import "C"

// http://blog.sina.com.cn/s/blog_538d55be01015h6g.html
//再次提醒大家：import "C" 一定要紧跟C语言代码注释结束的最后一行，绝对不能空出一行，也不能和其他包合并写到import小括号内。

func MySecret() string {
	return (C.GoString(C.MySecret()))
}
