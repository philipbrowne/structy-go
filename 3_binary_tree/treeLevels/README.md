# Tree Levels

Write a function, treeLevels, that takes in the root of a binary tree. The function should return a 2-Dimensional slice where each subslice represents a level of the tree.

```
Test 1:
a := &Node{Val: "a"}
b := &Node{Val: "b"}
c := &Node{Val: "c"}
d := &Node{Val: "d"}
e := &Node{Val: "e"}
f := &Node{Val: "f"}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f

//      a
//    /   \
//   b     c
//  / \     \
// d   e     f

treeLevels(a)
// [][]interface{}{{"a"},
// {"b", "c"},
// {"d", "e", "f"},
// }
```

```
Test 2:
a := &Node{Val: "a"}
b := &Node{Val: "b"}
c := &Node{Val: "c"}
d := &Node{Val: "d"}
e := &Node{Val: "e"}
f := &Node{Val: "f"}
g := &Node{Val: "g"}
h := &Node{Val: "h"}
i := &Node{Val: "i"}

a.Left = b
a.Right = c
b.Left = d
b.Right = e
c.Right = f

//         a
//      /    \
//     b      c
//   /  \      \
//  d    e      f
//      / \    /
//     g  h   i

treeLevels(a)
// [][]interface{}{{"a"},
// {"b", "c"},
// {"d", "e", "f"},
// {"g", "h", "i"},
// }
```

```
Test 3:
q := &Node{Val: "q"}
r := &Node{Val: "r"}
s := &Node{Val: "s"}
tt := &Node{Val: "t"}
u := &Node{Val: "u"}
v := &Node{Val: "v"}

q.Left = r
q.Right = s
r.Right = t
t.Left = u
u.Right = v

//      q
//    /   \
//   r     s
//    \
//     t
//    /
//   u
//  /
// v

treeLevels(q)
// [][]interface{}{{"q"},
// {"r", "s"},
// {"t"},
// {"u"},
// {"v"},
// }
```

```
Test 4:
treeLevels(nil)
// [][]interface{}{}
```
