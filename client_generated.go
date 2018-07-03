package artifact

import (
	"context"
	"github.com/faceit/go-artifact/events"
	"github.com/faceit/go-artifact/protocol"
)

// AbandonGauntlet abandons a gauntlet.
// Request ID: k_EMsgClientToGCAbandonGauntlet
// Response ID: k_EMsgClientToGCAbandonGauntletResponse
// Request type: CMsgClientToGCAbandonGauntlet
// Response type: CMsgClientToGCAbandonGauntletResponse
func (d *Artifact) AbandonGauntlet(
	ctx context.Context,
	gauntletID uint32,
) (*protocol.CMsgClientToGCAbandonGauntletResponse, error) {
	req := &protocol.CMsgClientToGCAbandonGauntlet{
		GauntletId: &gauntletID,
	}
	resp := &protocol.CMsgClientToGCAbandonGauntletResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCAbandonGauntlet),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCAbandonGauntletResponse),
		resp,
	)
}

// CreateLobbyPrivate creates a lobby private.
// Request ID: k_EMsgClientToGCPrivateLobbyCreate
// Response ID: k_EMsgClientToGCPrivateLobbyCreateResponse
// Request type: CMsgClientToGCPrivateLobbyCreate
// Response type: CMsgClientToGCPrivateLobbyCreateResponse
func (d *Artifact) CreateLobbyPrivate(
	ctx context.Context,
) (*protocol.CMsgClientToGCPrivateLobbyCreateResponse, error) {
	req := &protocol.CMsgClientToGCPrivateLobbyCreate{}
	resp := &protocol.CMsgClientToGCPrivateLobbyCreateResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyCreate),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyCreateResponse),
		resp,
	)
}

// GetAIVsAIMatchComplete gets a ai vs ai match complete.
// Request ID: k_EMsgClientToGCGetAIVsAIMatchComplete
// Request type: CMsgClientToGCGetAIVsAIMatchComplete
func (d *Artifact) GetAIVsAIMatchComplete(
	aiMatchID uint32,
	winningPlayer uint32,
) {
	req := &protocol.CMsgClientToGCGetAIVsAIMatchComplete{
		AiMatchId:     &aiMatchID,
		WinningPlayer: &winningPlayer,
	}
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetAIVsAIMatchComplete), req)
}

// GetAIVsAIMatchConfig gets a ai vs ai match config.
// Request ID: k_EMsgClientToGCGetAIVsAIMatchConfig
// Response ID: k_EMsgClientToGCGetAIVsAIMatchConfigResponse
// Request type: CMsgClientToGCGetAIVsAIMatchConfig
// Response type: CMsgClientToGCGetAIVsAIMatchConfigResponse
func (d *Artifact) GetAIVsAIMatchConfig(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetAIVsAIMatchConfigResponse, error) {
	req := &protocol.CMsgClientToGCGetAIVsAIMatchConfig{}
	resp := &protocol.CMsgClientToGCGetAIVsAIMatchConfigResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetAIVsAIMatchConfig),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetAIVsAIMatchConfigResponse),
		resp,
	)
}

// GetDeckRegistry gets a deck registry.
// Request ID: k_EMsgClientToGCGetDeckRegistry
// Response ID: k_EMsgClientToGCGetDeckRegistryResponse
// Request type: CMsgClientToGCGetDeckRegistry
// Response type: CMsgClientToGCGetDeckRegistryResponse
func (d *Artifact) GetDeckRegistry(
	ctx context.Context,
) (*protocol.CMsgClientToGCGetDeckRegistryResponse, error) {
	req := &protocol.CMsgClientToGCGetDeckRegistry{}
	resp := &protocol.CMsgClientToGCGetDeckRegistryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetDeckRegistry),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetDeckRegistryResponse),
		resp,
	)
}

// GetGauntletRankLeaderboard gets a gauntlet rank leaderboard.
// Request ID: k_EMsgClientToGCGetGauntletRankLeaderboard
// Response ID: k_EMsgClientToGCGetGauntletRankLeaderboardResponse
// Request type: CMsgClientToGCGetGauntletRankLeaderboard
// Response type: CMsgClientToGCGetGauntletRankLeaderboardResponse
func (d *Artifact) GetGauntletRankLeaderboard(
	ctx context.Context,
	seasonID uint32,
) (*protocol.CMsgClientToGCGetGauntletRankLeaderboardResponse, error) {
	req := &protocol.CMsgClientToGCGetGauntletRankLeaderboard{
		SeasonId: &seasonID,
	}
	resp := &protocol.CMsgClientToGCGetGauntletRankLeaderboardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetGauntletRankLeaderboard),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetGauntletRankLeaderboardResponse),
		resp,
	)
}

// GetLeagueDetails gets league details.
// Request ID: k_EMsgClientToGCGetLeagueDetails
// Response ID: k_EMsgClientToGCGetLeagueDetailsResponse
// Request type: CMsgClientToGCGetLeagueDetails
// Response type: CMsgClientToGCGetLeagueDetailsResponse
func (d *Artifact) GetLeagueDetails(
	ctx context.Context,
	leagueID uint32,
) (*protocol.CMsgClientToGCGetLeagueDetailsResponse, error) {
	req := &protocol.CMsgClientToGCGetLeagueDetails{
		LeagueId: &leagueID,
	}
	resp := &protocol.CMsgClientToGCGetLeagueDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetLeagueDetails),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetLeagueDetailsResponse),
		resp,
	)
}

