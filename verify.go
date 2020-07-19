package main

import (
	"log"
	"strconv"
	"time"
	"fmt"
	"errors"

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
			st, err := discord.GuildMembers(impactServer, prev, 1000)
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
				/*if IsUserStaff(member) {
					discord.GuildMemberNickname(impactServer, member.User.ID, "")
				}*/
				if IsUserStaff(member) {
					staffIDs[GetHighestStaffRole(member)] = append(staffIDs[GetHighestStaffRole(member)], member.User.ID)
				}
			}
		}
		log.Println("Processed", total, "members")
		log.Println("There are", donatorCount, "donators")
		log.Println("There are", verifiedCount, "verified")
		if DB == nil {
			return
		}
		for i := 1; i < 5; i++ {
			for _, id := range staffIDs[i] {
				var num int
				err := DB.QueryRow("SELECT nick FROM nicks WHERE id = $1", id).Scan(&num)
				if err != nil {
					err = DB.QueryRow("INSERT INTO nicks(id) VALUES ($1) RETURNING (nick)", id).Scan(&num)
					if err != nil {
						panic(err)
					}
				}
				str := strconv.Itoa(num)
				for len(str) < 2 {
					str = "0" + str
				}
				nick := "Volunteer "+str
				err = discord.GuildMemberNickname(impactServer, id, nick)
				if err != nil {
					log.Println(err)
				}
				nicknameENFORCEMENT[id] = nick
			}
		}
	}()
}

func onGuildMemberUpdate(discord *discordgo.Session, guildMemberUpdate *discordgo.GuildMemberUpdate) {
	memberSanityCheck(guildMemberUpdate.Member)
}

func memberSanityCheck(member *discordgo.Member) {
	if member.User.Bot {
		// Don't enforce roles/nicks on bots
		log.Printf("Skipping sanity check on bot user %s#%s (%s)\n", member.User.Username, member.User.Discriminator, member.Nick)
		return
	}
	if len(member.Roles) > 0 && !hasRole(member, Verified) {
		log.Println("Member", member.User.ID, "had roles not including verified")
		err := discord.GuildMemberRoleAdd(impactServer, member.User.ID, Verified.ID)
		if err != nil {
			log.Println(err)
		}
	}
	if !hasRole(member, Verified) && accountCreatedMoreThanSixMonthsAgo(member.User.ID) && joinedServerMoreThanOneMonthAgo(member) {
		log.Println("Member", member.User.ID, "has been in the server for a month, and has an account over 6 months old, but isn't verified")
		err := discord.GuildMemberRoleAdd(impactServer, member.User.ID, Verified.ID)
		if err != nil {
			log.Println(err)
		}
	}
	if hasRole(member, InVoice) && !checkDeservesInVoiceRole(member.User.ID) {
		log.Println("Member", member.User.ID, "had In Voice but isn't in voice")
		err := discord.GuildMemberRoleRemove(impactServer, member.User.ID, InVoice.ID)
		if err != nil {
			log.Println(err)
		}
	}
	enforceNickname(member)
}

func wantHandler(caller *discordgo.Member, msg *discordgo.Message, args []string) error {
	reply := discordgo.MessageEmbed{
		Color: prettyembedcolor,
	}

	switch len(args) {
	case 1:
		reply.Title = "no"
		reply.Description = "give number you want"
	case 2:
		want, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		sentBy := msg.Author.ID

		var curr int 
		err = DB.QueryRow("SELECT nick FROM nicks WHERE id = $1", sentBy).Scan(&curr)
		if err != nil {
			return err
		}
		if curr == want {
			return errors.New("No")
		}
		var already string
		err = DB.QueryRow("SELECT id FROM nicks WHERE nick = $1", want).Scan(&already)
		if err != nil {
			return err
		}
		// already held by already
		_, err = DB.Exec("INSERT INTO nicktrade (id, desirednick) VALUES ($1, $2) ON CONFLICT (id, desirednick) DO NOTHING", sentBy, want)
		if err != nil {
			return err
		}
		rows, err := DB.Query("SELECT nicktrade.desirednick AS desired, nicks.nick AS curr FROM nicks INNER JOIN nicktrade ON nicktrade.id = nicks.id")
		if err != nil {
			return err
		}
		defer rows.Close()
		edges := make(map[int][]int)
		for rows.Next() {
			var desired int
			var curr int
			err = rows.Scan(&desired, &curr)
			if err != nil {
				return err
			}
			edges[curr]=append(edges[curr], desired)
		}
		err = rows.Err()
		if err != nil {
			return err
		}
		path := DFS(edges, curr, curr)
		if len(path) < 2 {
			return errors.New("okay i added your request to the database but i cannot satisfy it at the moment")
		}
		path=path[:len(path)-1]
		reply.Title="yes"
		reply.Description=fmt.Sprintf("Based cycle", path)
		IDs := make([]string,0)
		for i := range path {
			oldNick := path[i]
			//newNick := path[(i+1)%len(path)]
			var person string
			err = DB.QueryRow("SELECT id FROM nicks WHERE nick = $1", oldNick).Scan(&person)
			if err != nil {
				return err
			}
			IDs = append(IDs, person)
		}
		for i := range path {
			//oldNick := path[i]
			newNick := path[(i+1)%len(path)]
			person := IDs[i]

			_, err := DB.Exec("UPDATE nicks SET nick = $1 WHERE id = $2", newNick, person)
			if err != nil {
				return err
			}

			_, err = DB.Exec("DELETE FROM nicktrade WHERE desirednick = $1 AND id = $2", newNick, person)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("incorrect number of arguments")
	}

	_, err := discord.ChannelMessageSendEmbed(msg.ChannelID, &reply)
	return err
}

func DFS(edges map[int][]int, start int, end int) []int {
	for _, str := range edges[start] {
		if str == end {
			return []int{start, end}
		}
		path := DFS(edges, str, end)
		if path != nil {
			return append([]int{start}, path...)
		}
	}
	return nil
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
