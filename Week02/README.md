作业题目：

## 1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


Dao层属于与数据库通讯的基础层，报错应该把错误原样抛给上层处理。
当Dao层遇到sql.ErrNoRows时，上层Service捕获到根据具体业务具体处理，或者再Warp一下抛给上一层