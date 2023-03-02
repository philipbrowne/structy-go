# Sum List

Write a function, sumList, that takes in the head of a linked list containing numbers as an argument. The function should return the total sum of all values in the linked list.

```
a := Node{Val: 2}
b := Node{Val: 8}
c := Node{Val: 3}
d := Node{Val: -1}
e := Node{Val: 7}
a.Next = &b
b.Next = &c
c.Next = &d
d.Next = &e

sumList(a); // 19
```
