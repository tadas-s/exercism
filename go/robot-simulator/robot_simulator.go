package robot

import (
	"fmt"
)

const (
	N Dir = iota
	E
	S
	W
)

// Right returns direction to the right from current direction d
func (d Dir) Right() Dir {
	return (d + 1) % 4
}

// Left returns direction to the left from current direction d
func (d Dir) Left() Dir {
	return (d + 3) % 4
}

func Right() {
	Step1Robot.Dir = Step1Robot.Dir.Right()
}

func Left() {
	Step1Robot.Dir = Step1Robot.Dir.Left()
}

func Advance() {
	nextPos := Pos{
		RU(Step1Robot.X),
		RU(Step1Robot.Y),
	}.Advance(Step1Robot.Dir)

	Step1Robot.X = int(nextPos.Easting)
	Step1Robot.Y = int(nextPos.Northing)
}

func (d Dir) String() string {
	switch d {
	case N:
		return "North"
	case E:
		return "East"
	case S:
		return "South"
	case W:
		return "West"
	default:
		panic("bad direction")
	}
}

// Within returns true if position is within boundaries of the given rectangle
func (pos Pos) Within(r Rect) bool {
	if r.Min.Northing <= pos.Northing && pos.Northing <= r.Max.Northing &&
		r.Min.Easting <= pos.Easting && pos.Easting <= r.Max.Easting {
		return true
	}

	return false
}

// Advance returns the next position when advancing towards the given direction
func (pos Pos) Advance(d Dir) Pos {
	switch d {
	case N:
		pos.Northing++
	case E:
		pos.Easting++
	case S:
		pos.Northing--
	case W:
		pos.Easting--
	}

	return pos
}

func (pos Pos) String() string {
	return fmt.Sprintf("[%d, %d]", pos.Easting, pos.Northing)
}

type Action byte

func StartRobot(command chan Command, action chan Action) {
	for cmd := range command {
		action <- Action(cmd)
	}

	close(action)
}

func Room(extent Rect, robot Step2Robot, action chan Action, report chan Step2Robot) {
	for act := range action {
		switch act {
		case 'A':
			nextPos := robot.Pos.Advance(robot.Dir)

			if nextPos.Within(extent) {
				robot.Pos = nextPos
			}
		case 'R':
			robot.Dir = robot.Dir.Right()
		case 'L':
			robot.Dir = robot.Dir.Left()
		}
	}

	report <- robot
}

type Action3 struct {
	Name   string
	Action Action
}

func StartRobot3(name, script string, action chan Action3, log chan string) {
	if name == "" {
		log <- "blank robot name"
	}

	for _, cmd := range []rune(script) {
		action <- Action3{name, Action(cmd)}
	}

	// Final "robot done" termination message
	action <- Action3{name, Action('!')}
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan []Step3Robot, log chan string) {
	robotsByName := map[string]*Step3Robot{}
	robotsByPosition := map[Pos]*Step3Robot{}
	robotComplete := map[string]bool{}

	defer func() { rep <- robots }()

	for i := range len(robots) {
		if _, exists := robotsByName[robots[i].Name]; exists {
			log <- fmt.Sprintf("duplicate robot with name '%s'", robots[i].Name)
			return
		}

		robotsByName[robots[i].Name] = &robots[i]

		if otherRobot, occupied := robotsByPosition[robots[i].Pos]; occupied {
			log <- fmt.Sprintf(
				"cannot place robot '%s' at %s, it is taken by '%s'",
				robots[i].Name, robots[i].Pos, otherRobot.Name,
			)
		}

		robotsByPosition[robots[i].Pos] = &robots[i]

		robotComplete[robots[i].Name] = false

		if !robots[i].Pos.Within(extent) {
			log <- fmt.Sprintf("robot with name '%s' outside of the given area", robots[i].Name)
			return
		}
	}

	for act := range action {
		robot, exists := robotsByName[act.Name]

		if !exists {
			log <- fmt.Sprintf("robot with name '%s' does not exist", act.Name)
			break
		}

		if act.Action == 'A' {
			nextPos := robot.Pos.Advance(robot.Dir)
			otherRobot, occupied := robotsByPosition[nextPos]

			if !nextPos.Within(extent) {
				log <- fmt.Sprintf("robot '%s' bumped into a wall", robot.Name)
			} else if occupied {
				log <- fmt.Sprintf("robot '%s' smashed into '%s'", robot.Name, otherRobot.Name)
			} else {
				delete(robotsByPosition, robot.Pos)
				robot.Pos = nextPos
				robotsByPosition[robot.Pos] = robot
			}
		} else if act.Action == 'R' {
			robot.Dir = robot.Dir.Right()
		} else if act.Action == 'L' {
			robot.Dir = robot.Dir.Left()
		} else if act.Action == '!' {
			robotComplete[robot.Name] = true
		} else {
			log <- fmt.Sprintf("bad command '%s'", string(act.Action))
			break
		}

		allComplete := true

		for _, complete := range robotComplete {
			if !complete {
				allComplete = false
			}
		}

		if allComplete {
			break
		}
	}
}
