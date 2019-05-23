package samedayofbirth

type Person struct{
    name string
    birthDay  int //Day of year - from 1 to 365
}

func (p *Person) generateBirthday() {
    p.birthDay = getRandom()
}

func MakePerson(name string) *Person {
    person := new(Person)
    person.name = name
    person.generateBirthday()
    return person
}