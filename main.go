package main

import(
  "fmt"
)

/* utils */

func GetAgentById(id AgentID, tab []Agent)  (x *Agent){
  for _,agent := range(tab){
    if agent.ID == id{
      return &agent
    }
  }
  return nil
}

func GetAgentByIdbis(id AgentID, tab []*Agent)  (x *Agent){
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

func (agent *Agent) PrefIterator () (id  AgentID){
  pref := agent.Prefs[0]
  if len(agent.Prefs) > 1{
    agent.Prefs = agent.Prefs[1:]
  }
  return pref
}
/* --- */

/*
type Appariement struct{
  a Agent
  b Agent
}*/

func main(){
  /* groupe proposant */
  Anames := [...]string{
		"Khaled",
		"Sylvain",
		"Emmanuel",
		"Bob",
	}
  /* groupe disposant */
  Bnames := [...]string{
    "Nathalie",
    "Annaïck",
    "Brigitte",
    "Anaelle",
  }


// Init agents
poolA := make([]*Agent, 0, len(Anames))
poolB := make([]Agent, 0, len(Bnames))


groupA_prefix := "a"
groupB_prefix := "b"

prefsA := make([]AgentID, len(Anames))
prefsB := make([]AgentID, len(Bnames))

for i := 0; i < len(Anames); i++ {
  prefsA[i] = AgentID(groupA_prefix + fmt.Sprintf("%d", i))
}

for i := 0; i < len(Bnames); i++ {
  prefsB[i] = AgentID(groupB_prefix + fmt.Sprintf("%d", i))
}

for i := 0; i < len(Anames); i++ {
  prefs := RandomPrefs(prefsB)
  a := Agent{prefsA[i], Anames[i], prefs,0}
  poolA = append(poolA, &a)
}

for i := 0; i < len(Bnames); i++ {
  prefs := RandomPrefs(prefsA)
  b := Agent{prefsB[i], Bnames[i], prefs,0}
  poolB = append(poolB, b)
}

 //x := get_agent_by_id("a1", poolA)
 //a :=  make(map[AgentID]AgentID)
 a2 :=  make(map[AgentID]AgentID)
 //AcceptationImmediate(poolA, poolB, a)
// AcceptationDifferee(poolA, poolB, a)
 TopTradingCycles(poolA, poolB, a2)
 fmt.Println(poolA)
 fmt.Println(poolB)
 fmt.Println(a2)
}

/* ---- Acceptation immediate ----- */

func AcceptationImmediate(agtA []Agent, agtB []Agent, a map[AgentID]AgentID){
  pref_tour :=  make(map[AgentID]AgentID)
  poolA := agtA
  c := 0
  for  len(poolA) != 0{
    for _, agentA := range(poolA){
      agentB := GetAgentById(agentA.Prefs[c], agtB)
      if _, ok := pref_tour[agentA.Prefs[c]]; !ok{
        pref_tour[agentA.Prefs[c]] = agentA.ID
      }else if agentB.Prefers(agentA, *GetAgentById(pref_tour[agentA.Prefs[c]], agtA)){
        pref_tour[agentA.Prefs[c]] = agentA.ID
      }
    }

  //  On retire les agents de poolA contenus dans prefTour et actualiser poolA, ajouter à l'Appariement
  blacklist_ids := make(map[AgentID]bool)
  for agentB, agentA := range(pref_tour){
  a[agentB] = agentA
  blacklist_ids[agentA] = true
  }
  //
  newpoolA := make([]Agent, 0, len(poolA))
  for _,agentA := range(poolA){
    if _, ok := blacklist_ids[agentA.ID]; !ok {
      newpoolA = append(newpoolA, agentA)
    }
  }
  poolA = newpoolA
  c++
  }
}

// Gale & Shapley

func AcceptationDifferee(agtA []*Agent, agtB []Agent, a map[AgentID]AgentID){
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
        if GetAgentById(preference, agtB).Prefers(*proposant,*GetAgentByIdbis(a[preference],agtA)){
            delete(engages,GetAgentByIdbis(a[preference],agtA).ID)
            a[preference] = proposant.ID
        }
      }
    }
  }
}


