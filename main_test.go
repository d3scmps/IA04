package main

import (
  "testing"
  "fmt"
)

func TestAcceptationDifferree(t *testing.T){

  // Verification de la bonne terminaison de l'algorithme pour n = 4
  Anames := [...]string{
		"Khaled",
		"Sylvain",
		"Emmanuel",
		"Bob",
	}
  Bnames := [...]string{
    "Nathalie",
    "Annaïck",
    "Brigitte",
    "Anaelle",
  }


// Init agents
poolA := make([]*Agent, 0, len(Anames))
poolB := make([]*Agent, 0, len(Bnames))


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
  poolB = append(poolB, &b)
}

var verif = map[AgentID]AgentID{
    "b0":"a3",
    "b1":"a2",
    "b2":"a1",
    "b3":"a0",
}
   resultat := make(Mariages)
   AcceptationDifferee(poolA, poolB, resultat)

   for key, value := range(resultat){
     if value != verif[key]{
       t.Errorf("L'algorithme d'acceptation différé 'est pas correct, l'appariement renvoyé n'est pas le bon.")
     }
   }

  // Verification de la stabilité de l'algorithme pour un n donné (test sur n=50 ici)

   test := ExtractPrenoms()
   teststab := make(Mariages)
   poolA2, poolB2 := GenerateAgents(test, 50)
   poolAbis := TableauAgTOtableauPointeursAg(poolA2)
   poolBbis := TableauAgTOtableauPointeursAg(poolB2)
   AcceptationDifferee(poolAbis, poolBbis, teststab)
   est_stable := teststab.IsStable(poolA2, poolB2)

   if !est_stable{
     t.Errorf("L'appariement ressorti par l'algorithme d'acceptation différé n'est pas stable ")
   }


}
