package repository

//import (
//	"HeadHunter/configs"
//	"github.com/kolesa-team/go-webp/decoder"
//	"github.com/kolesa-team/go-webp/webp"
//	"image/jpeg"
//	"log"
//	"os"
//)
//
//func GetImage(path, name string, cfg *configs.ImageConfig) (err error) {
//	file, err := os.Open("test_data/images/m4_q75.webp")
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	output, err := os.Create("example/output_decode.jpg")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer output.Close()
//
//	img, err := webp.Decode(file, &decoder.Options{})
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	if err = jpeg.Encode(output, img, &jpeg.Options{Quality: 75}); err != nil {
//		log.Fatalln(err)
//	}
//	return nil
//}
//
//func DeleteImage(path, name string, cfg *configs.ImageConfig) (err error) {
//	return nil
//}
//
//func CreateImage(path, name string, cfg *configs.ImageConfig) (err error) {
//	return nil
//}
