package main

import (
	"fmt"
)
type Mariages map[AgentID]AgentID

func (m Mariages) Debug(){
  for proposantID, disposantID := range m{
    fmt.Println(proposantID, "is married with", disposantID, ".")
  }
}

func (m Mariages) IsStable(proposants []Agent, disposants []Agent) bool{
  for proposantID := range m{
    for proposantID2  := range m{
      disposantID := m[proposantID]
      disposantID2 := m[proposantID2]

      if(proposantID != proposantID2 && disposantID != disposantID2){

        proposant := GetAgent(proposants, proposantID)
        disposant := GetAgent(disposants, disposantID)
        proposant2 := GetAgent(proposants, proposantID2)
        disposant2 := GetAgent(disposants, disposantID2)

        if proposant.Prefers(disposant2, disposant) && disposant2.Prefers(proposant, proposant2){ //le mariage n'est pas stable
          return false
        }
        if disposant.Prefers(proposant2, proposant) && proposant2.Prefers(disposant, disposant2){ //le mariage n'est pas stable
          return false
        }
      }
    }
  }
  return true
}
