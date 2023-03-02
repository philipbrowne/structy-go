# Linked List Values

Write a function, linkedListValues, that takes in the head of a linked list as an argument. The function should return a slice containing all values of the nodes in the linked list.

```
test_00:
a := &Node{Val: "A"}
b := &Node{Val: "B"}
c := &Node{Val: "C"}
d := &Node{Val: "D"}
a.Next = b
b.Next = c
c.Next = d
// a -> b -> c -> d
linkedListValues(a) // -> [ "a", "b", "c", "d" ]

test_01:
x := &Node{Val: "X"}
y := &Node{Val: "Y"}

x.next = y

// x -> y

linkedListValues(x) // -> [ "x", "y" ]

test_02:
q := &Node{Val: "Q"}

// q

linkedListValues(q) // -> [ "q" ]
```
