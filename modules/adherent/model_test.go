package adherent

import (
	"log"
	"testing"
)

func setUpModel() {
	NewAdherent(&Adherent{Firstname: "Test1", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false, University: "Nancy", Homeland: "Pologne", Speakabout: "Facebook", Newsletter: false})
	NewAdherent(&Adherent{Firstname: "Test2", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false, University: "UBFC", Homeland: "Mexique", Speakabout: "Twitter", Newsletter: false})
	NewAdherent(&Adherent{Firstname: "Test3", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false, University: "Lyon lumiere", Homeland: "Mexique", Speakabout: "Twitter", Newsletter: true})
	NewAdherent(&Adherent{Firstname: "Test4", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false, University: "MIT", Homeland: "USA", Speakabout: "Twitter", Newsletter: false})
}

func TestNewAdherent(t *testing.T) {
	NewAdherent(&Adherent{Firstname: "Test", Lastname: "Tutu", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false, University: "UBFC", Homeland: "Mexique", Speakabout: "Twitter", Newsletter: false})
}

func TestFindAdherentById(t *testing.T) {
	adh := FindAdherentById(4)

	if adh.Id != 4 {
		log.Fatal("Adherent is not him")
	}
}

func TestFindAdherentByName(t *testing.T) {
	adh := FindAdherentByName("Test4", "Tutu")

	if adh.Firstname != "Test4" || adh.Lastname != "Tutu" {
		log.Fatal("Adherent is not him")
	}
}

func TestAllAdherents(t *testing.T) {
	adhs := AllAdherents()

	//log.Println(adhs)
	if len(*adhs) == 0 {
		log.Fatal("Adherent is empty")
	}
}

func TestUpdateAdherent(t *testing.T) {
	adh := Adherent{Id: 3, Firstname: "Mario", Lastname: "Bros", Email: "toto@toto.fr", Dateofbirth: "24-04-1995", Student: false, University: "UBFC", Homeland: "Mexique", Speakabout: "Twitter", Newsletter: false}
	UpdateAdherent(&adh)

	adh_3 := FindAdherentById(3)
	if adh_3.Firstname != "Mario" {
		log.Fatal("Adh_3 didn't updated !")
	}
}

func TestDeleteAdherentById(t *testing.T) {
	err := DeleteAdherentById(3)

	if err != nil {
		log.Fatal("Delete had a problem : ", err)
	}

	adhs := AllAdherents()

	for _, adh := range *adhs {
		if adh.Firstname == "Mario" {
			log.Fatal("Adh_3 didn't be removed !")
		}
	}
}