// TTC - top trading cycles

func cycle_echange(proposant map[AgentID]AgentID, disposant map[AgentID]AgentID) { //map[AgentID]AgentID
  cycle := make(map[AgentID]AgentID)
  fin_recherche := make(map[AgentID]bool)
  pre_cycle := make([]AgentID, len(proposant))
  var dernier_element AgentID
  // init dernier_element
  for prop, _ := range proposant{
    dernier_element = prop
    break
  }
  fmt.Println(dernier_element)
  // condition de sortie ? fin_recherche contient deux fois AgentID
  c := 0
  for {
      if _, ok := fin_recherche[dernier_element]; ok {
        break
      }
      fin_recherche[dernier_element] = true
      if c % 2 == 0 { // disposant
        dernier_element = proposant[dernier_element]
      }else{
        dernier_element = disposant[dernier_element]
      }
      

  }
}

func TopTradingCycles (agtA []*Agent, agtB []Agent, a map[AgentID]AgentID){
  // sortie de boucle : len(a) != len(agtA) ça change pas.
  engages := make(map[AgentID]bool)
  engages_prop := make(map[AgentID]bool)
  for len(a) != len(agtA){
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
            pointeur_disp = GetAgentByIdbis(disposant.Prefs[disposant.iterateur], agtA).ID
            if _, ok := engages_prop[pointeur_disp]; !ok {
              has_pointeur = false
            }else{
              disposant.iteration()
            }
          }
          pointeurs_disp[disposant.ID] = pointeur_disp
        }
    }
    // on identifie un cycle d'echange, boucle sur pointeurs
    cycle_echange(pointeurs_prop, pointeurs_disp)
    //fmt.Println("ok", cycle)
    //
   }
  }
  // chaque agent de agtA pointe vers son pref du groupe B boucle agtA
      // on boucle sur agtB, on regarde ceux qui sont pointés par A.

  //

  // puis on apparie.

/*
func AcceptationDifferee(agtA []Agent, agtB []Agent, a map[AgentID]AgentID){
  // taille des clefs du a doit etre egale à la taille de agtA. -> condition de sortie de boucle.
  poolA := agtA
  pref_tour :=  make(map[AgentID]AgentID)
  // couple temporaire pour simuler le peut-etre
  // blacklist pour les nons.
  blacklist :=  make(map[AgentID][]AgentID)
  c := 0
  for len(a) != len(agtA) {
    for _,agentA := range(poolA){
        agentB := GetAgentById(agentA.Prefs[c], agtB)
        // verifie que agentA soit pas cancel
        isBlackliste := false
        for _,blackliste := range(blacklist[agentB.ID]){
          if blackliste == agentA.ID{
            isBlackliste = true
          }
        }
        if isBlackliste{
          continue
        }
        if  _, ok := a[agentA.Prefs[c]]; !ok{
          pref_tour[agentA.Prefs[c]] = agentA.ID
          a[agentB.ID] = agentA.ID
        }else if agentB.Prefers(*GetAgentById(pref_tour[agentA.Prefs[c]], agtA),agentA ) {
            // l'ancien d'agentB redevient celib -> on le remet dans le poolA ?
            delete(a, a[agentB.ID])
            poolA = append(poolA, *GetAgentById(pref_tour[agentB.ID], agtA))
            // agentA maque a agentB
            pref_tour[agentB.ID] = agentA.ID
            a[agentB.ID] = agentA.ID
        }else{
          // agentA est cancel par agentB
          blacklist[agentB.ID] = append(blacklist[agentB.ID], agentA.ID)
        }
      }
      fmt.Println(len(a))
      fmt.Println(len(agtA))
    }
    c++
  }
*/

/*
func AcceptationDifferee(agtA []Agent, agtB []Agent, a map[AgentID]AgentID){
  for len(agtA) != 0 {}
}*/

/* --- Instabilite  --- */
