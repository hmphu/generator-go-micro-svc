<%=licenseText%>
package state

import (
	"gopkg.in/mgo.v2"
)

// NewMongoStore creates a new MongoStore. Use an empty string for databaseName
// to use the database name that was provided in the connection string.
func NewMongoStore(session *mgo.Session, databaseName string) (*MongoStore, error) {
	return &MongoStore{
		session: session,
		db:      databaseName,
	}, nil
}

// MongoStore is an implementation of Store using Mongo as the database.
type MongoStore struct {
	session *mgo.Session
	db      string
}

var _ Store = (*MongoStore)(nil)

// TODO: Add methods here

// C get a Collection from sess by using the database defined on the store.
func (s *MongoStore) C(sess *mgo.Session, collectionName string) *mgo.Collection {
	return sess.DB(s.db).C(collectionName)
}

// Initialize will be called once during startup and should ensure any required
// indexes are created.
func (s *MongoStore) Initialize() error {
	return nil
}

// Healthy return nil if nothing is wrong. If it is unable to Ping Mongo it
// will try to refresh the session and will return the err.
func (s *MongoStore) Healthy() error {
	sess := s.session.Clone()
	defer sess.Close()

	err := sess.Ping()
	if err != nil {
		return err
	}

	return nil
}

// Close calls Close on the Mongo session.
func (s *MongoStore) Close() error {
	s.session.Close()
	return nil
}
