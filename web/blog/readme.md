# 博客代码

## 动机

我很想实现一个自己的产品，“博客”是一个成熟度很高的产品，网上能搜索到想到多的点子。我想试试，看看自己能不能独立完成一个产品的设计和实现。

说白了，我想要一个自己的东西。

另一个动机是，golang这门语言我是真的喜欢，简单又粗暴，基于函数的一套开发是非常对我的胃口的。我喜欢它，所以我愿意花时间用折腾他。


## 安装 tailwind.css

```sh

 npm install tailwindcss @tailwindcss/cli
 touch static/input.css
 echo '@import "tailwindcss";' > static/tailwind_input.css
 npx @tailwindcss/cli -i ./static/tailwind_input.css -o ./static/tailwind.css --watch --minify

```
