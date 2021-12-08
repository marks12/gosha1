package cmd

const usualRegionality = `package types

import (
    "{ms-name}/core"
    "{ms-name}/dbmodels"
)

type Regionality struct {
    RegionId   int
    LanguageId int

    region   dbmodels.Region
    language dbmodels.Language
}

func (r *Regionality) GetCurrentRegion() (res dbmodels.Region) {
    if r.RegionId != r.region.ID {
        core.Db.Where(dbmodels.Region{ID: r.RegionId}).First(&r.region)
    }
    return r.region
}

func (r *Regionality) GetCurrentLanguage() (res dbmodels.Language) {
    if r.LanguageId != r.language.ID {
        core.Db.Where(dbmodels.Language{ID: r.LanguageId}).First(&r.language)
    }
    return r.language
}

func (r Regionality) GetLanguageId() int {
    return r.LanguageId
}


`

var usualTemplateRegionality = template{
    Path:    "./types/regionality.go",
    Content: usualRegionality,
}
