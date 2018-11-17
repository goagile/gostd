package main

import (
	"fmt"
	"regexp"
)

func main() {

	pattern := "(П|п)([а-я]+)ка"
	example := "Помидорка"
	match, _ := regexp.MatchString(pattern, example)
	fmt.Printf("%s MatchString %s = %t\n", pattern, example, match)

	p, _ := regexp.Compile(pattern)
	fmt.Printf("%s MatchString %s = %t\n", pattern, example, p.MatchString(example))
	
	str := "Помидорка в корзинке"
	fmt.Printf("%s FindString %s = %s\n", pattern, example, p.FindString(str))
	fmt.Printf("%s FindStringIndex %s = %v\n", pattern, example, p.FindStringIndex(str))
	fmt.Printf("%s FindStringSubmatch %s = %v\n", pattern, example, p.FindStringSubmatch(str))

	str = "Помидорка в корзинке помидорка положилка"
	fmt.Printf("%s FindAllString %s = %v\n", pattern, example, p.FindAllString(str, 10))

	fmt.Printf("%s Match %s = %v\n", pattern, example, p.Match([]byte(str)))

}