// GetLeagueStandings gets league standings.
// Request ID: k_EMsgClientToGCGetLeagueStandings
// Response ID: k_EMsgClientToGCGetLeagueStandingsResponse
// Request type: CMsgClientToGCGetLeagueStandings
// Response type: CMsgClientToGCGetLeagueStandingsResponse
func (d *Artifact) GetLeagueStandings(
	ctx context.Context,
	leagueID uint32,
) (*protocol.CMsgClientToGCGetLeagueStandingsResponse, error) {
	req := &protocol.CMsgClientToGCGetLeagueStandings{
		LeagueId: &leagueID,
	}
	resp := &protocol.CMsgClientToGCGetLeagueStandingsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetLeagueStandings),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetLeagueStandingsResponse),
		resp,
	)
}

// GetMatchDetails gets match details.
// Request ID: k_EMsgClientToGCGetMatchDetails
// Response ID: k_EMsgClientToGCGetMatchDetailsResponse
// Request type: CMsgClientToGCGetMatchDetails
// Response type: CMsgClientToGCGetMatchDetailsResponse
func (d *Artifact) GetMatchDetails(
	ctx context.Context,
	matchID uint64,
) (*protocol.CMsgClientToGCGetMatchDetailsResponse, error) {
	req := &protocol.CMsgClientToGCGetMatchDetails{
		MatchId: &matchID,
	}
	resp := &protocol.CMsgClientToGCGetMatchDetailsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetMatchDetails),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetMatchDetailsResponse),
		resp,
	)
}

// GetMatchHistory gets a match history.
// Request ID: k_EMsgClientToGCGetMatchHistory
// Response ID: k_EMsgClientToGCGetMatchHistoryResponse
// Request type: CMsgClientToGCGetMatchHistory
// Response type: CMsgClientToGCGetMatchHistoryResponse
func (d *Artifact) GetMatchHistory(
	ctx context.Context,
	accountID uint32,
) (*protocol.CMsgClientToGCGetMatchHistoryResponse, error) {
	req := &protocol.CMsgClientToGCGetMatchHistory{
		AccountId: &accountID,
	}
	resp := &protocol.CMsgClientToGCGetMatchHistoryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetMatchHistory),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetMatchHistoryResponse),
		resp,
	)
}

// GetPackOpening gets a pack opening.
// Request ID: k_EMsgClientToGCGetPackOpening
// Response ID: k_EMsgClientToGCGetPackOpeningResponse
// Request type: CMsgClientToGCGetPackOpening
// Response type: CMsgClientToGCGetPackOpeningResponse
func (d *Artifact) GetPackOpening(
	ctx context.Context,
	packItemID uint64,
) (*protocol.CMsgClientToGCGetPackOpeningResponse, error) {
	req := &protocol.CMsgClientToGCGetPackOpening{
		PackItemId: &packItemID,
	}
	resp := &protocol.CMsgClientToGCGetPackOpeningResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetPackOpening),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetPackOpeningResponse),
		resp,
	)
}

// GetQuestMapStatus gets quest map status.
// Request ID: k_EMsgClientToGCGetQuestMapStatus
// Response ID: k_EMsgClientToGCGetQuestMapStatusResponse
// Request type: CMsgClientToGCGetQuestMapStatus
// Response type: CMsgClientToGCGetQuestMapStatusResponse
func (d *Artifact) GetQuestMapStatus(
	ctx context.Context,
	questID uint32,
) (*protocol.CMsgClientToGCGetQuestMapStatusResponse, error) {
	req := &protocol.CMsgClientToGCGetQuestMapStatus{
		QuestId: &questID,
	}
	resp := &protocol.CMsgClientToGCGetQuestMapStatusResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetQuestMapStatus),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGetQuestMapStatusResponse),
		resp,
	)
}

// GrantGauntlet grants a gauntlet.
// Request ID: k_EMsgClientToGCGauntletGrant
// Response ID: k_EMsgClientToGCGauntletGrantResponse
// Request type: CMsgClientToGCGauntletGrant
// Response type: CMsgClientToGCGauntletGrantResponse
func (d *Artifact) GrantGauntlet(
	ctx context.Context,
	gauntletID uint32,
	grantID uint32,
) (*protocol.CMsgClientToGCGauntletGrantResponse, error) {
	req := &protocol.CMsgClientToGCGauntletGrant{
		GauntletId: &gauntletID,
		GrantId:    &grantID,
	}
	resp := &protocol.CMsgClientToGCGauntletGrantResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGauntletGrant),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGauntletGrantResponse),
		resp,
	)
}

// GrantGauntletChoice grants a gauntlet choice.
// Request ID: k_EMsgClientToGCGauntletGrantChoice
// Response ID: k_EMsgClientToGCGauntletGrantChoiceResponse
// Request type: CMsgClientToGCGauntletGrantChoice
// Response type: CMsgClientToGCGauntletGrantChoiceResponse
func (d *Artifact) GrantGauntletChoice(
	ctx context.Context,
	gauntletID uint32,
	grantID uint32,
	choiceDefIndex []uint32,
) (*protocol.CMsgClientToGCGauntletGrantChoiceResponse, error) {
	req := &protocol.CMsgClientToGCGauntletGrantChoice{
		GauntletId:     &gauntletID,
		GrantId:        &grantID,
		ChoiceDefIndex: choiceDefIndex,
	}
	resp := &protocol.CMsgClientToGCGauntletGrantChoiceResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGauntletGrantChoice),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGauntletGrantChoiceResponse),
		resp,
	)
}

