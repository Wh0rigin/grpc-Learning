## Important file

`tree`
`   |-Client`
`   |    |- main.go       //提供客户端服务(函数调用方)`
`   |-pdfile`
`   |    |- hello.proto   //proto文件`  
`   |-services`
`   |    |- hello.go      //定义函数的实现`
`   |    |- hello.pb.go   //声明函数即其他的rgpc设置/由proto编译后生成`
`   |-main.go             //提供服务端函数(函数实现方)`
