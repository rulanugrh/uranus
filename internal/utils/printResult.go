package utils

import (
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/domain/web"
)

func PrintResultCreateCourse(req entity.Course) web.ResponseCreateCourse {
	response := web.ResponseCreateCourse {
		Name: req.Name,
		Description: req.Description,
		Category: req.Categories.Name,
		MaxParticipant: req.MaxParticipant,
		Waktu: req.Waktu,
	}

	return response
}

func PrintResultCourse(req entity.Course) web.ResponseFindCourse {
	var listParticipant []web.ListParticipant
	for _, data := range req.Participant {
		participant := web.ListParticipant {
			UserName: data.UserMod.Name,
			UserEmail: data.UserMod.Email,
		}

		listParticipant = append(listParticipant, participant)
	}

	response := web.ResponseFindCourse {
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
		Category: req.Categories.Name,
		MaxParticipant: req.MaxParticipant,
		Waktu: req.Waktu,
		Participant: listParticipant,
	}

	return response
}

func PrintResultAllCourse(req []entity.Course) []web.ResponseFindCourse {
	var response []web.ResponseFindCourse
	for _, data := range req {
		var listParticipant []web.ListParticipant
		for _, user := range data.Participant{
			users := web.ListParticipant{
				UserName: user.UserMod.Name,
				UserEmail: user.UserMod.Email,
			}

			listParticipant = append(listParticipant, users)
		}

		res := web.ResponseFindCourse{
			Name: data.Name,
			Price: data.Price,
			Category: data.Categories.Name,
			Description: data.Description,
			Participant: listParticipant,
			Waktu: data.Waktu,
			MaxParticipant: data.MaxParticipant,
		}

		response = append(response, res)
		
	}

	return response
}