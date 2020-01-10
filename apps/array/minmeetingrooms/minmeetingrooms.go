package main

import(
	"fmt"
	"container/heap"
)

type Event struct  {
    time int
    start bool
    // The index of the item in the heap
    index int 
}

// An EventQueue implements heap.Interface and holds Events.
type EventQueue []*Event

func (eq EventQueue) Len() int { return len(eq) }

// Less returns the smaller of two given events:
// * time(i) < time(j)
// * If both times are the same, then "end" events come before "start" events 
func (eq EventQueue) Less(i, j int) bool {
    if eq[i].time < eq[j].time {
        return true
    } else if eq[i].time == eq[j].time {
        // If two events end and start at the same time
        // then the end time shall occur before the next start time
        return !eq[i].start && eq[j].start
    } else {
        return false
    }
}

func (eq EventQueue) Swap(i, j int) {
	eq[i], eq[j] = eq[j], eq[i]
	eq[i].index = i
	eq[j].index = j
}

func (eq *EventQueue) Push(x interface{}) {
	n := len(*eq)
	event := x.(*Event)
	event.index = n
	*eq = append(*eq, event)
}

func (eq *EventQueue) Pop() interface{} {
	old := *eq
	n := len(old)
	event := old[n-1]
	old[n-1] = nil  // avoid memory leak
	event.index = -1 // for safety
	*eq = old[0 : n-1]
	return event
}

func initEventQueue(intervals [][]int) *EventQueue {
    eq := make(EventQueue, 2*len(intervals))
	i := 0
	for _, v := range intervals {
		eq[i] = &Event{
			time:  v[0],
			start: true,
			index: i,
		}
		i++
        eq[i] = &Event{
			time:  v[1],
			start: false,
			index: i,
		}
		i++
	}
	heap.Init(&eq)
    return &eq
}

func minMeetingRooms(intervals [][]int) int {
    eq := initEventQueue(intervals)
    maxRooms := 0
    rooms := 0
	nofEvents := len(*eq)
    for i := 0; i < nofEvents; i++ {
        event := heap.Pop(eq).(*Event)
        if event.start {
            rooms++
            if rooms > maxRooms {
                maxRooms = rooms
            }
        } else {
            rooms--
        }
    }
    return maxRooms
}

func main() {
	// intervals := [][]int{{9,10},{4,9},{4,17}}
	intervals := [][]int{{1,5},{8,9},{8,9}}
	//intervals := [][]int{{9,14},{5,6},{3,5},{12,19}}
	minRooms := minMeetingRooms(intervals)
	fmt.Println("Minimum rooms:", minRooms)
}