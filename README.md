# chinese-poetry-migrate
https://github.com/chinese-poetry/chinese-poetry
上面仓库的数据迁移工具，跑一跑程序就可以方便的将数据迁移到数据库。

## 目前完成的表
- Authors 作者
- LunYu  论语
- TangShi 唐诗
- SongCi 宋词
- YuanQu 元曲
- SiShuWuJing  四书五经
- YouMengYing 幽梦影 
## 使用
将 chinese-poetry 仓库克隆到和此工具同一个目录，运行迁移即可。
目前 幽梦影的 comment 字段不统一，需要修改，四书五经 并不是统一的 [] 结构，也需要修改