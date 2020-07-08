package main

import (
	"log"
	"strings"
	"sync"
	"time"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var censor = map[string]Censorship{
	"209785549010108416": {"Arisa", setup("loli", "smh")},
	"207833493651193856": {"XPHonza", setup("boomer")},
	"297773911158816769": {"leijurv", setup("not allowed to say this")},
}

var globalCensor = Censorship{
	"anyone",
	append([]Explained{
		Explained{"retard", "an ableist slur", ""},
		Explained{"nigg", "a racist slur", "685255238571130891"}, 
		Explained{"fabritone is better than baritone", "something no one is allowed to say", ""},
		Explained{"impact is better than future", "something no one is allowed to say", ""},
	}, setup("kami", "blue", "力ミ", "ブル", "salhack")...),
}

var bannedNicks = []string{
	"loli",
}

var censorCounts = make(map[string]int)
var censorCountsLock sync.Mutex

func init() {
	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for range ticker.C {
			censorCountsLock.Lock()
			censorCounts = make(map[string]int)
			censorCountsLock.Unlock()
		}
	}()
}

func setup(strs ...string) []Explained {
	ret := make([]Explained, 0)
	for _, str := range strs {
		ret = append(ret, Explained{str, "\"" + str + "\"", ""})
	}
	return ret
}

func onMessageSent3(session *discordgo.Session, m *discordgo.MessageCreate) {
	enforcement(session, m.Message)
}

func onMessageUpdate(session *discordgo.Session, m *discordgo.MessageUpdate) {
	enforcement(session, m.Message)
}

func enforceNickname(m *discordgo.Member) {
	if nick, ok := nicknameENFORCEMENT[m.User.ID]; ok && m.Nick != nick {
		discord.GuildMemberNickname(impactServer, m.User.ID, nick)
		return
	}
	for _, badNick := range bannedNicks {
		if strings.Contains(strings.ToLower(m.Nick), strings.ToLower(badNick)) {
			resp(impactBotLog, "Note: User "+m.User.Username+" tried to change their nick to \""+m.Nick+"\", which is ILLEGAL")
			err := discord.GuildMemberNickname(impactServer, m.User.ID, trash)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func incrementCensorCounts(authorID string, msg *discordgo.Message, reason string) int {
	censorCountsLock.Lock()
	defer censorCountsLock.Unlock()
	censorCounts[authorID] = censorCounts[authorID] + 1
	if censorCounts[authorID] < 5 {
		return censorCounts[authorID]
	}
	me, err := GetMember(myselfID)
	if err != nil {
		return censorCounts[authorID]
	}
	muteHandler(me, msg, []string{"tempmute", "<@" + authorID + ">", reason})
	return censorCounts[authorID]
}

func enforcement(session *discordgo.Session, msg *discordgo.Message) {
	if msg == nil || msg.Author == nil || msg.Type != discordgo.MessageTypeDefault {
		return // wtf
	}
	
	for _, censorship := range []Censorship{censor[msg.Author.ID], globalCensor} {
		for _, bannedWord := range censorship.bannedWords {
			if strings.Contains(strings.ToLower(msg.Content), strings.ToLower(bannedWord.str)) {
				if bannedWord.onlyIfTheyDontHave != "" {
					user, err := GetMember(msg.Author.ID)
					if err != nil {
						return
					}
					if includes(user.Roles, bannedWord.onlyIfTheyDontHave) {
						continue
					}
				}
				session.ChannelMessageDelete(msg.ChannelID, msg.ID)
				ret := "Note: a message containing "+bannedWord.explain+" from "+censorship.name+" was deleted."
				cnt := incrementCensorCounts(msg.Author.ID, msg, ret)
				if cnt > 2 {
					ret += " This has happened "+strconv.Itoa(cnt)+ "times in the last 10 minutes. Watch out!"
				}
				resp(msg.ChannelID, ret)
				return
			}
		}
	}
}

type Explained struct {
	str string
	explain string
	onlyIfTheyDontHave string
}

type Censorship struct {
	name        string
	bannedWords []Explained
}
