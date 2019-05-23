package samedayofbirth

import (
    "fmt"
)

/*
В классе 30 человек.
Какая вероятность того, что у двух одноклассиников будет дни
рождения в один день?
сезонность дней рождения не учитывать
*/


const EXPERIMENTS_NUMBER = 999
const MAX_CLASSLIMIT = 35


func GeneratePersonSetTestData() [MAX_CLASSLIMIT]*Person {
    var result [MAX_CLASSLIMIT]*Person

    for i:= 0; i<MAX_CLASSLIMIT; i++ {
        result[i] = &Person {name: "name-debug", birthDay: i}
    }

    //Artificially make repeat
    //result[MAX_CLASSLIMIT-2].birthDay = 4

    return result
}


func GeneratePersonSet() [MAX_CLASSLIMIT]*Person {
    var result [MAX_CLASSLIMIT]*Person

    for i:= 0; i<MAX_CLASSLIMIT; i++ {
        result[i] = MakePerson(fmt.Sprintf("Person #%d",i))
    }
    return result
}

/**
*  Return  1 if 2 persons are birthday in same day
*  Return 0 otherwise
*/
func CheckTwoPersonsAreBirthdayBoys(personSet [MAX_CLASSLIMIT]*Person) int {
    matches := 0

    for i, aPersonI :=range(personSet) {
        for j, aPersonJ :=range(personSet) {

            //fmt.Printf("i is %d, j is %d, compare %d with %d",i,j,aPersonI.birthDay,aPersonJ.birthDay)

            if j == i {
                continue
            }

            if aPersonI.birthDay == aPersonJ.birthDay {
                matches = 1
                break
            }
        }
    }
    return matches
}

func PerformExperiment() string {
    var result string = "Total Experiment Count: " + fmt.Sprintf("%d",EXPERIMENTS_NUMBER) + "\n"
    result += "Person Count: " +   fmt.Sprintf("%d",MAX_CLASSLIMIT)  + "\n"

    successCount := 0
    for i:=0 ;i < EXPERIMENTS_NUMBER; i++{
        personSet := GeneratePersonSet()
        //personSet := GeneratePersonSetTestData()

            /*  DEBUG DATA
            for _, aPerson := range(personSet){
                result += fmt.Sprintf("%+v\n",aPerson)
            }
            result += "\nExperiment result " + fmt.Sprintf("%d",CheckTwoPersonsAreBirthdayBoys(personSet))
            */

        successCount += CheckTwoPersonsAreBirthdayBoys(personSet)
    }
    result += "Success Experiment Results: " + fmt.Sprintf("%d",successCount)
    result += "\n\nProbability: " + fmt.Sprintf("%.6f",float64(successCount)/float64(EXPERIMENTS_NUMBER))

    return result;
}