# Lem-in

## Objectives
The goal of this project is to simulate a digital version of an ant farm. The program `lem-in` reads a file that describes the layout of an ant colony, including details about rooms, tunnels, and the ants themselves. The objective is to find the quickest way to move a specified number of ants from a starting room (`##start`) to an ending room (`##end`).

## How it works
1. **Ant Farm Setup**: 
   - You define an ant farm with rooms and tunnels.
   - All ants start in the room designated as `##start`.
   - The objective is to move all ants to the room designated as `##end`.

2. **Finding the Path**:
   - The program calculates the shortest path or paths from `##start` to `##end` using various graph traversal algorithms.

3. **Output Format**:
   - The program outputs the number of ants, the details of each room, and the connections between rooms.
   - It then displays each move made by the ants from room to room until they reach `##end`.

4. **Error Handling**:
   - The program handles various error conditions, such as invalid data formats, incorrect room configurations, or missing `##start` or `##end`.

5. **Constraints**:
   - Rooms are defined by "name coord_x coord_y", and links (tunnels) are defined by "name1-name2".
   - Each room can only contain one ant at a time, except `##start` and `##end`.
   - Each tunnel can only be used once per turn.
   - The program must ensure ants move efficiently without causing traffic jams.

## Example
Here's an example of a colony setup and the corresponding representation:

```      
        _________________
       /                 \
  ____[5]----[3]--[1]     |
 /            |    /      |
[6]---[0]----[4]  /       |
 \   ________/|  /        |
  \ /        [2]/________/
  [7]_________/
```

  
## Project Structure :

```
.
├── README.md
├── examples
│   ├── README.md
│   ├── badtest00.txt
│   ├── badtest01.txt
│   ├── test00.txt
│   ├── test01.txt
│   ├── test02.txt
│   ├── test03.txt
│   ├── test04.txt
│   ├── test05.txt
│   ├── test06.txt
│   └── test07.txt
├── go.mod
├── go.work
├── main.go
└── src
    ├── colony.go
    ├── go.mod
    ├── initFarm.go
    ├── input.go
    └── population.go
  
  ```

  - colony.go : zone, tunnels and rooms
  - initFarm.go : read file.txt and extract data
  - population.go : ants

### Exemple de Ants grid :

Knowing which paths to use :

P = Path
A = Ant

```  
          A3
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
        +---+
     A2 |   |
    +---+---+
 A1 |   |   |
+---+---+---+
|   |   |   |
+---+---+---+
  P1  P2  P3
  ```  

- Ant1 will take 1 steps to reach the end 
- Ant2 will take 2 steps 
- Ant3 will take 6 steps, which is significantly more

It might be more efficient to redirect ant3 to one of the shorter paths.

To determine the best time for this, we just need to do some basic calculations.

Similarly, now we have to decide whether to send ant2 to Path1 or Path2. We simply add the number of rooms in a path to the number of ants already in that path. If the total for Path1 exceeds the total for Path2, then we send the ant to Path2.

R1 : Number of room in path 1 
A1 : Number of ants already in path 1

R2 : Number of room in path 2 
A2 : Number of ants already in path 2

Example : 

Ants : A1, A2, A3, A4

Sum for Path1= R1+A1
Sum for Path2 = R2+A2

If R1 + A1 > R2 + A2 ; A2 go to path 2

So if R1 + A1 > R2 + A2 :

#### For A2 : 

- Rooms in Path1: 2 
- Ants in Path1: 1
- Rooms in Path2: 2 + 1 = 3

3 is not greater than 3. Put ant2 in Path1.
         
  ```         
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
 A2     +---+ <------ Add A2 in Path 1
 A1     |   |
+---+---+---+
|   |   |   |
+---+---+---+
|   |   |   |
+---+---+---+
  P1  P2  P3
```  

#### For A3 : 

- Rooms in Path1: 2 
- Ants in Path1: 2
- Ants in Path2: 0
- Rooms in Path2: 2

2 + 2 = 4 

4 is greater than 3. Put ant3 in Path2.

 ```           
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
 A2     +---+ 
 A1  A3 |   | <------ Add A3 in Path 2
+---+---+---+
|   |   |   |
+---+---+---+
|   |   |   |
+---+---+---+
  P1  P2  P3
```  

#### For A4 : 

- Rooms in Path1: 2 
- Ants in Path1: 2
- Ants in Path2: 1
- Rooms in Path2: 2

2 + 1 = 3 

3 not is greater than 3. Put ant4 in Path2.

  ```          
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
 A2  A4 +---+ <------ Add A4 in Path 2
 A1  A3 |   | 
+---+---+---+
|   |   |   |
+---+---+---+
|   |   |   |
+---+---+---+
  P1  P2  P3
```  

### Resolve :

1 :

  ```          
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
        +---+ 
 A2  A4 |   | 
+---+---+---+
|A1 |A3 |   |
+---+---+---+
|   |   |   |
+---+---+---+
  P1  P2  P3
```  

2 :

  ```          
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
        +---+ 
        |   | 
+---+---+---+
|A2 |A4 |   |
+---+---+---+
|A1 |A3 |   |
+---+---+---+
  P1  P2  P3
```  

3 :

  ```          
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
        +---+ 
        |   | 
+---+---+---+
|   |   |   |
+---+---+---+
|A2 |A4 |   |
+---+---+---+
  A1  A3     
``` 

4 :

3 :

  ```          
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
        +---+
        |   |
        +---+ 
        |   | } <----- Start
+---+---+---+
|   |   |   |
+---+---+---+
|   |   |   |
+---+---+---+
 A2  A4       
 A1  A3       Finish !
``` 

### Result Struct 

```
number_of_ants
the_rooms
the_links

Lx-y
Lx-y Lx-y
Lx-y Lx-y Lx-y
Lx-y Lx-y
Lx-y
```
