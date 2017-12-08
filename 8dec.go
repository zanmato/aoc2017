package aoc2017

import (
	"fmt"
	"strings"
)

type instruction struct {
	Register  string
	Modifier  string
	Value     int
	Reference string
	Operand   string
	Condition int
}

func (i *instruction) Run(v int) (int, error) {
	pass, err := i.evaluateCondition(v)
	if err != nil {
		return 0, err
	}

	if !pass {
		return 0, nil
	}

	switch i.Modifier {
	case "inc":
		return +i.Value, nil
	case "dec":
		return -i.Value, nil
	}

	return 0, fmt.Errorf("Unknown modifier: %s", i.Modifier)
}

func (i *instruction) evaluateCondition(v int) (bool, error) {
	c := i.Condition
	switch i.Operand {
	case ">":
		return v > c, nil
	case "<":
		return v < c, nil
	case ">=":
		return v >= c, nil
	case "<=":
		return v <= c, nil
	case "==":
		return v == c, nil
	case "!=":
		return v != c, nil
	}

	return false, fmt.Errorf("Unknown operand: %s", i.Operand)
}

func getLargestValueInRegister(input string) (int, int, error) {
	register := map[string]int{}
	inst := []*instruction{}
	rows := strings.Split(input, "\n")

	for _, row := range rows {
		i := &instruction{}
		_, err := fmt.Sscanf(
			row,
			"%s %s %d if %s %s %d",
			&i.Register,
			&i.Modifier,
			&i.Value,
			&i.Reference,
			&i.Operand,
			&i.Condition,
		)
		if err != nil {
			return 0, 0, err
		}

		register[i.Register] = 0

		inst = append(inst, i)
	}

	m := 0
	for _, i := range inst {
		v, exists := register[i.Reference]
		if !exists {
			return 0, 0, fmt.Errorf("Register '%s' doesn't exist", i.Reference)
		}

		res, err := i.Run(v)
		if err != nil {
			return 0, 0, err
		}

		register[i.Register] += res

		if register[i.Register] > m {
			m = register[i.Register]
		}
	}

	mv := 0
	for _, v := range register {
		if v > mv {
			mv = v
		}
	}

	return mv, m, nil
}