// JoinChatChannel joins a chat channel.
// Request ID: k_EMsgClientToGCJoinChatChannel
// Response ID: k_EMsgClientToGCJoinChatChannelResponse
// Request type: CMsgClientToGCJoinChatChannel
// Response type: CMsgClientToGCJoinChatChannelResponse
func (d *Artifact) JoinChatChannel(
	ctx context.Context,
	roomType protocol.EChatRoomType,
	roomKey uint64,
	subRoomIndex uint32,
	requestID uint32,
) (*protocol.CMsgClientToGCJoinChatChannelResponse, error) {
	req := &protocol.CMsgClientToGCJoinChatChannel{
		RoomType:     &roomType,
		RoomKey:      &roomKey,
		SubRoomIndex: &subRoomIndex,
		RequestId:    &requestID,
	}
	resp := &protocol.CMsgClientToGCJoinChatChannelResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCJoinChatChannel),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCJoinChatChannelResponse),
		resp,
	)
}

// JoinGauntlet joins a gauntlet.
// Request ID: k_EMsgClientToGCJoinGauntlet
// Response ID: k_EMsgClientToGCJoinGauntletResponse
// Request type: CMsgClientToGCJoinGauntlet
// Response type: CMsgClientToGCJoinGauntletResponse
func (d *Artifact) JoinGauntlet(
	ctx context.Context,
	gauntletID uint32,
	deckCode []byte,
) (*protocol.CMsgClientToGCJoinGauntletResponse, error) {
	req := &protocol.CMsgClientToGCJoinGauntlet{
		GauntletId: &gauntletID,
		DeckCode:   deckCode,
	}
	resp := &protocol.CMsgClientToGCJoinGauntletResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCJoinGauntlet),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCJoinGauntletResponse),
		resp,
	)
}

// JoinLobbyPrivate joins a lobby private.
// Request ID: k_EMsgClientToGCPrivateLobbyJoin
// Response ID: k_EMsgClientToGCPrivateLobbyJoinResponse
// Request type: CMsgClientToGCPrivateLobbyJoin
// Response type: CMsgClientToGCPrivateLobbyJoinResponse
func (d *Artifact) JoinLobbyPrivate(
	ctx context.Context,
	privateLobbyID uint64,
) (*protocol.CMsgClientToGCPrivateLobbyJoinResponse, error) {
	req := &protocol.CMsgClientToGCPrivateLobbyJoin{
		PrivateLobbyId: &privateLobbyID,
	}
	resp := &protocol.CMsgClientToGCPrivateLobbyJoinResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyJoin),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyJoinResponse),
		resp,
	)
}

// LeaveChatChannel leaves a chat channel.
// Request ID: k_EMsgClientToGCLeaveChatChannel
// Request type: CMsgClientToGCLeaveChatChannel
func (d *Artifact) LeaveChatChannel(
	chatRoomID uint64,
) {
	req := &protocol.CMsgClientToGCLeaveChatChannel{
		ChatRoomId: &chatRoomID,
	}
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCLeaveChatChannel), req)
}

// LeaveChatChannelByKey leaves a chat channel by key.
// Request ID: k_EMsgClientToGCLeaveChatChannelByKey
// Request type: CMsgClientToGCLeaveChatChannelByKey
func (d *Artifact) LeaveChatChannelByKey(
	roomType protocol.EChatRoomType,
	roomKey uint64,
) {
	req := &protocol.CMsgClientToGCLeaveChatChannelByKey{
		RoomType: &roomType,
		RoomKey:  &roomKey,
	}
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCLeaveChatChannelByKey), req)
}

// LeaveLobby leaves a lobby.
// Request ID: k_EMsgClientToGCLeaveLobby
// Response ID: k_EMsgClientToGCLeaveLobbyResponse
// Request type: CMsgClientToGCLeaveLobby
// Response type: CMsgClientToGCLeaveLobbyResponse
func (d *Artifact) LeaveLobby(
	ctx context.Context,
	lobbyID uint64,
) (*protocol.CMsgClientToGCLeaveLobbyResponse, error) {
	req := &protocol.CMsgClientToGCLeaveLobby{
		LobbyId: &lobbyID,
	}
	resp := &protocol.CMsgClientToGCLeaveLobbyResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCLeaveLobby),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCLeaveLobbyResponse),
		resp,
	)
}

// LeaveLobbyPrivate leaves a lobby private.
// Request ID: k_EMsgClientToGCPrivateLobbyLeave
// Response ID: k_EMsgClientToGCPrivateLobbyLeaveResponse
// Request type: CMsgClientToGCPrivateLobbyLeave
// Response type: CMsgClientToGCPrivateLobbyLeaveResponse
func (d *Artifact) LeaveLobbyPrivate(
	ctx context.Context,
	privateLobbyID uint64,
) (*protocol.CMsgClientToGCPrivateLobbyLeaveResponse, error) {
	req := &protocol.CMsgClientToGCPrivateLobbyLeave{
		PrivateLobbyId: &privateLobbyID,
	}
	resp := &protocol.CMsgClientToGCPrivateLobbyLeaveResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyLeave),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyLeaveResponse),
		resp,
	)
}

