package main

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/usecases/utils"
)

func main() {
	resume := models.ResumeInPDF{
		Image:             "1.webp",
		Title:             "Backend developer",
		ApplicantName:     "Владислав",
		ApplicantSurname:  "Кирпичов",
		Location:          "Moscow",
		ContactNumber:     "+7 (999) 999-99-99",
		Email:             "helloworld@mail.ru",
		Age:               20,
		ExperienceInYears: "1",
		Description:       "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam interdum purus vel felis scelerisque, ac interdum nibh posuere. Ut dapibus, dolor a accumsan dapibus, metus nibh rhoncus nibh, sit amet aliquam est odio a nisl. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Praesent aliquet sed turpis et pulvinar. Praesent tristique, mi non mattis tincidunt, felis nulla dictum eros, et elementum urna metus eu ante. Maecenas nec aliquet nibh. Fusce mattis diam urna, eu aliquet mi lacinia vel. Praesent nisi arcu, vehicula tristique ipsum ut, tincidunt placerat nisl. Nam in sem at nisi imperdiet placerat. Nunc blandit eget eros at ullamcorper. Aliquam imperdiet dictum euismod. Ut sed purus a ipsum imperdiet ultrices. Fusce in est sed sem pharetra porta sit amet eu quam. Donec a lacus dolor. Quisque ullamcorper velit dui, in eleifend sapien tempor id. Ut egestas tincidunt ullamcorper.",
	}

	utils.GenerateResumeInPDF(&resume, &configs.ImageConfig{
		Path:                   "./data/image/",
		DefaultApplicantAvatar: "./",
		DefaultAverageColor:    ",/.",
		DefaultEmployerAvatar:  "./",
	}, "serif_fonts_resume.html")
}
