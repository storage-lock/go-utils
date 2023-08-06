# Utils

# 一、这是什么

整个工具库比较通用的、基础的工具方法会下沉到这里

# 二、安装依赖

```bash
go get -u github.com/storage-lock/go-utils
```

# 三、API示例

## 3.1 utils.RandomID

```go
package main

import (
	"fmt"
	"github.com/storage-lock/go-utils"
)

func main() {

	// 生成一个随机的ID
	id := utils.RandomID()
	fmt.Println(id)
	// Output:
	// bd75553a04c9413c8550702c23026662

	// 生成一个随机ID，拼接给定的前缀
	id = utils.RandomID("foo-prefix-")
	fmt.Println(id)
	// Output:
	// foo-prefix-fcc5cec3599a4539abc7eebabc4491b4

}
```





