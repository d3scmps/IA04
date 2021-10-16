package main

/*
 *	Algorithme de DynamiqueLibre
 *
 *  Renvoie une appariement stable si l'algorithme termine (ce qui n'est pas toujours le cas).
 */

func stabilizeMariage(proposant Agent, disposant Agent, proposant2 Agent, disposant2 Agent) (proposantID AgentID, disposantID AgentID, proposantID2 AgentID, disposantID2 AgentID){
  if proposant.Prefers(disposant2, disposant) && disposant2.Prefers(proposant, proposant2){ //le mariage n'est pas stable
    return proposant.ID, disposant2.ID, proposant2.ID, disposant.ID //si le mariage n'est pas stable, on inverse l'ordre des AgentID
  }
  if disposant.Prefers(proposant2, proposant) && proposant2.Prefers(disposant, disposant2){ //le mariage n'est pas stable
    return proposant.ID, disposant2.ID, proposant2.ID, disposant.ID //si le mariage n'est pas stable, on inverse l'ordre des AgentID
  }
  return proposant.ID, disposant.ID, proposant2.ID, disposant2.ID //si le mariage est stage on renvoie les AgentID dans le bon ordre.
}

func DynamiqueLibreAlgorithm(proposants []Agent, disposants []Agent) Mariages{
  mariages := make(Mariages)  //on crée les mariages

  //Initialisement de l'algorithme avec un appariement parfait quelconque
  for i := 0; i < len(proposants); i++{   //on crée N mariages quelconques
    mariages[proposants[i].ID] = disposants[i].ID
  }

  for !mariages.IsStable(proposants, disposants){
    for proposantID := range mariages{
      for proposantID2 := range mariages{
        disposantID := mariages[proposantID]
        disposantID2 := mariages[proposantID2]

        if proposantID != proposantID2 && disposantID != disposantID2{
          pID, dID, pID2, dID2 := stabilizeMariage(
            GetAgent(proposants, proposantID),
            GetAgent(disposants, disposantID),
            GetAgent(proposants, proposantID2),
            GetAgent(disposants, disposantID2),
          )

          mariages[pID] = dID //on met à jour les mariages
          mariages[pID2] = dID2
        }
      }
    }
  }
  return mariages
}
