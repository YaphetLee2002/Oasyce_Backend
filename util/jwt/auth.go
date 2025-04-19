package jwt

import (
	"Oasyce-backend/util/json"
	"Oasyce-backend/util/rand"
	"crypto/ecdsa"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type Authenticator struct {
	publicKey  *ecdsa.PublicKey
	privateKey *ecdsa.PrivateKey
}

// NewAuthenticator returns an Authenticator with given keys.
func NewAuthenticator(publicKey, privateKey string) (*Authenticator, error) {
	// Load the ES256 keypair.
	parsedPublicKey, err := jwt.ParseECPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return nil, errors.WithMessage(err, "failed to parse ecdsa public key")
	}
	parsedPrivateKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return nil, errors.WithMessage(err, "failed to parse ecdsa private key")
	}
	return &Authenticator{
		publicKey:  parsedPublicKey,
		privateKey: parsedPrivateKey,
	}, nil
}

const (
	jwtIssuer   = "Code"
	jwtAudience = "Code"
)

type ClaimsScope []string

func (c ClaimsScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.Join(c, " "))
}

func (c *ClaimsScope) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*c = strings.Split(str, " ")
	return nil
}

type Claims struct {
	jwt.StandardClaims
	LFSAction *string     `json:"lfs_action,omitempty"`
	RepoID    *int64      `json:"repo_id,omitempty"`
	ClientID  *string     `json:"client_id,omitempty"`
	Scope     ClaimsScope `json:"scope,omitempty"`
	Sudo      *bool       `json:"sudo,omitempty"` // true means sign by username, false means sign by cloud jwt
}

type IssueOption struct {
	Subject    string
	Expiration time.Duration
	LFSAction  *string
	RepoID     *int64
	ClientID   *string // if an app requests on behalf of a user, client_id is the app's username (not app id or app name)
	Scope      ClaimsScope
	Sudo       *bool
}

func (a *Authenticator) Issue(opt *IssueOption) (string, error) {
	now := time.Now()
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Audience:  jwtAudience,
			ExpiresAt: 0,
			Id:        rand.RandomID("jti"),
			IssuedAt:  now.Unix(),
			Issuer:    jwtIssuer,
			NotBefore: now.Unix(),
			Subject:   opt.Subject,
		},
		LFSAction: opt.LFSAction,
		RepoID:    opt.RepoID,
		ClientID:  opt.ClientID,
		Sudo:      opt.Sudo,
	}
	if opt.Expiration != 0 {
		claims.ExpiresAt = now.Add(opt.Expiration).Unix()
	}
	if len(opt.Scope) > 0 {
		claims.Scope = opt.Scope
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	jwtToken, err := token.SignedString(a.privateKey)
	if err != nil {
		return "", errors.WithMessage(err, "failed to sign jwt token")
	}
	return jwtToken, nil
}

func (a *Authenticator) Parse(token string) (*Claims, error) {
	claims := new(Claims)
	keyFunc := func(_ *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	}
	if t, err := jwt.ParseWithClaims(token, claims, keyFunc); err != nil || !t.Valid {
		return claims, errors.WithMessage(err, "invalid jwt token")
	}
	if claims.Issuer != jwtIssuer || claims.Audience != jwtAudience {
		return claims, errors.New("invalid jwt token issuer or audience")
	}
	return claims, nil
}
