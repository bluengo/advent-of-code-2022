#!/usr/bin/env python3

def main():
    results = []
    total = 0
    forest = []
    #Open inputs
    with open("../inputs-day8.txt", "r") as f:
    #with open("example.txt", "r") as f:
        for line in f:
            newline = [int(i) for i in line[:-1]]
            forest.append(newline)

    num_rows = len(forest)
    num_cols = len(forest[0])

    # We sum all items in the edge:
    # Top and Bottom lines:
    total += (num_cols * 2) # First and last row
    total += (num_rows - 2) * 2 # First and last columns, without countint first and last row
    
    # Check if any element is greater than the elements to its edge
    for row in range(1, num_rows-1):
        for col in range(1, num_cols-1):
            print("\n")       
            left = forest[row][:col]
            right = forest[row][col:][1:]
            top = getColumn(forest, col)[:row]
            bottom = getColumn(forest, col)[row:][1:]

            print(forest[row][col])
            print(f"left: {left}")
            print(f"right: {right}")
            print(f"top: {top}")
            print(f"bottom: {bottom}")
            
            if forest[row][col] > max(left):
                print(f"\n{forest[row][col]} visible from left")
                results.append(f"{row}-{col}")
                total += 1
                continue

            elif forest[row][col] > max(right):
                print(f"\n{forest[row][col]} visible from right")
                results.append(f"{row}-{col}")
                total += 1
                continue

            elif forest[row][col] > max(top):
                print(f"\n{forest[row][col]} visible from top")
                results.append(f"{row}-{col}")
                total += 1
                continue

            elif forest[row][col] > max(bottom):
                print(f"\n{forest[row][col]} visible from bottom")
                results.append(f"{row}-{col}")
                total += 1
                continue
            #else:
                #print(f"{forest[row][col]} NOT VISIBLE")

    print(f"TOTAL: {total}")
    print(results)

def getColumn(arr, column_index):
    return [row[column_index] for row in arr]

if __name__ == "__main__":
    main()