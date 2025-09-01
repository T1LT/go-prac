package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
    for _, numBirds := range birdsPerDay {
        sum += numBirds
    }
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    sum := 0
    startIdx := (week - 1) * 7
	for i := startIdx; i < startIdx + 7; i++ {
        sum += birdsPerDay[i]
    }
    return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for idx := 0; idx < len(birdsPerDay); idx++ {
        if idx % 2 == 0 {
            birdsPerDay[idx]++
        }
    }
    return birdsPerDay
}
