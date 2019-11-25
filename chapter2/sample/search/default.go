package search

// defaultMatcher implements the default matcher.
type defaultMatcher struct{}

// init registers the default matcher with the program.
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implements the behavior for the default matcher.
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}


/**
	学习笔记：
	该文件位于search目录下，所使用的的包名称为为search，另外search文件夹下面的所有文件都使用同一个包名
	诸如feed.go、match.go、search.go，这个是按照go的规则来的
	可以理解为一个文件夹下面就是一个包，一个包定义一组编译后的代码，每段代码都描述包的一部分
 */
