# Reverse Linked List

Write a function, reverseList, that takes in the head of a linked list as an argument. The function should reverse the order of the nodes in the linked list in-place and return the new head of the reversed linked list.

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

// a -> b -> c -> d -> e -> f
reverseList(a); // f -> e -> d -> c -> b -> a
```
