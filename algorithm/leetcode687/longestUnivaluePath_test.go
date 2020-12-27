package leetcode687

import "testing"

func Test_longestUnivaluePath(t *testing.T) {
	cases := []struct {
		root     *TreeNode
		expected int
	}{
		// 		 5
		// 		/ \
		// 	 4   5
		// 	/ \   \
		// 1  1   5
		{&TreeNode{5,
			&TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}},
			&TreeNode{5, nil, &TreeNode{5, nil, nil}}},
			2},
		// [1,null,1,1,1,1,1,1]
		// 	 1
		// nil  1
		// 	 1   1
		//  1  1 1
		{&TreeNode{1, nil, &TreeNode{1, &TreeNode{1, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{1, &TreeNode{1, nil, nil}, nil}}}, 4},
	}
	for _, c := range cases {
		result := longestUnivaluePath(c.root)
		if result != c.expected {
			t.Fatalf("\n longestUnivaluePath: root is %v, expected is %v, result is %v \n", c.root, c.expected, result)
		}
	}
}
