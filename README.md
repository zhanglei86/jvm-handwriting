## jvm-handwriting
>动手实践jvm

`write your own java virtual machine.`整理自[张秀宏](https://github.com/zxh0)的书籍和源码，顺便练练go语言。


来自《深入理解java虚拟机》作者周志明
>想要了解java虚拟机的内部运行原理，阅读虚拟机规范、书籍、源码是一种常见的途径，而从0开始，自己动手编写一个实验性质的java虚拟机，也许会是一种更加有趣且有效的学习路径。
如果不考虑java庞大类库的实现和jvm的实际生产力需求，仅是去“正确地”实现一台java虚拟机，其实并不如大多数人所想的那样高深和困难。所做的只需正确读取class文件中每一条字节码指令，并且能正确执行这些指令所蕴含的操作即可。
xx, 一步步完成java虚拟机的各个组成部分，在动手的过程中了解jvm的运行原理。


### note

debug
```
ch2的
bin/jvmgo -Xjre /Library/Java/JavaVirtualMachines/jdk1.8.0_191.jdk/Contents/Home/jre java.lang.Object 123
ch03,04
bin/jvmgo -Xjre /Library/Java/JavaVirtualMachines/jdk1.8.0_191.jdk/Contents/Home/jre class/ch03/ClassFileTest
ch05
bin/jvmgo -Xjre /Library/Java/JavaVirtualMachines/jdk1.8.0_191.jdk/Contents/Home/jre class/ch05/GaussTest
ch06
bin/jvmgo -Xjre /Library/Java/JavaVirtualMachines/jdk1.8.0_191.jdk/Contents/Home/jre class/ch06/MyObject
jvmgo/book/ch06/MyObject
```


### READ
xx

- http://www.runoob.com/
