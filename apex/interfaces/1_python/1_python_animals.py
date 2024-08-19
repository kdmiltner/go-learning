# Class Dog implements sound method
class Dog:
    def sound(self):
        return "Woof!"


# Class Cat implements sound method
class Cat:
    def sound(self):
        return "Meow!"


# Class Cow implements sound method
class Cow:
    def sound(self):
        return "Moo!"


# make_sound function takes an argument of animal and will call the sound method
def make_sound(animal):
    print(animal.sound())


# Creating instances of each class
dog = Dog()
cat = Cat()
cow = Cow()

# Using the make_sound function to call the sound method on different objects
make_sound(dog)  # Output: Woof!
make_sound(cat)  # Output: Meow!
make_sound(cow)  # Output: Moo!
