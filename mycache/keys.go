package mycache

import "strconv"

func GetRecommendationKey(userID uint) string {
	return "recommandation:" + strconv.Itoa(int(userID))
}

func GetRecommendationQueryFlagKey(userID uint) string {
	return "recommandation:setflag:" + strconv.Itoa(int(userID))
}
