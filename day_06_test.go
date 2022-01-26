package main

import (
	. "christmas-challenge/day_06"
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

func toFlockOfFish(fishString string)FlockOfFish{
    retval := make([]Fish, 0)
    for _, strNum := range strings.Split(fishString, ","){
        strNum = strings.TrimSpace(strNum)
        i, e := strconv.Atoi(strNum)
        if e == nil {
            retval = append(retval, Fish{Day: i})
        }
    }
    return FlockOfFish{retval}
}

var flockOfFishtests = []struct {
    inDay int 
    inFlock FlockOfFish
    out FlockOfFish
}{
    {0, toFlockOfFish("3,4,3,1,2"), toFlockOfFish("3,4,3,1,2")},
    {1, toFlockOfFish("3,4,3,1,2"), toFlockOfFish("2,3,2,0,1")},
}

func TestFlockOfFishBehaviour(t *testing.T){
    for _, tt := range flockOfFishtests {
        t.Run(tt.inFlock.String()+ "|||" +tt.out.String(), func(t *testing.T) {
            result := tt.inFlock.Simulate(tt.inDay)
            if result.String() != tt.out.String(){
                t.Errorf("\nin\t%s\tday %d\n\nex\t%s\ngt\t%s\n",
                    tt.inFlock.String(), tt.inDay, 
                    tt.out.String(), result.String())

            }
        })
    }
}
