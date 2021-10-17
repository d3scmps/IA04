package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type AgentID string

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []AgentID
	iterateur int
}

func TableauAgTOtableauPointeursAg(tab []Agent) []*Agent {
  taille := len(tab)
  tabbis := make([]*Agent, taille)
  for i, _ := range(tab){
    tabbis[i] = &tab[i]
  }
  return tabbis
}

func GetAgentById(id AgentID, tab []*Agent)  (x *Agent){
  for _,agent := range(tab){
    if agent.ID == id{
      return agent
    }
  }
  return nil
}

func (agent *Agent) iteration() {
  agent.iterateur ++
}

func NewAgent(id AgentID, name string, prefs []AgentID) Agent {
	return Agent{
		id,
		name,
		prefs,
		0,
	}
}

func Equal(ag1 Agent, ag2 Agent) bool {
	if ag1.ID != ag2.ID {
		return false
	}

	// Pas obligatoire à partir de là, à discuter...
	if ag1.Name != ag2.Name {
		return false
	}

	if len(ag1.Prefs) != len(ag2.Prefs) {
		return false
	}

	for i := range ag1.Prefs {
		if ag1.Prefs[i] != ag1.Prefs[i] {
			return false
		}
	}

	return true
}

func (a Agent) String() string {
	return fmt.Sprintf("%s %s %v", a.ID, a.Name, a.Prefs)
}

func (a Agent) rank(b Agent) (int, error) {
	for i, v := range a.Prefs {
		if v == b.ID {
			return i, nil
		}
	}
	return -1, errors.New("Agent %s not found" + string(b.ID))
}

// renvoie vrai si ag préfère a à b
func (ag Agent) Prefers(a, b Agent) bool {
	r1, err1 := ag.rank(a)
	if err1 != nil {
		return false
	}

	r2, err2 := ag.rank(b)
	if err2 != nil {
		return false
	}

	return r1 < r2
}

func RandomPrefs(ids []AgentID) (res []AgentID) {
	res = make([]AgentID, len(ids))
	copy(res, ids)
	rand.Shuffle(len(res), func(i, j int) { res[i], res[j] = res[j], res[i] })
	return
}
