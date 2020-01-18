# [Mars Rover Kata](https://www.youtube.com/watch?v=24vzFAvOzo0)

Goal here was I was trying to practice classicist TDD style more.

## Rules

1. Initial starting point for a rover is 0,0,N (coordination is 0,0 and heading North)
2. The grid size id 10,10
3. Directions are N, E, S and W
4. L and R commands mean rotating Left and Right
5. Command M means Move (in current direction)
6. Rover receives a list of commands like "RMMLM" and returns final position like 2:1:N
7. Rover wraps around if reaches the end of the grid
8. The grid may have obstacles. If the sequence on commands end up in an obstacle, the rover moves to last possible coordinate and reports the obstacle Put an O (not a zero) at the beginning O:2:2:N
