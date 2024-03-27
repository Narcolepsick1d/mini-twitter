package models

const TableFollow = "followers"

type (
	Follow struct {
		LeadId     string `db:"lead_id" json:"lead_id"`
		FollowerId string `db:"follower" json:"follower_id"`
	}
)
