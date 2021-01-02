package ruleset

import (
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	log "github.com/sirupsen/logrus"
)

type Severity struct {
	level int
	name  string
}

var (
	Deprecation = Severity{3, "Deprecation"}
	Error       = Severity{0, "Error"}
	Info        = Severity{2, "Info"}
	Warning     = Severity{1, "Warning"}
)

type Rule struct {
	id             string
	description    string
	severity       Severity
	validationFunc func(command instructions.Command) RuleValidationResult
}

func NewRule(name string, description string, severity Severity,
	validationFunc func(command instructions.Command) RuleValidationResult) bool {
	log.Trace("NewRule called")

	all = append(all, Rule{
		name,
		description,
		severity,
		validationFunc,
	})

	log.Trace("New rule,", all[len(all)-1].Id, "added.")

	return true
}

func (rule *Rule) Id() string {
	return rule.id
}

func (rule *Rule) Severity() Severity {
	return rule.severity
}

func (rule *Rule) Description() string {
	return rule.description
}

func (rule *Rule) ValidationFunc(command instructions.Command) RuleValidationResult {
	return rule.validationFunc(command)
}

type RuleValidationResult struct {
	rule          Rule
	isViolated    bool
	LocationRange Range
}

func (ruleValidationResult RuleValidationResult) IsViolated() bool {
	return ruleValidationResult.isViolated
}

func (ruleValidationResult *RuleValidationResult) SetViolated() {
	ruleValidationResult.isViolated = true
}

func (ruleValidationResult *RuleValidationResult) Location() Range {
	return ruleValidationResult.LocationRange
}

func (ruleValidationResult *RuleValidationResult) SetLocation(startLineNUmber, startCharNumber, endLineNumber, endCharNumber int) {
	ruleValidationResult.LocationRange.start.lineNumber = startLineNUmber
	ruleValidationResult.LocationRange.start.charNumber = startCharNumber
	ruleValidationResult.LocationRange.end.lineNumber = endLineNumber
	ruleValidationResult.LocationRange.end.charNumber = endCharNumber
}

func (ruleValidationResult *RuleValidationResult) SetRule(rule Rule) {
	ruleValidationResult.rule = rule
}

type Range struct {
	start Location
	end   Location
}

type Location struct {
	lineNumber int
	charNumber int
}

func (location *Location) LineNumber() int {
	return location.lineNumber
}

func (location *Location) CharNumber() int {
	return location.charNumber
}

func (locationRange *Range) Start() Location {
	return locationRange.start
}

func (locationRange *Range) End() Location {
	return locationRange.end
}

func Get() []Rule {
	return all
}

var all []Rule