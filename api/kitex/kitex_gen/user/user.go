// Code generated by thriftgo (0.4.1). DO NOT EDIT.

package user

import (
	"Oasyce-backend/api/kitex/kitex_gen/property"
	"fmt"
)

const (
	UserTypePersonal = "personal"

	UserTypeApp = "app"

	UserStatusActive = "active"

	UserStatusBlocked = "blocked"

	ThemeModeLight = "light"

	ThemeModeDark = "dark"

	ThemeModeAuto = "auto"
)

type UserType = string

type UserStatus = string

type ThemeMode = string

type Email struct {
	Id         string  `thrift:"Id,1,required" frugal:"1,required,string" json:"Id"`
	Email      string  `thrift:"Email,2,required" frugal:"2,required,string" json:"Email"`
	VerifiedAt *string `thrift:"VerifiedAt,3,optional" frugal:"3,optional,string" json:"VerifiedAt,omitempty"`
	CreatedAt  string  `thrift:"CreatedAt,4,required" frugal:"4,required,string" json:"CreatedAt"`
	Primary    bool    `thrift:"Primary,5,required" frugal:"5,required,bool" json:"Primary"`
}

func NewEmail() *Email {
	return &Email{}
}

func (p *Email) InitDefault() {
}

func (p *Email) GetId() (v string) {
	return p.Id
}

func (p *Email) GetEmail() (v string) {
	return p.Email
}

var Email_VerifiedAt_DEFAULT string

func (p *Email) GetVerifiedAt() (v string) {
	if !p.IsSetVerifiedAt() {
		return Email_VerifiedAt_DEFAULT
	}
	return *p.VerifiedAt
}

func (p *Email) GetCreatedAt() (v string) {
	return p.CreatedAt
}

func (p *Email) GetPrimary() (v bool) {
	return p.Primary
}
func (p *Email) SetId(val string) {
	p.Id = val
}
func (p *Email) SetEmail(val string) {
	p.Email = val
}
func (p *Email) SetVerifiedAt(val *string) {
	p.VerifiedAt = val
}
func (p *Email) SetCreatedAt(val string) {
	p.CreatedAt = val
}
func (p *Email) SetPrimary(val bool) {
	p.Primary = val
}

func (p *Email) IsSetVerifiedAt() bool {
	return p.VerifiedAt != nil
}

func (p *Email) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Email(%+v)", *p)
}

var fieldIDToName_Email = map[int16]string{
	1: "Id",
	2: "Email",
	3: "VerifiedAt",
	4: "CreatedAt",
	5: "Primary",
}

type User struct {
	Id               string               `thrift:"Id,1,required" frugal:"1,required,string" json:"Id"`
	Username         string               `thrift:"Username,2,required" frugal:"2,required,string" json:"Username"`
	DisplayName      *property.I18nString `thrift:"DisplayName,3,optional" frugal:"3,optional,property.I18nString" json:"DisplayName,omitempty"`
	Email            string               `thrift:"Email,4,required" frugal:"4,required,string" json:"Email"`
	TenantId         string               `thrift:"TenantId,5,required" frugal:"5,required,string" json:"TenantId"`
	CreatedAt        string               `thrift:"CreatedAt,6,required" frugal:"6,required,string" json:"CreatedAt"`
	Emails           []*Email             `thrift:"Emails,7,optional" frugal:"7,optional,list<Email>" json:"Emails,omitempty"`
	Type             UserType             `thrift:"Type,8,required" frugal:"8,required,string" json:"Type"`
	Status           UserStatus           `thrift:"Status,9,required" frugal:"9,required,string" json:"Status"`
	AvatarURL        *string              `thrift:"AvatarURL,10,optional" frugal:"10,optional,string" json:"AvatarURL,omitempty"`
	CanPinRepository bool                 `thrift:"CanPinRepository,11" frugal:"11,default,bool" json:"CanPinRepository"`
	Features         map[string]bool      `thrift:"Features,12,optional" frugal:"12,optional,map<string:bool>" json:"Features,omitempty"`
	Bio              *string              `thrift:"Bio,13,optional" frugal:"13,optional,string" json:"Bio,omitempty"`
	Location         *string              `thrift:"Location,14,optional" frugal:"14,optional,string" json:"Location,omitempty"`
	ThemeMode        *ThemeMode           `thrift:"ThemeMode,15,optional" frugal:"15,optional,string" json:"ThemeMode,omitempty"`
	Language         *property.Locale     `thrift:"Language,16,optional" frugal:"16,optional,string" json:"Language,omitempty"`
	Properties       []*property.Property `thrift:"Properties,17,optional" frugal:"17,optional,list<property.Property>" json:"Properties,omitempty"`
}

