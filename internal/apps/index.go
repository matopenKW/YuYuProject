package apps

import (
	"YuYuProject/internal/dto"
	"log"
)

func ShowMainPage() (map[string]interface{}, error) {
	list, eastList, westList := getTeamList()
	log.Print(list)

	return map[string]interface{}{
		"allList":  list,
		"eastList": eastList,
		"westList": westList,
	}, nil

}

func getTeamList() (allList, eastList, westList []*dto.Team) {

	ALL := 30
	WEST := 15
	EAST := 15

	A_WEST := 3
	A_EAST := 4

	B_WEST := 3
	B_EAST := 6

	C_WEST := 9
	C_EAST := 5

	A := &dto.Team{
		"A", "a_team", (A_WEST + A_EAST) * 100 / ALL, "",
	}

	B := &dto.Team{
		"B", "b_team", (B_WEST + B_EAST) * 100 / ALL, "",
	}

	C := &dto.Team{
		"C", "c_team", (C_WEST + C_EAST) * 100 / ALL, "",
	}
	allList = []*dto.Team{A, B, C}

	for _, v := range allList {
		log.Println(v)
	}

	A = &dto.Team{
		"A", "a_team", A_WEST * 100 / WEST, "",
	}

	B = &dto.Team{
		"B", "b_team", B_WEST * 100 / WEST, "",
	}

	C = &dto.Team{
		"C", "c_team", C_WEST * 100 / WEST, "",
	}
	eastList = []*dto.Team{A, B, C}

	A = &dto.Team{
		"A", "a_team", A_EAST * 100 / EAST, "",
	}

	B = &dto.Team{
		"B", "b_team", B_EAST * 100 / EAST, "",
	}

	C = &dto.Team{
		"C", "c_team", C_EAST * 100 / EAST, "",
	}
	westList = []*dto.Team{A, B, C}

	return
}

func getRate(all int, num ...int) int {

	return 0
}
