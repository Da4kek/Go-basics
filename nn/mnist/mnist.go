package mnist

import (
	"gonum.org/v1/hdf5"
)

const (
	mnistFilePath      = "path/to/mnist/dataset/file.h5"
	mnistImagesDataset = "images"
	mnistLabelsDataset = "labels"
)

func ReadMNIST() ([][]float64, []uint8, error) {
	file, err := hdf5.OpenFile(mnistFilePath, hdf5.F_ACC_RDONLY)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	imagesDataset, err := file.OpenDataset(mnistImagesDataset)
	if err != nil {
		return nil, nil, err
	}
	defer imagesDataset.Close()

	labelsDataset, err := file.OpenDataset(mnistLabelsDataset)
	if err != nil {
		return nil, nil, err
	}
	defer labelsDataset.Close()

	imagesDims, _, err := imagesDataset.Space().SimpleExtentDims()
	if err != nil {
		return nil, nil, err
	}
	numImages := imagesDims[0]
	imageWidth := imagesDims[1]
	imageHeight := imagesDims[2]
	imagesData := make([]float64, numImages*imageWidth*imageHeight)
	err = imagesDataset.Read(&imagesData)
	if err != nil {
		return nil, nil, err
	}

	labelsData := make([]uint8, numImages)
	err = labelsDataset.Read(&labelsData)
	if err != nil {
		return nil, nil, err
	}

	images := make([][]float64, numImages)
	for i := range images {
		images[i] = imagesData[i*int(imageWidth)*int(imageHeight) : (i+1)*int(imageWidth)*int(imageHeight)]
	}

	return images, labelsData, nil
}
