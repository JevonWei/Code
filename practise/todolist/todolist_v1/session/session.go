package session

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Session struct {
	contents    map[string]interface{}
	expiredTime time.Time
	lock        sync.RWMutex
}

func NewSession(expiredTime time.Time) *Session {
	return &Session{
		contents:    make(map[string]interface{}),
		expiredTime: expiredTime,
	}
}

func (s *Session) Set(key string, value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.contents[key] = value

}

func (s *Session) Get(key string) (interface{}, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	value, ok := s.contents[key]
	return value, ok
}

func (s *Session) Delete(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.contents, key)
}

type Manager struct {
	sessions   map[string]*Session
	cookieName string // sid
	timeout    int
	lock       sync.Mutex
}

func NewManager(cookieName string, timeout int) *Manager {
	// timeout：session过期时间
	manager := Manager{
		sessions:   make(map[string]*Session),
		cookieName: cookieName,
		timeout:    timeout,
	}
	go manager.SessionGC()
	return &manager
}

func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) *Session {
	var session *Session
	var sid string
	cookie, err := r.Cookie(m.cookieName)
	m.lock.Lock()
	defer m.lock.Unlock()

	// 检查并生成sid
	if err != nil || cookie.Value == "" {
		sid = uuid.NewV4().String()
		session = NewSession(time.Now().Add(time.Duration(m.timeout) * time.Second))
		m.sessions[sid] = session
	} else if _, ok := m.sessions[cookie.Value]; !ok {
		sid = uuid.NewV4().String()
		session = NewSession(time.Now().Add(time.Duration(m.timeout) * time.Second))
		m.sessions[sid] = session
	} else {
		return m.sessions[cookie.Value]
	}

	cookie = &http.Cookie{
		Name:     m.cookieName,
		Value:    sid,
		Path:     "/",
		MaxAge:   m.timeout,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	return session

}

func (m *Manager) SessionDestory(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)

	if err == nil && cookie.Value != "" {
		m.lock.Lock()
		delete(m.sessions, cookie.Value)
		m.lock.Unlock()
	}

	cookie = &http.Cookie{
		Name:     m.cookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}

func (m *Manager) SessionGC() {
	for now := range time.Tick(2 * time.Second) {
		m.lock.Lock()
		deleted := []string{}
		for sid, session := range m.sessions {
			if session.expiredTime.Before(now) {
				deleted = append(deleted, sid)
			}
		}
		for _, sid := range deleted {
			fmt.Println(sid)
			delete(m.sessions, sid)
		}
		m.lock.Unlock()
	}
}

var DefaultSessionManager *Manager

func init() {
	DefaultSessionManager = NewManager("sid", 120)
}
