package favor

import "testing"

func TestFoodRatings(t *testing.T) {
	foods := []string{"pasta", "brisket", "sausage", "gumbo", "jambalaya", "frito pie", "pizza", "meatballs"}
	cuisines := []string{"italian", "bbq", "bbq", "cajun", "cajun", "bbq", "italian", "italian"}
	ratings := []int{2, 9, 10, 5, 7, 12, 8, 15}

	mockFrs := FoodRatingConstructor(foods, cuisines, ratings)

	res1 := mockFrs.HighestRated("italian")
	if res1 != "meatballs" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", res1, "meatballs")
	}
	res2 := mockFrs.HighestRated("cajun")
	if res2 != "jambalaya" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", res2, "jambalaya")
	}
	mockFrs.ChangeRating("gumbo", 11)
	res3 := mockFrs.HighestRated("cajun")
	if res3 != "gumbo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", res2, "gumbo")
	}
	mockFrs.ChangeRating("jambalaya", 11)
	res4 := mockFrs.HighestRated("cajun")
	if res4 != "gumbo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", res2, "gumbo")
	}

}
