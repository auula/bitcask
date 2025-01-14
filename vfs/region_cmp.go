package vfs

import "os"

// 垃圾回收压缩器工作原理
// 1. 如果磁盘上没有索引快照，就全局扫描恢复索引
// 2. 当索引恢复之后，运行了一段时间出发垃圾回收
// 3. 启动 GC 扫描磁盘数据文件和内存索引最新版比较
// 4. 如果磁盘上的文件记录和索引记录对上就迁移到新文件
// 5. 如果没有对上说明是旧文件，不要管它重复这个过程
// 6. 直到 GC 扫描完整个数据文件完成，最后删除这个文件
// 7. PS：重点是反向扫描，通过磁盘数据文件中 Key 名找内存到记录比较
// 8. 如果通过内存索引来找，会出现无法确定一个文件是否扫描干净
// 9. 因为内存索引的对应的数据记录会分配在不同数据文件中
type RegionCompressor struct {
	DirtyReginos []*os.File
}
