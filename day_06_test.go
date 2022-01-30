package main

import (
	. "christmas-challenge/day_06"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)
var fishtests = []struct {
    in  Fish
    out string
}{
    {Fish{0}, "day:000,cycle:000,timer:8"},
    {Fish{1}, "day:001,cycle:000,timer:7"},
    {Fish{2}, "day:002,cycle:000,timer:6"},
    {Fish{3}, "day:003,cycle:000,timer:5"},
    {Fish{4}, "day:004,cycle:000,timer:4"},
    {Fish{5}, "day:005,cycle:000,timer:3"},
    {Fish{6}, "day:006,cycle:000,timer:2"},
    {Fish{7}, "day:007,cycle:000,timer:1"},
    {Fish{8}, "day:008,cycle:000,timer:0"},
    {Fish{9}, "day:009,cycle:001,timer:6"},
    {Fish{10}, "day:010,cycle:001,timer:5"},
    {Fish{11}, "day:011,cycle:001,timer:4"},
    {Fish{12}, "day:012,cycle:001,timer:3"},
    {Fish{13}, "day:013,cycle:001,timer:2"},
    {Fish{14}, "day:014,cycle:001,timer:1"},
    {Fish{15}, "day:015,cycle:001,timer:0"},
    {Fish{16}, "day:016,cycle:002,timer:6"},
    {Fish{17}, "day:017,cycle:002,timer:5"},
}

func TestDay06(t *testing.T){
	inputs := make([]Fish, 0)
	inputs = append(inputs, Fish{2})

    got := GetStateDayX(inputs, 1)
    want := []Fish{{2}}

	for i, _ := range got{
		if got[i] != want[i] {
			t.Errorf("got %s, wanted %s", got[i], want[i])
		}
	}
}

func TestFishBehaviour(t *testing.T){
    for _, tt := range fishtests {
        t.Run(tt.in.String(), func(t *testing.T) {
            if tt.in.String() != tt.out{
                t.Errorf("Fish(%d) \nexpect\t%s\t\nretval\t%s",
                    tt.in.Day, tt.out, tt.in.String())

            }
        })
    }
}

func toFlockOfFish(fishTimeToGestationString string)FlockOfFish{
    // we assume that all fish in the zeroth cycle in the initial state
    retval := make([]Fish, 0)
    for _, strNum := range strings.Split(fishTimeToGestationString, ","){
        strNum = strings.TrimSpace(strNum)
        i, e := strconv.Atoi(strNum)
        if e == nil {
            retval = append(retval, Fish{Day: FIRST_CYCLE_COUNTER - i})
        }
    }
    return FlockOfFish{State:retval}
}

func getDiffFromTwoStringsWithNumbers(s1 string, s2 string) string{
    a1 := convertToNumbersSlice(s1)
    a2 := convertToNumbersSlice(s2)
    f1 := getNumbersFrequency(a1)
    f2 := getNumbersFrequency(a2)
    
    f1f2Delta := printNumbersInFirstButNotInSecond(f1, f2)
    f2f1Delta := printNumbersInFirstButNotInSecond(f2, f1)
    if len(f1f2Delta) == 0 && len(f2f1Delta) == 0{
        return ""
    } else{
        return fmt.Sprintf("First second delta: '%s'\nSecond first delta: '%s'", f1f2Delta, f2f1Delta)
    }
}

func sortNumbersString(s string)string{
    retval := ""
    ns := convertToNumbersSlice(s)
    sort.Ints(ns)
    for _, n := range ns{
        retval += fmt.Sprintf("%d,", n)
    }
    return retval
}

func printNumbersInFirstButNotInSecond(f1 map[int]int, f2 map[int]int) string{
    deltaF1F2 := ""
    for f1val, f1freq := range f1{
        f2freq, exist := f2[f1val]
        var numberOfValuesNotPresent int
        if !exist{
            numberOfValuesNotPresent = f1freq
        }else{
            numberOfValuesNotPresent = f1freq - f2freq
        }

        if numberOfValuesNotPresent > 0{
            for i:=0;i<numberOfValuesNotPresent;i++{
                deltaF1F2 += fmt.Sprintf("%d,", f1val)
            }
        }        
    }
    return strings.TrimSuffix(deltaF1F2, ",")
}

