# Tree Value Count

Write a function, tree_value_count, that takes in the root of a binary tree and a target value. The function should return the number of times that the target occurs in the tree.

```
Test 1:
a := &Node{Val: 12}
b := &Node{Val: 6}
c := &Node{Val: 6}
d := &Node{Val: 4}
e := &Node{Val: 6}
f := &Node{Val: 12}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f

//      12
//    /   \
//   6     6
//  / \     \
// 4   6     12

treeValueCount(a,  6) // 3
```
