package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Advent of Code 2023 - Day 19")

	inputBytes, _ := os.ReadFile("day_19/part_1/input.txt")
	input := string(inputBytes)

	total := 0

	type part struct {
		x, m, a, s int
		total      int
		approved   bool
	}

	type step struct {
		destination string
		field       string
		comparator  string
		value       int
	}

	type workflow struct {
		name  string
		steps []*step
	}

	workflows := map[string]*workflow{}
	parts := []*part{}

	doingWorkflows := true
	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			if len(workflows) > 0 {
				doingWorkflows = false
			}
			continue
		}

		if doingWorkflows {

			lineParts := strings.Split(line, "{")
			name := lineParts[0]
			steps := strings.Split(strings.TrimSuffix(lineParts[1], "}"), ",")

			thisWorflow := &workflow{name: name}

			for _, s := range steps {
				bits := strings.Split(s, ":")

				if len(bits) == 1 {
					thisWorflow.steps = append(thisWorflow.steps, &step{destination: bits[0]})
					continue
				}

				var field, comparator string
				var value int
				if strings.Contains(bits[0], "<") {
					comparator = "<"
					fv := strings.Split(bits[0], "<")
					field = fv[0]
					value, _ = strconv.Atoi(fv[1])
				}

				if strings.Contains(bits[0], ">") {
					comparator = ">"
					fv := strings.Split(bits[0], ">")
					field = fv[0]
					value, _ = strconv.Atoi(fv[1])
				}

				thisWorflow.steps = append(thisWorflow.steps, &step{
					destination: bits[len(bits)-1],
					field:       field,
					comparator:  comparator,
					value:       value,
				})
			}
			workflows[name] = thisWorflow
		} else {

			l := strings.TrimSuffix(strings.TrimPrefix(line, "{"), "}")
			partsVals := strings.Split(l, ",")

			thisPart := &part{}

			for _, pv := range partsVals {

				ppv := strings.Split(pv, "=")

				val, _ := strconv.Atoi(ppv[1])

				switch ppv[0] {
				case "x":
					thisPart.x = val
				case "m":
					thisPart.m = val
				case "a":
					thisPart.a = val
				case "s":
					thisPart.s = val
				}
			}

			thisPart.total = thisPart.x + thisPart.m + thisPart.a + thisPart.s

			parts = append(parts, thisPart)
		}

	}

A:
	for _, p := range parts {

		wf := "in"
		approvedOrRejected := false

	B:
		for !approvedOrRejected {
			currentWorkflow := workflows[wf]

			for _, s := range currentWorkflow.steps {
				var newDestination string
				switch s.field {
				case "x":
					switch s.comparator {
					case "<":
						if p.x < s.value {
							newDestination = s.destination
						}
					case ">":
						if p.x > s.value {
							newDestination = s.destination
						}
					}
				case "m":
					switch s.comparator {
					case "<":
						if p.m < s.value {
							newDestination = s.destination
						}
					case ">":
						if p.m > s.value {
							newDestination = s.destination
						}
					}
				case "a":
					switch s.comparator {
					case "<":
						if p.a < s.value {
							newDestination = s.destination
						}
					case ">":
						if p.a > s.value {
							newDestination = s.destination
						}
					}
				case "s":
					switch s.comparator {
					case "<":
						if p.s < s.value {
							newDestination = s.destination
						}
					case ">":
						if p.s > s.value {
							newDestination = s.destination
						}
					}
				case "":
					{
						newDestination = s.destination
					}
				}

				if newDestination == "R" {
					approvedOrRejected = true
					continue A
				}

				if newDestination == "A" {
					p.approved = true
					approvedOrRejected = true
					continue A
				}

				if newDestination != "" {
					wf = newDestination
					continue B
				}

			}

		}

	}

	for _, p := range parts {
		if p.approved {
			total += p.total
		}
	}

	fmt.Println("Total:", total)

}
