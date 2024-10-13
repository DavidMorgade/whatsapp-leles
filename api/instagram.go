package api

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var url = "https://instagram-bulk-scraper-latest.p.rapidapi.com/media_info_from_shortcode/CwqI-QTpUG2"

var apiKey string

type Response struct {
	Data    Data    `json:"data"`
	Status  string  `json:"status"`
	Message *string `json:"message"`
}

type Data struct {
	ID                          string                     `json:"id"`
	Shortcode                   string                     `json:"shortcode"`
	ThumbnailSrc                string                     `json:"thumbnail_src"`
	Dimensions                  Dimensions                 `json:"dimensions"`
	GatingInfo                  *string                    `json:"gating_info"`
	FactCheckOverallRating      *string                    `json:"fact_check_overall_rating"`
	FactCheckInformation        *string                    `json:"fact_check_information"`
	SensitivityFrictionInfo     *string                    `json:"sensitivity_friction_info"`
	SharingFrictionInfo         SharingFrictionInfo        `json:"sharing_friction_info"`
	MediaOverlayInfo            *string                    `json:"media_overlay_info"`
	MediaPreview                string                     `json:"media_preview"`
	DisplayURL                  string                     `json:"display_url"`
	DisplayResources            []DisplayResource          `json:"display_resources"`
	AccessibilityCaption        *string                    `json:"accessibility_caption"`
	DashInfo                    DashInfo                   `json:"dash_info"`
	HasAudio                    bool                       `json:"has_audio"`
	VideoURL                    string                     `json:"video_url"`
	VideoViewCount              int                        `json:"video_view_count"`
	VideoPlayCount              int                        `json:"video_play_count"`
	EncodingStatus              *string                    `json:"encoding_status"`
	IsPublished                 bool                       `json:"is_published"`
	ProductType                 string                     `json:"product_type"`
	Title                       string                     `json:"title"`
	VideoDuration               float64                    `json:"video_duration"`
	ClipsMusicAttributionInfo   ClipsMusicAttributionInfo  `json:"clips_music_attribution_info"`
	IsVideo                     bool                       `json:"is_video"`
	TrackingToken               string                     `json:"tracking_token"`
	UpcomingEvent               *string                    `json:"upcoming_event"`
	EdgeMediaToTaggedUser       EdgeMediaToTaggedUser      `json:"edge_media_to_tagged_user"`
	Owner                       Owner                      `json:"owner"`
	EdgeMediaToCaption          EdgeMediaToCaption         `json:"edge_media_to_caption"`
	CanSeeInsightsAsBrand       bool                       `json:"can_see_insights_as_brand"`
	CaptionIsEdited             bool                       `json:"caption_is_edited"`
	HasRankedComments           bool                       `json:"has_ranked_comments"`
	LikeAndViewCountsDisabled   bool                       `json:"like_and_view_counts_disabled"`
	EdgeMediaToParentComment    EdgeMediaToParentComment   `json:"edge_media_to_parent_comment"`
	EdgeMediaToHoistedComment   EdgeMediaToHoistedComment  `json:"edge_media_to_hoisted_comment"`
	EdgeMediaPreviewComment     EdgeMediaPreviewComment    `json:"edge_media_preview_comment"`
	CommentsDisabled            bool                       `json:"comments_disabled"`
	CommentingDisabledForViewer bool                       `json:"commenting_disabled_for_viewer"`
	TakenAtTimestamp            int64                      `json:"taken_at_timestamp"`
	EdgeMediaPreviewLike        EdgeMediaPreviewLike       `json:"edge_media_preview_like"`
	EdgeMediaToSponsorUser      EdgeMediaToSponsorUser     `json:"edge_media_to_sponsor_user"`
	IsAffiliate                 bool                       `json:"is_affiliate"`
	IsPaidPartnership           bool                       `json:"is_paid_partnership"`
	Location                    *string                    `json:"location"`
	NftAssetInfo                *string                    `json:"nft_asset_info"`
	ViewerHasLiked              bool                       `json:"viewer_has_liked"`
	ViewerHasSaved              bool                       `json:"viewer_has_saved"`
	ViewerHasSavedToCollection  bool                       `json:"viewer_has_saved_to_collection"`
	ViewerInPhotoOfYou          bool                       `json:"viewer_in_photo_of_you"`
	ViewerCanReshare            bool                       `json:"viewer_can_reshare"`
	IsAd                        bool                       `json:"is_ad"`
	EdgeWebMediaToRelatedMedia  EdgeWebMediaToRelatedMedia `json:"edge_web_media_to_related_media"`
	CoauthorProducers           []CoauthorProducer         `json:"coauthor_producers"`
	PinnedForUsers              []string                   `json:"pinned_for_users"`
}

type Dimensions struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type SharingFrictionInfo struct {
	ShouldHaveSharingFriction bool    `json:"should_have_sharing_friction"`
	BloksAppUrl               *string `json:"bloks_app_url"`
}

type DisplayResource struct {
	Src          string `json:"src"`
	ConfigWidth  int    `json:"config_width"`
	ConfigHeight int    `json:"config_height"`
}

type DashInfo struct {
	IsDashEligible    bool    `json:"is_dash_eligible"`
	VideoDashManifest *string `json:"video_dash_manifest"`
	NumberOfQualities int     `json:"number_of_qualities"`
}