func NewUser() *User {
	return &User{}
}

func (p *User) InitDefault() {
}

func (p *User) GetId() (v string) {
	return p.Id
}

func (p *User) GetUsername() (v string) {
	return p.Username
}

var User_DisplayName_DEFAULT *property.I18nString

func (p *User) GetDisplayName() (v *property.I18nString) {
	if !p.IsSetDisplayName() {
		return User_DisplayName_DEFAULT
	}
	return p.DisplayName
}

func (p *User) GetEmail() (v string) {
	return p.Email
}

func (p *User) GetTenantId() (v string) {
	return p.TenantId
}

func (p *User) GetCreatedAt() (v string) {
	return p.CreatedAt
}

var User_Emails_DEFAULT []*Email

func (p *User) GetEmails() (v []*Email) {
	if !p.IsSetEmails() {
		return User_Emails_DEFAULT
	}
	return p.Emails
}

func (p *User) GetType() (v UserType) {
	return p.Type
}

func (p *User) GetStatus() (v UserStatus) {
	return p.Status
}

var User_AvatarURL_DEFAULT string

func (p *User) GetAvatarURL() (v string) {
	if !p.IsSetAvatarURL() {
		return User_AvatarURL_DEFAULT
	}
	return *p.AvatarURL
}

func (p *User) GetCanPinRepository() (v bool) {
	return p.CanPinRepository
}

var User_Features_DEFAULT map[string]bool

func (p *User) GetFeatures() (v map[string]bool) {
	if !p.IsSetFeatures() {
		return User_Features_DEFAULT
	}
	return p.Features
}

var User_Bio_DEFAULT string

func (p *User) GetBio() (v string) {
	if !p.IsSetBio() {
		return User_Bio_DEFAULT
	}
	return *p.Bio
}

var User_Location_DEFAULT string

func (p *User) GetLocation() (v string) {
	if !p.IsSetLocation() {
		return User_Location_DEFAULT
	}
	return *p.Location
}

var User_ThemeMode_DEFAULT ThemeMode

func (p *User) GetThemeMode() (v ThemeMode) {
	if !p.IsSetThemeMode() {
		return User_ThemeMode_DEFAULT
	}
	return *p.ThemeMode
}

var User_Language_DEFAULT property.Locale

func (p *User) GetLanguage() (v property.Locale) {
	if !p.IsSetLanguage() {
		return User_Language_DEFAULT
	}
	return *p.Language
}

var User_Properties_DEFAULT []*property.Property

func (p *User) GetProperties() (v []*property.Property) {
	if !p.IsSetProperties() {
		return User_Properties_DEFAULT
	}
	return p.Properties
}
func (p *User) SetId(val string) {
	p.Id = val
}
func (p *User) SetUsername(val string) {
	p.Username = val
}
func (p *User) SetDisplayName(val *property.I18nString) {
	p.DisplayName = val
}
func (p *User) SetEmail(val string) {
	p.Email = val
}
func (p *User) SetTenantId(val string) {
	p.TenantId = val
}
func (p *User) SetCreatedAt(val string) {
	p.CreatedAt = val
}
func (p *User) SetEmails(val []*Email) {
	p.Emails = val
}
func (p *User) SetType(val UserType) {
	p.Type = val
}
func (p *User) SetStatus(val UserStatus) {
	p.Status = val
}
func (p *User) SetAvatarURL(val *string) {
	p.AvatarURL = val
}
func (p *User) SetCanPinRepository(val bool) {
	p.CanPinRepository = val
}
func (p *User) SetFeatures(val map[string]bool) {
	p.Features = val
}
func (p *User) SetBio(val *string) {
	p.Bio = val
}
func (p *User) SetLocation(val *string) {
	p.Location = val
}
func (p *User) SetThemeMode(val *ThemeMode) {
	p.ThemeMode = val
}
func (p *User) SetLanguage(val *property.Locale) {
	p.Language = val
}
func (p *User) SetProperties(val []*property.Property) {
	p.Properties = val
}

