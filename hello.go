package main  // 代码包声明语句。
import "fmt" //系统包用来输出的
import "time"

func main() {
    // 打印函数调用语句。用于打印输出信息。
   for {
       time.Sleep(time.Duration(1)* time.Second)
       fmt.Println(time.Now())
   }
}
