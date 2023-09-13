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