func (p *User) IsSetDisplayName() bool {
	return p.DisplayName != nil
}

func (p *User) IsSetEmails() bool {
	return p.Emails != nil
}

func (p *User) IsSetAvatarURL() bool {
	return p.AvatarURL != nil
}

func (p *User) IsSetFeatures() bool {
	return p.Features != nil
}

func (p *User) IsSetBio() bool {
	return p.Bio != nil
}

func (p *User) IsSetLocation() bool {
	return p.Location != nil
}

func (p *User) IsSetThemeMode() bool {
	return p.ThemeMode != nil
}

func (p *User) IsSetLanguage() bool {
	return p.Language != nil
}

func (p *User) IsSetProperties() bool {
	return p.Properties != nil
}

func (p *User) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("User(%+v)", *p)
}

var fieldIDToName_User = map[int16]string{
	1:  "Id",
	2:  "Username",
	3:  "DisplayName",
	4:  "Email",
	5:  "TenantId",
	6:  "CreatedAt",
	7:  "Emails",
	8:  "Type",
	9:  "Status",
	10: "AvatarURL",
	11: "CanPinRepository",
	12: "Features",
	13: "Bio",
	14: "Location",
	15: "ThemeMode",
	16: "Language",
	17: "Properties",
}

type UserBrief struct {
	Id          string               `thrift:"Id,1,required" frugal:"1,required,string" json:"Id"`
	Username    string               `thrift:"Username,2,required" frugal:"2,required,string" json:"Username"`
	DisplayName *property.I18nString `thrift:"DisplayName,3,optional" frugal:"3,optional,property.I18nString" json:"DisplayName,omitempty"`
	Email       string               `thrift:"Email,4,required" frugal:"4,required,string" json:"Email"`
	TenantId    string               `thrift:"TenantId,5,required" frugal:"5,required,string" json:"TenantId"`
	Type        UserType             `thrift:"Type,6,required" frugal:"6,required,string" json:"Type"`
	Status      UserStatus           `thrift:"Status,7,required" frugal:"7,required,string" json:"Status"`
	AvatarURL   *string              `thrift:"AvatarURL,8,optional" frugal:"8,optional,string" json:"AvatarURL,omitempty"`
}

func NewUserBrief() *UserBrief {
	return &UserBrief{}
}

func (p *UserBrief) InitDefault() {
}

func (p *UserBrief) GetId() (v string) {
	return p.Id
}

func (p *UserBrief) GetUsername() (v string) {
	return p.Username
}

var UserBrief_DisplayName_DEFAULT *property.I18nString

func (p *UserBrief) GetDisplayName() (v *property.I18nString) {
	if !p.IsSetDisplayName() {
		return UserBrief_DisplayName_DEFAULT
	}
	return p.DisplayName
}

func (p *UserBrief) GetEmail() (v string) {
	return p.Email
}

func (p *UserBrief) GetTenantId() (v string) {
	return p.TenantId
}

func (p *UserBrief) GetType() (v UserType) {
	return p.Type
}

func (p *UserBrief) GetStatus() (v UserStatus) {
	return p.Status
}

var UserBrief_AvatarURL_DEFAULT string