// ListRollLootDev lists a roll loot dev.
// Request ID: k_EMsgClientToGCDevRollLootList
// Response ID: k_EMsgClientToGCDevRollLootListResponse
// Request type: CMsgClientToGCDevRollLootList
// Response type: CMsgClientToGCDevRollLootListResponse
func (d *Artifact) ListRollLootDev(
	ctx context.Context,
	lootListName string,
	itemDefIndex uint32,
	numRolls uint32,
	randomSeed uint32,
) (*protocol.CMsgClientToGCDevRollLootListResponse, error) {
	req := &protocol.CMsgClientToGCDevRollLootList{
		LootListName: &lootListName,
		ItemDefIndex: &itemDefIndex,
		NumRolls:     &numRolls,
		RandomSeed:   &randomSeed,
	}
	resp := &protocol.CMsgClientToGCDevRollLootListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCDevRollLootList),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCDevRollLootListResponse),
		resp,
	)
}

// OpenPhantomPack opens a phantom pack.
// Request ID: k_EMsgClientToGCOpenPhantomPack
// Response ID: k_EMsgClientToGCOpenPhantomPackResponse
// Request type: CMsgClientToGCOpenPhantomPack
// Response type: CMsgClientToGCOpenPhantomPackResponse
func (d *Artifact) OpenPhantomPack(
	ctx context.Context,
	phantomItemID uint64,
) (*protocol.CMsgClientToGCOpenPhantomPackResponse, error) {
	req := &protocol.CMsgClientToGCOpenPhantomPack{
		PhantomItemId: &phantomItemID,
	}
	resp := &protocol.CMsgClientToGCOpenPhantomPackResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCOpenPhantomPack),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCOpenPhantomPackResponse),
		resp,
	)
}

// RecycleCards recycles cards.
// Request ID: k_EMsgClientToGCRecycleCards
// Response ID: k_EMsgClientToGCRecycleCardsResponse
// Request type: CMsgClientToGCRecycleCards
// Response type: CMsgClientToGCRecycleCardsResponse
func (d *Artifact) RecycleCards(
	ctx context.Context,
	itemIDs []uint64,
	expectedReturn uint32,
) (*protocol.CMsgClientToGCRecycleCardsResponse, error) {
	req := &protocol.CMsgClientToGCRecycleCards{
		ItemIds:        itemIDs,
		ExpectedReturn: &expectedReturn,
	}
	resp := &protocol.CMsgClientToGCRecycleCardsResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCRecycleCards),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCRecycleCardsResponse),
		resp,
	)
}

// RedeemCurrencyForItem redeems a currency for item.
// Request ID: k_EMsgClientToGCRedeemCurrencyForItem
// Response ID: k_EMsgClientToGCRedeemCurrencyForItemResponse
// Request type: CMsgClientToGCRedeemCurrencyForItem
// Response type: CMsgClientToGCRedeemCurrencyForItemResponse
func (d *Artifact) RedeemCurrencyForItem(
	ctx context.Context,
	redeemForDefIndex uint32,
	useCurrencyID uint32,
	expectedCost uint32,
) (*protocol.CMsgClientToGCRedeemCurrencyForItemResponse, error) {
	req := &protocol.CMsgClientToGCRedeemCurrencyForItem{
		RedeemForDefIndex: &redeemForDefIndex,
		UseCurrencyId:     &useCurrencyID,
		ExpectedCost:      &expectedCost,
	}
	resp := &protocol.CMsgClientToGCRedeemCurrencyForItemResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCRedeemCurrencyForItem),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCRedeemCurrencyForItemResponse),
		resp,
	)
}

// RedeemCurrencyForItemList redeems a currency for item list.
// Request ID: k_EMsgClientToGCRedeemCurrencyForItemList
// Response ID: k_EMsgClientToGCRedeemCurrencyForItemListResponse
// Request type: CMsgClientToGCRedeemCurrencyForItemList
// Response type: CMsgClientToGCRedeemCurrencyForItemListResponse
func (d *Artifact) RedeemCurrencyForItemList(
	ctx context.Context,
	redeemForDefIndex []uint32,
	expectedCost []*protocol.CMsgClientToGCRedeemCurrencyForItemList_Currency,
) (*protocol.CMsgClientToGCRedeemCurrencyForItemListResponse, error) {
	req := &protocol.CMsgClientToGCRedeemCurrencyForItemList{
		RedeemForDefIndex: redeemForDefIndex,
		ExpectedCost:      expectedCost,
	}
	resp := &protocol.CMsgClientToGCRedeemCurrencyForItemListResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCRedeemCurrencyForItemList),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCRedeemCurrencyForItemListResponse),
		resp,
	)
}

// ReportDebug reports a debug.
// Request ID: k_EMsgClientToGCDebugReport
// Request type: CMsgClientToGCDebugReport
func (d *Artifact) ReportDebug(
	buildVersion uint32,
	eventList []*protocol.CMsgClientToGCDebugReport_Event,
) {
	req := &protocol.CMsgClientToGCDebugReport{
		BuildVersion: &buildVersion,
		EventList:    eventList,
	}
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCDebugReport), req)
}

