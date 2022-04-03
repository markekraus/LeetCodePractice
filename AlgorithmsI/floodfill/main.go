package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	recurseFill(image, sr, sc, image[sr][sc], newColor)
	return image
}

func recurseFill(image [][]int, sr int, sc int, originalColor int, newColor int) {
	maxR := len(image)
	maxC := len(image[0])
	if (image)[sr][sc] == originalColor {
		image[sr][sc] = newColor
	}
	if sr-1 >= 0 && image[sr-1][sc] == originalColor {
		image[sr-1][sc] = newColor
		recurseFill(image, sr-1, sc, originalColor, newColor)
	}
	if sr+1 < maxR && image[sr+1][sc] == originalColor {
		image[sr+1][sc] = newColor
		recurseFill(image, sr+1, sc, originalColor, newColor)
	}
	if sc-1 >= 0 && image[sr][sc-1] == originalColor {
		image[sr][sc-1] = newColor
		recurseFill(image, sr, sc-1, originalColor, newColor)
	}
	if sc+1 < maxC && image[sr][sc+1] == originalColor {
		image[sr][sc+1] = newColor
		recurseFill(image, sr, sc+1, originalColor, newColor)
	}
}
