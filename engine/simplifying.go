package engine

// {"age": {"$gte": 5, "$in": [0, 5]}} -> {"$or": {"age": {"$eq": 5, "$gt": 5}}}
// func Simplify(filter map[string]any) map[string]any {
// 	simplifiedFilter := make(map[string]any, len(filter))

// 	for key, value := range filter {
// 		if IsTopLevelFilterOp(key) {
// 			subfilter, _ := ToFilter(value)
// 			simplifiedFilter[key] = Simplify(subfilter)
// 		} else if IsFilterOp(key) {
// 			switch key {
// 			case "$gte":
// 				simplifiedFilter["$or"] = map[string]any{}
// 			}
// 		}
// 	}
// }