// SendAIMatchStarted sends a ai match started.
// Request ID: k_EMsgClientToGCAIMatchStarted
// Request type: CMsgClientToGCAIMatchStarted
func (d *Artifact) SendAIMatchStarted(
	difficulty uint32,
	startTime uint32,
	playerDeck []byte,
	aiDeck []byte,
) {
	req := &protocol.CMsgClientToGCAIMatchStarted{
		Difficulty: &difficulty,
		StartTime:  &startTime,
		PlayerDeck: playerDeck,
		AiDeck:     aiDeck,
	}
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCAIMatchStarted), req)
}

// SendChatMessage sends a chat message.
// Request ID: k_EMsgClientToGCSendChatMessage
// Request type: CMsgClientToGCSendChatMessage
func (d *Artifact) SendChatMessage(
	chatRoomID uint64,
	chatMsg string,
) {
	req := &protocol.CMsgClientToGCSendChatMessage{
		ChatRoomId: &chatRoomID,
		ChatMsg:    &chatMsg,
	}
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCSendChatMessage), req)
}

// SendChatMessageRoll sends a chat message roll.
// Request ID: k_EMsgClientToGCSendChatMessageRoll
// Request type: CMsgClientToGCSendChatMessageRoll
func (d *Artifact) SendChatMessageRoll(
	chatRoomID uint64,
	rollMin uint32,
	rollMax uint32,
) {
	req := &protocol.CMsgClientToGCSendChatMessageRoll{
		ChatRoomId: &chatRoomID,
		RollMin:    &rollMin,
		RollMax:    &rollMax,
	}
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCSendChatMessageRoll), req)
}

// SendClaimQuestMapNode sends a claim quest map node.
// Request ID: k_EMsgClientToGCClaimQuestMapNode
// Response ID: k_EMsgClientToGCClaimQuestMapNodeResponse
// Request type: CMsgClientToGCClaimQuestMapNode
// Response type: CMsgClientToGCClaimQuestMapNodeResponse
func (d *Artifact) SendClaimQuestMapNode(
	ctx context.Context,
	questID uint32,
	nodeID uint32,
	deferredNode bool,
) (*protocol.CMsgClientToGCClaimQuestMapNodeResponse, error) {
	req := &protocol.CMsgClientToGCClaimQuestMapNode{
		QuestId:      &questID,
		NodeId:       &nodeID,
		DeferredNode: &deferredNode,
	}
	resp := &protocol.CMsgClientToGCClaimQuestMapNodeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCClaimQuestMapNode),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCClaimQuestMapNodeResponse),
		resp,
	)
}

// SendDeckBuilderStats sends deck builder stats.
// Request ID: k_EMsgClientToGCDeckBuilderStats
// Request type: CMsgClientToGCDeckBuilderStats
func (d *Artifact) SendDeckBuilderStats(
	timeSpentS uint32,
) {
	req := &protocol.CMsgClientToGCDeckBuilderStats{
		TimeSpentS: &timeSpentS,
	}
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCDeckBuilderStats), req)
}

// SendDownloadReplay sends a download replay.
// Request ID: k_EMsgClientToGCDownloadReplay
// Response ID: k_EMsgClientToGCDownloadReplayResponse
// Request type: CMsgClientToGCDownloadReplay
// Response type: CMsgClientToGCDownloadReplayResponse
func (d *Artifact) SendDownloadReplay(
	ctx context.Context,
	matchID uint64,
	replayType protocol.CMsgClientToGCDownloadReplay_EType,
	requireVersion uint32,
) (*protocol.CMsgClientToGCDownloadReplayResponse, error) {
	req := &protocol.CMsgClientToGCDownloadReplay{
		MatchId:        &matchID,
		ReplayType:     &replayType,
		RequireVersion: &requireVersion,
	}
	resp := &protocol.CMsgClientToGCDownloadReplayResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCDownloadReplay),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCDownloadReplayResponse),
		resp,
	)
}

// SendGauntletScoreboard sends a gauntlet scoreboard.
// Request ID: k_EMsgClientToGCGauntletScoreboard
// Response ID: k_EMsgClientToGCGauntletScoreboardResponse
// Request type: CMsgClientToGCGauntletScoreboard
// Response type: CMsgClientToGCGauntletScoreboardResponse
func (d *Artifact) SendGauntletScoreboard(
	ctx context.Context,
	gauntletID uint32,
) (*protocol.CMsgClientToGCGauntletScoreboardResponse, error) {
	req := &protocol.CMsgClientToGCGauntletScoreboard{
		GauntletId: &gauntletID,
	}
	resp := &protocol.CMsgClientToGCGauntletScoreboardResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGauntletScoreboard),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCGauntletScoreboardResponse),
		resp,
	)
}

// SendIsInMatchmaking sends a is in matchmaking.
// Request ID: k_EMsgClientToGCIsInMatchmaking
// Response ID: k_EMsgClientToGCIsInMatchmakingResponse
// Request type: CMsgClientToGCIsInMatchmaking
// Response type: CMsgClientToGCIsInMatchmakingResponse
func (d *Artifact) SendIsInMatchmaking(
	ctx context.Context,
) (*protocol.CMsgClientToGCIsInMatchmakingResponse, error) {
	req := &protocol.CMsgClientToGCIsInMatchmaking{}
	resp := &protocol.CMsgClientToGCIsInMatchmakingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCIsInMatchmaking),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCIsInMatchmakingResponse),
		resp,
	)
}

