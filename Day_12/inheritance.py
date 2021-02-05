##### WARNING #######
##### WARNING #######
### The code for the class person was provided for HackerRank
# My code was the wrote in Student Class

class Person:
	def __init__(self, firstName, lastName, idNumber):
		self.firstName = firstName
		self.lastName = lastName
		self.idNumber = idNumber
	def printPerson(self):
		print("Name:", self.lastName + ",", self.firstName)
		print("ID:", self.idNumber)

class Student(Person):
    def __init__(self, firstName, lastName, idNumber, scores):
        self.scores = scores
        super().__init__(firstName, lastName, idNumber)
    
    def calculate(self):
        average=sum(self.scores)/len(self.scores)
        if (average <= 100) & (average >= 90):
            return "O"
        elif (average >= 80) & (average < 90):
            return "E"
        elif (average >= 70) & (average < 80):
            return "A"
        elif (average >= 55) & (average < 70):
            return "P"
        elif (average >= 40) & (average < 55):
            return "D"
        else:
            return "T" 


##### WARNING #######
##### WARNING #######

## The following code was provided for Hacker Rank,            
line = input().split()
firstName = line[0]
lastName = line[1]
idNum = line[2]
numScores = int(input()) # not needed for Python
scores = list( map(int, input().split()) )
s = Student(firstName, lastName, idNum, scores)
s.printPerson()
print("Grade:", s.calculate())