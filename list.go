package go_leda

import (
    "fmt"
)

type ItemType int

type List struct {
    item ItemType
    next *List
}

func SearchList(l *List, x ItemType) *List {
    if l == nil {
        return nil
    }

    if l.item == x {
        return l
    } else {
        return (SearchList(l.next, x))
    }
}

func InsertList(l *List, x ItemType) {
    var p *List

    p = new(List)
    p.item = x
    p.next = l
    l = p
}


