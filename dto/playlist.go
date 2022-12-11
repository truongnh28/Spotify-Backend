package dto

type PlayList struct {
	PlayListID uint   `json:"play_list_id"`
	Name       string `json:"name"`
	CoverImg   string `json:"cover_img"`
	UserID     uint   `json:"user_id"`
}

type PlayListResponse struct {
	PlayLists []PlayList `json:"play_lists"`
	StatusError
}
