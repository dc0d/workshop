package model

import (
	"strings"
	"time"
)

// Statement .
type Statement struct {
	Lines []StatementLine
}

// NewStatement .
func NewStatement() *Statement { return &Statement{} }

// AddStatementLine .
func (s *Statement) AddStatementLine(line StatementLine) {
	s.Lines = append(s.Lines, line)
}

func (s *Statement) String() string {
	var sb strings.Builder
	sb.WriteString("date || credit || debit || balance\n")
	if len(s.Lines) > 0 {
		for i := len(s.Lines) - 1; i >= 0; i = i - 1 {
			sb.WriteString(s.Lines[i].String() + "\n")
		}
	}
	return strings.TrimRight(sb.String(), "\n")
}

// StatementLine .
type StatementLine struct {
	Date    time.Time
	Credit  Amount
	Debit   Amount
	Balance Amount
}

func (sl StatementLine) String() string {
	parts := []string{
		sl.Date.Format("02/01/2006"),
		prependSpace(sl.Credit.String()),
		prependSpace(sl.Debit.String()),
		prependSpace(sl.Balance.String()),
	}
	return strings.Join(parts, " ||")
}

func prependSpace(s string) string {
	if s == "" {
		return s
	}
	return " " + s
}