// SendMarketPrices sends market prices.
// Request ID: k_EMsgClientToGCMarketPrices
// Request type: CMsgClientToGCMarketPrices
func (d *Artifact) SendMarketPrices(
	currencyID uint32,
	priorTimeStamp uint32,
) {
	req := &protocol.CMsgClientToGCMarketPrices{
		CurrencyId:     &currencyID,
		PriorTimeStamp: &priorTimeStamp,
	}
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCMarketPrices), req)
}

// SendMatchResultScreenAnalytics sends match result screen analytics.
// Request ID: k_EMsgClientToGCMatchResultScreenAnalytics
// Request type: CMsgClientToGCMatchResultScreenAnalytics
func (d *Artifact) SendMatchResultScreenAnalytics(
	startingGraph uint32,
	matchID uint64,
	graphStats []*protocol.CMsgClientToGCMatchResultScreenAnalytics_GraphStats,
) {
	req := &protocol.CMsgClientToGCMatchResultScreenAnalytics{
		StartingGraph: &startingGraph,
		MatchId:       &matchID,
		GraphStats:    graphStats,
	}
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCMatchResultScreenAnalytics), req)
}

// SendPrivateLobbyAction sends a private lobby action.
// Request ID: k_EMsgClientToGCPrivateLobbyAction
// Response ID: k_EMsgClientToGCPrivateLobbyActionResponse
// Request type: CMsgClientToGCPrivateLobbyAction
// Response type: CMsgClientToGCPrivateLobbyActionResponse
func (d *Artifact) SendPrivateLobbyAction(
	ctx context.Context,
	req *protocol.CMsgClientToGCPrivateLobbyAction,
) (*protocol.CMsgClientToGCPrivateLobbyActionResponse, error) {
	resp := &protocol.CMsgClientToGCPrivateLobbyActionResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyAction),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyActionResponse),
		resp,
	)
}

// SendPrivateLobbyChallenge sends a private lobby challenge.
// Request ID: k_EMsgClientToGCPrivateLobbyChallenge
// Response ID: k_EMsgClientToGCPrivateLobbyChallengeResponse
// Request type: CMsgClientToGCPrivateLobbyChallenge
// Response type: CMsgClientToGCPrivateLobbyChallengeResponse
func (d *Artifact) SendPrivateLobbyChallenge(
	ctx context.Context,
	challengeAccountID uint32,
) (*protocol.CMsgClientToGCPrivateLobbyChallengeResponse, error) {
	req := &protocol.CMsgClientToGCPrivateLobbyChallenge{
		ChallengeAccountId: &challengeAccountID,
	}
	resp := &protocol.CMsgClientToGCPrivateLobbyChallengeResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyChallenge),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyChallengeResponse),
		resp,
	)
}

// SendPrivateLobbyClientVersion sends a private lobby client version.
// Request ID: k_EMsgClientToGCPrivateLobbyClientVersion
// Request type: CMsgClientToGCPrivateLobbyClientVersion
func (d *Artifact) SendPrivateLobbyClientVersion(
	privateLobbyID uint64,
) {
	req := &protocol.CMsgClientToGCPrivateLobbyClientVersion{
		PrivateLobbyId: &privateLobbyID,
	}
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyClientVersion), req)
}

// SendPrivateLobbyInviteUser sends a private lobby invite user.
// Request ID: k_EMsgClientToGCPrivateLobbyInviteUser
// Response ID: k_EMsgClientToGCPrivateLobbyInviteUserResponse
// Request type: CMsgClientToGCPrivateLobbyInviteUser
// Response type: CMsgClientToGCPrivateLobbyInviteUserResponse
func (d *Artifact) SendPrivateLobbyInviteUser(
	ctx context.Context,
	privateLobbyID uint64,
	inviteAccountID uint32,
) (*protocol.CMsgClientToGCPrivateLobbyInviteUserResponse, error) {
	req := &protocol.CMsgClientToGCPrivateLobbyInviteUser{
		PrivateLobbyId:  &privateLobbyID,
		InviteAccountId: &inviteAccountID,
	}
	resp := &protocol.CMsgClientToGCPrivateLobbyInviteUserResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyInviteUser),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyInviteUserResponse),
		resp,
	)
}

// SendQuestDevCommand sends a quest dev command.
// Request ID: k_EMsgClientToGCQuestDevCommand
// Request type: CMsgClientToGCQuestDevCommand
func (d *Artifact) SendQuestDevCommand(
	questID uint32,
	resetProgress bool,
	grantCurrency int32,
) {
	req := &protocol.CMsgClientToGCQuestDevCommand{
		QuestId:       &questID,
		ResetProgress: &resetProgress,
		GrantCurrency: &grantCurrency,
	}
	d.write(uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCQuestDevCommand), req)
}

