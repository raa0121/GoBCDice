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
		{"C(-1)", "C(-1)"},
		{"C(1-(-1))", "C(1-(-1))"},
		{"C(-1+2)", "C(-1+2)"},
		{"C((-1)*2)", "C(-1*2)"},
		{"C((-1)*(-2))", "C(-1*(-2))"},
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
		{"C(1/2u)", "C(1/2U)"},
		{"C(1/2r)", "C(1/2R)"},
		{"C(100/(1+2)u)", "C(100/(1+2)U)"},
		{"C(100/(1+2)r)", "C(100/(1+2)R)"},
		{"2D6", "2D6"},
		{"12D60", "12D60"},
		{"-2D6", "-2D6"},
		{"+2D6", "2D6"},
		{"2D6+1", "2D6+1"},
		{"1+2D6", "1+2D6"},
		{"-2D6+1", "-2D6+1"},
		{"+2D6+1", "2D6+1"},
		{"2d6+1-1-2-3-4", "2D6+1-1-2-3-4"},
		{"2D6+4D10", "2D6+4D10"},
		{"(2D6)", "2D6"},
		{"-(2D6)", "-2D6"},
		{"+(2D6)", "2D6"},
		{"2d6*3", "2D6*3"},
		{"2d6/2", "2D6/2"},
		{"2d6/2u", "2D6/2U"},
		{"2d6/2r", "2D6/2R"},
		{"100/2d6+1", "100/2D6+1"},
		{"100/2d6u+1", "100/2D6U+1"},
		{"100/2d6r+1", "100/2D6R+1"},
		{"100/(2d6+1)+4*5", "100/(2D6+1)+4*5"},
		{"100/(2d6+1)u+4*5", "100/(2D6+1)U+4*5"},
		{"100/(2d6+1)r+4*5", "100/(2D6+1)R+4*5"},
		{"4d10/2d6+1", "4D10/2D6+1"},
		{"4d10/2d6u+1", "4D10/2D6U+1"},
		{"4d10/2d6r+1", "4D10/2D6R+1"},
		{"2d10+3-4", "2D10+3-4"},
		{"2d10+3*4", "2D10+3*4"},
		{"2d10/3+4*5-6", "2D10/3+4*5-6"},
		{"2d10/3u+4*5-6", "2D10/3U+4*5-6"},
		{"2d10/3r+4*5-6", "2D10/3R+4*5-6"},
		{"2d6*3-1d6+1", "2D6*3-1D6+1"},
		{"(2+3)d6-1+3d6+2", "(2+3)D6-1+3D6+2"},
		{"(2*3-4)d6-1d4+1", "(2*3-4)D6-1D4+1"},
		{"((2+3)*4/3)d6*2+5", "((2+3)*4/3)D6*2+5"},
		{"2d(1+5)", "2D(1+5)"},
		{"(8/2)D(4+6)", "(8/2)D(4+6)"},
		{"(2-1)d(8/2)*(1+1)d(3*4/2)+2*3", "(2-1)D(8/2)*(1+1)D(3*4/2)+2*3"},
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
				t.Fatalf("中置表記生成エラー: %s", notationErr)
				return
			}

			if actual != test.expected {
				t.Fatalf("got %q, want %q", actual, test.expected)
			}
		})
	}
}
