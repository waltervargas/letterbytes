package letterbytes

func Run(scanner string, fileName string) error {
	img, err := scan(scanner)
	if err != nil {
		return err
	}

	err = toPNG(img, fileName)
	if err != nil {
		return err
	}

	err = toTIFF(img, fileName)
	if err != nil {
		return err
	}

	err = toTXT(img, fileName)
	if err != nil {
		return err
	}
	return nil
}
