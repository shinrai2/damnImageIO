## About damnImageIO

damnImageIO 是一个使用 Golang 编写的图像处理库。
项目中处理像素数据的底层调用 [gonum](http://www.gonum.org/) 封装的矩阵库（~~原为自我封装的简易矩阵结构~~），而作为数据入口的上层从图片读入到矩阵转换的代码完全是自我编写。

本项目的最终目的是成为一个自用的图片处理库，编译成 C-shared library 供 Ruby 调用。一来弥补了 Ruby 中的 RMagick 库难以安装、MiniMagick 缺少教程的问题；二是可以更加深入地了解图片的编码和存储原理。
