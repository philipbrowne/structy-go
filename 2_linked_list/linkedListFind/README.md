# Linked List Find

Write a function, linkedListFind, that takes in the head of a linked list and a target value. The function should return a boolean indicating whether or not the linked list contains the target.

```
a := Node{Val: "A"}
b := Node{Val: "B"}
c := Node{Val: "C"}
d := Node{Val: "D"}
a.Next = &b
b.Next = &c
c.Next = &d

linkedListFind(a, "c"); // true
```
