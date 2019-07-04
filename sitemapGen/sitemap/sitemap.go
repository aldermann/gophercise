package sitemap

import (
    "sync"
)

var instance SiteMap

var lock = &sync.Mutex{}

type SiteMap struct {
    Adjacent map[string]map[string]bool
    Visited  map[string]bool
}

func (s SiteMap) IsNil() bool {
    return s.Adjacent == nil
}

func GetSiteMap() SiteMap {
    lock.Lock()
    defer lock.Unlock()
    if instance.IsNil() {
        instance = SiteMap{Adjacent: make(map[string]map[string]bool), Visited: make(map[string]bool)}
    }
    return instance
}

func (s SiteMap) MarkEdge(from, to string) {
    lock.Lock()
    defer lock.Unlock()
    if s.Adjacent[from] == nil {
        s.Adjacent[from] = make(map[string]bool)
    }
    s.Adjacent[from][to] = true
}

func (s SiteMap) MarkVisited(node string) {
    lock.Lock()
    defer lock.Unlock()
    s.Visited[node] = true
}

func (s SiteMap) IsVisited(node string) bool {
    lock.Lock()
    defer lock.Unlock()
    r, ok := s.Visited[node]
    if !ok {
        return false
    }
    return r
}