type ClipsMusicAttributionInfo struct {
	ArtistName            string `json:"artist_name"`
	SongName              string `json:"song_name"`
	UsesOriginalAudio     bool   `json:"uses_original_audio"`
	ShouldMuteAudio       bool   `json:"should_mute_audio"`
	ShouldMuteAudioReason string `json:"should_mute_audio_reason"`
	AudioID               string `json:"audio_id"`
}

type EdgeMediaToTaggedUser struct {
	Edges []TaggedUserEdge `json:"edges"`
}

type TaggedUserEdge struct {
	Node TaggedUserNode `json:"node"`
}

type TaggedUserNode struct {
	User User   `json:"user"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
	ID   string `json:"id"`
}

type User struct {
	FullName         string `json:"full_name"`
	FollowedByViewer bool   `json:"followed_by_viewer"`
	ID               string `json:"id"`
	IsVerified       bool   `json:"is_verified"`
	ProfilePicUrl    string `json:"profile_pic_url"`
	Username         string `json:"username"`
}

type Owner struct {
	ID                        string                   `json:"id"`
	Username                  string                   `json:"username"`
	IsVerified                bool                     `json:"is_verified"`
	ProfilePicUrl             string                   `json:"profile_pic_url"`
	BlockedByViewer           bool                     `json:"blocked_by_viewer"`
	RestrictedByViewer        *string                  `json:"restricted_by_viewer"`
	FollowedByViewer          bool                     `json:"followed_by_viewer"`
	FullName                  string                   `json:"full_name"`
	HasBlockedViewer          bool                     `json:"has_blocked_viewer"`
	IsEmbedsDisabled          bool                     `json:"is_embeds_disabled"`
	IsPrivate                 bool                     `json:"is_private"`
	IsUnpublished             bool                     `json:"is_unpublished"`
	RequestedByViewer         bool                     `json:"requested_by_viewer"`
	PassTieringRecommendation bool                     `json:"pass_tiering_recommendation"`
	EdgeOwnerToTimelineMedia  EdgeOwnerToTimelineMedia `json:"edge_owner_to_timeline_media"`
	EdgeFollowedBy            EdgeFollowedBy           `json:"edge_followed_by"`
}

type EdgeOwnerToTimelineMedia struct {
	Count int `json:"count"`
}

type EdgeFollowedBy struct {
	Count int `json:"count"`
}

type EdgeMediaToCaption struct {
	Edges []CaptionEdge `json:"edges"`
}

type CaptionEdge struct {
	Node CaptionNode `json:"node"`
}

type CaptionNode struct {
	CreatedAt int64  `json:"created_at"`
	Text      string `json:"text"`
	ID        string `json:"id"`
}

type EdgeMediaToParentComment struct {
	Count    int           `json:"count"`
	PageInfo PageInfo      `json:"page_info"`
	Edges    []CommentEdge `json:"edges"`
}

type PageInfo struct {
	HasNextPage bool   `json:"has_next_page"`
	EndCursor   string `json:"end_cursor"`
}

type CommentEdge struct {
	Node CommentNode `json:"node"`
}

type CommentNode struct {
	ID                   string               `json:"id"`
	Text                 string               `json:"text"`
	CreatedAt            int64                `json:"created_at"`
	DidReportAsSpam      bool                 `json:"did_report_as_spam"`
	Owner                User                 `json:"owner"`
	ViewerHasLiked       bool                 `json:"viewer_has_liked"`
	EdgeLikedBy          EdgeLikedBy          `json:"edge_liked_by"`
	IsRestrictedPending  bool                 `json:"is_restricted_pending"`
	EdgeThreadedComments EdgeThreadedComments `json:"edge_threaded_comments"`
}

type EdgeLikedBy struct {
	Count int `json:"count"`
}

type EdgeThreadedComments struct {
	Count    int           `json:"count"`
	PageInfo PageInfo      `json:"page_info"`
	Edges    []CommentEdge `json:"edges"`
}

type EdgeMediaToHoistedComment struct {
	Edges []CommentEdge `json:"edges"`
}

type EdgeMediaPreviewComment struct {
	Count int           `json:"count"`
	Edges []CommentEdge `json:"edges"`
}

type EdgeMediaPreviewLike struct {
	Count int        `json:"count"`
	Edges []LikeEdge `json:"edges"`
}

type LikeEdge struct {
	Node LikeNode `json:"node"`
}

type LikeNode struct {
	ID string `json:"id"`
}

type EdgeMediaToSponsorUser struct {
	Edges []SponsorUserEdge `json:"edges"`
}

type SponsorUserEdge struct {
	Node SponsorUserNode `json:"node"`
}

type SponsorUserNode struct {
	ID string `json:"id"`
}

type EdgeWebMediaToRelatedMedia struct {
	Edges []RelatedMediaEdge `json:"edges"`
}

type RelatedMediaEdge struct {
	Node RelatedMediaNode `json:"node"`
}

type RelatedMediaNode struct {
	ID string `json:"id"`
}

type CoauthorProducer struct {
	ID            string `json:"id"`
	IsVerified    bool   `json:"is_verified"`
	ProfilePicUrl string `json:"profile_pic_url"`
	Username      string `json:"username"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	apiKey = os.Getenv("RAPID_API_KEY")
}
