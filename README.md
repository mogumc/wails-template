# wails-template
wails2 + Vue3 + Element-Plus 通用模板

## 关于使用该模板
若需要修改程序名,请注意检查所有文件的引用名均已修改  

## 使用方法
在``service``目录下新建需要的go文件并编写函数  
```Go
func greet(str string) string {
	return fmt.Sprintf("%s", str)
}
```
在``service``目录下找到``router.go``文件，按照wails官方方式注册api  
```Go
func (a *App) Greet(str string) string {
	return greet(str)
}
```
在前端直接调用对应代码
```Vue
<script setup>
import { reactive } from 'vue'
import { Greet } from '../../wailsjs/go/service/App'

const data = reactive({
  resultText: '默认文本',
})

function greet() {
  Greet(data.name).then((result) => {
    data.resultText = result
  })
}
</script>
  
<template>
  <div id="result" class="result">{{ data.resultText }}</div>
</template>
```

## 兼容性信息
[wails](https://github.com/wailsapp/wails) v2.10.1

## License
GPL-3.0 license
