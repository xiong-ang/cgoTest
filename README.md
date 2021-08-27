# cgoTest

> 实例代码实现了用cgo使用c静态链接库，在Windows平台上测试通过。

## 具体实现步骤
1. 查看cpp可执行程序的flags.make，获取编译器选项
2. 查看cpp可执行程序的link.txt，获取链接库的编译项
3. 如果使用c++，还需要增加链接-lstdc++
4. 包含相应头文件，即可使用c接口的库函数

## 参考资料
1. https://toutiao.io/posts/493728/app_preview
