# Insert Node

Write a function, insertNode, that takes in the head of a linked list, a value, and an index. The function should insert a new node with the value into the list at the specified index. Consider the head of the linked list as index 0. The function should return the head of the resulting linked list.

Do this in-place.

You may assume that the input list is non-empty and the index is not greater than the length of the input list.

```
Test 1:
a := &Node{Val: "A"}
b := &Node{Val: "B"}
c := &Node{Val: "C"}
d := &Node{Val: "D"}
a.Next = b
b.Next = c
c.Next = d
insertNode(a, "X", 2)
A -> B -> X -> C -> D

Test 2:
a := &Node{Val: "A"}
b := &Node{Val: "B"}
c := &Node{Val: "C"}
d := &Node{Val: "D"}
a.Next = b
b.Next = c
c.Next = d
insertNode(a, "V", 3)
A -> B -> C -> -> V -> D

Test 3:
a := &Node{Val: "A"}
b := &Node{Val: "B"}
c := &Node{Val: "C"}
d := &Node{Val: "D"}
a.Next = b
b.Next = c
c.Next = d
insertNode(a, "M", 4)
A -> B -> C -> -> D -> M
```
