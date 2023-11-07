package helper

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"time"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"google.golang.org/api/option"
)

func EnvCloudName() string {
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	return os.Getenv("CLOUDINARY_API_SECRET")
}

func EnvCloudUploadFolder() string {
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}

func ImageUploadHelper(input []byte) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Printf("CLOUDINARY_CLOUD_NAME: %s, CLOUDINARY_API_KEY: %s, CLOUDINARY_API_SECRET: %s, CLOUDINARY_UPLOAD_FOLDER: %s", EnvCloudName(), EnvCloudAPIKey(), EnvCloudAPISecret(), EnvCloudUploadFolder())

	// create cloudinary instance
	cld, err := cloudinary.NewFromParams(EnvCloudName(), EnvCloudAPIKey(), EnvCloudAPISecret())
	if err != nil {
		fmt.Println("error on create cloudinary instance", err.Error())
		return "", err
	}

	readerBytes := bytes.NewReader(input)

	// upload file
	uploadParam, err := cld.Upload.Upload(ctx, readerBytes, uploader.UploadParams{Folder: EnvCloudUploadFolder()})
	if err != nil {
		fmt.Println("error on upload file", err.Error())
		return "", err
	}
	return uploadParam.SecureURL, nil
}

func OggToTranscript(oggData []byte) (string, error) {
	ctx := context.Background()

	credsFile := `{  "type": "service_account",  "project_id": "suaquadra-352518",  "private_key_id": "740ded3a2f0dcafbc1f66457bcf497abf1fce358",  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCurqyTiLtrsNye\nppWJ8wh6gX6qDH44oxKKv9Nb1I8qnUsdZIfMo6DrHku3FJZx3DppqI+aOrnhF32h\nfCY9i2J373xYUbNKJcJblMC8CenYsmTfKkZ24Ewuc4XaXVfnHIc1PyOoJHU3hZnx\nBvLHHI8hAu9qRgBMgb0vayfQAGScGWOhbLsu/U6mL3sIdhao8VP454i6XVYKNF+I\nmPzZgtaorZJPFK05+YnwbU+sZZp2kFT0eRF2i3GZ7JzY5advSINzXlGORMjv7rta\nQF/9ISCVkTUqhQLQP1ldz6t/LP22RcE/wxSt3DUUdd9+gITke8I8B3FL9zmcpd0X\nX2v6BCadAgMBAAECggEABoCTBPTggRw2wiMSSu3EgYbjd6H6atJLJOYKEI+DesMb\nIi91TJ1EpqvchqaaCQf5FqjDG6sW8zWEJCgyZjUTh8Je3wy/f0GTkAQj/nvh/AJ7\n9cClmdQ0kcAUGfJCjORHdihxA9fDkzsCZXHsRJQgEcsBrXOInFByAdtbwobZN+Q8\nPxDw/LrHOpPobFH09IwYicscS57eKX4bmJK6I1Zxwpan4V4oUZH6sypaBjxeY6Ba\nar+FRuR/7CxrbdM7PtpGPmplvvWoXu1KVcgiEg4ITf2TdyVMI43EUsQ8r7mdRYaC\nVW3Ov0ICXlqV0DzflOlH9j8ycVhSv76nzuaL4eIukQKBgQDmu2i3ehuWinSaIxkT\nyNi08ndppeH71wJ5pt9q9eH3rs41PgfiN24yN0h5XZ+23aMiSKJHzcl6tmqJWV+K\njmsCYw5VL45nCRCBc2sruUauUh9lnT1SXShXEsyworzie0gB7ZTCAhQS5DLp+o3M\noPpkxpSko/ehJPTBEdK+JTgdOQKBgQDBz+f9cdmXOZIpcA2zVK27NomB6KcHrf+R\n9q3LY6Vfhgcw/VE/GgT94TmzP6xFS+2PiDf1iHWxWRwC/zfs6YVK167IIAVrcUQz\nLZC/If/fPaII+6Ya37ZR8Ej6wpdaptbvO3/KPiVBDhvMjhuPaY9wtlY8aZV96SQ4\n+qasoby4hQKBgG0r5XtS4nTTZCJ+UuJCmQ7c3tV1MLz5Wel3pKS2XMnVwyn8BLzr\nn09RSxBp1SUwL30MQwSYgSl34GbGi+dCRa2mcuSpkMZ9ynqFwwK4MpJOtx5cTOSI\nwYqZkZJOHfNHg6Wt5UH6u8bIhLKi468bx/4g27oe/w8XLluf9EgV0jw5AoGAZC/r\nRrKRvoC+M7l++5Lsc/iPQJ5Zqbiignu3/4m1NRn9oa9xTmNO7UZ+I1Do5rGHTkkm\nVECERnc/6bSw3kEg7D1uVnlnE9FrZeFKD+Otd2NO3cHobb/zaYwCzc3Fm9DfDq/9\nMTjK3URDzowvZwU0Zxl6nqQd6QaZ+PJpMpgxDFkCgYEA1nWVwOnP7ORzJG2HxV5f\n+bD0O824yY/lurrqUjnjie/AdB/h2LQL7Hirq5WZ/pvBKRIuWv+/CFgV9/PJuMn9\nUzRKDYyIPlkWjKRHNxtdJC47OFamBUfSSp4CJSrQQ1H4BIa/yQ0c3ZJ1WomhT8Rp\nUQgtdrTDWyhjLFD7hEFsh20=\n-----END PRIVATE KEY-----\n",  "client_email": "transcript-audio@suaquadra-352518.iam.gserviceaccount.com",  "client_id": "114891471909474350465",  "auth_uri": "https://accounts.google.com/o/oauth2/auth",  "token_uri": "https://oauth2.googleapis.com/token",  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/transcript-audio%40suaquadra-352518.iam.gserviceaccount.com"}`

	// Creates a client.

	client, err := speech.NewClient(ctx, option.WithCredentialsJSON([]byte(credsFile)))
	if err != nil {
		return "", err
	}

	// Builds the recognition request.
	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_OGG_OPUS,
			SampleRateHertz: 16000,
			LanguageCode:    "pt-BR",
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{Content: oggData},
		},
	}

	// Performs speech recognition on the audio data.
	resp, err := client.Recognize(ctx, req)
	if err != nil {
		return "", err
	}

	// Concatenates the transcription results.
	var transcript string
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			transcript += alt.Transcript
		}
	}

	return transcript, nil
}
