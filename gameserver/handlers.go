package gameserver

import (
	"github.com/pangliang/MirServer-Go/protocol"
	"strings"
	"strconv"
	"fmt"
)

var gameHandlers = map[uint16]func(session *Session, request *protocol.Packet, server *GameServer) (err error){
	CM_NEWCHR : func(session *Session, request *protocol.Packet, server *GameServer) (err error) {
		const (
			WrongName = 0
			NameExist = 2
			MaxPlayers = 3
			SystemErr = 4
		)
		params := strings.Split(request.Data, "/")
		if len(params) < 5 {
			protocol.NewPacket(SM_NEWCHR_FAIL).SendTo(session.socket)
			return nil
		}

		server.env.RLock()
		user, ok := server.env.users[params[0]]
		server.env.RUnlock()

		if !ok {
			protocol.NewPacket(SM_NEWCHR_FAIL).SendTo(session.socket)
			return nil
		}

		player := Player{
			UserId:user.Id,
		}
		player.Name = params[1]
		player.Hair, _ = strconv.Atoi(params[2])
		player.Job, _ = strconv.Atoi(params[3])
		player.Gender, _ = strconv.Atoi(params[4])

		_, err = server.db.Save(player)
		if err != nil {
			protocol.NewPacket(SM_NEWCHR_FAIL).SendTo(session.socket)
			return
		}

		protocol.NewPacket(SM_NEWCHR_SUCCESS).SendTo(session.socket)
		return nil
	},
	CM_QUERYCHR : func(session *Session, request *protocol.Packet, server *GameServer) (err error) {
		params := strings.Split(request.Data, "/")
		if len(params) < 1 {
			protocol.NewPacket(SM_QUERYCHR_FAIL).SendTo(session.socket)
			return
		}
		username := params[0]
		server.env.RLock()
		loginUser, ok := server.env.users[username]
		server.env.RUnlock()

		if !ok {
			protocol.NewPacket(SM_QUERYCHR_FAIL).SendTo(session.socket)
			return
		}

		cert, err := strconv.Atoi(params[1])
		if err != nil || int32(cert) != loginUser.Cert {
			protocol.NewPacket(SM_QUERYCHR_FAIL).SendTo(session.socket)
			return
		}
		session.attr["user"] = loginUser

		var playerList []Player
		err = server.db.List(&playerList, "where userId=?", loginUser.Id)
		if err != nil {
			protocol.NewPacket(SM_QUERYCHR_FAIL).SendTo(session.socket)
			return err
		}

		if len(playerList) > 0 {
			resp := protocol.NewPacket(SM_QUERYCHR)
			resp.Header.Recog = int32(len(playerList))
			for _, player := range playerList {
				resp.Data += fmt.Sprintf("%s/%d/%d/%d/%d/",
					player.Name,
					player.Job,
					player.Hair,
					player.Level,
					player.Gender,
				)
			}

			resp.SendTo(session.socket)
		}

		return nil
	},
}
