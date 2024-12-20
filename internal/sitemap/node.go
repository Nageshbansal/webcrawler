package sitemap

type Node struct {
	URL   string
	Links map[string]*Node
}

func NewNode(url string) *Node {
	return &Node{URL: url, Links: make(map[string]*Node)}
}
