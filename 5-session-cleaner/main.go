//////////////////////////////////////////////////////////////////////
//
// Given is a SessionManager that stores session information in
// memory. The SessionManager itself is working, however, since we
// keep on adding new sessions to the manager our program will
// eventually run out of memory.
//
// Your task is to implement a session cleaner routine that runs
// concurrently in the background and cleans every session that
// hasn't been updated for more than 5 seconds (of course usually
// session times are much longer).
//
// Note that we expect the session to be removed anytime between 5 and
// 7 seconds after the last update. Also, note that you have to be
// very careful in order to prevent race conditions.
//

package main

import (
	"errors"
	"log"
	"sync"
    "time"
)

// SessionManager keeps track of all sessions from creation, updating
// to destroying.
type SessionManager struct {
	sessions map[string]Session
    currentTime int
    mux sync.RWMutex
}

// Session stores the session's data
type Session struct {
	Data map[string]interface{}
    LastUpdateTime int
}

func (m *SessionManager) CheckAndClean() {
    m.mux.Lock()
    defer m.mux.Unlock()
    m.currentTime += 1

    for key, session := range m.sessions {

        if m.currentTime - session.LastUpdateTime > 5 {

            //delete session
            delete (m.sessions, key)
        }
    }
}

func (m *SessionManager) CleanupSession() (string, error) {

    tickChan := time.NewTicker(time.Second * 1).C
    for {
        select {
            case <- tickChan:
                m.CheckAndClean()
        }
    }
}

// NewSessionManager creates a new sessionManager
func NewSessionManager() *SessionManager {
	m := &SessionManager{
		sessions: make(map[string]Session),
	}

    go m.CleanupSession()

	return m
}

// CreateSession creates a new session and returns the sessionID
func (m *SessionManager) CreateSession() (string, error) {
	sessionID, err := MakeSessionID()
	if err != nil {
		return "", err
	}

    m.mux.Lock()
    defer m.mux.Unlock()
	m.sessions[sessionID] = Session{
		Data: make(map[string]interface{}),
        LastUpdateTime: m.currentTime,
	}

	return sessionID, nil
}

// ErrSessionNotFound returned when sessionID not listed in
// SessionManager
var ErrSessionNotFound = errors.New("SessionID does not exists")

// GetSessionData returns data related to session if sessionID is
// found, errors otherwise
func (m *SessionManager) GetSessionData(sessionID string) (map[string]interface{}, error) {
    m.mux.RLock()
    defer m.mux.RUnlock()
	session, ok := m.sessions[sessionID]
	if !ok {
		return nil, ErrSessionNotFound
	}
	return session.Data, nil
}

// UpdateSessionData overwrites the old session data with the new one
func (m *SessionManager) UpdateSessionData(sessionID string, data map[string]interface{}) error {
	_, ok := m.sessions[sessionID]
	if !ok {
		return ErrSessionNotFound
	}

    m.mux.Lock()
    defer m.mux.Unlock()

	// Hint: you should renew expiry of the session here
	m.sessions[sessionID] = Session{
		Data: data,
        LastUpdateTime: m.currentTime,
	}

	return nil
}

func main() {
	// Create new sessionManager and new session
	m := NewSessionManager()
	sID, err := m.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created new session with ID", sID)

	// Update session data
	data := make(map[string]interface{})
	data["website"] = "longhoang.de"

	err = m.UpdateSessionData(sID, data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update session data, set website to longhoang.de")

	// Retrieve data from manager again
	updatedData, err := m.GetSessionData(sID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get session data:", updatedData)
}
