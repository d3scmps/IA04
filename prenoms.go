package main

import (
  "os"
	"bufio"
  "strconv"
  "strings"
  "fmt"
	"math/rand"
)

type Prenom struct{
  prenom    string
  genre     string
  langages  []string
  frequence float64
}

/*
 *  Constructeur Prenom
 *
 *  Permet de créer un nouvel objet prenom
 */
func NewPrenom(prenom string, genre string, langages []string, frequence float64) Prenom{
  return Prenom{
    prenom,
    genre,
    langages,
    frequence,
  }
}

//Dans notre ensemble de prénoms, les sexes sont soient masculins soient féminins. Ce qui est discutable en 2021...

/*
 *  Fonction isMale:
 *
 *  Retourne True si le prénom est masculin, False sinon
 */
func (p Prenom) IsMale() bool{
  return p.genre == "m"
}

/*
 *  Fonction isFemale:
 *
 *  Retourne True si le prénom est masculin, False sinon
 */
func (p Prenom) IsFemale() bool{
  return p.genre == "f"
}

/*
 *  Fonction ExtractPrenoms
 *
 *  Retourne un tableau de Prenom en fonction des prénoms présents dans le fichier Prenoms.csv
 */
func ExtractPrenoms() []Prenom{
  csvFile, err := os.Open("Prenoms.csv")  //ouverture du fichier
  if err != nil {
     fmt.Println(err.Error() + " Prenoms.csv")
     return nil
  }
  defer csvFile.Close() //on fermera le fichier à la fin de la fonction

  scanner := bufio.NewScanner(csvFile)

  var prenoms []Prenom
  for scanner.Scan() {
      line := scanner.Text()
      elements := strings.Split(line, ";")  //on découpe la ligne

      frequence, error := strconv.ParseFloat(elements[3], 64) //la fréquence n'est pas un string, on cast

      languages := strings.Split(elements[2], ", ") //Un prénom peut exister dans plusieurs langues

      if(error == nil){ //la première ligne renverra nil car il s'agit des headers
        prenomAlreadyPresent := false//Il est possible d'avoir des homonymes, dans ce cas on ajoute les langues et la fréquence

        for i := 0; i < len(prenoms); i++{
          if(prenoms[i].prenom == strings.Title(elements[0])){
            prenomAlreadyPresent = true

            prenoms[i].frequence = prenoms[i].frequence + frequence//on ajoute la fréquence

            for k := 0; k < len(languages); k++{  //mise à jour des langues de l'ancien prénom
              languageAlreadyPresent := false
              for j := 0; j < len(prenoms[i].langages); j++{
                if languages[k] == prenoms[i].langages[j]{
                  languageAlreadyPresent = true
                }
              }
              if !languageAlreadyPresent{
                prenoms[i].langages = append(prenoms[i].langages, languages[k])
              }
            }
          }
        }

        if !prenomAlreadyPresent{ //sinon on ajoute le prénom
          prenoms = append(prenoms, NewPrenom(
            strings.Title(elements[0]),
            elements[1],
            languages,
            frequence,
          ))
        }
      }
  }
  return prenoms
}

/*
 *  Fonction ShufflePrenoms
 *
 *  Retourne un tableau de Prenom mélangé de taille n
 */
func ShufflePrenoms(prenoms []Prenom, n int) []Prenom{
  for i := range prenoms {
      j := rand.Intn(i + 1)
      prenoms[i], prenoms[j] = prenoms[j], prenoms[i]
  }
  return prenoms
}

/*
 *  Fonction GetPrenomsLanguages
 *
 *  Retourne la liste des langues disponibles pour une liste de prénoms donnée.
 */
func GetPrenomsLanguages(prenoms []Prenom) []string{
  var languages []string;
  for i := 0; i < len(prenoms); i++{  //Pour chaque prénom
    for k := 0; k < len(prenoms[i].langages); k++{
      languageAlreadyPresent := false
      for j := 0; j < len(languages); j++{  //On regarde si la langue n'est pas déjà dans la liste des langues connues
        if(languages[j] == prenoms[i].langages[k]){
          languageAlreadyPresent = true
        }
      }

      if !languageAlreadyPresent{ //si la langue n'est pas déjà connue, on l'ajoute à la liste des langues connues
        languages = append(languages, prenoms[i].langages[k])
      }
    }
  }

  return languages  //on retourne la liste des langues
}

/*
 * Fonction FilterPrenoms
 *
 * Retourne une liste contenant tous les prénoms en fonction d'un genre donné ainsi qu'une liste de langues.
 */

func FilterPrenoms(prenoms []Prenom, genre string, languages []string) []Prenom{
  var newPrenoms []Prenom

  for i := 0; i < len(prenoms); i++{  //Pour chaque prénom
    if prenoms[i].genre == genre{ //Si le genre est valide
      validLanguage := false
      for j := 0; j < len(languages); j++{  //On regarde si la langue est valide
        for k := 0; k < len(prenoms[i].langages); k++{
          if prenoms[i].langages[k] == languages[j]{
            validLanguage = true
          }
        }
      }
      if validLanguage{ //Si la langue est valide on ajoute le prénom à la liste des prénoms filtrés
        newPrenoms = append(newPrenoms, prenoms[i])
      }
    }
  }

  return newPrenoms
}

/*
 * Fonction GetAllMales
 *
 * Retourne une liste contenant tous les prénoms masculins d'une liste de prénoms.
 */
func GetAllMales(prenoms []Prenom) []Prenom{
  return FilterPrenoms(prenoms, "m", GetPrenomsLanguages(prenoms))
}

/*
 * Fonction GetAllFemales
 *
 * Retourne une liste contenant tous les prénoms féminin d'une liste de prénoms.
 */
func GetAllFemales(prenoms []Prenom) []Prenom{
  return FilterPrenoms(prenoms, "f", GetPrenomsLanguages(prenoms))
}

/*
 * Fonction GetMalesByLang
 *
 * Retourne une liste contenant tous les prénoms masculins d'une liste de prénoms en fonction d'une langue donnée.
 */
func GetMalesByLang(prenoms []Prenom, language string) []Prenom{
  return FilterPrenoms(prenoms, "m", []string{language})
}

/*
 * Fonction GetFemalesByLang
 *
 * Retourne une liste contenant tous les prénoms féminins d'une liste de prénoms en fonction d'une langue donnée.
 */
func GetFemalesByLang(prenoms []Prenom, language string) []Prenom{
  return FilterPrenoms(prenoms, "f", []string{language})
}
