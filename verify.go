package main

import (
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

var staffIDs [5][]string

var nicknameENFORCEMENT = make(map[string]string)

func onReady2(discord *discordgo.Session, ready *discordgo.Ready) {
	go func() {
		prev := ""
		total := 0
		donatorCount := 0
		verifiedCount := 0
		for {
			log.Println("Fetching starting at", prev)
			st, err := discord.GuildMembers(IMPACT_SERVER, prev, 1000)
			if err != nil {
				log.Println(err)
				break
			}
			log.Println("Fetched", len(st), "more members, total so far is", total)
			if len(st) == 0 {
				log.Println("No more members!")
				break
			}
			prev = st[len(st)-1].User.ID
			for _, member := range st {
				total++
				memberSanityCheck(member)
				if hasRole(member, Donator) {
					donatorCount++
				}
				if hasRole(member, Verified) {
					verifiedCount++
				}
				if IsUserStaff(member) {
					staffIDs[GetHighestStaffRole(member)] = append(staffIDs[GetHighestStaffRole(member)], member.User.ID)
				}
			}
		}
		log.Println("Processed", total, "members")
		log.Println("There are", donatorCount, "donators")
		log.Println("There are", verifiedCount, "verified")
		meanEntity := 0
		for i := 0; i < 5; i++ {
			for _, id := range staffIDs[i] {
				str := strconv.Itoa(meanEntity)
				for len(str) < 2 {
					str = "0" + str
				}
				nick := "Mean Entity " + str
				if i == 3 {
					nick = "Red Entity " + str
				}
				if i == 4 {
					nick = "Purple Entity " + str
				}
				nick = `
͔`
				nick = "﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽﷽"
				nick = str
				err := discord.GuildMemberNickname(IMPACT_SERVER, id, nick)
				if err != nil {
					log.Println(err)
				}
				nicknameENFORCEMENT[id] = nick
				meanEntity++
			}
		}
	}()
}

func onGuildMemberUpdate(discord *discordgo.Session, guildMemberUpdate *discordgo.GuildMemberUpdate) {
	memberSanityCheck(guildMemberUpdate.Member)
}

func memberSanityCheck(member *discordgo.Member) {
	if len(member.Roles) > 0 && !hasRole(member, Verified) {
		log.Println("Member", member.User.ID, "had roles not including verified")
		err := discord.GuildMemberRoleAdd(IMPACT_SERVER, member.User.ID, Verified.ID)
		if err != nil {
			log.Println(err)
		}
	}
	if !hasRole(member, Verified) && accountCreatedMoreThanSixMonthsAgo(member.User.ID) && joinedServerMoreThanOneMonthAgo(member) {
		log.Println("Member", member.User.ID, "has been in the server for a month, and has an account over 6 months old, but isn't verified")
		err := discord.GuildMemberRoleAdd(IMPACT_SERVER, member.User.ID, Verified.ID)
		if err != nil {
			log.Println(err)
		}
	}
	if hasRole(member, InVoice) && !checkDeservesInVoiceRole(member.User.ID) {
		log.Println("Member", member.User.ID, "had In Voice but isn't in voice")
		err := discord.GuildMemberRoleRemove(IMPACT_SERVER, member.User.ID, InVoice.ID)
		if err != nil {
			log.Println(err)
		}
	}
	enforceNickname(member)
}

func accountCreatedMoreThanSixMonthsAgo(discordID string) bool {
	u, err := strconv.ParseUint(discordID, 10, 64)
	if err != nil {
		return false
	}
	nowMS := uint64(time.Now().Unix()) * 1000
	createdAtMS := u/(1<<22) + 1420070400000
	ageMS := nowMS - createdAtMS
	ageDays := ageMS / 1000 / 86400
	return ageDays > 6*30
}

func joinedServerMoreThanOneMonthAgo(member *discordgo.Member) bool {
	joinedAt, err := member.JoinedAt.Parse()
	if err != nil {
		return false
	}
	duration := time.Now().Sub(joinedAt)
	return duration > 30*24*time.Hour
}
