// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PipelineBodyV0031 struct {
	Coordinates               []string `multipartForm:"name=coordinates"`
	Encoding                  []string `multipartForm:"name=encoding"`
	Files                     [][]byte `multipartForm:"name=files,json"`
	GzUncompressedContentType *string  `multipartForm:"name=gz_uncompressed_content_type"`
	HiResModelName            []string `multipartForm:"name=hi_res_model_name"`
	OcrLanguages              []string `multipartForm:"name=ocr_languages"`
	OutputFormat              *string  `multipartForm:"name=output_format"`
	PdfInferTableStructure    []string `multipartForm:"name=pdf_infer_table_structure"`
	Strategy                  []string `multipartForm:"name=strategy"`
	XMLKeepTags               []string `multipartForm:"name=xml_keep_tags"`
}

func (o *PipelineBodyV0031) GetCoordinates() []string {
	if o == nil {
		return nil
	}
	return o.Coordinates
}

func (o *PipelineBodyV0031) GetEncoding() []string {
	if o == nil {
		return nil
	}
	return o.Encoding
}

func (o *PipelineBodyV0031) GetFiles() [][]byte {
	if o == nil {
		return nil
	}
	return o.Files
}

func (o *PipelineBodyV0031) GetGzUncompressedContentType() *string {
	if o == nil {
		return nil
	}
	return o.GzUncompressedContentType
}

func (o *PipelineBodyV0031) GetHiResModelName() []string {
	if o == nil {
		return nil
	}
	return o.HiResModelName
}

func (o *PipelineBodyV0031) GetOcrLanguages() []string {
	if o == nil {
		return nil
	}
	return o.OcrLanguages
}

func (o *PipelineBodyV0031) GetOutputFormat() *string {
	if o == nil {
		return nil
	}
	return o.OutputFormat
}

func (o *PipelineBodyV0031) GetPdfInferTableStructure() []string {
	if o == nil {
		return nil
	}
	return o.PdfInferTableStructure
}

func (o *PipelineBodyV0031) GetStrategy() []string {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *PipelineBodyV0031) GetXMLKeepTags() []string {
	if o == nil {
		return nil
	}
	return o.XMLKeepTags
}