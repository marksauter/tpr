package main

import (
	"bytes"
	"crypto/rand"
	"errors"
	"github.com/jackc/tpr/backend/data"
	"golang.org/x/crypto/scrypt"
	"io"
	"time"
)

var notFound = errors.New("not found")

type repository interface {
	CreateSession(id []byte, userID int32) (err error)
	DeleteSession(id []byte) (err error)

	CreatePasswordReset(*data.PasswordReset) error
	GetPasswordReset(token string) (*data.PasswordReset, error)
	UpdatePasswordReset(string, *data.PasswordReset) error

	GetFeedsUncheckedSince(since time.Time) (feeds []data.Feed, err error)
	UpdateFeedWithFetchSuccess(feedID int32, update *parsedFeed, etag data.String, fetchTime time.Time) error
	UpdateFeedWithFetchUnchanged(feedID int32, fetchTime time.Time) error
	UpdateFeedWithFetchFailure(feedID int32, failure string, fetchTime time.Time) (err error)

	CopyUnreadItemsAsJSONByUserID(w io.Writer, userID int32) error
	CopySubscriptionsForUserAsJSON(w io.Writer, userID int32) error

	MarkItemRead(userID, itemID int32) error

	CreateSubscription(userID int32, feedURL string) (err error)
	GetSubscriptions(userID int32) ([]Subscription, error)
	DeleteSubscription(userID, feedID int32) (err error)
}

func SetPassword(u *data.User, password string) error {
	salt := make([]byte, 8)
	_, err := rand.Read(salt)
	if err != nil {
		return err
	}

	digest, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
	if err != nil {
		return err
	}

	u.PasswordDigest = data.NewBytes(digest)
	u.PasswordSalt = data.NewBytes(salt)

	return nil
}

func IsPassword(u *data.User, password string) bool {
	digest, err := scrypt.Key([]byte(password), u.PasswordSalt.Value, 16384, 8, 1, 32)
	if err != nil {
		return false
	}

	return bytes.Equal(digest, u.PasswordDigest.Value)
}

type Subscription struct {
	FeedID              data.Int32
	Name                data.String
	URL                 data.String
	LastFetchTime       data.Time
	LastFailure         data.String
	LastFailureTime     data.Time
	FailureCount        data.Int32
	ItemCount           data.Int64
	LastPublicationTime data.Time
}

type staleFeed struct {
	id   int32
	url  string
	etag string
}
