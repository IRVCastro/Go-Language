`name := ""`
 - This is the quick/implicit assigment operator 
 - Define a variable where its type is implied based on the value that you put on the right side 

# Var types

## Var int128, uint(Unsigned int)
 - In Go you should Define how many bytes you will need for this type of variable
 - int16, 8, 64
 -float32, float64
 - uint can only store 

## Boolean, byte

## While
    - In go we don't have the while statment we have the:
    for true 

# func getSpin(reel[]string)[][] string{}
We will return a slice of 2 dimensions it's because each slice it's gonna represent a row that we gonna display to the user,
Something like {{A B C}, {A A A}, {A B D}}

# func getSpin(reel[]string, rows int, cols int)[][] string

We get our reel this contain all of the symbols 
We have the cols and rows that we want to generate
We return a two dimensions slice that represent the grid the user see

# Receiver Function
No Go não existem classes no sentido tradiciona, em vez disso o Go associa funções a tipos de dados (estruturas) usando o que chamamos de Receiver que fica logo antes do nome da função

    (g *Game)
