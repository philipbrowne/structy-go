# Remove Node

Write a function, removeNode, that takes in the head of a linked list and a target value as arguments. The function should delete the node containing the target value from the linked list and return the head of the resulting linked list. If the target appears multiple times in the linked list, only remove the first instance of the target in the list.

Do this in-place.

You may assume that the input list is non-empty.

You may assume that the input list contains the target.

```
a := &Node{Val: "A"}
b := &Node{Val: "B"}
c := &Node{Val: "C"}
d := &Node{Val: "D"}
e := &Node{Val: "E"}
f := &Node{Val: "F"}
g := &Node{Val: "G"}
h := &Node{Val: "H"}

a.Next = b
b.Next = c
c.Next = d
d.Next = e
e.Next = f
f.Next = g
g.Next = h

removeNode(a, "C")
// a -> b -> d -> e -> ...

```
