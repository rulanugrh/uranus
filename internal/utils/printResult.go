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

func PrintResultUser(req entity.User) web.ResponseUser {
	response := web.ResponseUser {
		Name: req.Name,
		Email: req.Email,
		Avatar: req.Avatar,
		Address: req.Avatar,
		Notelp: req.Notelp,
	}

	return response
}

func PrintResultCreateCategory(req entity.Category) web.ResponseCreateCategory {
	response := web.ResponseCreateCategory{
		Name: req.Name,
		Description: req.Description,
	}

	return response
}

func PrintResultCategoryByID(req entity.Category) web.ResponseFindCategory {
	var course []web.ListCourse
	for _, data := range req.Courses {
		courses := web.ListCourse{
			Name: data.Name,
			Description: data.Description,
			Waktu: data.Waktu,
			MaxParticipant: data.MaxParticipant,
			Price: data.Price,
		}

		course = append(course, courses)
	}

	response := web.ResponseFindCategory{
		Name: req.Name,
		Description: req.Description,
		Course: course,
	}

	return response
}

func PrintResultCategory(req []entity.Category) []web.ResponseFindCategory {
	var response []web.ResponseFindCategory
	for _, data := range req{
		var courseList []web.ListCourse
		for _, courses := range data.Courses{
			course := web.ListCourse{
				Name: courses.Name,
				Price: courses.Price,
				MaxParticipant: courses.MaxParticipant,
				Waktu: courses.Waktu,
				Description: courses.Description,
			}

			courseList = append(courseList, course)
		}

		res := web.ResponseFindCategory{
			Name: data.Name,
			Description: data.Description,
			Course: courseList,
		}

		response = append(response, res)
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