// SendRegisterGauntletDeck sends a register gauntlet deck.
// Request ID: k_EMsgClientToGCRegisterGauntletDeck
// Response ID: k_EMsgClientToGCRegisterGauntletDeckResponse
// Request type: CMsgClientToGCRegisterGauntletDeck
// Response type: CMsgClientToGCRegisterGauntletDeckResponse
func (d *Artifact) SendRegisterGauntletDeck(
	ctx context.Context,
	gauntletID uint32,
	deckBytes []byte,
) (*protocol.CMsgClientToGCRegisterGauntletDeckResponse, error) {
	req := &protocol.CMsgClientToGCRegisterGauntletDeck{
		GauntletId: &gauntletID,
		DeckBytes:  deckBytes,
	}
	resp := &protocol.CMsgClientToGCRegisterGauntletDeckResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCRegisterGauntletDeck),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCRegisterGauntletDeckResponse),
		resp,
	)
}

// SendRegisterTournamentDeck sends a register tournament deck.
// Request ID: k_EMsgClientToGCRegisterTournamentDeck
// Response ID: k_EMsgClientToGCRegisterTournamentDeckResponse
// Request type: CMsgClientToGCRegisterTournamentDeck
// Response type: CMsgClientToGCRegisterTournamentDeckResponse
func (d *Artifact) SendRegisterTournamentDeck(
	ctx context.Context,
	deckCode string,
) (*protocol.CMsgClientToGCRegisterTournamentDeckResponse, error) {
	req := &protocol.CMsgClientToGCRegisterTournamentDeck{
		DeckCode: &deckCode,
	}
	resp := &protocol.CMsgClientToGCRegisterTournamentDeckResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCRegisterTournamentDeck),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCRegisterTournamentDeckResponse),
		resp,
	)
}

// SendReplacementSDRTicket sends a replacement sdr ticket.
// Request ID: k_EMsgClientToGCReplacementSDRTicket
// Response ID: k_EMsgClientToGCReplacementSDRTicketResponse
// Request type: CMsgClientToGCReplacementSDRTicket
// Response type: CMsgClientToGCReplacementSDRTicketResponse
func (d *Artifact) SendReplacementSDRTicket(
	ctx context.Context,
	lobbyID uint64,
) (*protocol.CMsgClientToGCReplacementSDRTicketResponse, error) {
	req := &protocol.CMsgClientToGCReplacementSDRTicket{
		LobbyId: &lobbyID,
	}
	resp := &protocol.CMsgClientToGCReplacementSDRTicketResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCReplacementSDRTicket),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCReplacementSDRTicketResponse),
		resp,
	)
}

// SendStopMatchmaking sends a stop matchmaking.
// Request ID: k_EMsgClientToGCStopMatchmaking
// Response ID: k_EMsgClientToGCStopMatchmakingResponse
// Request type: CMsgClientToGCStopMatchmaking
// Response type: CMsgClientToGCStopMatchmakingResponse
func (d *Artifact) SendStopMatchmaking(
	ctx context.Context,
) (*protocol.CMsgClientToGCStopMatchmakingResponse, error) {
	req := &protocol.CMsgClientToGCStopMatchmaking{}
	resp := &protocol.CMsgClientToGCStopMatchmakingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCStopMatchmaking),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCStopMatchmakingResponse),
		resp,
	)
}

// SendUpdateDeckRegistry sends a update deck registry.
// Request ID: k_EMsgClientToGCUpdateDeckRegistry
// Response ID: k_EMsgClientToGCUpdateDeckRegistryResponse
// Request type: CMsgClientToGCUpdateDeckRegistry
// Response type: CMsgClientToGCUpdateDeckRegistryResponse
func (d *Artifact) SendUpdateDeckRegistry(
	ctx context.Context,
	deckName string,
	deckTags string,
	deckCode string,
	deleteDeck bool,
) (*protocol.CMsgClientToGCUpdateDeckRegistryResponse, error) {
	req := &protocol.CMsgClientToGCUpdateDeckRegistry{
		DeckName:   &deckName,
		DeckTags:   &deckTags,
		DeckCode:   &deckCode,
		DeleteDeck: &deleteDeck,
	}
	resp := &protocol.CMsgClientToGCUpdateDeckRegistryResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCUpdateDeckRegistry),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCUpdateDeckRegistryResponse),
		resp,
	)
}

// SetGlobalPhantomLeague sets a global phantom league.
// Request ID: k_EMsgClientToGCSetGlobalPhantomLeague
// Response ID: k_EMsgClientToGCSetGlobalPhantomLeagueResponse
// Request type: CMsgClientToGCSetGlobalPhantomLeague
// Response type: CMsgClientToGCSetGlobalPhantomLeagueResponse
func (d *Artifact) SetGlobalPhantomLeague(
	ctx context.Context,
	phantomLeagueID uint32,
) (*protocol.CMsgClientToGCSetGlobalPhantomLeagueResponse, error) {
	req := &protocol.CMsgClientToGCSetGlobalPhantomLeague{
		PhantomLeagueId: &phantomLeagueID,
	}
	resp := &protocol.CMsgClientToGCSetGlobalPhantomLeagueResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCSetGlobalPhantomLeague),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCSetGlobalPhantomLeagueResponse),
		resp,
	)
}

