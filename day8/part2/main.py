#!/usr/bin/env python3

import numpy as np

def main():
    trees = []
    results = []
    
    #Open inputs
    with open("../inputs-day8.txt", "r") as f:
    #with open("../example.txt", "r") as f:
        for line in f.readlines():
            trees.append([int(i) for i in line[:-1]])
    f.close()

    # Create a NumPy array to manipulate better columns
    forest = np.array(trees)

    num_rows = forest.shape[0]
    num_cols = forest.shape[1]
    
    for row in range(0, num_rows):
        for col in range(0, num_cols): 
 
            left = forest[row,:col]
            right = forest[row,(col+1):]
            top = forest[:row,col]
            bottom = forest[(row+1):,col]

            # Calculate trees
            result = checkVisibility(
                forest[row][col],
                left,
                right,
                top,
                bottom)
            
            results.append(result)

    # Print highest score:
    print(f"The max is: {max(results)}")


def checkVisibility(n, lft, rgt, top, btm):
    
    # Check left
    lcount = 0
    for l in np.flip(lft):
        if l < n:
            lcount += 1
        else:
            lcount += 1
            break
    
    # Check right
    rcount = 0
    for r in rgt:
        if r < n:
            rcount += 1
        else:
            rcount += 1
            break
    
    # Check top
    tcount = 0
    for t in np.flip(top):
        if t < n:
            tcount += 1
        else:
            tcount += 1
            break
    
    # Check bottom
    bcount = 0
    for b in btm:
        if b < n:
            bcount += 1
        else:
            bcount += 1
            break

    return lcount * rcount * tcount * bcount

if __name__ == "__main__":
    main()