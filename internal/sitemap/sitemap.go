package sitemap

import (
	"github.com/Nageshbansal/web-crawler/internal/logger"
	"strings"
)

type Sitemap struct {
	Root *Node
}

func NewSitemap(urls []string) *Sitemap {

	logger.Info("[NewSitemap]: NewSitemap started")
	root := NewNode("root")
	for _, url := range urls {
		parts := strings.Split(strings.Trim(url, "/"), "/")
		currentNode := root
		for _, part := range parts {
			if _, exists := currentNode.Links[part]; !exists {
				currentNode.Links[part] = NewNode(part)
			}
			currentNode = currentNode.Links[part]
		}
	}
	return &Sitemap{Root: root}
}

func (s *Sitemap) String() string {
	var sb strings.Builder
	s.buildString(&sb, s.Root, 0)
	return sb.String()
}

func (s *Sitemap) buildString(sb *strings.Builder, node *Node, depth int) {
	if node.URL != "root" {
		indent := strings.Repeat("    ", depth)
		sb.WriteString(indent + node.URL + "\n")
	}
	for _, link := range node.Links {
		s.buildString(sb, link, depth+1)
	}
}
