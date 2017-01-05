package gameserver

const (
	// For Select Player Process
	CM_QUERYCHR = 100
	CM_NEWCHR = 101
	CM_DELCHR = 102
	CM_SELCHR = 103

	SM_QUERYCHR = 520
	SM_NEWCHR_SUCCESS = 521
	SM_NEWCHR_FAIL = 522
	SM_DELCHR_SUCCESS = 523
	SM_DELCHR_FAIL = 524
	SM_STARTPLAY = 525
	SM_STARTFAIL = 526
	SM_QUERYCHR_FAIL = 527

	// For game process
	// Client To Server Commands
	CM_GAMELOGIN = 65001
	CM_QUERYUSERNAME = 80
	CM_QUERYBAGITEMS = 81
	CM_QUERYUSERSTATE = 82

	CM_DROPITEM = 1000
	CM_PICKUP = 1001
	CM_TAKEONITEM = 1003
	CM_TAKEOFFITEM = 1004
	CM_BUTCH = 1007
	CM_MAGICKEYCHANGE = 1008
	CM_LOGINNOTICEOK = 1018

	CM_EAT = 1006

	CM_CLICKNPC = 1010
	CM_MERCHANTDLGSELECT = 1011
	CM_USERSELLITEM = 1013
	CM_WANTMINIMAP = 1033

	CM_TURN = 3010
	CM_WALK = 3011
	CM_SITDOWN = 3012
	CM_RUN = 3013
	CM_HIT = 3014
	CM_HEAVYHIT = 3015
	CM_BIGHIT = 3016
	CM_SPELL = 3017
	CM_POWERHIT = 3018
	CM_LONGHIT = 3019
	CM_WIDEHIT = 3024
	CM_FIREHIT = 3025
	CM_SAY = 3030
	CM_RIDE = 3031

	// Server to Client Commands
	SM_RUSH = 6
	SM_FIREHIT = 8
	SM_BACKSTEP = 9
	SM_TURN = 10
	SM_WALK = 11
	SM_SITDOWN = 12
	SM_RUN = 13
	SM_HIT = 14
	SM_SPELL = 17
	SM_POWERHIT = 18
	SM_LONGHIT = 19
	SM_DIGUP = 20
	SM_DIGDOWN = 21
	SM_FLYAXE = 22
	SM_LIGHTING = 23
	SM_WIDEHIT = 24
	SM_DISAPPEAR = 30
	SM_STRUCK = 31
	SM_DEATH = 32
	SM_SKELETON = 33
	SM_NOWDEATH = 34

	SM_HEAR = 40
	SM_FEATURECHANGED = 41
	SM_USERNAME = 42
	SM_WINEXP = 44
	SM_LEVELUP = 45
	SM_DAYCHANGING = 46
	SM_LOGON = 50
	SM_NEWMAP = 51
	SM_ABILITY = 52
	SM_HEALTHSPELLCHANGED = 53
	SM_MAPDESCRIPTION = 54
	SM_GAMEGOLDNAME = 55
	SM_SPELL2 = 117

	SM_SYSMESSAGE = 100
	SM_GROUPMESSAGE = 101
	SM_CRY = 102
	SM_WHISPER = 103
	SM_GUILDMESSAGE = 104

	SM_ADDITEM = 200
	SM_BAGITEMS = 201
	SM_DELITEM = 202

	SM_ADDMAGIC = 210
	SM_SENDMYMAGIC = 211
	SM_DELMAGIC = 212

	SM_DROPITEM_SUCCESS = 600
	SM_DROPITEM_FAIL = 601
	SM_ITEMSHOW = 610
	SM_ITEMHIDE = 611
	SM_DOOROPEN = 612
	SM_TAKEON_OK = 615
	SM_TAKEON_FAIL = 616
	SM_TAKEOFF_OK = 619
	SM_TAKEOFF_FAIL = 620
	SM_SENDUSEITEMS = 621
	SM_WEIGHTCHANGED = 622
	SM_CLEAROBJECTS = 633
	SM_CHANGEMAP = 634
	SM_EAT_OK = 635
	SM_EAT_FAIL = 636
	SM_BUTCH = 637
	SM_MAGICFIRE = 638
	SM_MAGIC_LVEXP = 640
	SM_DURACHANGE = 642

	SM_MERCHANTSAY = 643
	SM_SENDUSERSELL = 646
	SM_USERSELLITEM_OK = 648
	SM_USERSELLITEM_FAIL = 649

	SM_GOLDCHANGED = 653
	SM_CHANGELIGHT = 654
	SM_CHANGENAMECOLOR = 656
	SM_CHARSTATUSCHANGED = 657
	SM_SENDNOTICE = 658

	SM_AREASTATE = 708
	SM_DELITEMS = 709
	SM_READMINIMAP_OK = 710
	SM_READMINIMAP_FAIL = 711
	SM_SUBABILITY = 752

	SM_MENU_OK = 767

	SM_SHOWEVENT = 804
	SM_HIDEEVENT = 805
	SM_OPENHEALTH = 1100
	SM_CLOSEHEALTH = 1101
	SM_CHANGEFACE = 1104
	SM_VERSION_FAIL = 1106
	SM_ITEMUPDATE = 1500
	SM_MONSTERSAY = 1501
	UNKNOWN_1000 = 11088
)