func getNumbersFrequency(a1 []int)map[int]int {
    frequencyMap := make(map[int]int, 0)
    for _, n := range a1{
        value, exist := frequencyMap[n]
        if !exist{
            frequencyMap[n] = 0
        }
        frequencyMap[n] = value + 1
    }
    return frequencyMap
}

func convertToNumbersSlice(s1 string)[]int {
    retval := make([]int, 0)
    for _, strNum := range strings.Split(s1, ","){
        strNum = strings.TrimSpace(strNum)
        i, e := strconv.Atoi(strNum)
        if e == nil {
            retval = append(retval, i)
        }
    }
    return retval
}

var flockOfFishtests = []struct {
    inDay int 
    inFlock FlockOfFish
    out string
}{
{0, toFlockOfFish("3,4,3,1,2"), "3,4,3,1,2"},
{1, toFlockOfFish("3,4,3,1,2"), "2,3,2,0,1"},
{2, toFlockOfFish("3,4,3,1,2"), "1,2,1,6,0,8"},
{3, toFlockOfFish("3,4,3,1,2"), "0,1,0,5,6,7,8"},
{4, toFlockOfFish("3,4,3,1,2"), "6,0,6,4,5,6,7,8,8"},
{5, toFlockOfFish("3,4,3,1,2"), "5,6,5,3,4,5,6,7,7,8"},
{6, toFlockOfFish("3,4,3,1,2"), "4,5,4,2,3,4,5,6,6,7"},
{7, toFlockOfFish("3,4,3,1,2"), "3,4,3,1,2,3,4,5,5,6"},
{8, toFlockOfFish("3,4,3,1,2"), "2,3,2,0,1,2,3,4,4,5"},
{9, toFlockOfFish("3,4,3,1,2"), "1,2,1,6,0,1,2,3,3,4,8"},
{10, toFlockOfFish("3,4,3,1,2"), "0,1,0,5,6,0,1,2,2,3,7,8"},
{11, toFlockOfFish("3,4,3,1,2"), "6,0,6,4,5,6,0,1,1,2,6,7,8,8,8"},
{12, toFlockOfFish("3,4,3,1,2"), "5,6,5,3,4,5,6,0,0,1,5,6,7,7,7,8,8"},
{13, toFlockOfFish("3,4,3,1,2"), "4,5,4,2,3,4,5,6,6,0,4,5,6,6,6,7,7,8,8"},
{14, toFlockOfFish("3,4,3,1,2"), "3,4,3,1,2,3,4,5,5,6,3,4,5,5,5,6,6,7,7,8"},
{15, toFlockOfFish("3,4,3,1,2"), "2,3,2,0,1,2,3,4,4,5,2,3,4,4,4,5,5,6,6,7"},
{16, toFlockOfFish("3,4,3,1,2"), "1,2,1,6,0,1,2,3,3,4,1,2,3,3,3,4,4,5,5,6,8"},
{17, toFlockOfFish("3,4,3,1,2"), "0,1,0,5,6,0,1,2,2,3,0,1,2,2,2,3,3,4,4,5,7,8"},
{18, toFlockOfFish("3,4,3,1,2"), "6,0,6,4,5,6,0,1,1,2,6,0,1,1,1,2,2,3,3,4,6,7,8,8,8,8"},
}

func TestFlockOfFishBehaviour(t *testing.T){
    for _, tt := range flockOfFishtests {
        t.Run(fmt.Sprintf("day %d__%s", tt.inDay, tt.inFlock.String(), ), func(t *testing.T) {
            result := tt.inFlock.Simulate(tt.inDay)
            compareResult := getDiffFromTwoStringsWithNumbers(result.String(), tt.out)

            if compareResult != ""{
                t.Errorf("\nin\t%s\tday %d\n\nex\t%s\ngt\t%s\n%s\n",
                    tt.inFlock.String(), tt.inDay, 
                    sortNumbersString(tt.out), sortNumbersString(result.String()), compareResult)
                // t.Errorf("\nin\t%s\tday %d\n\nex\t%s\ngt\t%s\n",
                //     tt.inFlock.String(), tt.inDay, 
                //     tt.out, result.String())
            }

        })
    }
}
