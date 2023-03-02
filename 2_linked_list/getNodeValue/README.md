# Get Node Value

Write a function, `getNodeValue`, that takes in the head of a linked list and an index. The function should return the value of the linked list at the specified index.

If there is no node at the given index, then return `nil`.

```
a := Node{Val: "A"}
b := Node{Val: "B"}
c := Node{Val: "C"}
d := Node{Val: "D"}
a.Next = &b
b.Next = &c
c.Next = &d

// a -> b -> c -> d

getNodeValue(a, 2); // "c"
```