func (p *UserBrief) GetAvatarURL() (v string) {
	if !p.IsSetAvatarURL() {
		return UserBrief_AvatarURL_DEFAULT
	}
	return *p.AvatarURL
}
func (p *UserBrief) SetId(val string) {
	p.Id = val
}
func (p *UserBrief) SetUsername(val string) {
	p.Username = val
}
func (p *UserBrief) SetDisplayName(val *property.I18nString) {
	p.DisplayName = val
}
func (p *UserBrief) SetEmail(val string) {
	p.Email = val
}
func (p *UserBrief) SetTenantId(val string) {
	p.TenantId = val
}
func (p *UserBrief) SetType(val UserType) {
	p.Type = val
}
func (p *UserBrief) SetStatus(val UserStatus) {
	p.Status = val
}
func (p *UserBrief) SetAvatarURL(val *string) {
	p.AvatarURL = val
}

func (p *UserBrief) IsSetDisplayName() bool {
	return p.DisplayName != nil
}

func (p *UserBrief) IsSetAvatarURL() bool {
	return p.AvatarURL != nil
}

func (p *UserBrief) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserBrief(%+v)", *p)
}

var fieldIDToName_UserBrief = map[int16]string{
	1: "Id",
	2: "Username",
	3: "DisplayName",
	4: "Email",
	5: "TenantId",
	6: "Type",
	7: "Status",
	8: "AvatarURL",
}

type UserGroup struct {
	Id string `thrift:"Id,1,required" frugal:"1,required,string" json:"Id"`
}

func NewUserGroup() *UserGroup {
	return &UserGroup{}
}

func (p *UserGroup) InitDefault() {
}

func (p *UserGroup) GetId() (v string) {
	return p.Id
}
func (p *UserGroup) SetId(val string) {
	p.Id = val
}

func (p *UserGroup) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserGroup(%+v)", *p)
}

var fieldIDToName_UserGroup = map[int16]string{
	1: "Id",
}

type UserSelector struct {
	UserType            *UserType  `thrift:"UserType,1,optional" frugal:"1,optional,string" json:"UserType,omitempty"`
	UserId              *string    `thrift:"UserId,2,optional" frugal:"2,optional,string" json:"UserId,omitempty"`
	RequiredPermissions []string   `thrift:"RequiredPermissions,3,optional" frugal:"3,optional,list<string>" json:"RequiredPermissions,omitempty"`
	UserGroupId         *string    `thrift:"UserGroupId,4,optional" frugal:"4,optional,string" json:"UserGroupId,omitempty"`
	User                *User      `thrift:"User,5,optional" frugal:"5,optional,User" json:"User,omitempty"`
	UserGroup           *UserGroup `thrift:"UserGroup,6,optional" frugal:"6,optional,UserGroup" json:"UserGroup,omitempty"`
}

func NewUserSelector() *UserSelector {
	return &UserSelector{}
}

func (p *UserSelector) InitDefault() {
}

var UserSelector_UserType_DEFAULT UserType

func (p *UserSelector) GetUserType() (v UserType) {
	if !p.IsSetUserType() {
		return UserSelector_UserType_DEFAULT
	}
	return *p.UserType
}

var UserSelector_UserId_DEFAULT string

func (p *UserSelector) GetUserId() (v string) {
	if !p.IsSetUserId() {
		return UserSelector_UserId_DEFAULT
	}
	return *p.UserId
}

var UserSelector_RequiredPermissions_DEFAULT []string

func (p *UserSelector) GetRequiredPermissions() (v []string) {
	if !p.IsSetRequiredPermissions() {
		return UserSelector_RequiredPermissions_DEFAULT
	}
	return p.RequiredPermissions
}

var UserSelector_UserGroupId_DEFAULT string

func (p *UserSelector) GetUserGroupId() (v string) {
	if !p.IsSetUserGroupId() {
		return UserSelector_UserGroupId_DEFAULT
	}
	return *p.UserGroupId
}

var UserSelector_User_DEFAULT *User

