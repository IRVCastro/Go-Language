# Creating a Chess Game

## Steps:
1 - Setting up TheChess Board 
2 - Displaying the Board 
3 - Taking input
4 - Move Validations

https://code.likeagirl.io/building-a-simple-chess-game-using-go-e003f187bbcb#229c

https://www.youtube.com/watch?v=_B4dzWMsVQY&list=PLHtpHBMEnTArfduVS6JpevZInP6kQGfT9&index=5

# 1 - Setting Up The board
 - Represent the board using a 2D array, Hold a string  like:
    "Wp" = White Pawn
    "Bk" = balck king

- At this point we are setting up our board and putting the pices on there, show this on the screen to the viewer

# 2 - Displaying the board // Ebiten engine
board/ — your Board type, your Piece type, initBoard, move validation, check detection. Pure data and rules. It should be possible to play a whole game through a terminal with only this package. It must never import Ebiten.
ui/ — your Game struct with Update, Draw, Layout. Square sizes, colors, sprite loading, mouse clicks. This is the only folder that imports Ebiten.
main.go — creates a board, hands it to the ui, calls RunGame. Should be ~10 lines and contain no decisions.

o's export rule: a name is visible outside its package only if it starts with a capital letter. Not a keyword, not an annotation — just capitalization. Lowercase means "private to this package."


Part of documentation that will help me to draw shapes
https://pkg.go.dev/github.com/erparts/go-shapes#Renderer.DrawRect

Important concept Understanding the layout, how a screen layout displayes work for video games

`I need to create a separete shape for each waure on the board and send them to a specific coordinates`
Because at this momento i only have a screen with some draws

## Board Coordinates
    - The coordinates are shifted by one Square
    640px Screen and 80px squares
    Last row strat at x = 0 , y = 560

    Definning pieces at our board we use a thing called iota on go, check wich one are balck and white 

    Create the function new Board and define the initial positions of the board:
        b[0][f] = back[f]     // White pieces on rank 1
        b[1][f] = piece.WPawn // White pawns on rank 2
        b[6][f] = piece.BPawn // Black pawns on rank 7
        b[7][f] = back[f] + 6 // Black pieces on rank 8

## Adding Images to our project
    https://ebitengine.org/en/documents/matrix.html

    Now i could display the images,

    Now i have to find a way to display the images at the right positions and make an associate the image with the position

    In chess Rank are the numbers
    Files are the characters

    `GeoM Important Bible that we are gonan ahve to check in the future`

`In Go, the _ (underscore) is called the Blank Identifier.`

It acts as an "anonymous placeholder" or a code sinkhole: it tells the Go compiler to ignore or discard a value returned by an assignment or a function call.

## Geo M
Para posicionar cada peça na casa correta e ajustar a escala (tamanho) dos seus PNGs para que caibam exatamente na casa de $80 \times 80\text{px}$, você vai usar a matriz geométrica GeoM do ebiten.DrawImageOptions

Primeiro a Escala (Scale): Redimensiona a imagem para $80 \times 80\text{px}$.Depois a Translação (Translate): Move a imagem redimensionada para a posição $(X, Y)$ correta da casa no tabuleiro.

## Move Creations /  Variations
http://chessprogramming.org/Encoding_Moves
https://www.boristhebrave.com/2022/10/22/defining-chess-piece-moves-using-regular-expressions/

The Pieces Fall into two groups:
Jumpers:
 - Knight King 
 - one step in each directions, then stop

Sliders:
 - Rook bishop queen
 - Keep walking in each direction until they hit the edge or piece

Declare a second type — call it Position or GameState — that is a struct
  containing a Board field plus SideToMove, castling flags, and the en-passant square. Board stays exactly as it is (still an
  array, still indexable, your NewBoard untouched), and the new type holds everything a grid can't express.

## Taking Inputs
