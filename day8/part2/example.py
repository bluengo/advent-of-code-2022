#!/usr/bin/env python3

import numpy as np

def main():
    trees = []
    
    #Open inputs
    with open("../example.txt", "r") as f:
        for line in f.readlines():
            trees.append([int(i) for i in line[:-1]])
    f.close()

    # Create a NumPy array to manipulate better columns
    forest = np.array(trees)

    num_rows = forest.shape[0]
    num_cols = forest.shape[1]
    
    print(forest[:,1]) # [0 5 5 3 5]
    print(forest[:3,1]) # [0 5 5]
    print(forest[1,:]) # [2 5 5 1 2]
    print(forest[:2,1]) # [0 5] # Column with index 1, from row 0 (:) and take 2
    print(forest[:3,1]) # [0 5 5]
    
    for row in range(0, num_rows):
        for col in range(0, num_cols):
            
            print(f"\nNumber: {forest[row][col]}")
            
            left = forest[row,:col]
            print(f"left: {left}")

            right = forest[row,(col+1):]
            print(f"right: {right}")

            top = forest[:row,col]
            print(f"top: {top}")

            bottom = forest[(row+1):,col]
            print(f"bottom: {bottom}")


if __name__ == "__main__":
    main()