func (p *UserSelector) GetUser() (v *User) {
	if !p.IsSetUser() {
		return UserSelector_User_DEFAULT
	}
	return p.User
}

var UserSelector_UserGroup_DEFAULT *UserGroup

func (p *UserSelector) GetUserGroup() (v *UserGroup) {
	if !p.IsSetUserGroup() {
		return UserSelector_UserGroup_DEFAULT
	}
	return p.UserGroup
}
func (p *UserSelector) SetUserType(val *UserType) {
	p.UserType = val
}
func (p *UserSelector) SetUserId(val *string) {
	p.UserId = val
}
func (p *UserSelector) SetRequiredPermissions(val []string) {
	p.RequiredPermissions = val
}
func (p *UserSelector) SetUserGroupId(val *string) {
	p.UserGroupId = val
}
func (p *UserSelector) SetUser(val *User) {
	p.User = val
}
func (p *UserSelector) SetUserGroup(val *UserGroup) {
	p.UserGroup = val
}

func (p *UserSelector) IsSetUserType() bool {
	return p.UserType != nil
}

func (p *UserSelector) IsSetUserId() bool {
	return p.UserId != nil
}

func (p *UserSelector) IsSetRequiredPermissions() bool {
	return p.RequiredPermissions != nil
}

func (p *UserSelector) IsSetUserGroupId() bool {
	return p.UserGroupId != nil
}

func (p *UserSelector) IsSetUser() bool {
	return p.User != nil
}

func (p *UserSelector) IsSetUserGroup() bool {
	return p.UserGroup != nil
}

func (p *UserSelector) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserSelector(%+v)", *p)
}

var fieldIDToName_UserSelector = map[int16]string{
	1: "UserType",
	2: "UserId",
	3: "RequiredPermissions",
	4: "UserGroupId",
	5: "User",
	6: "UserGroup",
}

type GetUserRequest struct {
	Username   *string        `thrift:"Username,3,optional" frugal:"3,optional,string" json:"Username,omitempty"`
	Id         *string        `thrift:"Id,4,optional" frugal:"4,optional,string" json:"Id,omitempty"`
	WithEmails *bool          `thrift:"WithEmails,1,optional" frugal:"1,optional,bool" json:"WithEmails,omitempty"`
	Selector   *FieldSelector `thrift:"Selector,2,optional" frugal:"2,optional,FieldSelector" json:"Selector,omitempty"`
}

func NewGetUserRequest() *GetUserRequest {
	return &GetUserRequest{}
}

func (p *GetUserRequest) InitDefault() {
}

var GetUserRequest_Username_DEFAULT string

func (p *GetUserRequest) GetUsername() (v string) {
	if !p.IsSetUsername() {
		return GetUserRequest_Username_DEFAULT
	}
	return *p.Username
}

var GetUserRequest_Id_DEFAULT string

func (p *GetUserRequest) GetId() (v string) {
	if !p.IsSetId() {
		return GetUserRequest_Id_DEFAULT
	}
	return *p.Id
}

var GetUserRequest_WithEmails_DEFAULT bool

func (p *GetUserRequest) GetWithEmails() (v bool) {
	if !p.IsSetWithEmails() {
		return GetUserRequest_WithEmails_DEFAULT
	}
	return *p.WithEmails
}

var GetUserRequest_Selector_DEFAULT *FieldSelector

func (p *GetUserRequest) GetSelector() (v *FieldSelector) {
	if !p.IsSetSelector() {
		return GetUserRequest_Selector_DEFAULT
	}
	return p.Selector
}
func (p *GetUserRequest) SetUsername(val *string) {
	p.Username = val
}
func (p *GetUserRequest) SetId(val *string) {
	p.Id = val
}
func (p *GetUserRequest) SetWithEmails(val *bool) {
	p.WithEmails = val
}
func (p *GetUserRequest) SetSelector(val *FieldSelector) {
	p.Selector = val
}

func (p *GetUserRequest) IsSetUsername() bool {
	return p.Username != nil
}

