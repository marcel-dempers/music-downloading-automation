package models
type Test struct {
	Name string
}

type Configuration struct {
	TestCollection []Test
	AnotherTestItem AnotherTest
}

type AnotherTest struct {
	test string
}