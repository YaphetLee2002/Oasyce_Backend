include "property.thrift"

struct Email {
    /** Email ID. */
    1: required string Id,
    /** Email address. */
    2: required string Email,
    /** Verification timestamp. Empty if not verified. In RFC3339 format (e.g. `2006-01-02T15:04:05Z07:00`). */
    3: optional string VerifiedAt,
    /** Time when the email was created. In RFC3339 format (e.g. `2006-01-02T15:04:05Z07:00`). */
    4: required string CreatedAt,
    /** Whether this is the primary email. */
    5: required bool Primary,
}

typedef string UserType
const UserType UserTypePersonal = "personal"
const UserType UserTypeApp = "app"

typedef string UserStatus
const UserStatus UserStatusActive = "active"
const UserStatus UserStatusBlocked = "blocked"

typedef string ThemeMode
const ThemeMode ThemeModeLight = "light"
const ThemeMode ThemeModeDark = "dark"
const ThemeMode ThemeModeAuto = "auto"

struct User {
    /** User ID. */
    1: required string Id,
    /** Username (unique, typically the email prefix). */
    2: required string Username,
    /** Display name (to be shown). */
    3: optional property.I18nString DisplayName,
    /** Primary email address. */
    4: required string Email,
    /** Tenant or main account ID the user belongs to. */
    5: required string TenantId,
    /** Time when the user was created. In RFC3339 format (e.g. `2006-01-02T15:04:05Z07:00`). */
    6: required string CreatedAt,
    /** List of all user emails, including verified and unverified. */
    7: optional list<Email> Emails,
    /** Account type. Enum values: `personal` / `app`. */
    8: required UserType Type,
    /** Account status. Enum values: `active` / `blocked`. */
    9: required UserStatus Status,
    /** Avatar URL. */
    10: optional string AvatarURL,
    /** Whether the user can pin repositories. */
    11: bool CanPinRepository,
    /** User's feature settings. */
    12: optional map<string, bool> Features,
    /** User biography or brief introduction. */
    13: optional string Bio,
    /** User's timezone location. (e.g. `Asia/Shanghai`). */
    14: optional string Location,
    /** User's preferred theme mode. Enum values: `light` / `dark` / `auto`. */
    15: optional ThemeMode ThemeMode,
    /** User's preferred language. Enum values: `zh-CN` / `en-US`. */
    16: optional property.Locale Language,
    /** Custom properties. */
    17: optional list<property.Property> Properties,
}

struct UserBrief {
    /** User ID. */
    1: required string Id,
    /** Username (unique, typically the email prefix). */
    2: required string Username,
    /** Display name (to be shown). */
    3: optional property.I18nString DisplayName,
    /** Primary email address. */
    4: required string Email,
    /** Tenant or main account ID the user belongs to. */
    5: required string TenantId,
    /** Account type. Enum values: `personal` / `app`. */
    6: required UserType Type,
    /** Account status. Enum values: `active` / `blocked`. */
    7: required UserStatus Status,
    /** Avatar URL. */
    8: optional string AvatarURL,
}

struct UserGroup {
    /** User group ID. */
    1: required string Id,
}

struct UserSelector {
    /** Account type. Enum values: `personal` / `app`. */
    1: optional UserType UserType,
    2: optional string UserId,
    /** Required permissions. The account must have all of these permissions. */
    3: optional list<string> RequiredPermissions,
    4: optional string UserGroupId,
    5: optional User User,
    6: optional UserGroup UserGroup,
}

struct GetUserRequest {
    /** Username of the user to fetch. It will take the caller as the user requested if both Id and Username are empty. */
    3: optional string Username,
    /** ID of the user to fetch. It will take the caller as the user requested if both Id and Username are empty. */
    4: optional string Id,
    /** Deprecated. Whether to return all user emails. */
    1: optional bool WithEmails,
    /** Deprecated. Whether to return all SSH public keys. */
    2: optional FieldSelector Selector,
}

struct FieldSelector {
    /** Whether to return the user's feature settings. */
    1: optional bool Features,
    /** Whether to return the user's emails. */
    2: optional bool Emails,
    /** Whether to fetch user-defined properties. */
    3: optional bool Properties,
    /** Valid only when `Properties` is true. List of property name identifiers to fetch. Fetches all if unspecified. */
    4: optional list<string> PropertyNames,
}

struct GetUserResponse {
    /** User information. */
    1: required User User,
}