func (p *GetUserRequest) IsSetId() bool {
	return p.Id != nil
}

func (p *GetUserRequest) IsSetWithEmails() bool {
	return p.WithEmails != nil
}

func (p *GetUserRequest) IsSetSelector() bool {
	return p.Selector != nil
}

func (p *GetUserRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetUserRequest(%+v)", *p)
}

var fieldIDToName_GetUserRequest = map[int16]string{
	3: "Username",
	4: "Id",
	1: "WithEmails",
	2: "Selector",
}

type FieldSelector struct {
	Features      *bool    `thrift:"Features,1,optional" frugal:"1,optional,bool" json:"Features,omitempty"`
	Emails        *bool    `thrift:"Emails,2,optional" frugal:"2,optional,bool" json:"Emails,omitempty"`
	Properties    *bool    `thrift:"Properties,3,optional" frugal:"3,optional,bool" json:"Properties,omitempty"`
	PropertyNames []string `thrift:"PropertyNames,4,optional" frugal:"4,optional,list<string>" json:"PropertyNames,omitempty"`
}

func NewFieldSelector() *FieldSelector {
	return &FieldSelector{}
}

func (p *FieldSelector) InitDefault() {
}

var FieldSelector_Features_DEFAULT bool

func (p *FieldSelector) GetFeatures() (v bool) {
	if !p.IsSetFeatures() {
		return FieldSelector_Features_DEFAULT
	}
	return *p.Features
}

var FieldSelector_Emails_DEFAULT bool

func (p *FieldSelector) GetEmails() (v bool) {
	if !p.IsSetEmails() {
		return FieldSelector_Emails_DEFAULT
	}
	return *p.Emails
}

var FieldSelector_Properties_DEFAULT bool

func (p *FieldSelector) GetProperties() (v bool) {
	if !p.IsSetProperties() {
		return FieldSelector_Properties_DEFAULT
	}
	return *p.Properties
}

var FieldSelector_PropertyNames_DEFAULT []string

func (p *FieldSelector) GetPropertyNames() (v []string) {
	if !p.IsSetPropertyNames() {
		return FieldSelector_PropertyNames_DEFAULT
	}
	return p.PropertyNames
}
func (p *FieldSelector) SetFeatures(val *bool) {
	p.Features = val
}
func (p *FieldSelector) SetEmails(val *bool) {
	p.Emails = val
}
func (p *FieldSelector) SetProperties(val *bool) {
	p.Properties = val
}
func (p *FieldSelector) SetPropertyNames(val []string) {
	p.PropertyNames = val
}

func (p *FieldSelector) IsSetFeatures() bool {
	return p.Features != nil
}

func (p *FieldSelector) IsSetEmails() bool {
	return p.Emails != nil
}

func (p *FieldSelector) IsSetProperties() bool {
	return p.Properties != nil
}

func (p *FieldSelector) IsSetPropertyNames() bool {
	return p.PropertyNames != nil
}

func (p *FieldSelector) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FieldSelector(%+v)", *p)
}

var fieldIDToName_FieldSelector = map[int16]string{
	1: "Features",
	2: "Emails",
	3: "Properties",
	4: "PropertyNames",
}

type GetUserResponse struct {
	User *User `thrift:"User,1,required" frugal:"1,required,User" json:"User"`
}

func NewGetUserResponse() *GetUserResponse {
	return &GetUserResponse{}
}

func (p *GetUserResponse) InitDefault() {
}

var GetUserResponse_User_DEFAULT *User

func (p *GetUserResponse) GetUser() (v *User) {
	if !p.IsSetUser() {
		return GetUserResponse_User_DEFAULT
	}
	return p.User
}
func (p *GetUserResponse) SetUser(val *User) {
	p.User = val
}

func (p *GetUserResponse) IsSetUser() bool {
	return p.User != nil
}

func (p *GetUserResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetUserResponse(%+v)", *p)
}

var fieldIDToName_GetUserResponse = map[int16]string{
	1: "User",
}
