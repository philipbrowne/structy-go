# Longest Streak

Write a function, longestStreak, that takes in the head of a linked list as an argument. The function should return the length of the longest consecutive streak of the same value within the list.

```
Test 1:
a := &Node{Val: 5}
b := &Node{Val: 5}
c := &Node{Val: 7}
d := &Node{Val: 7}
e := &Node{Val: 7}
f := &Node{Val: 6}
a.Next = b
b.Next = c
c.Next = d
d.Next = e
e.Next = f
longestStreak(a) // 3

Test 2:
a := &Node{Val: 3}
b := &Node{Val: 3}
c := &Node{Val: 3}
d := &Node{Val: 3}
e := &Node{Val: 9}
f := &Node{Val: 9}
a.Next = b
b.Next = c
c.Next = d
d.Next = e
e.Next = f
longestStreak(a) // 4
```
