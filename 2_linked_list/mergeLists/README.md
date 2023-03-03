# Merge Lists

Write a function, mergeLists, that takes in the head of two sorted linked lists as arguments. The function should merge the two lists together into single sorted linked list. The function should return the head of the merged linked list.

Do this in-place, by mutating the original Nodes.

You may assume that both input lists are non-empty and contain increasing sorted numbers.

```
    a := &Node{
		Val: 5,
	}
	b := &Node{
		Val: 7,
	}
	c := &Node{
		Val: 10,
	}
	d := &Node{
		Val: 12,
	}
	e := &Node{
		Val:20,
	}
	f := &Node{
		Val: 28,
	}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f

    // 5 -> 7 -> 10 -> 12 -> 20 -> 28

	q := &Node{
		Val: 6,
	}
	r := &Node{
		Val: 8,
	}
	s := &Node{
		Val: 9,
	}
	t := &Node{
		Val: 25,
	}
	q.Next = r
	r.Next = s
	s.Next = t

    // 6 -> 8 -> 9 -> 25
    mergeLists(a, q)
    // 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 12 -> 20 -> 25 -> 28
```
