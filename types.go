package opensea

import (
	"fmt"
)

type Address string

func ParseAddress(addr string) (Address, error) {
	if addr == "" {
		return "", fmt.Errorf("empty address")
	}
	return Address(addr), nil
}

const NullAddress Address = ""

type Collection struct {
	// todo: Support commented fields in Collection struct for /collections GET request
	BannerImageUrl              string      `json:"banner_image_url" bson:"banner_image_url"`
	ChatUrl                     string      `json:"chat_url" bson:"chat_url"`
	CreatedDate                 string      `json:"created_date" bson:"created_date"`
	DefaultToFiat               bool        `json:"default_to_fiat" bson:"default_to_fiat"`
	Description                 string      `json:"description" bson:"description"`
	DevBuyerFeeBasisPoints      string      `json:"dev_buyer_fee_basis_points" bson:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     string      `json:"dev_seller_fee_basis_points" bson:"dev_seller_fee_basis_points"`
	DiscordUrl                  string      `json:"discord_url" bson:"discord_url"`
	DisplayData                 interface{} `json:"display_data" bson:"display_data"`
	ExternalUrl                 string      `json:"external_url" bson:"external_url"`
	Featured                    bool        `json:"featured" bson:"featured"`
	FeaturedImageUrl            string      `json:"featured_image_url" bson:"featured_image_url"`
	Hidden                      bool        `json:"hidden" bson:"hidden"`
	SafelistRequestStatus       string      `json:"safelist_request_status" bson:"safelist_request_status"`
	ImageUrl                    string      `json:"image_url" bson:"image_url"`
	IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist" bson:"is_subject_to_whitelist"`
	LargeImageUrl               string      `json:"large_image_url" bson:"large_image_url"`
	MediumUsername              string      `json:"medium_username" bson:"medium_username"`
	Name                        string      `json:"name" bson:"name"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers" bson:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points" bson:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points" bson:"opensea_seller_fee_basis_points"`
	PayoutAddress               string      `json:"payout_address" bson:"payout_address"`
	RequireEmail                bool        `json:"require_email" bson:"require_email"`
	ShortDescription            string      `json:"short_description" bson:"short_description"`
	Slug                        string      `json:"slug" bson:"slug"`
	TelegramUrl                 string      `json:"telegram_url" bson:"telegram_url"`
	TwitterUsername             string      `json:"twitter_username" bson:"twitter_username"`
	InstagramUsername           string      `json:"instagram_username" bson:"instagram_username"`
	WikiUrl                     string      `json:"wiki_url" bson:"wiki_url"`
}

type User struct {
	Username string `json:"username" bson:"username"`
}

type Account struct {
	User          User    `json:"user" bson:"user"`
	ProfileImgURL string  `json:"profile_img_url" bson:"profile_img_url"`
	Address       Address `json:"address" bson:"address"`
	Config        string  `json:"config" bson:"config"`
	DiscordID     string  `json:"discord_id" bson:"discord_id"`
}

type AssetContract struct {
	Address                     Address     `json:"address" bson:"address"`
	AssetContractType           string      `json:"asset_contract_type" bson:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date" bson:"created_date"`
	Name                        string      `json:"name" bson:"name"`
	NftVersion                  string      `json:"nft_version" bson:"nft_version"`
	OpenseaVersion              interface{} `json:"opensea_version" bson:"opensea_version"`
	Owner                       int64       `json:"owner" bson:"owner"`
	SchemaName                  string      `json:"schema_name" bson:"schema_name"`
	Symbol                      string      `json:"symbol" bson:"symbol"`
	TotalSupply                 interface{} `json:"total_supply" bson:"total_supply"`
	Description                 string      `json:"description" bson:"description"`
	ExternalLink                string      `json:"external_link" bson:"external_link"`
	ImageURL                    string      `json:"image_url" bson:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat" bson:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int64       `json:"dev_buyer_fee_basis_points" bson:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int64       `json:"dev_seller_fee_basis_points" bson:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers" bson:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int64       `json:"opensea_buyer_fee_basis_points" bson:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int64       `json:"opensea_seller_fee_basis_points" bson:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int64       `json:"buyer_fee_basis_points" bson:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int64       `json:"seller_fee_basis_points" bson:"seller_fee_basis_points"`
	PayoutAddress               Address     `json:"payout_address" bson:"payout_address"`
}

type Asset struct {
	// todo: Support commented fields in Asset struct
	ID                   int64          `json:"id" bson:"id"`
	TokenID              string         `json:"token_id" bson:"token_id"`
	NumSales             int64          `json:"num_sales" bson:"num_sales"`
	BackgroundColor      string         `json:"background_color" bson:"background_color"`
	ImageURL             string         `json:"image_url" bson:"image_url"`
	ImagePreviewURL      string         `json:"image_preview_url" bson:"image_preview_url"`
	ImageThumbnailURL    string         `json:"image_thumbnail_url" bson:"image_thumbnail_url"`
	ImageOriginalURL     string         `json:"image_original_url" bson:"image_original_url"`
	AnimationURL         string         `json:"animation_url" bson:"animation_url"`
	AnimationOriginalURL string         `json:"animation_original_url" bson:"animation_original_url"`
	Name                 string         `json:"name" bson:"name"`
	Description          string         `json:"description" bson:"description"`
	ExternalLink         string         `json:"external_link" bson:"external_link"`
	AssetContract        *AssetContract `json:"asset_contract" bson:"asset_contract"`
	Permalink            string         `json:"permalink" bson:"permalink"`
	Collection           *Collection    `json:"collection" bson:"collection"`
	Decimals             int64          `json:"decimals" bson:"decimals"`
	TokenMetadata        string         `json:"token_metadata" bson:"token_metadata"`
	Owner *Account `json:"owner" bson:"owner"`
	Traits interface{} `json:"traits" bson:"traits"`
}


func (a Address) String() string {
	return string(a)
}

type TimeNano int64

func (t TimeNano) String() string {
	return fmt.Sprintf("%d", t)
}

type Number string

func (n Number) String() string {
	return string(n)
}

type Bytes []byte

func (b Bytes) String() string {
	return string(b)
} 