package notation

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/parser"
	"testing"
)

// TestInfixNotationはノードの中置表記をテストする
func TestInfixNotation(t *testing.T) {
	testcase := []struct {
		input    string
		expected string
	}{
		{"C(1)", "C(1)"},
		{"C(1-(-1))", "C(1-(-1))"},
		{"C(((1+2)+3)+4)", "C(1+2+3+4)"},
		{"C(((1+2)+3)-4)", "C(1+2+3-4)"},
		{"C(((1+2)+3)*4)", "C((1+2+3)*4)"},
		{"C(((1+2)+3)/4)", "C((1+2+3)/4)"},
		{"C(((1+2)-3)+4)", "C(1+2-3+4)"},
		{"C(((1+2)-3)-4)", "C(1+2-3-4)"},
		{"C(((1+2)-3)*4)", "C((1+2-3)*4)"},
		{"C(((1+2)-3)/4)", "C((1+2-3)/4)"},
		{"C(((1+2)*3)+4)", "C((1+2)*3+4)"},
		{"C(((1+2)*3)-4)", "C((1+2)*3-4)"},
		{"C(((1+2)*3)*4)", "C((1+2)*3*4)"},
		{"C(((1+2)*3)/4)", "C((1+2)*3/4)"},
		{"C(((1+2)/3)+4)", "C((1+2)/3+4)"},
		{"C(((1+2)/3)-4)", "C((1+2)/3-4)"},
		{"C(((1+2)/3)*4)", "C((1+2)/3*4)"},
		{"C(((1+2)/3)/4)", "C((1+2)/3/4)"},
		{"C(((1-2)+3)+4)", "C(1-2+3+4)"},
		{"C(((1-2)+3)-4)", "C(1-2+3-4)"},
		{"C(((1-2)+3)*4)", "C((1-2+3)*4)"},
		{"C(((1-2)+3)/4)", "C((1-2+3)/4)"},
		{"C(((1-2)-3)+4)", "C(1-2-3+4)"},
		{"C(((1-2)-3)-4)", "C(1-2-3-4)"},
		{"C(((1-2)-3)*4)", "C((1-2-3)*4)"},
		{"C(((1-2)-3)/4)", "C((1-2-3)/4)"},
		{"C(((1-2)*3)+4)", "C((1-2)*3+4)"},
		{"C(((1-2)*3)-4)", "C((1-2)*3-4)"},
		{"C(((1-2)*3)*4)", "C((1-2)*3*4)"},
		{"C(((1-2)*3)/4)", "C((1-2)*3/4)"},
		{"C(((1-2)/3)+4)", "C((1-2)/3+4)"},
		{"C(((1-2)/3)-4)", "C((1-2)/3-4)"},
		{"C(((1-2)/3)*4)", "C((1-2)/3*4)"},
		{"C(((1-2)/3)/4)", "C((1-2)/3/4)"},
		{"C(((1*2)+3)+4)", "C(1*2+3+4)"},
		{"C(((1*2)+3)-4)", "C(1*2+3-4)"},
		{"C(((1*2)+3)*4)", "C((1*2+3)*4)"},
		{"C(((1*2)+3)/4)", "C((1*2+3)/4)"},
		{"C(((1*2)-3)+4)", "C(1*2-3+4)"},
		{"C(((1*2)-3)-4)", "C(1*2-3-4)"},
		{"C(((1*2)-3)*4)", "C((1*2-3)*4)"},
		{"C(((1*2)-3)/4)", "C((1*2-3)/4)"},
		{"C(((1*2)*3)+4)", "C(1*2*3+4)"},
		{"C(((1*2)*3)-4)", "C(1*2*3-4)"},
		{"C(((1*2)*3)*4)", "C(1*2*3*4)"},
		{"C(((1*2)*3)/4)", "C(1*2*3/4)"},
		{"C(((1*2)/3)+4)", "C(1*2/3+4)"},
		{"C(((1*2)/3)-4)", "C(1*2/3-4)"},
		{"C(((1*2)/3)*4)", "C(1*2/3*4)"},
		{"C(((1*2)/3)/4)", "C(1*2/3/4)"},
		{"C(((1/2)+3)+4)", "C(1/2+3+4)"},
		{"C(((1/2)+3)-4)", "C(1/2+3-4)"},
		{"C(((1/2)+3)*4)", "C((1/2+3)*4)"},
		{"C(((1/2)+3)/4)", "C((1/2+3)/4)"},
		{"C(((1/2)-3)+4)", "C(1/2-3+4)"},
		{"C(((1/2)-3)-4)", "C(1/2-3-4)"},
		{"C(((1/2)-3)*4)", "C((1/2-3)*4)"},
		{"C(((1/2)-3)/4)", "C((1/2-3)/4)"},
		{"C(((1/2)*3)+4)", "C(1/2*3+4)"},
		{"C(((1/2)*3)-4)", "C(1/2*3-4)"},
		{"C(((1/2)*3)*4)", "C(1/2*3*4)"},
		{"C(((1/2)*3)/4)", "C(1/2*3/4)"},
		{"C(((1/2)/3)+4)", "C(1/2/3+4)"},
		{"C(((1/2)/3)-4)", "C(1/2/3-4)"},
		{"C(((1/2)/3)*4)", "C(1/2/3*4)"},
		{"C(((1/2)/3)/4)", "C(1/2/3/4)"},
		{"C((1+(2+3))+4)", "C(1+2+3+4)"},
		{"C((1+(2+3))-4)", "C(1+2+3-4)"},
		{"C((1+(2+3))*4)", "C((1+2+3)*4)"},
		{"C((1+(2+3))/4)", "C((1+2+3)/4)"},
		{"C((1+(2-3))+4)", "C(1+2-3+4)"},
		{"C((1+(2-3))-4)", "C(1+2-3-4)"},
		{"C((1+(2-3))*4)", "C((1+2-3)*4)"},
		{"C((1+(2-3))/4)", "C((1+2-3)/4)"},
		{"C((1+(2*3))+4)", "C(1+2*3+4)"},
		{"C((1+(2*3))-4)", "C(1+2*3-4)"},
		{"C((1+(2*3))*4)", "C((1+2*3)*4)"},
		{"C((1+(2*3))/4)", "C((1+2*3)/4)"},
		{"C((1+(2/3))+4)", "C(1+2/3+4)"},
		{"C((1+(2/3))-4)", "C(1+2/3-4)"},
		{"C((1+(2/3))*4)", "C((1+2/3)*4)"},
		{"C((1+(2/3))/4)", "C((1+2/3)/4)"},
		{"C((1-(2+3))+4)", "C(1-(2+3)+4)"},
		{"C((1-(2+3))-4)", "C(1-(2+3)-4)"},
		{"C((1-(2+3))*4)", "C((1-(2+3))*4)"},
		{"C((1-(2+3))/4)", "C((1-(2+3))/4)"},
		{"C((1-(2-3))+4)", "C(1-(2-3)+4)"},
		{"C((1-(2-3))-4)", "C(1-(2-3)-4)"},
		{"C((1-(2-3))*4)", "C((1-(2-3))*4)"},
		{"C((1-(2-3))/4)", "C((1-(2-3))/4)"},
		{"C((1-(2*3))+4)", "C(1-2*3+4)"},
		{"C((1-(2*3))-4)", "C(1-2*3-4)"},
		{"C((1-(2*3))*4)", "C((1-2*3)*4)"},
		{"C((1-(2*3))/4)", "C((1-2*3)/4)"},
		{"C((1-(2/3))+4)", "C(1-2/3+4)"},
		{"C((1-(2/3))-4)", "C(1-2/3-4)"},
		{"C((1-(2/3))*4)", "C((1-2/3)*4)"},
		{"C((1-(2/3))/4)", "C((1-2/3)/4)"},
		{"C((1*(2+3))+4)", "C(1*(2+3)+4)"},
		{"C((1*(2+3))-4)", "C(1*(2+3)-4)"},
		{"C((1*(2+3))*4)", "C(1*(2+3)*4)"},
		{"C((1*(2+3))/4)", "C(1*(2+3)/4)"},
		{"C((1*(2-3))+4)", "C(1*(2-3)+4)"},
		{"C((1*(2-3))-4)", "C(1*(2-3)-4)"},
		{"C((1*(2-3))*4)", "C(1*(2-3)*4)"},
		{"C((1*(2-3))/4)", "C(1*(2-3)/4)"},
		{"C((1*(2*3))+4)", "C(1*2*3+4)"},
		{"C((1*(2*3))-4)", "C(1*2*3-4)"},
		{"C((1*(2*3))*4)", "C(1*2*3*4)"},
		{"C((1*(2*3))/4)", "C(1*2*3/4)"},
		{"C((1*(2/3))+4)", "C(1*2/3+4)"},
		{"C((1*(2/3))-4)", "C(1*2/3-4)"},
		{"C((1*(2/3))*4)", "C(1*2/3*4)"},
		{"C((1*(2/3))/4)", "C(1*2/3/4)"},
		{"C((1/(2+3))+4)", "C(1/(2+3)+4)"},
		{"C((1/(2+3))-4)", "C(1/(2+3)-4)"},
		{"C((1/(2+3))*4)", "C(1/(2+3)*4)"},
		{"C((1/(2+3))/4)", "C(1/(2+3)/4)"},
		{"C((1/(2-3))+4)", "C(1/(2-3)+4)"},
		{"C((1/(2-3))-4)", "C(1/(2-3)-4)"},
		{"C((1/(2-3))*4)", "C(1/(2-3)*4)"},
		{"C((1/(2-3))/4)", "C(1/(2-3)/4)"},
		{"C((1/(2*3))+4)", "C(1/(2*3)+4)"},
		{"C((1/(2*3))-4)", "C(1/(2*3)-4)"},
		{"C((1/(2*3))*4)", "C(1/(2*3)*4)"},
		{"C((1/(2*3))/4)", "C(1/(2*3)/4)"},
		{"C((1/(2/3))+4)", "C(1/(2/3)+4)"},
		{"C((1/(2/3))-4)", "C(1/(2/3)-4)"},
		{"C((1/(2/3))*4)", "C(1/(2/3)*4)"},
		{"C((1/(2/3))/4)", "C(1/(2/3)/4)"},
		{"C((1+2)+(3+4))", "C(1+2+3+4)"},
		{"C((1+2)+(3-4))", "C(1+2+3-4)"},
		{"C((1+2)+(3*4))", "C(1+2+3*4)"},
		{"C((1+2)+(3/4))", "C(1+2+3/4)"},
		{"C((1+2)-(3+4))", "C(1+2-(3+4))"},
		{"C((1+2)-(3-4))", "C(1+2-(3-4))"},
		{"C((1+2)-(3*4))", "C(1+2-3*4)"},
		{"C((1+2)-(3/4))", "C(1+2-3/4)"},
		{"C((1+2)*(3+4))", "C((1+2)*(3+4))"},
		{"C((1+2)*(3-4))", "C((1+2)*(3-4))"},
		{"C((1+2)*(3*4))", "C((1+2)*3*4)"},
		{"C((1+2)*(3/4))", "C((1+2)*3/4)"},
		{"C((1+2)/(3+4))", "C((1+2)/(3+4))"},
		{"C((1+2)/(3-4))", "C((1+2)/(3-4))"},
		{"C((1+2)/(3*4))", "C((1+2)/(3*4))"},
		{"C((1+2)/(3/4))", "C((1+2)/(3/4))"},
		{"C((1-2)+(3+4))", "C(1-2+3+4)"},
		{"C((1-2)+(3-4))", "C(1-2+3-4)"},
		{"C((1-2)+(3*4))", "C(1-2+3*4)"},
		{"C((1-2)+(3/4))", "C(1-2+3/4)"},
		{"C((1-2)-(3+4))", "C(1-2-(3+4))"},
		{"C((1-2)-(3-4))", "C(1-2-(3-4))"},
		{"C((1-2)-(3*4))", "C(1-2-3*4)"},
		{"C((1-2)-(3/4))", "C(1-2-3/4)"},
		{"C((1-2)*(3+4))", "C((1-2)*(3+4))"},
		{"C((1-2)*(3-4))", "C((1-2)*(3-4))"},
		{"C((1-2)*(3*4))", "C((1-2)*3*4)"},
		{"C((1-2)*(3/4))", "C((1-2)*3/4)"},
		{"C((1-2)/(3+4))", "C((1-2)/(3+4))"},
		{"C((1-2)/(3-4))", "C((1-2)/(3-4))"},
		{"C((1-2)/(3*4))", "C((1-2)/(3*4))"},
		{"C((1-2)/(3/4))", "C((1-2)/(3/4))"},
		{"C((1*2)+(3+4))", "C(1*2+3+4)"},
		{"C((1*2)+(3-4))", "C(1*2+3-4)"},
		{"C((1*2)+(3*4))", "C(1*2+3*4)"},
		{"C((1*2)+(3/4))", "C(1*2+3/4)"},
		{"C((1*2)-(3+4))", "C(1*2-(3+4))"},
		{"C((1*2)-(3-4))", "C(1*2-(3-4))"},
		{"C((1*2)-(3*4))", "C(1*2-3*4)"},
		{"C((1*2)-(3/4))", "C(1*2-3/4)"},
		{"C((1*2)*(3+4))", "C(1*2*(3+4))"},
		{"C((1*2)*(3-4))", "C(1*2*(3-4))"},
		{"C((1*2)*(3*4))", "C(1*2*3*4)"},
		{"C((1*2)*(3/4))", "C(1*2*3/4)"},
		{"C((1*2)/(3+4))", "C(1*2/(3+4))"},
		{"C((1*2)/(3-4))", "C(1*2/(3-4))"},
		{"C((1*2)/(3*4))", "C(1*2/(3*4))"},
		{"C((1*2)/(3/4))", "C(1*2/(3/4))"},
		{"C((1/2)+(3+4))", "C(1/2+3+4)"},
		{"C((1/2)+(3-4))", "C(1/2+3-4)"},
		{"C((1/2)+(3*4))", "C(1/2+3*4)"},
		{"C((1/2)+(3/4))", "C(1/2+3/4)"},
		{"C((1/2)-(3+4))", "C(1/2-(3+4))"},
		{"C((1/2)-(3-4))", "C(1/2-(3-4))"},
		{"C((1/2)-(3*4))", "C(1/2-3*4)"},
		{"C((1/2)-(3/4))", "C(1/2-3/4)"},
		{"C((1/2)*(3+4))", "C(1/2*(3+4))"},
		{"C((1/2)*(3-4))", "C(1/2*(3-4))"},
		{"C((1/2)*(3*4))", "C(1/2*3*4)"},
		{"C((1/2)*(3/4))", "C(1/2*3/4)"},
		{"C((1/2)/(3+4))", "C(1/2/(3+4))"},
		{"C((1/2)/(3-4))", "C(1/2/(3-4))"},
		{"C((1/2)/(3*4))", "C(1/2/(3*4))"},
		{"C((1/2)/(3/4))", "C(1/2/(3/4))"},
		{"C(1+((2+3)+4))", "C(1+2+3+4)"},
		{"C(1+((2+3)-4))", "C(1+2+3-4)"},
		{"C(1+((2+3)*4))", "C(1+(2+3)*4)"},
		{"C(1+((2+3)/4))", "C(1+(2+3)/4)"},
		{"C(1+((2-3)+4))", "C(1+2-3+4)"},
		{"C(1+((2-3)-4))", "C(1+2-3-4)"},
		{"C(1+((2-3)*4))", "C(1+(2-3)*4)"},
		{"C(1+((2-3)/4))", "C(1+(2-3)/4)"},
		{"C(1+((2*3)+4))", "C(1+2*3+4)"},
		{"C(1+((2*3)-4))", "C(1+2*3-4)"},
		{"C(1+((2*3)*4))", "C(1+2*3*4)"},
		{"C(1+((2*3)/4))", "C(1+2*3/4)"},
		{"C(1+((2/3)+4))", "C(1+2/3+4)"},
		{"C(1+((2/3)-4))", "C(1+2/3-4)"},
		{"C(1+((2/3)*4))", "C(1+2/3*4)"},
		{"C(1+((2/3)/4))", "C(1+2/3/4)"},
		{"C(1-((2+3)+4))", "C(1-(2+3+4))"},
		{"C(1-((2+3)-4))", "C(1-(2+3-4))"},
		{"C(1-((2+3)*4))", "C(1-(2+3)*4)"},
		{"C(1-((2+3)/4))", "C(1-(2+3)/4)"},
		{"C(1-((2-3)+4))", "C(1-(2-3+4))"},
		{"C(1-((2-3)-4))", "C(1-(2-3-4))"},
		{"C(1-((2-3)*4))", "C(1-(2-3)*4)"},
		{"C(1-((2-3)/4))", "C(1-(2-3)/4)"},
		{"C(1-((2*3)+4))", "C(1-(2*3+4))"},
		{"C(1-((2*3)-4))", "C(1-(2*3-4))"},
		{"C(1-((2*3)*4))", "C(1-2*3*4)"},
		{"C(1-((2*3)/4))", "C(1-2*3/4)"},
		{"C(1-((2/3)+4))", "C(1-(2/3+4))"},
		{"C(1-((2/3)-4))", "C(1-(2/3-4))"},
		{"C(1-((2/3)*4))", "C(1-2/3*4)"},
		{"C(1-((2/3)/4))", "C(1-2/3/4)"},
		{"C(1*((2+3)+4))", "C(1*(2+3+4))"},
		{"C(1*((2+3)-4))", "C(1*(2+3-4))"},
		{"C(1*((2+3)*4))", "C(1*(2+3)*4)"},
		{"C(1*((2+3)/4))", "C(1*(2+3)/4)"},
		{"C(1*((2-3)+4))", "C(1*(2-3+4))"},
		{"C(1*((2-3)-4))", "C(1*(2-3-4))"},
		{"C(1*((2-3)*4))", "C(1*(2-3)*4)"},
		{"C(1*((2-3)/4))", "C(1*(2-3)/4)"},
		{"C(1*((2*3)+4))", "C(1*(2*3+4))"},
		{"C(1*((2*3)-4))", "C(1*(2*3-4))"},
		{"C(1*((2*3)*4))", "C(1*2*3*4)"},
		{"C(1*((2*3)/4))", "C(1*2*3/4)"},
		{"C(1*((2/3)+4))", "C(1*(2/3+4))"},
		{"C(1*((2/3)-4))", "C(1*(2/3-4))"},
		{"C(1*((2/3)*4))", "C(1*2/3*4)"},
		{"C(1*((2/3)/4))", "C(1*2/3/4)"},
		{"C(1/((2+3)+4))", "C(1/(2+3+4))"},
		{"C(1/((2+3)-4))", "C(1/(2+3-4))"},
		{"C(1/((2+3)*4))", "C(1/((2+3)*4))"},
		{"C(1/((2+3)/4))", "C(1/((2+3)/4))"},
		{"C(1/((2-3)+4))", "C(1/(2-3+4))"},
		{"C(1/((2-3)-4))", "C(1/(2-3-4))"},
		{"C(1/((2-3)*4))", "C(1/((2-3)*4))"},
		{"C(1/((2-3)/4))", "C(1/((2-3)/4))"},
		{"C(1/((2*3)+4))", "C(1/(2*3+4))"},
		{"C(1/((2*3)-4))", "C(1/(2*3-4))"},
		{"C(1/((2*3)*4))", "C(1/(2*3*4))"},
		{"C(1/((2*3)/4))", "C(1/(2*3/4))"},
		{"C(1/((2/3)+4))", "C(1/(2/3+4))"},
		{"C(1/((2/3)-4))", "C(1/(2/3-4))"},
		{"C(1/((2/3)*4))", "C(1/(2/3*4))"},
		{"C(1/((2/3)/4))", "C(1/(2/3/4))"},
		{"C(1+(2+(3+4)))", "C(1+2+3+4)"},
		{"C(1+(2+(3-4)))", "C(1+2+3-4)"},
		{"C(1+(2+(3*4)))", "C(1+2+3*4)"},
		{"C(1+(2+(3/4)))", "C(1+2+3/4)"},
		{"C(1+(2-(3+4)))", "C(1+2-(3+4))"},
		{"C(1+(2-(3-4)))", "C(1+2-(3-4))"},
		{"C(1+(2-(3*4)))", "C(1+2-3*4)"},
		{"C(1+(2-(3/4)))", "C(1+2-3/4)"},
		{"C(1+(2*(3+4)))", "C(1+2*(3+4))"},
		{"C(1+(2*(3-4)))", "C(1+2*(3-4))"},
		{"C(1+(2*(3*4)))", "C(1+2*3*4)"},
		{"C(1+(2*(3/4)))", "C(1+2*3/4)"},
		{"C(1+(2/(3+4)))", "C(1+2/(3+4))"},
		{"C(1+(2/(3-4)))", "C(1+2/(3-4))"},
		{"C(1+(2/(3*4)))", "C(1+2/(3*4))"},
		{"C(1+(2/(3/4)))", "C(1+2/(3/4))"},
		{"C(1-(2+(3+4)))", "C(1-(2+3+4))"},
		{"C(1-(2+(3-4)))", "C(1-(2+3-4))"},
		{"C(1-(2+(3*4)))", "C(1-(2+3*4))"},
		{"C(1-(2+(3/4)))", "C(1-(2+3/4))"},
		{"C(1-(2-(3+4)))", "C(1-(2-(3+4)))"},
		{"C(1-(2-(3-4)))", "C(1-(2-(3-4)))"},
		{"C(1-(2-(3*4)))", "C(1-(2-3*4))"},
		{"C(1-(2-(3/4)))", "C(1-(2-3/4))"},
		{"C(1-(2*(3+4)))", "C(1-2*(3+4))"},
		{"C(1-(2*(3-4)))", "C(1-2*(3-4))"},
		{"C(1-(2*(3*4)))", "C(1-2*3*4)"},
		{"C(1-(2*(3/4)))", "C(1-2*3/4)"},
		{"C(1-(2/(3+4)))", "C(1-2/(3+4))"},
		{"C(1-(2/(3-4)))", "C(1-2/(3-4))"},
		{"C(1-(2/(3*4)))", "C(1-2/(3*4))"},
		{"C(1-(2/(3/4)))", "C(1-2/(3/4))"},
		{"C(1*(2+(3+4)))", "C(1*(2+3+4))"},
		{"C(1*(2+(3-4)))", "C(1*(2+3-4))"},
		{"C(1*(2+(3*4)))", "C(1*(2+3*4))"},
		{"C(1*(2+(3/4)))", "C(1*(2+3/4))"},
		{"C(1*(2-(3+4)))", "C(1*(2-(3+4)))"},
		{"C(1*(2-(3-4)))", "C(1*(2-(3-4)))"},
		{"C(1*(2-(3*4)))", "C(1*(2-3*4))"},
		{"C(1*(2-(3/4)))", "C(1*(2-3/4))"},
		{"C(1*(2*(3+4)))", "C(1*2*(3+4))"},
		{"C(1*(2*(3-4)))", "C(1*2*(3-4))"},
		{"C(1*(2*(3*4)))", "C(1*2*3*4)"},
		{"C(1*(2*(3/4)))", "C(1*2*3/4)"},
		{"C(1*(2/(3+4)))", "C(1*2/(3+4))"},
		{"C(1*(2/(3-4)))", "C(1*2/(3-4))"},
		{"C(1*(2/(3*4)))", "C(1*2/(3*4))"},
		{"C(1*(2/(3/4)))", "C(1*2/(3/4))"},
		{"C(1/(2+(3+4)))", "C(1/(2+3+4))"},
		{"C(1/(2+(3-4)))", "C(1/(2+3-4))"},
		{"C(1/(2+(3*4)))", "C(1/(2+3*4))"},
		{"C(1/(2+(3/4)))", "C(1/(2+3/4))"},
		{"C(1/(2-(3+4)))", "C(1/(2-(3+4)))"},
		{"C(1/(2-(3-4)))", "C(1/(2-(3-4)))"},
		{"C(1/(2-(3*4)))", "C(1/(2-3*4))"},
		{"C(1/(2-(3/4)))", "C(1/(2-3/4))"},
		{"C(1/(2*(3+4)))", "C(1/(2*(3+4)))"},
		{"C(1/(2*(3-4)))", "C(1/(2*(3-4)))"},
		{"C(1/(2*(3*4)))", "C(1/(2*3*4))"},
		{"C(1/(2*(3/4)))", "C(1/(2*3/4))"},
		{"C(1/(2/(3+4)))", "C(1/(2/(3+4)))"},
		{"C(1/(2/(3-4)))", "C(1/(2/(3-4)))"},
		{"C(1/(2/(3*4)))", "C(1/(2/(3*4)))"},
		{"C(1/(2/(3/4)))", "C(1/(2/(3/4)))"},
	}

	for _, test := range testcase {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			ast, parseErr := parser.Parse(test.input)
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			actual, notationErr := InfixNotation(ast)
			if notationErr != nil {
				t.Fatalf("中置記法表記生成エラー: %s", notationErr)
				return
			}

			if actual != test.expected {
				t.Fatalf("got %q, want %q", actual, test.expected)
			}
		})
	}
}
