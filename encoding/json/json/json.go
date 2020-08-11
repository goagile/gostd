package main

import (
	"encoding/json"
	"fmt"
	// "os"
)

type Person struct {
	Name string
	Age int
}

type PersonMap struct {
	Name string `json: name`
	Age int `json: age`
}

func (p *PersonMap) Marshal() ([]byte, error) {
	return json.Marshal(&p)
}

func (p *PersonMap) Unmarshal(jsonPerson []byte) PersonMap {
	personMap := PersonMap{}
	if err := json.Unmarshal(jsonPerson, &personMap); err != nil {
		panic(err)
	}
	return personMap
}

func main() {

	// Запаковываем объекты различного типа
	// Логический тип
	boolPack, _ := json.Marshal(true)
	fmt.Println("boolPack:", string(boolPack))

	// Целые числа
	intPack, _ := json.Marshal(1)
	fmt.Println("intPack:", string(intPack))

	// Вещественные числа
	floatPack, _ := json.Marshal(2.5)
	fmt.Println("floatPack:", string(floatPack))

	// Строковые переменные
	stringPack, _ := json.Marshal("Hello")
	fmt.Println("stringPack:", string(stringPack))

	// Срезы (Массивы)
	slice := []string{"a", "b", "c"}
	slicePack, _ := json.Marshal(slice)
	fmt.Println("slicePack:", string(slicePack))

	// Хэш-мап (Словарь)
	mappa := map[string]int{"a": 1, "b": 2}
	mapPack, _ := json.Marshal(mappa)
	fmt.Println("mapPack:", string(mapPack))

	// Структура
	ivan := &Person{Name: "Ivan", Age: 32}
	ivanPack, _ := json.Marshal(ivan)
	fmt.Println("ivanPack:", string(ivanPack))

	// Структура с определением маппинга
	boris := &PersonMap{Name: "Boris", Age: 50}
	borisPack, _ := json.Marshal(boris)
	fmt.Println("borisPack:", string(borisPack))


	// Распаковываем объекты различного типа
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(`{"a": 2, "b": ["z", "x"]}`), &result); err != nil {
		panic(err)
	}
	fmt.Println("result:", result)

	// 
	personMap := PersonMap{}
	personByt := []byte(`{"name": "Alice", "age": 23}`)
	if err := json.Unmarshal(personByt, &personMap); err != nil {
		panic(err)
	}
	fmt.Println("personMap:", personMap)

	// PersonMap с методом Marshal
	kate := PersonMap{Name: "Kate", Age: 27}
	kateByt, err := kate.Marshal()
	if err != nil {
		panic(err)
	}
	fmt.Println("kate json:", string(kateByt))

	zulu := PersonMap{}
	r := zulu.Unmarshal([]byte(`{"name": "Zulu", "age": 58}`))
	fmt.Println("r:", r)

}
