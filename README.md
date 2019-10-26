#Induce
##一、rightfight 
###1. Right（num int）
- 备注：函数 Right 用于按 << （右移运算符）规则解析 num 返回一个[]int 
>例如：
arr := Right(5) //[1,4]

###2.HasRight（count，num int）
- 备注：函数 HasRight 用于将参数 num 按 <<（右移运算符）规则解析后检查参数 count 是否包含在其中
>例如：
isHas := HasRight(1,5) //true