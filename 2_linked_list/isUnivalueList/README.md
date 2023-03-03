# Is Univalue List

Write a function, isUnivalueList, that takes in the head of a linked list as an argument. The function should return a boolean indicating whether or not the linked list contains exactly one unique value.

You may assume that the input list is non-empty.

```
Test 1:
a := &Node{Val: 7}
b := &Node{Val: 7}
c := &Node{Val: 7}
a.Next = b
b.Next = c
isUnivalueList(a) // true

Test 2:
a := &Node{Val: 7}
b := &Node{Val: 7}
c := &Node{Val: 4}
a.Next = b
b.Next = c
isUnivalueList(a) // false

Test 3:
a := &Node{Val: 7}
isUnivalueList(a) // true

```
