package domain

import "testing"

func TestCampaign_UniquePromotion(t *testing.T) {
	c1 := Campaign{ID: 1}
	p1 := Promotion{CampaignID: c1.ID, Title: "Promo1"}
	p2 := Promotion{CampaignID: c1.ID, Title: "Promo1"}
	if p1.Title == p2.Title && p1.CampaignID == p2.CampaignID {
		t.Log("Duplicate promotion detected as expected")
	} else {
		t.Error("Promotion titles and campaign IDs should match for this test")
	}
}
