# Add Lists

This problem is on Leetcode at https://leetcode.com/problems/add-two-numbers/

Write a function, addLists, that takes in the head of two linked lists, each representing a number. The nodes of the linked lists contain digits as values. The nodes in the input lists are reversed; this means that the least significant digit of the number is the head. The function should return the head of a new linked listed representing the sum of the input lists. The output list should have its digits reversed as well.

```
Say we wanted to compute 621 + 354 normally. The sum is 975:

   621
 + 354
 -----
   975

Then, the reversed linked list format of this problem would appear as:

    1 -> 2 -> 6
 +  4 -> 5 -> 3
 --------------
    5 -> 7 -> 9
```

```
Test 1:
a1 := &Node{Val: 1}
a2 := &Node{Val: 2}
a3 := &Node{Val: 6}
a1.Next = a2
a2.Next = a3

b1 := &Node{Val: 4}
b2 := &Node{Val: 5}
b3 := &Node{Val: 3}
b1.Next = b2
b2.Next = b3

addLists(a1, b1)
// 5 -> 7 -> 9

Test 2:
a1 := &Node{Val: 1}
a2 := &Node{Val: 4}
a3 := &Node{Val: 5}
a4 := &Node{Val: 7}
a1.Next = a2
a2.Next = a3
a3.Next = a4

b1 := &Node{Val: 2}
b2 := &Node{Val: 3}
b1.Next = b2

addLists(a1, b1)
// 3 -> 7 -> 5 -> 7

Test 3:
a1 := &Node{Val: 9}
a2 := &Node{Val: 3}
a1.Next = a2

b1 := &Node{Val: 7}
b2 := &Node{Val: 4}
b1.Next = b2

addLists(a1, b1)
// 6 -> 8

Test 4:
a1 := &Node{Val: 9}
a2 := &Node{Val: 8}
a1.Next = a2

b1 := &Node{Val: 7}
b2 := &Node{Val: 4}
b1.Next = b2

addLists(a1, b1)
// 6 -> 3 -> 1

Test 5:
a1 := &Node{Val: 9}
a2 := &Node{Val: 9}
a3 := &Node{Val: 9}
a1.Next = a2
a2.Next = a3

b1 := &Node{Val: 6}

addLists(a1, b1)
// 5 -> 0 -> 0 -> 1
```
