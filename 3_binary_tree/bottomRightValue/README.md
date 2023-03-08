# Bottom Right Value

Write a function, bottomRightValue, that takes in the root of a binary tree. The function should return the right-most value in the bottom-most level of the tree.

You may assume that the input tree is non-empty.

```
Test 1:
a := &Node{Val: 3}
b := &Node{Val: 11}
c := &Node{Val: 10}
d := &Node{Val: 4}
e := &Node{Val: -2}
f := &Node{Val: 1}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f

//       3
//    /    \
//   11     10
//  / \      \
// 4   -2     1

bottomRightValue(a) // 1
```

```
Test 2:
a := &Node{Val: -1}
b := &Node{Val: -6}
c := &Node{Val: -5}
d := &Node{Val: -3}
e := &Node{Val: -4}
f := &Node{Val: -13}
g := &Node{Val: -2}
h := &Node{Val: 6}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f
e.Left = g
e.Right = h

//        -1
//      /   \
//    -6    -5
//   /  \     \
// -3   -4   -13
//     / \       
//    -2  6

bottomRightValue(a) // 6
```

```
Test 3:
a := &Node{Val: -1}
b := &Node{Val: -6}
c := &Node{Val: -5}
d := &Node{Val: -3}
e := &Node{Val: -4}
f := &Node{Val: -13}
g := &Node{Val: -2}
h := &Node{Val: 6}
i := &Node{Val: 7}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f
e.Left = g
e.Right = h
f.Left = i

//        -1
//      /   \
//    -6    -5
//   /  \     \
// -3   -4   -13
//     / \    /   
//    -2  6  7 

bottomRightValue(a) // 7
```

```
Test 4:
a := &Node{Val: "a"}
b := &Node{Val: "b"}
c := &Node{Val: "c"}
d := &Node{Val: "d"}
e := &Node{Val: "e"}
f := &Node{Val: "f"}

a.Left = b
a.Right = c
b.Right = d
d.Left = e
e.Right = f

//      a
//    /   \ 
//   b     c
//    \
//     d
//    /
//   e
//  /
// f

bottomRightValue(a) // "f"
```

```
Test 5:
a := &Node{Val: 42}

bottomRightValue(a) // 42
```