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

/*
 *	Algorithme de AcceptationImmediate
 *
 *  Renvoie une appariement parfait.
 */

 func AcceptationImmediateAlgorithm(proposants []Agent, disposants []Agent) Mariages {
 	mariages := make(Mariages)  //on crée les mariages
	_p := proposants	//on copie les tableaux
	_d := disposants
	for len(_p) > 0{ //tant qu'il reste des proposants
		propositions := make(Propositions)	//on remet les propositions à 0 à chaque tour
		for i := 0; i < len(_p); i++{	//les proposants proposent
			propositions[_p[i].Prefs[0]] = append(propositions[_p[i].Prefs[0]], _p[i].ID)	//le proposant envoie une proposition à son disposant préféré
		}
		for i := 0; i < len(_d); i++{ //les disposants disposent
			disposantID := _d[i].ID
			if len(propositions[disposantID]) > 0{	//si le disposant a reçu des propositions
				preferedProposant := propositions[disposantID][0]
				for j := 1; j < len(propositions[disposantID]); j++{
					if _d[i].Prefers(GetAgent(proposants, propositions[disposantID][j]), GetAgent(proposants, preferedProposant)){
						preferedProposant = propositions[disposantID][j]
					}
				}
				mariages[preferedProposant] = disposantID	//on crée le mariage
				_p, _d = deleteAgents(preferedProposant, disposantID, _p, _d)	//on met à jour //les agents sont déjà mariés, ils ne seront plus disponibles
			}
		}
	}
 	return mariages	//on retourne tous les mariages
 }

/*	Fonction deleteAgents
 *
 * Permet de supprimer deux agents des groupes d'agents diponibles pour se marier.
 */
 func deleteAgents(proposantID AgentID, disposantID AgentID, proposants []Agent, disposants []Agent) (_p []Agent, _d []Agent){
	 for j := 0; j < len(proposants); j++{
		 if proposantID == proposants[j].ID{
			 proposants = RemoveIndex(proposants, j)
		 }
	 }
	 for j := 0; j < len(disposants); j++{
		 if disposantID == disposants[j].ID{
			 disposants = RemoveIndex(disposants, j)
		 }
	 }
	 for j := 0; j < len(proposants); j++{
		 for i, v := range proposants[j].Prefs {	//on supprime le disposant des préférences des proposants
			 if v == disposantID {
		     proposants[j].Prefs = append(proposants[j].Prefs[:i], proposants[j].Prefs[i+1:]...)
		     break
		   }
		 }

	 	 for i, v := range disposants[j].Prefs {	//on supprime le proposant des préférences des disposants
			 if v == proposantID {
		     disposants[j].Prefs = append(disposants[j].Prefs[:i], disposants[j].Prefs[i+1:]...)
		     break
		   }
		 }
	 }
	 return proposants, disposants
 }

func RemoveIndex(s []Agent, index int) []Agent {
	return append(s[:index], s[index+1:]...)
}


/*
 *	Algorithme d'Acceptation Différée (Gale & Shapley, 1962)
 *
 *  
 */

func AcceptationDifferee(agtA []*Agent, agtB []*Agent, a Mariages){
  engages := make(map[AgentID]bool)
  for len(a) != len(agtA){
    for _,proposant := range agtA{ // recreation des agents a chaque boucle
      // si le proposant est engage pour le moment on passe l'iteration de boucle.
      if _, ok := engages[proposant.ID] ; ok {
          continue
        }
      // sinon on regarde sa preference actuelle, et on la supprime.
      preference := proposant.Prefs[proposant.iterateur]
      proposant.iteration()
      // si le disposant n'est pas deja apparie, on ajoute l'appariement le couple prop, disp
      if _, ok := a[preference] ; !ok {
        a[preference] = proposant.ID
        engages[proposant.ID] = true
      }else{
        // sinon on compare, si il y a preference on change l'appariement et on remet l'ancien prop dans les celibataires.
        if GetAgentById(preference, agtB).Prefers(*proposant,*GetAgentById(a[preference],agtA)){
            delete(engages,GetAgentById(a[preference],agtA).ID)
            a[preference] = proposant.ID
            engages[proposant.ID] = true
        }
      }
    }
  }
}

func cycle_echange(proposant map[AgentID]AgentID, disposant map[AgentID]AgentID) map[AgentID]AgentID { //map[AgentID]AgentID
  cycle := make(map[AgentID]AgentID)
  for disp, prop := range(disposant){
    if _, ok := proposant[prop] ; ok {
      cycle[prop] = disp
  }

}
return cycle
}

func TopTradingCycles (agtA []*Agent, agtB []*Agent, a Mariages){
  // sortie de boucle : len(a) != len(agtA) ça change pas.
  engages := make(map[AgentID]bool)
  engages_prop := make(map[AgentID]bool)
  for len(a) != len(agtA){
    fmt.Println(a)
    //cycle := make(map[AgentID]AgentID)
    pointeurs_prop := make(map[AgentID]AgentID)
    pointeurs_disp := make(map[AgentID]AgentID)
    dispos_pointes := make(map[AgentID]bool)
    var pointeur_prop AgentID
    var pointeur_disp AgentID
    for _,proposant := range agtA{
      // on passe si proposant deja dans l'appariement
      if _, ok := engages[proposant.ID] ; ok {
          continue
        }
      // Pointeur // possibilité de reduire ça correctement
      has_pointeur := true
      for has_pointeur{
        pointeur_prop = GetAgentById(proposant.Prefs[proposant.iterateur], agtB).ID
        // si le dispo est deja engage, on ne peut pas lui proposer.
        if _, ok := engages[pointeur_prop]; !ok {
          has_pointeur = false
        }else{
          proposant.iteration()
        }
      }
      pointeurs_prop[proposant.ID] = pointeur_prop
      dispos_pointes[pointeur_prop] = true
    }

    for _,disposant := range agtB{
        // pour ameliorer la complexite il faudrait faire comme pour les proposants et verifier l'appartenance au map.
        if _, ok := engages[disposant.ID] ; ok {
            continue
        }
        if  _, ok := dispos_pointes[disposant.ID]; ok{
          has_pointeur := true

          for has_pointeur{
            pointeur_disp = GetAgentById(disposant.Prefs[disposant.iterateur], agtA).ID
            if _, ok := engages_prop[pointeur_disp]; !ok {
              has_pointeur = false
            }else{
              disposant.iteration()
            }
          }
          pointeurs_disp[disposant.ID] = pointeur_disp
        }
    }
    // on identifie un/des cycles d'echange, boucle sur pointeurs
    fmt.Println("pointeurs_prop", pointeurs_prop)
    fmt.Println("pointeurs_disp", pointeurs_disp)
    cycle := cycle_echange(pointeurs_prop, pointeurs_disp)

    for prop, disp := range(cycle){
      a[disp] = prop
      engages[GetAgentById(disp, agtB).ID] = true
      engages[GetAgentById(prop, agtA).ID] = true
      engages_prop[GetAgentById(prop, agtA).ID] = true
    }
    //fmt.Println("ok", cycle)
    //
   }
  }

