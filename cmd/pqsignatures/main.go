package main

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fogleman/gg"
	"github.com/kevinburke/nacl/secretbox"

	"fmt"
	"time"

	pqtime "github.com/PastelNetwork/pqsignatures/internal/time"
	"github.com/PastelNetwork/pqsignatures/legroast"
	"github.com/PastelNetwork/pqsignatures/qr"

	"encoding/base64"
	"encoding/hex"

	"golang.org/x/crypto/sha3"

	"github.com/auyer/steganography"
)

const (
	PastelIdSignatureFilesFolder = "pastel_id_signature_files"
)

func get_image_hash_from_image_file_path_func(sample_image_file_path string) (string, error) {
	f, err := os.Open(sample_image_file_path)
	if err != nil {
		return "", err
	}

	defer f.Close()
	hash := sha3.New256()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func pastel_id_keypair_generation_func() (string, string) {
	fmt.Println("\nGenerating LegRoast keypair now...")
	pk, sk := legroast.Keygen()
	pastel_id_private_key_b16_encoded := base64.StdEncoding.EncodeToString(sk)
	pastel_id_public_key_b16_encoded := base64.StdEncoding.EncodeToString(pk)
	return pastel_id_private_key_b16_encoded, pastel_id_public_key_b16_encoded
}

func write_pastel_public_and_private_key_to_file_func(pastel_id_public_key_b16_encoded string, pastel_id_private_key_b16_encoded string, box_key_file_path string) {
	pastel_id_public_key_export_format := "-----BEGIN LEGROAST PUBLIC KEY-----\n" + pastel_id_public_key_b16_encoded + "\n-----END LEGROAST PUBLIC KEY-----"
	pastel_id_private_key_export_format := "-----BEGIN LEGROAST PRIVATE KEY-----\n" + pastel_id_private_key_b16_encoded + "\n-----END LEGROAST PRIVATE KEY-----"
	box_key, err := naclBoxKeyFromFile(box_key_file_path)
	if err != nil {
		log.Fatal(err)
	}
	var key [32]byte
	copy(key[:], box_key)
	encrypted := secretbox.EasySeal(([]byte)(pastel_id_private_key_export_format[:]), &key)

	key_file_path := "pastel_id_key_files"
	if _, err := os.Stat(key_file_path); os.IsNotExist(err) {
		os.MkdirAll(key_file_path, 0770)
	}
	err = os.WriteFile(key_file_path+"/pastel_id_legroast_public_key.pem", []byte(pastel_id_public_key_export_format), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(key_file_path+"/pastel_id_legroast_private_key.pem", []byte(encrypted), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func generateKeypairQRs(pk string, sk string) ([]qr.Image, error) {
	key_file_path := "pastel_id_key_files"
	pkPngs, err := qr.Encode(pk, key_file_path, "Pastel Public Key", "pastel_id_legroast_public_key_qr_code", "")
	if err != nil {
		return nil, err
	}
	_, err = qr.Encode(sk, key_file_path, "", "pastel_id_legroast_private_key_qr_code", "")
	if err != nil {
		return nil, err
	}
	return pkPngs, nil
}

func pastel_id_write_signature_on_data_func(input_data_or_string string, pastel_id_private_key_b16_encoded string, pastel_id_public_key_b16_encoded string) string {
	fmt.Printf("\nGenerating LegRoast signature now...")
	defer pqtime.Measure(time.Now())

	pastel_id_private_key, _ := base64.StdEncoding.DecodeString(pastel_id_private_key_b16_encoded)
	pastel_id_public_key, _ := base64.StdEncoding.DecodeString(pastel_id_public_key_b16_encoded)
	pqtime.Sleep()
	pastel_id_signature := legroast.Sign(pastel_id_public_key, pastel_id_private_key, ([]byte)(input_data_or_string[:]))
	pastel_id_signature_b16_encoded := base64.StdEncoding.EncodeToString(pastel_id_signature)
	pqtime.Sleep()
	return pastel_id_signature_b16_encoded
}

func pastel_id_verify_signature_with_public_key_func(input_data_or_string string, pastel_id_signature_b16_encoded string, pastel_id_public_key_b16_encoded string) int {
	fmt.Printf("\nVerifying LegRoast signature now...")
	defer pqtime.Measure(time.Now())

	pastel_id_signature, _ := base64.StdEncoding.DecodeString(pastel_id_signature_b16_encoded)
	pastel_id_public_key, _ := base64.StdEncoding.DecodeString(pastel_id_public_key_b16_encoded)
	pqtime.Sleep()
	verified := legroast.Verify(pastel_id_public_key, ([]byte)(input_data_or_string[:]), pastel_id_signature)
	pqtime.Sleep()
	if verified > 0 {
		fmt.Printf("\nSignature is valid!")
	} else {
		fmt.Printf("\nWarning! Signature was NOT valid!")
	}
	return verified
}

func import_pastel_public_and_private_keys_from_pem_files_func(box_key_file_path string) (string, string) {
	key_file_path := "pastel_id_key_files"
	if info, err := os.Stat(key_file_path); os.IsNotExist(err) || !info.IsDir() {
		fmt.Printf("\nCan't find key storage directory, trying to use current working directory instead!")
		key_file_path = "."
	}
	pastel_id_public_key_pem_filepath := key_file_path + "/pastel_id_legroast_public_key.pem"
	pastel_id_private_key_pem_filepath := key_file_path + "/pastel_id_legroast_private_key.pem"

	pastel_id_public_key_b16_encoded := ""
	pastel_id_private_key_b16_encoded := ""
	pastel_id_public_key_export_format := ""
	pastel_id_private_key_export_format__encrypted := ""
	otp_correct := false

	infoPK, errPK := os.Stat(pastel_id_public_key_pem_filepath)
	infoSK, errSK := os.Stat(pastel_id_private_key_pem_filepath)
	if !os.IsNotExist(errPK) && !infoPK.IsDir() && !os.IsNotExist(errSK) && !infoSK.IsDir() {
		pastel_id_public_key_export_data, errPK := ioutil.ReadFile(pastel_id_public_key_pem_filepath)
		pastel_id_private_key_export_data_encrypted, errSK := ioutil.ReadFile(pastel_id_private_key_pem_filepath)
		if errPK == nil && errSK == nil {
			pastel_id_public_key_export_format = string(pastel_id_public_key_export_data)
			pastel_id_private_key_export_format__encrypted = string(pastel_id_private_key_export_data_encrypted)

			otp_string := generateCurrentOtpString()
			if otp_string == "" {
				otp_string = generateCurrentOtpStringFromUserInput()
			}
			fmt.Println("\n\nPlease Enter your pastel Google Authenticator Code:")
			fmt.Println(otp_string)

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			otp_from_user_input := scanner.Text()

			if len(otp_from_user_input) != 6 {
				return "", ""
			}
			otp_correct = (otp_from_user_input == otp_string)
		}
	}

	if otp_correct {
		box_key, _ := naclBoxKeyFromFile(box_key_file_path)
		var key [32]byte
		copy(key[:], box_key)
		pastel_id_private_key_export_format, _ := secretbox.EasyOpen(([]byte)(pastel_id_private_key_export_format__encrypted[:]), &key)
		pastel_id_private_key := strings.ReplaceAll(string(pastel_id_private_key_export_format), "-----BEGIN LEGROAST PRIVATE KEY-----\n", "")
		pastel_id_private_key_b16_encoded = strings.ReplaceAll(string(pastel_id_private_key), "\n-----END LEGROAST PRIVATE KEY-----", "")

		pastel_id_public_key_export_format = strings.ReplaceAll(pastel_id_public_key_export_format, "-----BEGIN LEGROAST PUBLIC KEY-----\n", "")
		pastel_id_public_key_b16_encoded = strings.ReplaceAll(pastel_id_public_key_export_format, "\n-----END LEGROAST PUBLIC KEY-----", "")
	}
	return pastel_id_public_key_b16_encoded, pastel_id_private_key_b16_encoded
}

func hide_signature_image_in_sample_image_func(sample_image_file_path string, signature_layer_image_output_filepath string, signed_image_output_path string) {
	img, err := gg.LoadImage(sample_image_file_path)
	if err != nil {
		log.Fatal(err)
	}

	signature_layer_image_data, err := ioutil.ReadFile(signature_layer_image_output_filepath)
	if err != nil {
		panic(err)
	}

	w := new(bytes.Buffer)
	err = steganography.Encode(w, img, signature_layer_image_data)
	if err != nil {
		panic(err)
	}
	outFile, _ := os.Create(signed_image_output_path)
	w.WriteTo(outFile)
	outFile.Close()
}

func extract_signature_image_in_sample_image_func(signed_image_output_path string, extracted_signature_layer_image_output_filepath string) {
	img, err := gg.LoadImage(signed_image_output_path)
	if err != nil {
		panic(err)
	}

	sizeOfMessage := steganography.GetMessageSizeFromImage(img)

	decodedData := steganography.Decode(sizeOfMessage, img)
	err = os.WriteFile(extracted_signature_layer_image_output_filepath, decodedData, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	if _, err := os.Stat(OTPSecretFile); os.IsNotExist(err) {
		if err := setupGoogleAuthenticatorForPrivateKeyEncryption(); err != nil {
			panic(err)
		}
	}

	box_key_file_path := "box_key.bin"
	if _, err := os.Stat(box_key_file_path); os.IsNotExist(err) {
		if err := generateAndStoreKeyForNacl(); err != nil {
			panic(err)
		}
	}

	sample_image_file_path := "sample_image2.png"

	fmt.Printf("\nApplying signature to file %v", sample_image_file_path)
	sha256_hash_of_image_to_sign, err := get_image_hash_from_image_file_path_func(sample_image_file_path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nSHA256 Hash of Image File: %v", sha256_hash_of_image_to_sign)
	pastel_id_public_key_b16_encoded, pastel_id_private_key_b16_encoded := import_pastel_public_and_private_keys_from_pem_files_func(box_key_file_path)
	if pastel_id_public_key_b16_encoded == "" {
		pastel_id_private_key_b16_encoded, pastel_id_public_key_b16_encoded = pastel_id_keypair_generation_func()
		write_pastel_public_and_private_key_to_file_func(pastel_id_public_key_b16_encoded, pastel_id_private_key_b16_encoded, box_key_file_path)
	}
	keypairImgs, err := generateKeypairQRs(pastel_id_public_key_b16_encoded, pastel_id_private_key_b16_encoded)
	if err != nil {
		panic(err)
	}

	pastel_id_signature_b16_encoded := pastel_id_write_signature_on_data_func(sha256_hash_of_image_to_sign, pastel_id_private_key_b16_encoded, pastel_id_public_key_b16_encoded)
	_ = pastel_id_verify_signature_with_public_key_func(sha256_hash_of_image_to_sign, pastel_id_signature_b16_encoded, pastel_id_public_key_b16_encoded)

	timestamp := time.Now().Format("Jan_02_2006_15_04_05")
	signatureImags, err := qr.Encode(pastel_id_signature_b16_encoded, PastelIdSignatureFilesFolder, "Pastel Signature", "pastel_id_legroast_signature_qr_code", timestamp)
	if err != nil {
		panic(err)
	}

	imgsToMap := append(keypairImgs, signatureImags...)

	signatureLayerImageOutputFilepath := filepath.Join(PastelIdSignatureFilesFolder, fmt.Sprintf("Complete_Signature_Image_Layer__%v.png", timestamp))
	sample_image, err := gg.LoadImage(sample_image_file_path)
	if err != nil {
		panic(err)
	}
	err = qr.MapImages(imgsToMap, sample_image.Bounds().Size(), signatureLayerImageOutputFilepath)
	if err != nil {
		panic(err)
	}

	signed_image_output_path := "final_watermarked_image.png"
	hide_signature_image_in_sample_image_func(sample_image_file_path, signatureLayerImageOutputFilepath, signed_image_output_path)

	extracted_signature_layer_image_output_filepath := "extracted_signature_image.png"
	extract_signature_image_in_sample_image_func(signed_image_output_path, extracted_signature_layer_image_output_filepath)
}
