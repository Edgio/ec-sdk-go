ACCESSRULE=addAccessRule
MANAGEDRULE=managed_rules

addAccessRule:
	go build -o ./bin/${ACCESSRULE} ./example/waf/addAccessRule.go
managedRule:
	go build -o ./bin/${MANAGEDRULE} ./example/waf/managed_rules/${MANAGEDRULE}.go