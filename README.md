# Soccer Simulation Deep Learning Multithreading

This project aims to experiment with Reinforced Deep Q-Learning algorithms and techniques in combination with multithreading to see if it is feasible, efficient, and effective. Model building is computationally extensive and time consuming. Thus, multithreading possibly introduces a means to improve this issue. 

To build the game environment and neural network, I decided to use Golang. Python, being the most popular, is quite slow. C/C++ is fast, but not memory safe. To implement multithreading in a productive, fast, and memory safe way made Golang optimal.

### Environment (WIP)

The environment for which models will be generated for is a soccer shooting simulation. Given any location on the soccer field, the model needs to score a goal to be rewarded. 

##### Action Space (WIP)

Horizontal Angle (radians)
Vertical Angle   (radians)
Power            (feet per second)

##### Observation Space (WIP)

Field Position X (feet)
Field Position Y (feet)

##### Rewards (WIP)

If a goal is scored, 10.0 points are given.
If the shot hits the crossbar or post, -1.0 points are given.
If the shot misses entirely, -20.0 points are given.

### Neural Network (WIP)

DEEP Q-LEARNING INFO HERE

##### Layers (WIP)

FC layer, flattened layer
