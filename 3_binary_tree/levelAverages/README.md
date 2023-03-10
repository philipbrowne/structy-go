# Level Averages

Write a function, levelAverages, that takes in the root of a binary tree that contains number values. The function should return a slice containing the average value of each level.

```
Test 1:
a := &Node{Val: 3}
b := &Node{Val: 11}
c := &Node{Val: 4}
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
//   11     4
//  / \      \
// 4   -2     1

levelAverages(a)
// []float64{3, 7.5, 1}
```

```
Test 2:
a := &Node{Val: 5}
b := &Node{Val: 11}
c := &Node{Val: 54}
d := &Node{Val: 20}
e := &Node{Val: 15}
f := &Node{Val: 1}
g := &Node{Val:3}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
e.Left = f
e.Right = g

//        5
//     /    \
//    11    54
//  /   \
// 20   15
//      / \
//     1  3

levelAverages(a)
// []float64{5, 32.5, 17.5, 2}
```

```
Test 3
a := &Node{Val: -1}
b := &Node{Val: -6}
c := &Node{Val: -5}
d := &Node{Val: -3}
e := &Node{Val: 0}
f := &Node{Val: 45}
g := &Node{Val:-1}
h := &Node{Val:-2}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f
e.Left = g
f.Right = h

//        -1
//      /   \
//    -6    -5
//   /  \     \
// -3   0     45
//     /       \
//    -1       -2

levelAverages(a)
// []float64{-1, -5.5, 14, -1.5}
```

```
Test 4:
q := &Node{Val: 13}
r := &Node{Val: 4}
s := &Node{Val: 2}
tt := &Node{Val: 9}
u := &Node{Val: 2}
v := &Node{Val: 42}
	
q.Left = r
q.Right = s
r.Right = tt
tt.Left = u
u.Right = v

//        13
//      /   \
//     4     2
//      \
//       9
//      /
//     2
//    /
//   42

res := levelAverages(q)
// []float64{13, 3, 9, 2, 42}
```

```
Test 5:
levelAverages(nil)
// []float64{}
```