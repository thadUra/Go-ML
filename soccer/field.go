package soccer

import (
	"errors"
	"fmt"
	"math"
)

/**
 *  Field Struct
 *  Information about the field, net, and ball dimensions
 *  Distances are in feet and grams
 */
type Field struct {
	FIELD_WIDTH   float64
	FIELD_HEIGHT  float64
	NET_WIDTH     float64
	NET_HEIGHT    float64
	NET_DIAMETER  float64
	BALL_WEIGHT   float64
	BALL_DIAMETER float64
}

/**
 *  GenerateField() Func
 *  Generate the field via parameters
 *  Default to fifa regulations
 */
func GenerateField(
	field_width float64,
	field_height float64,
	net_width float64,
	net_height float64,
	net_diameter float64,
	weight float64,
	ball_diameter float64,
	def bool,
) Field {
	var f Field
	if def {
		f.FIELD_HEIGHT = 345.0
		f.FIELD_WIDTH = 224.0
		f.NET_WIDTH = 24.0
		f.NET_HEIGHT = 8.0
		f.NET_DIAMETER = 1.0 / 3.0
		f.BALL_WEIGHT = 430.0
		f.BALL_DIAMETER = 0.75
	} else {
		f.FIELD_HEIGHT = field_height
		f.FIELD_WIDTH = field_width
		f.NET_WIDTH = net_width
		f.NET_HEIGHT = net_height
		f.NET_DIAMETER = net_diameter
		f.BALL_WEIGHT = weight
		f.BALL_DIAMETER = ball_diameter
	}
	return f
}

/**
 *  GetShotParameterLimits() Func
 *  Return the range of horizontal/vertical angles and power of the shot
 */
func (f Field) GetShotParameterLimits() [][]float64 {
	// Min and max values for parameter
	horizontal_angle := 80.0 * math.Pi / 180.0 // 160 degrees
	vertical_angle := 90.0 * math.Pi / 180.0   // 90 degrees
	power := 150.0                             // feet per second

	// Generate limitations
	var limitations [][]float64
	limitations = append(limitations, []float64{-horizontal_angle, horizontal_angle})
	limitations = append(limitations, []float64{0, vertical_angle})
	limitations = append(limitations, []float64{0, power})
	return limitations
}

/**
 *  Shoot() Func
 *  Perform shooting action of ball given position, horizontal angle, vertical angle, and power
 *  Accounts for bouncing, rolling, friction, and energy loss
 *  Calculates if shot is a goal, miss, or hits the post/crossbar
 */
