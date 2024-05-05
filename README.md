# Project Description: Flappy Bird Terminal Game
![Screencast from 2024-05-06 02-39-44(1)](https://github.com/icryez/jumpingGame/assets/35337801/e13d42ea-a36b-4ba2-bfdb-20232bf66696)


## Overview:
This project aims to recreate the classic game Flappy Bird using terminal-based graphics and Go programming language. The game features a bird that the player must navigate through a series of obstacles by jumping at the right time.

## Features:
1. **Terminal Graphics:** The game utilizes terminal graphics to create a simple yet engaging gameplay experience.
2. **Obstacle Generation:** Random obstacles are generated horizontally across the screen at different heights to challenge the player.
3. **Player Controls:** The player controls the bird's jumps using the spacebar key.
4. **Gravity Simulation:** Gravity affects the bird, causing it to fall if not supported by an obstacle.
5. **Game Over Condition:** The game ends when the bird collides with an obstacle or falls below the visible area.

## Code Structure:
- **Main Package:** Controls the game flow, including initialization, game loop, and game over handling.
- **Player Package:** Manages the player's position and movement within the game grid.
- **Map Module Package:** Handles the generation of obstacles and maintains the game grid.
- **Animation Package:** Manages the animation of game elements, including rendering the game grid and player movement.

## How to Play:
1. Run the program.
2. Press the spacebar to make the bird jump and avoid obstacles.
3. Navigate through the obstacles to score points.
4. The game ends if the bird collides with an obstacle or falls below the visible area.

## Development:
- The game leverages goroutines to manage concurrent tasks such as player movement and obstacle generation.
- Terminal manipulation commands are used to achieve graphical output within the terminal window.
- The code is structured into packages for better modularity and maintainability.

## Future Improvements:
1. Implement scoring and high-score tracking.
2. Add more diverse obstacle patterns and difficulty levels.
3. Enhance graphics and animations for a more immersive experience.
4. Introduce sound effects and music to enhance gameplay atmosphere.

## Conclusion:
This Flappy Bird Terminal Game project provides a fun and challenging experience for players within the confines of a terminal window. It demonstrates the capabilities of Go programming language in creating simple games and showcases the potential of terminal-based graphics for game development.

Note: Currently only working on Linux and WSL2.
