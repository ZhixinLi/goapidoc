// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package api_detail

import "github.com/gogf/gf/frame/g"

// Fill with you ideas below.
func FindAllOrd(where g.Map, ord string) *Entity {
	var apiDetail *Entity
	db := g.DB().Table("api_detail").Safe()
	err := db.Where(where).Order(ord).FindScan(&apiDetail)
	if err != nil {
		return nil
	}
	return apiDetail
}
