class Meodinger {

    public int[][] floodFill(int[][] image, int sr, int sc, int newColor) {
        return image[sr][sc] == newColor ? image : changes(image, sr, sc, image[sr][sc], newColor);
    }

    public int[][] changes(int[][] image, int i, int j, int a, int newColor) {
        if(i < 0 || i == image.length || j < 0 || j == image[0].length || image[i][j] != a) return image;
        
        image[i][j] = newColor;
        changes(image, i + 1, j, a, newColor);
        changes(image, i - 1, j, a, newColor);
        changes(image, i, j + 1, a, newColor);
        changes(image, i, j - 1, a, newColor);

        return image;
    }

}