package game

import (
	"fmt"
	"math"
)

type Environment struct {
	args              Soccer
	ACTION_SPACE      [][]float64
	OBSERVATION_SPACE [][]float64
}

func InitEnvironment(args Soccer) Environment {

	horizontal_angle := 90.0 * math.Pi / 180.0 // radians
	vertical_angle := 90.0 * math.Pi / 180.0   // radians
	power := 145.0                             // feet per second

	var env Environment
	env.args = args
	env.ACTION_SPACE = append(env.ACTION_SPACE, []float64{0, vertical_angle})
	env.ACTION_SPACE = append(env.ACTION_SPACE, []float64{-horizontal_angle, horizontal_angle})
	env.ACTION_SPACE = append(env.ACTION_SPACE, []float64{0, power})
	env.OBSERVATION_SPACE = append(env.OBSERVATION_SPACE, []float64{0, float64(env.args.FIELD_WIDTH)})
	env.OBSERVATION_SPACE = append(env.OBSERVATION_SPACE, []float64{0, float64(env.args.FIELD_HEIGHT)})
	return env
}

func (env Environment) Step(action []float64, location Shot, print bool) int {

	// Params for shot
	horizontal_angle := action[0]
	vertical_angle := action[1]
	power := action[2]

	// Calculate ball position along goal line
	gravity := -32.17
	length := location.DISTANCE_Y / math.Cos(horizontal_angle)
	velocity_x := math.Cos(vertical_angle) * power
	velocity_y := math.Cos((90*math.Pi/180)-vertical_angle) * power
	duration := length / velocity_x
	height := (velocity_y * duration) + (0.5 * gravity * duration * duration)
	width := location.DISTANCE_X + (math.Sin(horizontal_angle) * length)

	// Determine physics for bounce on ground (WIP)
	coeff_restitution := 0.8
	current_vel_y := velocity_y
	goal_bounce_time := 0.0
	if height < 0 && duration < 5 {
		fmt.Println("===BOUNCE CALCULATIONS===")
		bounceDuration := 0.0
		count := 0
		peak := env.args.BALL_DIAMETER + 1
		for (bounceDuration < duration) && (peak > env.args.BALL_DIAMETER*0.6) {
			time := (-1 * current_vel_y) / (0.5 * gravity)
			bounceDuration += time
			if bounceDuration >= duration {
				goal_bounce_time = bounceDuration - duration
			}
			peak = (current_vel_y * time / 2) + (0.5 * gravity * time * time / 4)
			fmt.Printf("	Bounce %d with time: %f with peak: %f\n", count, bounceDuration, peak)
			fmt.Printf("	Path Equation: (%ft+%f,(%ft - 16.085t^2)) %f\n", velocity_x, (bounceDuration-time)*velocity_x, current_vel_y, time)
			current_vel_y *= coeff_restitution
			count++
		}
		bounceHeight := (current_vel_y * goal_bounce_time) + (0.5 * gravity * goal_bounce_time * goal_bounce_time)
		fmt.Printf("	New height with bounces: %f\n", bounceHeight)
	}

	// Determine physics for rolling (WIP)

	// Determine dimensions for goal with Soccer struct params
	max_y := env.args.NET_HEIGHT - (env.args.BALL_DIAMETER / 2)
	min_y := (env.args.BALL_DIAMETER / 2)
	max_x := (env.args.FIELD_WIDTH / 2) + (env.args.NET_WIDTH / 2) - (env.args.BALL_DIAMETER / 2)
	min_x := (env.args.FIELD_WIDTH / 2) - (env.args.NET_WIDTH / 2) + (env.args.BALL_DIAMETER / 2)

	// Print statements
	if print {
		fmt.Println("===SHOT CALCULATIONS===")
		fmt.Printf("	Len     : %f\n", length)
		fmt.Printf("	Vel_X   : %f\n", velocity_x)
		fmt.Printf("	Vel_Y   : %f\n", velocity_y)
		fmt.Printf("	Duration: %f\n", duration)
		fmt.Printf("	Height  : %f\n", height)
		fmt.Printf("	Width   : %f\n\n", width)
		fmt.Println("===NET CALCULATIONS===")
		fmt.Printf("	Max_Y   : %f\n", max_y)
		fmt.Printf("	Min_Y   : %f\n", min_y)
		fmt.Printf("	Max_X   : %f\n", max_x)
		fmt.Printf("	Min_X   : %f\n", min_x)
		fmt.Printf("	Path Equation: (%ft,(%ft - 16.085t^2)) %f\n\n", velocity_x, velocity_y, duration)
	}

	// Determine reward (WIP)
	// Include hitting left or right post + crossbar
	if height >= min_y && height <= max_y && width >= min_x && width <= max_x {
		fmt.Println("GOAL")
		return 10
	} else {
		fmt.Println("MISS")
		return -20
	}
}
