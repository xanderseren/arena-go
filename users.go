package arena

// MyAnimeList API docs: http://myanimelist.net/modules.php?go=api
type UsersService struct {
	client *Client
}

type UserStruct struct {
	ID          int    `json:"id"`
	Slug        string `json:"slug"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	FullName    string `json:"full_name"`
	Avatar      string `json:"avatar"`
	AvatarImage struct {
		Thumb   string `json:"thumb"`
		Display string `json:"display"`
	} `json:"avatar_image"`
	ChannelCount   int            `json:"channel_count"`
	FollowingCount int            `json:"following_count"`
	ProfileId      int            `json:"profile_id"`
	FollowerCount  int            `json:"follower_count"`
	Initials       string         `json:"initials"`
	CanIndex       bool           `json:"can_index"`
	IsPremium      bool           `json:"is_premium"`
	Metadata       MetadataStruct `json:"metadata"`
	BaseClass      string         `json:"base_class"`
	Class          string         `json:"class"`
}

// Search performs a search of users based on the search query q
func (s *UsersService) Search(args Arguments) (searchStruct *SearchStruct, err error) {
	path := "search/users"
	err = s.client.Get(path, args, &searchStruct)
	if searchStruct != nil {
		searchStruct.client = s.client
	}
	return
}
