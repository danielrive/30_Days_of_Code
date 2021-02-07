import sys

S = input().strip()

try:
    newInt= int(S)
    print(newInt)
except ValueError:
    print("Bad String")