package oidc

import (
	"errors"
	"sort"

	casbinlog "github.com/casbin/casbin/v2/log"
	"github.com/casbin/casbin/v2/rbac"
	"github.com/labstack/gommon/log"
)

type RoleManager struct {
	allRoles          map[string]*SessionRole
	maxHierarchyLevel int
}

// AddDomainMatchingFunc implements rbac.RoleManager.
func (RoleManager) AddDomainMatchingFunc(name string, fn rbac.MatchingFunc) {
	panic("unimplemented")
}

// AddLink implements rbac.RoleManager.
func (rm *RoleManager) AddLink(name1 string, name2 string, domain ...string) error {
	role1 := rm.createRole(name1)
	role2 := rm.createRole(name2)

	session := Session{role2}
	role1.addSession(session)
	return nil
}

// AddMatchingFunc implements rbac.RoleManager.
func (RoleManager) AddMatchingFunc(name string, fn rbac.MatchingFunc) {
	panic("unimplemented")
}

// BuildRelationship implements rbac.RoleManager.
func (RoleManager) BuildRelationship(name1 string, name2 string, domain ...string) error {
	panic("unimplemented")
}

// GetAllDomains implements rbac.RoleManager.
func (RoleManager) GetAllDomains() ([]string, error) {
	panic("unimplemented")
}

// GetDomains implements rbac.RoleManager.
func (RoleManager) GetDomains(name string) ([]string, error) {
	panic("unimplemented")
}

// Match implements rbac.RoleManager.
func (RoleManager) Match(str string, pattern string) bool {
	panic("unimplemented")
}

// NewRoleManager is the constructor for creating an instance of the
// SessionRoleManager implementation.
func NewRoleManager(maxHierarchyLevel int) rbac.RoleManager {
	rm := &RoleManager{}
	rm.allRoles = make(map[string]*SessionRole)
	rm.maxHierarchyLevel = maxHierarchyLevel
	return rm
}

func (rm *RoleManager) hasRole(name string) bool {
	_, ok := rm.allRoles[name]
	return ok
}

func (rm *RoleManager) createRole(name string) *SessionRole {
	if !rm.hasRole(name) {
		rm.allRoles[name] = newSessionRole(name)
	}
	return rm.allRoles[name]
}

// Clear clears all stored data and resets the role manager to the initial state.
func (rm *RoleManager) Clear() error {
	rm.allRoles = make(map[string]*SessionRole)
	return nil
}

// DeleteLink deletes the inheritance link between role: name1 and role: name2.
// aka role: name1 does not inherit role: name2 any more.
// unused is not used.
func (rm *RoleManager) DeleteLink(name1 string, name2 string, domain ...string) error {
	if !rm.hasRole(name1) || !rm.hasRole(name2) {
		return errors.New("error: name1 or name2 does not exist")
	}

	role1 := rm.createRole(name1)
	role2 := rm.createRole(name2)

	role1.deleteSessions(role2.name)
	return nil
}

// HasLink determines whether role: name1 inherits role: name2.
// requestTime is the querying time for the role inheritance link.
func (rm *RoleManager) HasLink(name1 string, name2 string, domain ...string) (bool, error) {

	if name1 == name2 {
		return true, nil
	}

	if !rm.hasRole(name1) || !rm.hasRole(name2) {
		return false, nil
	}

	role1 := rm.createRole(name1)
	return role1.hasValidSession(name2, rm.maxHierarchyLevel), nil
}

// GetRoles gets the roles that a subject inherits.
// currentTime is the querying time for the role inheritance link.
func (rm *RoleManager) GetRoles(name string, domain ...string) ([]string, error) {

	if !rm.hasRole(name) {
		return nil, errors.New("error: name does not exist")
	}

	sessionRoles := rm.createRole(name).getSessionRoles()
	return sessionRoles, nil
}

// GetUsers gets the users that inherits a subject.
// currentTime is the querying time for the role inheritance link.
func (rm *RoleManager) GetUsers(name string, domain ...string) ([]string, error) {

	users := []string{}
	for _, role := range rm.allRoles {
		if role.hasDirectRole(name) {
			users = append(users, role.name)
		}
	}
	sort.Strings(users)
	return users, nil
}

// PrintRoles prints all the roles to log.
func (rm *RoleManager) PrintRoles() error {
	for _, role := range rm.allRoles {
		log.Infof(role.toString())
	}
	return nil
}

// SetLogger sets role manager's logger.
func (rm *RoleManager) SetLogger(logger casbinlog.Logger) {

}

// SessionRole is a modified version of the default role.
// A SessionRole not only has a name, but also a list of sessions.
type SessionRole struct {
	name     string
	sessions []Session
}

func newSessionRole(name string) *SessionRole {
	sr := SessionRole{name: name}
	return &sr
}

func (sr *SessionRole) addSession(s Session) {
	sr.sessions = append(sr.sessions, s)
}

func (sr *SessionRole) deleteSessions(sessionName string) {
	// Delete sessions from an array while iterating it
	index := 0
	for _, srs := range sr.sessions {
		if srs.role.name != sessionName {
			sr.sessions[index] = srs
			index++
		}
	}
	sr.sessions = sr.sessions[:index]
}

//
//func (sr *SessionRole) getSessions() []Session {
//	return sr.sessions
//}

func (sr *SessionRole) getSessionRoles() []string {
	names := []string{}
	for _, session := range sr.sessions {

		if !contains(names, session.role.name) {
			names = append(names, session.role.name)
		}

	}
	return names
}

func (sr *SessionRole) hasValidSession(name string, hierarchyLevel int) bool {
	if hierarchyLevel == 1 {
		return sr.name == name
	}

	for _, s := range sr.sessions {

		if s.role.name == name {
			return true
		}
		if s.role.hasValidSession(name, hierarchyLevel-1) {
			return true
		}

	}
	return false
}

func (sr *SessionRole) hasDirectRole(name string) bool {
	for _, session := range sr.sessions {
		if session.role.name == name {

			return true

		}
	}
	return false
}

func (sr *SessionRole) toString() string {
	sessions := ""
	for i, session := range sr.sessions {
		if i == 0 {
			sessions += session.role.name
		} else {
			sessions += ", " + session.role.name
		}

	}
	return sr.name + " < " + sessions
}

// Session represents the activation of a role inheritance for a
// specified time. A role inheritance is always bound to its temporal validity.
// As soon as a session loses its validity, the corresponding role inheritance
// becomes invalid too.
type Session struct {
	role *SessionRole
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
