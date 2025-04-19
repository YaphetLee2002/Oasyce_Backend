struct Property {
    /** Identifier of the property. */
    1: required string Name,
    /** Display name of the property. Supports internationalization. */
    5: optional I18nString DisplayName,
    /** Data type of the property value. Enum values: `text` / `markdown` / `user` / `list`. */
    2: required PropertyType Type,
    /** Type of elements in the list. Only present when `Type` is `list`. */
    3: optional PropertyType ItemType,
    /** Value of the property. See `PropertyValue`. */
    4: required PropertyValue Value,
    /**
     * Raw value for certain properties.
     * - `bytetree`: Represents the service tree node ID.
     * - `department`: Represents the department ID.
     * - `ecc`: Represents the ECC tag.
     * - `risk_level`: Represents the repository security level.
     */
    6: optional string RawValue,
}

typedef string PropertyType
const PropertyType PropertyTypeText = "text"
const PropertyType PropertyTypeMarkdown = "markdown"
const PropertyType PropertyTypeUser = "user"
const PropertyType PropertyTypeList = "list"

union PropertyValue {
    /** Plain text. */
    1: I18nString Text,
    /** Markdown supported text. */
    2: I18nString Markdown,
    3: User User,
    /** A list of property values. */
    4: list<PropertyValue> List,
}

struct User {
    /** User ID. */
    1: required string Id,
    /** Username of the user. */
    2: required string Username,
    /** Full name of the user. */
    3: required string Name,
    /** Display name of the user. Supports internationalization. */
    5: optional I18nString DisplayName,
    /** URL of the user's avatar. */
    4: optional string AvatarURL,
}

struct I18nString {
    /** Default or original value of the string. */
    1: required string Content,
    /**
     * Map for localized values.
     * - Key: Locale identifier (e.g. `zh-CN`, `en-US`).
     * - Value: Localized string for the specified locale.
     * If a locale is not specified, the `Content` value should be used.
     */
    2: optional map<Locale, string> I18n,
}

typedef string Locale
const Locale LocaleZhCN = "zh-CN"
const Locale LocaleEnUS = "en-US"