func (f Field) Shoot(
	pos Position,
	parameters []float64,
	debug bool,
) (string, error) {
	// Parameters for shot
	horizontal_angle := parameters[0]
	vertical_angle := parameters[1]
	power := parameters[2]

	// Check shot parameters
	limitations := f.GetShotParameterLimits()
	if horizontal_angle < limitations[0][0] || horizontal_angle > limitations[0][1] {
		return "", errors.New("field.Shoot: horizontal angle not in range")
	} else if vertical_angle < limitations[1][0] || vertical_angle > limitations[1][1] {
		return "", errors.New("field.Shoot: vertical angle not in range")
	} else if power < limitations[2][0] || power > limitations[2][1] {
		return "", errors.New("field.Shoot: power not in range")
	}

	// Parameters for goal
	max_y := f.NET_HEIGHT - (f.BALL_DIAMETER / 2)
	min_y := (f.BALL_DIAMETER / 2)
	max_x := (f.FIELD_WIDTH / 2) + (f.NET_WIDTH / 2) - (f.BALL_DIAMETER / 2)
	min_x := (f.FIELD_WIDTH / 2) - (f.NET_WIDTH / 2) + (f.BALL_DIAMETER / 2)
	max_duration := 2.1 // 0.6 seconds to react + 1.5 seconds to act and save shot

	// Calculate ball position at goal line
	gravity := -32.17
	length := pos.DISTANCE_Y / math.Cos(horizontal_angle)
	velocity_x := math.Cos(vertical_angle) * power
	velocity_y := math.Cos((90*math.Pi/180)-vertical_angle) * power
	duration := length / velocity_x
	height := (velocity_y * duration) + (0.5 * gravity * duration * duration)
	width := pos.DISTANCE_X + (math.Sin(horizontal_angle) * length)

	// Print statements
	if debug {
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

	// Determine if ball bounces or rolls during shot (WIP)
	if height < min_y && duration < max_duration*3 {

		// Helper vars for bounces and rolling
		new_duration := 0.0
		current_distance := 0.0
		rolling := false
		coeff_restitution_y := 0.8
		coeff_restitution_x := 0.95
		coeff_friction := 0.55

		// Calculate bounces
		if debug {
			fmt.Println("===BOUNCE CALCULATIONS===")
		}
		for (current_distance < length) && !rolling {
			time := (-1 * velocity_y) / (0.5 * gravity)
			peak := (velocity_y * time / 2) + (0.5 * gravity * time * time / 4)
			new_duration += time
			current_distance += velocity_x * time
			if current_distance >= length {
				// Change so its not using duration
				goal_bounce_time := (length - (current_distance - velocity_x*time)) / velocity_x
				new_duration += goal_bounce_time - time
				height = (velocity_y * goal_bounce_time) + (0.5 * gravity * goal_bounce_time * goal_bounce_time)
			}
			if peak <= f.BALL_DIAMETER*0.6 {
				rolling = true
				height = f.BALL_DIAMETER / 2
			}
			if debug {
				fmt.Printf("	Path Equation: (%ft+%f,(%ft - 16.085t^2)) %f\n", velocity_x, current_distance-(velocity_x*time), velocity_y, time)
			}
			// Update velocity after bounce
			velocity_y *= coeff_restitution_y
			velocity_x *= coeff_restitution_x
		}
		if debug {
			fmt.Printf("	New height with bounces  : %f\n", height)
			fmt.Printf("	Current dist with bounces: %f\n\n", current_distance)
		}

		// Calculate duration for rolling
		if (current_distance < length) && rolling {
			/* Derivation for time
			 * vf^2 = vi^2 + 2ad
			 * vf = sqrt(vi^2 + 2ad)
			 * vf = vi + at
			 * t = (sqrt(vi^2 + 2ad)-vi / a)
			 */
			roll_distance := length - current_distance
			final_vel := math.Sqrt(velocity_x*velocity_x + (2 * coeff_friction * gravity * roll_distance))
			test_time := (final_vel - velocity_x) / (coeff_friction * gravity)
			new_duration += test_time
			if debug {
				fmt.Println("===ROLL CALCULATIONS===")
				fmt.Printf("	ROLL DIST: %f\n", roll_distance)
				fmt.Printf("	FINAL VEL: %f\n", final_vel)
				fmt.Printf("	ROLL TIME: %f\n\n", test_time)
			}
		}

		// Update duration after bouncing and rolling
		if debug {
			fmt.Println("===DURATION CALCULATIONS===")
			fmt.Printf("	Old Duration: %f\n", duration)
			fmt.Printf("	New Duration: %f\n\n", new_duration)
		}
		duration = new_duration
	}

	// Determine result if GOAL, SAVED, CROSSBAR, POST, or MISS
	if height >= min_y &&
		height <= max_y &&
		width >= min_x &&
		width <= max_x {
		if duration < max_duration {
			return "GOAL", nil
		} else {
			return "SAVED", nil
		}
	} else if width >= min_x &&
		width <= max_x &&
		height >= min_y &&
		height <= (max_y+f.NET_DIAMETER+f.BALL_DIAMETER) {
		return "CROSSBAR", nil
	} else if height >= min_y &&
		height <= max_y &&
		((width <= max_x && width >= (min_x-f.NET_DIAMETER-f.BALL_DIAMETER)) ||
			(width >= min_x && width <= (max_x+f.NET_DIAMETER+f.BALL_DIAMETER))) {
		return "POST", nil
	} else {
		return "MISS", nil
	}
}
