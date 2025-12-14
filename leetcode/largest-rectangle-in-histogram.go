package main

func main() {
	println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	maxArea := 0
	stack := make([]int, 0)

	for i := 0; i <= n; i++ {
		for len(stack) > 0 && (i == n || heights[stack[len(stack)-1]] >= heights[i]) {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		if i < n {
			stack = append(stack, i)
		}
	}
	return maxArea
}
