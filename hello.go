package favor

import "container/heap"

type Food struct {
	Name   string
	Rating int
}

type FoodRatings struct {
	FoodRating     map[string]int
	FoodCuisine    map[string]string
	CuisineHeapMap map[string]*FoodHeap
}

type FoodHeap []Food

func (h FoodHeap) Len() int { return len(h) }

func (h FoodHeap) Less(i, j int) bool {
	if h[i].Rating == h[j].Rating {
		return h[i].Name < h[j].Name
	}
	return h[i].Rating > h[j].Rating
}

func (h FoodHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *FoodHeap) Push(heapintf interface{}) {
	*h = append(*h, heapintf.(Food))
}

func (h *FoodHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func FoodRatingConstructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodRating := make(map[string]int)
	foodCuisine := make(map[string]string)
	cuisineHeapMap := make(map[string]*FoodHeap)

	for i := 0; i < len(foods); i++ {
		foodRating[foods[i]] = ratings[i]
		foodCuisine[foods[i]] = cuisines[i]

		// need a FoodHeap mapped to each cuisine
		if _, ok := cuisineHeapMap[cuisines[i]]; !ok {
			cuisineHeapMap[cuisines[i]] = &FoodHeap{}
		}

		heap.Push(cuisineHeapMap[cuisines[i]], Food{foods[i], ratings[i]})

	}

	return FoodRatings{
		FoodRating:     foodRating,
		FoodCuisine:    foodCuisine,
		CuisineHeapMap: cuisineHeapMap,
	}

}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	// heap push
	this.FoodRating[food] = newRating
	c := this.FoodCuisine[food]
	heap.Push(this.CuisineHeapMap[c], Food{food, newRating})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	// top value
	highestRatedFood := (*this.CuisineHeapMap[cuisine])[0]
	// heap pop
	for highestRatedFood.Rating != this.FoodRating[highestRatedFood.Name] {
		heap.Pop(this.CuisineHeapMap[cuisine]) // remove incorrect values
		highestRatedFood = (*this.CuisineHeapMap[cuisine])[0]
	}

	return highestRatedFood.Name
}
