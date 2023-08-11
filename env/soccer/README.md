# Go-ML env/soccer

[![Documentation](https://img.shields.io/badge/documentation-GoDoc-blue.svg)](https://pkg.go.dev/github.com/thadUra/Go-ML/env/soccer)

Package soccer is a custom environment built using the Environment interface. As of now, this environment is not stable for actual utilization.

The goal of this environment is to attempt to train several different models synchronously to mimic an actual soccer match. Methods include using reinforcement learning algorithms and other machine learning techniques to achive this. As of now, environment capabilities consist of just dribbling in either 8 directions on a field and shooting the ball. The shooting calculation takes account for rolling, bouncing, and other related kinematic calculations to determine a goal or not.

In order to train an environment with such a high observation space will require a deep reinforcement learning algorithm like deep q-learning, but it has not been implemented yet in this library.

<!-- ## Action Space

WIP

## Observation Space

WIP

## Rewards

WIP

## Arguments

WIP -->