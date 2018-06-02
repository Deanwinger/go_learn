# GO基础
## 定义变量

    //定义一个名称为“variableName”，类型为"type"的变量
    var variableName type

    //初始化“variableName”的变量为“value”值，类型是“type”
    var variableName type = value


	//定义三个类型都是"type"的变量,并且分别初始化为相应的值
    var vname1, vname2, vname3 type= v1, v2, v3

        //可以直接忽略类型声明
        var vname1, vname2, vname3 = v1, v2, v3
        //可以继续简化
        vname1, vname2, vname3 := v1, v2, v3

`_`（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。在这个例子中，我们将值35赋予b，并同时丢弃34：

    _, b := 34, 35

## 内置基础类型
### Boolean
    //示例代码
    var isActive bool // 全局变量声明
    var enabled, disabled = true, false //忽略类型的声明
    func test(){
        var available bool  //一般声明
        valid := false      // 简短声明
        available = true    // 赋值操作
    }

### 数值类型
* 整数类型有无符号和带符号两种。Go同时支持int和uint，这两种类型的长度相同，但具体长度取决于不同编译器的实现。Go里面也有直接定义好位数的类型：`int8, int16, int32, int64`和`uint8, uint16, uint32, uint64`。其中`rune`是`int32`的别称，`byte`是`uint8`的别称

* 浮点数的类型有`float32`和`float64`两种（没有`float`类型），默认是`float64`    


## Pointer
~~~
    Unlike C, Go has no pointer arithmetic.
~~~

## Slices
~~~
    A slice does not store any data, it just describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array.
    不新创对象，直接在上面操纵
    https://tour.golang.org/moretypes/11
    上面的slice处有点无法理解
~~~