// SpectateUser spectates a user.
// Request ID: k_EMsgClientToGCSpectateUser
// Response ID: k_EMsgClientToGCSpectateUserResponse
// Request type: CMsgClientToGCSpectateUser
// Response type: CMsgClientToGCSpectateUserResponse
func (d *Artifact) SpectateUser(
	ctx context.Context,
	spectateAccountID uint32,
) (*protocol.CMsgClientToGCSpectateUserResponse, error) {
	req := &protocol.CMsgClientToGCSpectateUser{
		SpectateAccountId: &spectateAccountID,
	}
	resp := &protocol.CMsgClientToGCSpectateUserResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCSpectateUser),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCSpectateUserResponse),
		resp,
	)
}

// StartLobbyPrivateMatch starts a lobby private match.
// Request ID: k_EMsgClientToGCPrivateLobbyStartMatch
// Response ID: k_EMsgClientToGCPrivateLobbyStartMatchResponse
// Request type: CMsgClientToGCPrivateLobbyStartMatch
// Response type: CMsgClientToGCPrivateLobbyStartMatchResponse
func (d *Artifact) StartLobbyPrivateMatch(
	ctx context.Context,
	privateLobbyID uint64,
) (*protocol.CMsgClientToGCPrivateLobbyStartMatchResponse, error) {
	req := &protocol.CMsgClientToGCPrivateLobbyStartMatch{
		PrivateLobbyId: &privateLobbyID,
	}
	resp := &protocol.CMsgClientToGCPrivateLobbyStartMatchResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyStartMatch),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCPrivateLobbyStartMatchResponse),
		resp,
	)
}

// StartMatchmaking starts a matchmaking.
// Request ID: k_EMsgClientToGCStartMatchmaking
// Response ID: k_EMsgClientToGCStartMatchmakingResponse
// Request type: CMsgClientToGCStartMatchmaking
// Response type: CMsgClientToGCStartMatchmakingResponse
func (d *Artifact) StartMatchmaking(
	ctx context.Context,
	matchInfo protocol.CMsgStartFindingMatchInfo,
) (*protocol.CMsgClientToGCStartMatchmakingResponse, error) {
	req := &protocol.CMsgClientToGCStartMatchmaking{
		MatchInfo: &matchInfo,
	}
	resp := &protocol.CMsgClientToGCStartMatchmakingResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCStartMatchmaking),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCStartMatchmakingResponse),
		resp,
	)
}

// SubmitUserFeedback submits a user feedback.
// Request ID: k_EMsgClientToGCSubmitUserFeedback
// Response ID: k_EMsgClientToGCSubmitUserFeedbackResponse
// Request type: CMsgClientToGCSubmitUserFeedback
// Response type: CMsgClientToGCSubmitUserFeedbackResponse
func (d *Artifact) SubmitUserFeedback(
	ctx context.Context,
	messageText string,
	feedbackType protocol.CMsgClientToGCSubmitUserFeedback_EType,
	sourceInfo string,
) (*protocol.CMsgClientToGCSubmitUserFeedbackResponse, error) {
	req := &protocol.CMsgClientToGCSubmitUserFeedback{
		MessageText:  &messageText,
		FeedbackType: &feedbackType,
		SourceInfo:   &sourceInfo,
	}
	resp := &protocol.CMsgClientToGCSubmitUserFeedbackResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCSubmitUserFeedback),
		req,
		uint32(protocol.EGCDCGClientMessages_k_EMsgClientToGCSubmitUserFeedbackResponse),
		resp,
	)
}

// registerGeneratedHandlers registers the auto-generated event handlers.
func (d *Artifact) registerGeneratedHandlers() {
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientAvailableGauntlets)] = d.getEventEmitter(func() events.Event {
		return &events.AvailableGauntlets{}
	})
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientChatChannelJoined)] = d.getEventEmitter(func() events.Event {
		return &events.ChatChannelJoined{}
	})
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientChatMessage)] = d.getEventEmitter(func() events.Event {
		return &events.ChatMessage{}
	})
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientDevLobbyStatus)] = d.getEventEmitter(func() events.Event {
		return &events.DevLobbyStatus{}
	})
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientGlobalPhantomLeagues)] = d.getEventEmitter(func() events.Event {
		return &events.GlobalPhantomLeagues{}
	})
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientMarketPrices)] = d.getEventEmitter(func() events.Event {
		return &events.MarketPrices{}
	})
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientMatchmakingStatus)] = d.getEventEmitter(func() events.Event {
		return &events.MatchmakingStatus{}
	})
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientMatchmakingStopped)] = d.getEventEmitter(func() events.Event {
		return &events.MatchmakingStopped{}
	})
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientMessageOfTheDay)] = d.getEventEmitter(func() events.Event {
		return &events.MessageOfTheDay{}
	})
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientPrivateLobbyEvent)] = d.getEventEmitter(func() events.Event {
		return &events.PrivateLobbyEvent{}
	})
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientSDRTicket)] = d.getEventEmitter(func() events.Event {
		return &events.SDRTicket{}
	})
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientTournamentState)] = d.getEventEmitter(func() events.Event {
		return &events.TournamentState{}
	})
	d.handlers[uint32(protocol.EGCDCGClientMessages_k_EMsgGCToClientUserJoinedChatChannel)] = d.getEventEmitter(func() events.Event {
		return &events.UserJoinedChatChannel{}
	})
}
