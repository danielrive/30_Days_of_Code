##### WARNING #######

## The following code was provided for Hacker Rank,

class Difference:
    def __init__(self, a):
        self.__elements = a

    # The following code is mine    
    def computeDifference(self):
        maxValue = max(self.__elements)
        minValue = min(self.__elements)
        self.maximumDifference = maxValue - minValue

##### WARNING #######

## The following code was provided for Hacker Rank,
_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)
