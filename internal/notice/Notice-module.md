# Notice Module

## 业务分析

![image-20230612103907066](https://raw.githubusercontent.com/Anxiu0101/PicgoImg/master/202306121039034.png)

需要 Notice Title，Notice Category

## Basic Structure

```go
package notice

type (
	Notice struct {
		gorm.Model{
		    id
			create at
			delete at
			update at
        }
		issuer
		desc
		content
		signature
	}
)

```