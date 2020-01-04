package main

import "github.com/bwmarrin/discordgo"

type Role struct {
	// Discord id
	ID   string
	Name string
}

var (
	HeadDev   = Role{"209817890713632768", "headDeveloper"}
	Developer = Role{"221655083748687873", "developer"}
	SeniorMod = Role{"663065117738663938", "seniorMod"}
	Moderator = Role{"210377982731223040", "moderator"}
	Support   = Role{"245682967546953738", "support"}
)

// Higher (lower index) is better
var staffRoles = []Role{
	HeadDev,
	Developer,
	SeniorMod,
	Moderator,
	Support,
}

func (r Role) IsAtLeast(role Role) bool {
	for _, testRole := range staffRoles {
		switch testRole {
		case role:
			return true
		case r:
			return false
		}
	}
	return false
}

func (r Role) IsHigherThan(role Role) bool {
	for _, testRole := range staffRoles {
		switch testRole {
		case r:
			return false
		case role:
			return true
		}
	}
	return false
}

func GetRolesAtLeast(role Role) []Role {
	return staffRoles[:GetRank(role)+1]
}

func GetRolesHigherThan(role Role) []Role {
	return staffRoles[:GetRank(role)]
}

// Lower is better, -1 is not found
func GetRank(role Role) int {
	for i, r := range staffRoles {
		if r == role {
			return i
		}
	}
	return -1
}

func IsUserStaff(user *discordgo.Member) bool {
	return hasRole(user, staffRoles...)
}

func IsUserAtLeast(user *discordgo.Member, role Role) bool {
	return hasRole(user, GetRolesAtLeast(role)...)
}

func IsUserHigherThan(user *discordgo.Member, role Role) bool {
	return hasRole(user, GetRolesHigherThan(role)...)
}
