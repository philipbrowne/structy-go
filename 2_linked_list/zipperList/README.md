# Zipper Lists

Write a function, zipperLists, that takes in the head of two linked lists as arguments. The function should zipper the two lists together into single linked list by alternating nodes. If one of the linked lists is longer than the other, the resulting list should terminate with the remaining nodes. The function should return the head of the zippered linked list.

Do this in-place, by mutating the original Nodes.

You may assume that both input lists are non-empty.

```

a := &Node{Val: "A"}
b := &Node{Val: "B"}
c := &Node{Val: "C"}
d := &Node{Val: "D"}

a.Next = b
b.Next = c
c.Next = d

x := &Node{Val: "X"}
y := &Node{Val: "Y"}
z := &Node{Val: "Z"}

x.Next = y
y.Next = z

zipperList(a, y)
// "A" -> "X" -> "B" -> "Y" -> "Z" -> "D"
```
